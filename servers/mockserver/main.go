package main

import (
	"log"
	"net/http"

	"github.com/leedsjb/capstone2k18/servers/mockserver/handlers"
)

func main() {
	addr := ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("/missions", handlers.MissionsHandler)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
