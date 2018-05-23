package parsers

import (
	"database/sql"

	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
)

//HandlerContext provides handlers with extra information
type ParserContext struct {
	AircraftTrie *indexes.Trie
	GroupsTrie   *indexes.Trie
	PeopleTrie   *indexes.Trie
	DB           *sql.DB
	Notifier     *handlers.Notifier
}

//NewHandlerContext creates a new HandlerContext
func NewParserContext(aircraftTrie *indexes.Trie, groupsTrie *indexes.Trie, peopleTrie *indexes.Trie, db *sql.DB, notifier *handlers.Notifier) *ParserContext {
	if aircraftTrie == nil || groupsTrie == nil || peopleTrie == nil || db == nil || notifier == nil {
		panic("Missing aircraft trie, groups trie, peopleTrie, db, or notifier")
	}
	return &ParserContext{aircraftTrie, groupsTrie, peopleTrie, db, notifier}
}
