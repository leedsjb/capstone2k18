/*
Filename: main.go
Created: Wednesday May 23, 2018
Modified:
Author: J. Benjamin Leeds
License: None
Purpose: This Go program serves as the HTTPS web server for ALNW elevate. It serves the Elevate
	React Web App.
*/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		panic("Please set environment variable: ENVIRONMENT")
	}

	var addr string

	if environment == "local-dev" {
		addr = ":3000"
	} else if environment == "local-docker-dev" {
		addr = ":80"
	} else if environment == "kubernetes" {
		addr = ":80"
	} else {

	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./build")))

	fmt.Println("About to listen on: " + addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		panic(fmt.Errorf("Error starting web server: %v", err))
	}
}
