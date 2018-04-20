package handlers

import (
	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/indexes"
	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/models/users"
	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/sessions"
)

//HandlerContext provides handlers with extra information
type HandlerContext struct {
	SigningKey   string
	SessionStore *sessions.RedisStore
	UserStore    *users.MongoStore
	Trie         *indexes.Trie
}

//NewHandlerContext creates a new HandlerContext
func NewHandlerContext(signingKey string, sessionStore *sessions.RedisStore,
	userStore *users.MongoStore, trie *indexes.Trie) *HandlerContext {
	if signingKey == "" || sessionStore == nil || userStore == nil || trie == nil {
		panic("Missing either a signing key, session store, user store or trie")
	}
	return &HandlerContext{signingKey, sessionStore, userStore, trie}
}
