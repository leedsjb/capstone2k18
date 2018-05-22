package handlers

import (
	"database/sql"

	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
)

//HandlerContext provides handlers with extra information
type HandlerContext struct {
	AircraftTrie *indexes.Trie
	// PersonnelTrie *indexes.Trie
	GroupsTrie *indexes.Trie
	PeopleTrie *indexes.Trie
	DB         *sql.DB
}

//NewHandlerContext creates a new HandlerContext
func NewHandlerContext(aircraftTrie *indexes.Trie, groupsTrie *indexes.Trie, peopleTrie *indexes.Trie, db *sql.DB) *HandlerContext {
	if aircraftTrie == nil || groupsTrie == nil || peopleTrie == nil || db == nil {
		panic("Missing aircraft trie, groups trie, people trie, or db")
	}
	return &HandlerContext{aircraftTrie, groupsTrie, peopleTrie, db}
}
