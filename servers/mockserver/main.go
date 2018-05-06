package main

import (
	"log"
	"net/http"

	"github.com/leedsjb/capstone2k18/servers/mockserver/handlers"
	"github.com/leedsjb/capstone2k18/servers/mockserver/indexes"
)

func main() {
	addr := ":4000"

	// This (and general trie usage in this mock server), is a
	// terrible way of doing things. Do not try this at home.
	var aircraftTrie = indexes.NewTrie()
	if err := handlers.LoadAircraftTrie(aircraftTrie); err != nil {
		log.Fatalf("Error loading aircraft trie")
	}

	handlerCtx := handlers.NewHandlerContext(aircraftTrie)

	mux := http.NewServeMux()
	mux.HandleFunc("/aircraft", handlerCtx.AircraftHandler)
	mux.HandleFunc("/aircraft/", handlers.AircraftDetailHandler)
	mux.HandleFunc("/people", handlers.PeopleHandler)
	mux.HandleFunc("/people/me", handlers.PeopleMeHandler)
	mux.HandleFunc("/people/", handlers.PersonDetailHandler)
	mux.HandleFunc("/groups", handlers.GroupsHandler)
	mux.HandleFunc("/groups/", handlers.GroupDetailHandler)
	mux.HandleFunc("/resources/", handlers.ResourcesHandler)

	wrappedMux := handlers.NewCORSHandler(mux)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
