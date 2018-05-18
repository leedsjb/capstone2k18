package parsers

import (
	"database/sql"

	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
)

//HandlerContext provides handlers with extra information
type ParserContext struct {
	AircraftTrie  *indexes.Trie
	PersonnelTrie *indexes.Trie
	DB            *sql.DB
}

//NewHandlerContext creates a new HandlerContext
func NewParserContext(aircraftTrie *indexes.Trie, personnelTrie *indexes.Trie, db *sql.DB) *ParserContext {
	if aircraftTrie == nil || personnelTrie == nil || db == nil {
		panic("Missing aircraft trie, personnel trie, or db")
	}
	return &ParserContext{aircraftTrie, personnelTrie, db}
}
