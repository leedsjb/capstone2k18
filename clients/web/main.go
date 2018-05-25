/*
Filename: main.go
Created:
Modified: Thursday May 24, 2018
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
	// samlspsecure "github.com/edaniels/go-saml"
)

/*
	Generate certs:
	openssl req -x509 -newkey rsa:2048 -keyout myservice.key -out myservice.crt -days 365 -nodes -subj "/CN=myservice.example.com"
*/

// trivial protected page resource
// must be signed in via SAML SSO to access
// prints information about signed is session to page
// from the request object
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", samlsp.Token(r.Context()).Attributes.Get("cn"))

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

func RootPathHandler(w http.ResponseWriter, r *http.Request) {

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

	if r.Method != "GET" {
		http.Error(w, "must be a get request", http.StatusBadRequest)
		return
	}
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

func main() {

	// rs, err := http.Get("https://google.com")
	// // Process response
	// if err != nil {
	// 	panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
	// }
	// defer rs.Body.Close()

	// bodyBytes, err := ioutil.ReadAll(rs.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// bodyString := string(bodyBytes)

	// idpMetadataURLTEST, err := url.Parse("https://idp.u.washington.edu/metadata/idp-metadata.xml")
	// idpMetadataURLTEST, err := url.Parse("https://google.com")

	// rootCAs, _ := x509.SystemCertPool() // retrieves RootCA list from system
	// if rootCAs == nil {
	// 	fmt.Println("new ******")
	// 	rootCAs = x509.NewCertPool()
	// }

	// fmt.Println(rootCAs)

	// config := &tls.Config{
	// 	InsecureSkipVerify: false,
	// 	RootCAs:            rootCAs,
	// }
	// tr := &http.Transport{TLSClientConfig: config}
	// client := &http.Client{Transport: tr}

	// req, err := http.NewRequest("GET", idpMetadataURLTEST.String(), nil)
	// if err != nil {
	// 	errString := fmt.Errorf("error: %v", err)
	// 	log.Fatalln(errString)
	// }

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	errString := fmt.Errorf("error on line 146: %v", err)
	// 	log.Fatalln(errString)
	// }
	// // defer resp.Body.Close()

	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	PWD := os.Getenv("PWD")
	ENV := os.Getenv("ENVIRONMENT")

	if ENV == "" {
		panic("error: set environment environment variable")
	}

	var tlscert string
	var tlskey string
	var host string
	var port string

	if ENV == "local-dev" { // local dev
		log.Printf("In: local-dev")
		tlscert = PWD + "/tls/fullchain1.pem"
		tlskey = PWD + "/tls/privkey1.pem"
		host = ""
		port = "4430"
	} else if ENV == "local-docker-dev" {
		log.Printf("In: local-docker-dev")
		tlscert = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain1.pem"
		tlskey = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey1.pem"
		host = ""
		port = "443"
	} else if ENV == "do" { // digital ocean
		log.Printf("In: do")
		tlscert = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/fullchain.pem"
		tlskey = "/etc/letsencrypt/live/crewjam-saml.test.elevate.emeloid.co/privkey.pem"
		host = ""
		port = "443"
	} else if ENV == "kubernetes" {
		log.Printf("In: kubernetes")
		tlscert = "/etc/crewjam-secret-volume/test.elevate.airliftnw.org-fullchain1.pem"
		tlskey = "/etc/crewjam-secret-volume/test.elevate.airliftnw.org-privkey1.pem"
		host = ""
		port = "80"
	}

	log.Printf("tlscert is: %v", tlscert)
	log.Printf("tlskey is: %v", tlskey)

	// files, err := ioutil.ReadDir("/etc/crewjam-secret-volume/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	// temp, err := ioutil.ReadFile(tlscert)
	// if err != nil {
	// 	log.Printf("no read tlscert")
	// }
	// fmt.Print(string(temp))

	keyPair, err := tls.LoadX509KeyPair(tlscert, tlskey)
	if err != nil {
		log.Printf("Error loading keyPair: %v", err)
		// panic(err)
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		log.Printf("Error loading keyPair leaf: %v", err)
		// panic(err)
	}

	// idpMetadataURL, err := url.Parse("https://www.testshib.org/metadata/testshib-providers.xml")
	idpMetadataURL, err := url.Parse("https://idp.u.washington.edu/metadata/idp-metadata.xml")
	if err != nil {
		panic(err)
	}

	rootURL := &url.URL{} // ****** correct usage of &?

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

	mux := http.NewServeMux()

	// mux.HandleFunc("/", RootPathHandler)
	mux.Handle("/", http.FileServer(http.Dir("./build")))

	mux.HandleFunc("/testing/", testPathHandler)

	app := http.HandlerFunc(hello)                      // define handler to call for requests to protected path
	mux.Handle("/sign-in/", samlSP.RequireAccount(app)) // direct requests to /hello/* to hello to app and ensures clients are authenticated with SSO

	mux.Handle("/saml/", samlSP) // direct requests to the /saml/ route to samlSP

	addr := host + ":" + port
	fmt.Println("current set addr is : " + addr)

	var listenServeErr error

	if ENV == "kubernetes" {
		fmt.Println("crewjam listening at: " + addr)
		listenServeErr = http.ListenAndServe(addr, mux)

	} else {
		fmt.Println("crewjam listening at: " + addr)
		// listenServeErr = http.ListenAndServeTLS(addr, tlscert, tlskey, mux)
		listenServeErr = http.ListenAndServe(addr, mux)
	}

	if listenServeErr != nil {
		log.Printf("Unable to listen and serve: %v", listenServeErr)
		// panic(listenServeErr)
	}
}
