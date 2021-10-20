/*
Filename: main.go
Created:
Modified: Saturday May 26, 2018
Last Change: HTTPS redirect no longer applies to load balancer
Author: J. Benjamin Leeds
License: None
Purpose: This Go program serves as the HTTPS web server for ALNW elevate. It serves the Elevate
	React Web App and handles authentication via Shibboleth SAML with the UW IdP.
*/

package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	// "github.com/crewjam/saml"
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
	} else if ENV == "compute-engine" {
		log.Println("In: compute engine env")
		tlscert = "/etc/letsencrypt/live/elevate.benjaminleeds.com/fullchain.pem"
		tlskey = "/etc/letsencrypt/live/elevate.benjaminleeds.com/privkey.pem"
	}

	// generate tls.Certificate type from public and private key pair
	// the tls.Certificate type is from the go/crypto/tls library
	keyPair, err := tls.LoadX509KeyPair(tlscert, tlskey)
	if err != nil {
		log.Fatalf("Error loading keyPair: %v", err)
	}

	// generate the leaf certificate and store in the keyPair struct
	// the leaf certificate can reduce per handshake processing
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Fatalf("Error loading keyPair leaf: %v", err)
	}

	// UW IdP SSO Setup
	// -------------------------------------------------------------------------

	samlSP := setUpSaml(ENV, &keyPair)

	// Web Server Setup
	// -------------------------------------------------------------------------

	mux := http.NewServeMux() // create new mux instead of using default

	// handle "/" route
	// the "/" route serves a react application from the /build directory
	// the react application handles its own internal routing
	mux.Handle("/", http.FileServer(http.Dir("./build"))) // serve application

	// UW NetID Auth Components:
	mux.HandleFunc("/testing/", testPathHandler)

	// adapts hello func to serve as a request handler
	app := http.HandlerFunc(hello)

	// when requests arrive at the /sign-in/ route they are routed to the
	// samlSP middleware. Middleware ensures request has a valid session
	// associate with it. If not, it routes the user to the handler
	// associated w/ app to login the user
	mux.Handle("/sign-in/", samlSP.RequireAccount(app))

	mux.Handle("/saml/", samlSP) // direct requests to the /saml/ route to samlSP

	addr := host + ":" + port
	var listenServeErr error

	// create an EnsureHTTPS type which wraps the mux provided as an arg
	// this is like a constructor **
	wrappedMux := NewEnsureHTTPS(mux)

	// create an http.Server type from the address and wrappedMux
	srv := &http.Server{
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 5 * time.Second,
		Addr:    addr,
		Handler: wrappedMux,
	}

	if ENV == "kubernetes" {
		// in the kubernetes environment we use an HTTP server since
		// TLS connections are terminated at the GCP Load Balancer,
		// not at this application server
		fmt.Println("client server listening at: " + addr)
		listenServeErr = srv.ListenAndServe()
	} else {
		fmt.Println("client server listening at: " + addr)
		// listenServeErr = http.ListenAndServeTLS(addr, tlscert, tlskey, mux)
		listenServeErr = http.ListenAndServe(addr, mux)
	}

	if listenServeErr != nil {
		log.Fatalf("Unable to listen and serve: %v", listenServeErr)
	}
}

// End of Main function
// Helper methods below:
// -----------------------------------------------------------------------------

// define middleware handler
type EnsureHTTPS struct {
	handler http.Handler
}

/* This method is invoked by the http.Server type which listens for incoming
* requests and passes them to this handler when appropriate/
* Concrete implementation of the ServeHTTP method from the http.Handler
* interface. ServeHTTP takes an htttp.Request and http.ResponseWriter as
* params, reads the requests and responds via http.ResponseWriter
* This particular implementation ensures requests to client web server are HTTPS
* only with the exception of internal Google Cloud Load Balancer health checks
* which use HTTP
 */
func (ea *EnsureHTTPS) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Google Cloud Load Balancer is between client and this server. Therefore,
	// all TLS connections are terminated at the load balancer. We must check
	// whether the client connected via HTTP or HTTPS using the
	// X-Forwarded-Proto header the load balancer sets.

	// print entire request to log
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(requestDump))

	reqConnType := r.Header.Get("X-Forwarded-Proto")

	// empty string allows for health checks to bypass HTTPS redirect
	if reqConnType != "https" && reqConnType != "" {
		log.Printf("Non-https connection detected. Redirecting. ReqConnType: %v", reqConnType)

		w.Header().Set("Connection", "close")
		url := "https://" + r.Host + r.URL.String()
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return // critical: must return to prevent access to app via HTTP
	}

	log.Println("not redirecting")

	// this function operates on the http.Handler stored in the handler field
	// of the EnsureHTTPS type, ea, which is provided to this method
	// via reflection **??
	ea.handler.ServeHTTP(w, r)
}

/*
* Method takes an http.Handler as a parameter and returns pointer to EnsureHTTPS
* user type with the http.Handler stored in the handler field
 */
func NewEnsureHTTPS(handlerToWrap http.Handler) *EnsureHTTPS {
	return &EnsureHTTPS{handlerToWrap}
}

/* HTTP request handler for requests to the
* trivial protected page resource must be signed in via SAML SSO to access
// prints information about signed is session to page from the request object
*/
func hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %s", samlsp.Token(r.Context()).Attributes.Get("ePPN"))

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
	fmt.Printf("http_x_forwarded_proto header: %v \n", r.Header.Get("X-Forwarded-Proto"))
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
	w.Write([]byte("http_x_forwarded_proto header: " + r.Header.Get("X-Forwarded-Proto") + "\n"))
	w.Write([]byte("Google Cloud HTTPS L7 Load Balancer Health Check"))

	if r.Method != "GET" {
		http.Error(w, "must be a get request", http.StatusBadRequest)
		return
	}
}

func setUpSaml(ENV string, keyPair *tls.Certificate) *samlsp.Middleware {

	test := &http.Client{}

	// metadataServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Write()
	// })

	idpMetadataURL, err := url.Parse("https://idp.u.washington.edu/metadata/idp-metadata.xml")
	if err != nil {
		log.Fatalf("error parsing idpMetadataURL: %v", err)
	}

	samlEntityDescriptor, err := samlsp.FetchMetadata(context.TODO(), test, *idpMetadataURL)

	log.Println("*")
	log.Println(samlEntityDescriptor)

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
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: samlEntityDescriptor,
		// IDPMetadataURL: idpMetadataURL,
	})

	return samlSP

}
