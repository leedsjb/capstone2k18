package main

import (
	"log"
	"net/http"
)

// HelloHandler ...
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

func main() {
	addr := ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
