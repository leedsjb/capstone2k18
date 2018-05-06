package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", samlsp.Token(r.Context()).Attributes.Get("cn"))
}

/*

	Generate certs:
	openssl req -x509 -newkey rsa:2048 -keyout myservice.key -out myservice.cert -days 365 -nodes -subj "/CN=myservice.example.com"

*/

func main() {
	fmt.Println("crewjam")

	keyPair, err := tls.LoadX509KeyPair("myservice.cert", "myservice.key")
	if err != nil {
		panic(err)
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err)
	}

	idpMetadataURL, err := url.Parse("https://www.testshib.org/metadata/testshib-providers.xml")
	if err != nil {
		panic(err)
	}
	rootURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		panic(err)
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
	http.ListenAndServe(":8000", nil)                 // 2nd arg: handler, nil -> default serve mux
}
