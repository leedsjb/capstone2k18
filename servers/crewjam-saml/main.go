package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/crewjam/saml/samlsp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", samlsp.Token(r.Context()).Attributes.Get("cn"))
}

/*

	Generate certs:
	openssl req -x509 -newkey rsa:2048 -keyout myservice.key -out myservice.crt -days 365 -nodes -subj "/CN=myservice.example.com"

*/

func main() {

	PWD := os.Getenv("PWD")

	ENV := os.Getenv("environment")

	if ENV == "" {
		panic("error: set environment environment variable")
	}

	var tlscert string
	var tlskey string
	var host string
	var port string

	if ENV == "local-dev" { // local dev
		tlscert = PWD + "/tls/fullchain1.pem"
		tlskey = PWD + "/tls/privkey1.pem"
		host = ""
		port = "4430"
	} else if ENV == "local-docker-dev" {
		tlscert = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain1.pem"
		tlskey = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey1.pem"
		host = ""
		port = "443"
	} else if ENV == "do" { // digital ocean
		tlscert = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain.pem"
		tlskey = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey.pem"
		host = ""
		port = "443"
	} else if ENV == "kubernetes" {
		tlscert = "" // TODO: determine location for kubernetes
		tlskey = ""  // TODO: determine location for kubernetes
		host = ""
		port = "80"
	}

	keyPair, err := tls.LoadX509KeyPair(tlscert, tlskey)
	if err != nil {
		panic(err)
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err)
	}

	idpMetadataURL, err := url.Parse("http://www.testshib.org/metadata/testshib-providers.xml") //idp.u.washington.edu
	if err != nil {
		panic(err)
	}

	rootURL := &url.URL{} // ****** correct usage of &?

	if ENV == "kubernetes" {
		rootURL, err = url.Parse("https://test.elevate.airliftnw.org")
		if err != nil {
			panic(err)
		}
	} else {
		rootURL, err = url.Parse("https://crewjam-saml.test.elevate.emeloid.co")
		if err != nil {
			panic(err)
		}
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:            *rootURL,
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
		IDPMetadataURL: idpMetadataURL,
	})

	app := http.HandlerFunc(hello)                    // define function to call for requests to http endpoint
	http.Handle("/hello", samlSP.RequireAccount(app)) // direct requests to /hello to app
	http.Handle("/saml/", samlSP)                     // direct requests to the /saml/ route to samlSP

	addr := host + ":" + port
	fmt.Println("current set addr is : " + addr)

	var listenServeErr error

	if ENV == "kubernetes" {
		listenServeErr = http.ListenAndServe(addr, nil) // 2nd arg: handler, nil -> default serve mux

		fmt.Println("crewjam listening on port 80")
	} else {
		listenServeErr = http.ListenAndServeTLS(addr, tlscert, tlskey, nil)
		fmt.Println("crewjam listening on port 443")
	}

	if listenServeErr != nil {
		panic(listenServeErr)
	}

}
