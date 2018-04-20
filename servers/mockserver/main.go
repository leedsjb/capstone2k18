package main

import (
	"log"
	"net/http"

	"github.com/leedsjb/capstone2k18/servers/mockserver/handlers"
)

func main() {
	addr := ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("/aircraft", handlers.AircraftHandler)
	mux.HandleFunc("/aircraft/", handlers.AircraftDetailHandler)
	mux.HandleFunc("/people", handlers.PeopleHandler)
	mux.HandleFunc("/people/", handlers.PersonDetailHandler)
	// change to /users/me and handlers.UsersMeHandler
	mux.HandleFunc("/people/me", handlers.PeopleMeHandler)
	mux.HandleFunc("/groups", handlers.GroupsHandler)
	mux.HandleFunc("/groups/", handlers.GroupDetailHandler)
	mux.HandleFunc("/resources/", handlers.ResourcesHandler)

	wrappedMux := handlers.NewCORSHandler(mux)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, wrappedMux))
}
