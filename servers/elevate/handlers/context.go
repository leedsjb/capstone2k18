package handlers

import (
	"database/sql"

	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
)

//HandlerContext provides handlers with extra information
type HandlerContext struct {
	AircraftTrie  *indexes.Trie
	PersonnelTrie *indexes.Trie
	DB            *sql.DB
}

//NewHandlerContext creates a new HandlerContext
func NewHandlerContext(aircraftTrie *indexes.Trie, personnelTrie *indexes.Trie, db *sql.DB) *HandlerContext {
	if aircraftTrie == nil || personnelTrie == nil || db == nil {
		panic("Missing aircraft trie, personnel trie, or db")
	}
	return &HandlerContext{aircraftTrie, personnelTrie, db}
}
