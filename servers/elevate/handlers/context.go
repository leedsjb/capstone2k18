package handlers

import "github.com/leedsjb/capstone2k18/servers/mockserver/indexes"

//HandlerContext provides handlers with extra information
type HandlerContext struct {
	AircraftTrie *indexes.Trie
}

//NewHandlerContext creates a new HandlerContext
func NewHandlerContext(aircraftTrie *indexes.Trie, signingKey string) *HandlerContext {
	if aircraftTrie == nil {
		panic("Missing aircraft trie")
	}
	return &HandlerContext{aircraftTrie}
}
