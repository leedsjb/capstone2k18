package parsers

import (
	"database/sql"

	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
)

//HandlerContext provides handlers with extra information
type ParserContext struct {
	AircraftTrie  *indexes.Trie
	PersonnelTrie *indexes.Trie
	DB            *sql.DB
	Notifier      *handlers.Notifier
}

//NewHandlerContext creates a new HandlerContext
func NewParserContext(aircraftTrie *indexes.Trie, personnelTrie *indexes.Trie, db *sql.DB, notifier *handlers.Notifier) *ParserContext {
	if aircraftTrie == nil || personnelTrie == nil || db == nil || notifier == nil {
		panic("Missing aircraft trie, personnel trie, db, or notifier")
	}
	return &ParserContext{aircraftTrie, personnelTrie, db, notifier}
}
