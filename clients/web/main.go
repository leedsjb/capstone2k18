/*
Filename: main.go
Created:
Modified: Saturday May 26, 2018
Author: J. Benjamin Leeds
License: None
Purpose: This Go program serves as the HTTPS web server for ALNW elevate. It serves the Elevate
	React Web App and handles authentication via Shibboleth SAML with the UW IdP.
*/

package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/crewjam/saml/samlsp"
)

func main() {

	PWD := os.Getenv("PWD")
	ENV := os.Getenv("ENVIRONMENT")

	if ENV == "" {
		log.Fatal("error: set environment environment variable")
	}

	var tlscert string
	var tlskey string
	var host string
	var port string

	if ENV == "local-dev" { // local dev
		log.Println("In: local-dev")
		tlscert = PWD + "/tls/fullchain1.pem"
		tlskey = PWD + "/tls/privkey1.pem"
		host = ""
		port = "4430"
	} else if ENV == "local-docker-dev" {
		log.Println("In: local-docker-dev")
		tlscert = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain1.pem"
		tlskey = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey1.pem"
		host = ""
		port = "443"
	} else if ENV == "do" { // digital ocean
		log.Println("In: do")
		tlscert = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain.pem"
		tlskey = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey.pem"
		host = ""
		port = "443"
	} else if ENV == "kubernetes" {
		log.Println("In: kubernetes")
		tlscert = "/etc/crewjam-secret-volume/test.elevate.airliftnw.org-fullchain1.pem"
		tlskey = "/etc/crewjam-secret-volume/test.elevate.airliftnw.org-privkey1.pem"
		host = ""
		port = "80"
	}

	keyPair, err := tls.LoadX509KeyPair(tlscert, tlskey)
	if err != nil {
		log.Fatalf("Error loading keyPair: %v", err)
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Fatalf("Error loading keyPair leaf: %v", err)
	}

	idpMetadataURL, err := url.Parse("https://idp.u.washington.edu/metadata/idp-metadata.xml")
	if err != nil {
		log.Fatalf("error parsing idpMetadataURL: %v", err)
	}

	rootURL := &url.URL{}

	if ENV == "kubernetes" {
		log.Println("K8S: setting url to https://test.elevate.airliftnw.org")
		rootURL, err = url.Parse("https://test.elevate.airliftnw.org")
		if err != nil {
			panic(err)
		}
	} else if ENV == "local-docker-dev" {
		rootURL, err = url.Parse("http://localhost:80")
		if err != nil {
			panic(err)
		}
	} else {
		rootURL, err = url.Parse("https://localhost:4430")
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

	mux := http.NewServeMux()                             // create new mux instead of using default
	mux.Handle("/", http.FileServer(http.Dir("./build"))) // serve application

	// UW NetID Auth Components:
	mux.HandleFunc("/testing/", testPathHandler)
	app := http.HandlerFunc(hello)
	mux.Handle("/sign-in/", samlSP.RequireAccount(app))
	mux.Handle("/saml/", samlSP) // direct requests to the /saml/ route to samlSP

	addr := host + ":" + port
	var listenServeErr error

	if ENV == "kubernetes" {
		fmt.Println("client server listening at: " + addr)
		listenServeErr = http.ListenAndServe(addr, mux)
	} else {
		fmt.Println("client server listening at: " + addr)
		// listenServeErr = http.ListenAndServeTLS(addr, tlscert, tlskey, mux)
		listenServeErr = http.ListenAndServe(addr, mux)
	}

	if listenServeErr != nil {
		log.Fatalf("Unable to listen and serve: %v", listenServeErr)
	}
}

// trivial protected page resource must be signed in via SAML SSO to access
// prints information about signed is session to page from the request object
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", samlsp.Token(r.Context()).Attributes.Get("ePPN"))

	fmt.Println("request URI: " + r.RequestURI)
	fmt.Println("request Method:" + r.Method)
	fmt.Println("request URL:" + r.URL.String())
	fmt.Println(r)
	fmt.Println("************")

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))

	w.Write([]byte("request URI: " + r.RequestURI + "\n"))
	w.Write([]byte("request Method:" + r.Method + "\n"))
	w.Write([]byte("request URL:" + r.URL.String() + "\n"))
	w.Write([]byte("request raw: " + string(requestDump) + "\n"))
	w.Write([]byte("Google Cloud HTTPS L7 Load Balancer Health Check"))
}

func testPathHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("testPathHandler")
	fmt.Println("request URI: " + r.RequestURI)
	fmt.Println("request Method:" + r.Method)
	fmt.Println("request URL:" + r.URL.String())
	fmt.Println(r)
	fmt.Println("************")

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(requestDump))

	w.Write([]byte("testPathHandler\n"))
	w.Write([]byte("request URI: " + r.RequestURI + "\n"))
	w.Write([]byte("request Method:" + r.Method + "\n"))
	w.Write([]byte("request URL:" + r.URL.String() + "\n"))
	w.Write([]byte("request raw: " + string(requestDump) + "\n"))
	w.Write([]byte("Google Cloud HTTPS L7 Load Balancer Health Check"))

	if r.Method != "GET" {
		http.Error(w, "must be a get request", http.StatusBadRequest)
		return
	}
}
