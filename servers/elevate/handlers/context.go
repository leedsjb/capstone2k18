package handlers

import (
	"database/sql"

	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
)

//HandlerContext provides handlers with extra information
type HandlerContext struct {
	AircraftTrie *indexes.Trie
	DB           *sql.DB
}

//NewHandlerContext creates a new HandlerContext
func NewHandlerContext(aircraftTrie *indexes.Trie, db *sql.DB) *HandlerContext {
	if aircraftTrie == nil || db == nil {
		panic("Missing aircraft trie or db")
	}
	return &HandlerContext{aircraftTrie, db}
}
