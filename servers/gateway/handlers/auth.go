package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/sessions"

	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/models/users"
)

//UsersHandler handles requests for the "users" resource, and allows clients
//to create new user accounts.
func (ctx *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//The request body must contain
		//JSON that can be decoded into a users.NewUser struct.
		ctype := r.Header.Get(headerContentType)
		if !strings.HasPrefix(ctype, contentTypeJSON) {
			http.Error(w, "Content type must be application/json", http.StatusBadRequest)
			return
		}
		//Decode the request body into a users.NewUser struct
		nu := &users.NewUser{}
		if err := json.NewDecoder(r.Body).Decode(nu); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		//Validate the NewUser
		if err := nu.Validate(); err != nil {
			http.Error(w, "Error validating user", http.StatusBadRequest)
			return
		}
		//Ensure there isn't already a user in the user store with the same email address
		if _, err := ctx.UserStore.GetByEmail(nu.Email); err == nil {
			http.Error(w, "Email is already in use", http.StatusBadRequest)
			return
		}
		//Ensure there isn't already a user in the user store with the same user name
		if _, err := ctx.UserStore.GetByUserName(nu.UserName); err == nil {
			http.Error(w, "Username is already in use", http.StatusBadRequest)
			return
		}
		//Insert the new user into the user store
		ru, err := ctx.UserStore.Insert(nu)
		if err != nil {
			http.Error(w, "Error creating account", http.StatusInternalServerError)
			return
		}
		//Insert the new user into the trie
		if err := users.AddUserToTrie(ctx.Trie, ru); err != nil {
			http.Error(w, "Error indexing user", http.StatusInternalServerError)
			return
		}
		sessionState := &SessionState{time.Now(), ru}
		//Begin a new session
		if _, err = sessions.BeginSession(ctx.SigningKey, ctx.SessionStore,
			sessionState, w); err != nil {
			http.Error(w, fmt.Sprintf("Error starting session: %v", err), http.StatusInternalServerError)
			return
		}
		//Respond to the client with an http.StatusCreated status code,
		//and the users.User struct returned from the user store insert
		//method encoded as a JSON object
		w.WriteHeader(http.StatusCreated)
		respond(w, ru)

	case "GET":
		//Get the current user from the session state
		sessionState := &SessionState{}
		//If there is an error getting the session state, respond with an
		//http.StatusUnauthorized error.
		_, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, sessionState)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting session state: %v", err), http.StatusUnauthorized)
			return
		}
		term := r.URL.Query().Get("q")
		//Check if the user typed in a key
		if len(term) == 0 {
			respond(w, []bson.ObjectId{})
		} else {
			userIDS := ctx.Trie.GetUsers(strings.ToLower(term), 20)
			users, err := ctx.UserStore.GetUsers(userIDS)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error getting users: %v", err), http.StatusInternalServerError)
				return
			}
			respond(w, users)
		}

	default:
		http.Error(w, "Method must be POST", http.StatusMethodNotAllowed)
		return
	}
}

//UsersMeHandler handles requests for the "current user" resource.
func (ctx *HandlerContext) UsersMeHandler(w http.ResponseWriter, r *http.Request) {
	//Get the current user from the session state
	sessionState := &SessionState{}
	//If there is an error getting the session state, respond with an
	//http.StatusUnauthorized error.
	sid, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, sessionState)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting session state: %v", err), http.StatusUnauthorized)
		return
	}
	ru := sessionState.AuthenticatedUser

	//Support two different HTTP methods: GET and PATCH
	switch r.Method {
	case "GET":
		//Respond with the current user encoded as JSON object.
		respond(w, ru)

	case "PATCH":
		//PATCH: update the current user with the JSON in the request body
		updates := &users.Updates{}
		if err := json.NewDecoder(r.Body).Decode(updates); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		oldFirstName := ru.FirstName
		oldLastName := ru.LastName

		if err := ru.ApplyUpdates(updates); err != nil {
			http.Error(w, fmt.Sprintf("Error updating user: %v", err), http.StatusBadRequest)
			return
		}
		err := ctx.UserStore.Update(ru.ID, updates)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating user: %v", err), http.StatusInternalServerError)
			return
		}
		uu, err := ctx.UserStore.GetByID(ru.ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving updated user: %v", err), http.StatusInternalServerError)
			return
		}
		//Unindex user
		if err := ctx.Trie.RemoveUser(strings.ToLower(ru.Email), ru.ID); err != nil {
			http.Error(w, fmt.Sprintf("Error unindexing user: %v", err), http.StatusInternalServerError)
			return
		}
		if err := ctx.Trie.RemoveUser(strings.ToLower(ru.UserName), ru.ID); err != nil {
			http.Error(w, fmt.Sprintf("Error unindexing user: %v", err), http.StatusInternalServerError)
			return
		}
		if err := ctx.Trie.RemoveUser(strings.ToLower(oldFirstName), ru.ID); err != nil {
			http.Error(w, fmt.Sprintf("Error unindexing user: %v", err), http.StatusInternalServerError)
			return
		}
		if err := ctx.Trie.RemoveUser(strings.ToLower(oldLastName), ru.ID); err != nil {
			http.Error(w, fmt.Sprintf("Error unindexing user: %v", err), http.StatusInternalServerError)
			return
		}
		//Reindex user
		if err := users.AddUserToTrie(ctx.Trie, uu); err != nil {
			http.Error(w, fmt.Sprintf("Error indexing user: %v", err), http.StatusInternalServerError)
			return
		}
		//Update current user data in your session store
		sessionState.AuthenticatedUser = uu
		if err = ctx.SessionStore.Save(sid, sessionState); err != nil {
			http.Error(w, fmt.Sprintf("Error updating session: %v", err), http.StatusInternalServerError)
			return
		}
		//Respond with the newly updated user, encoded as a JSON object
		respond(w, uu)

	default:
		http.Error(w, "Method must be GET or PATCH", http.StatusMethodNotAllowed)
		return
	}
}

//SessionsHandler handles requests for the "sessions" resource, and allows clients
//to begin a new session using an existing user's credentials.
func (ctx *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//The method must be POST
	case "POST":
		//The request and the body must contain JSON that can be decoded
		//into a users.Credentials struct.
		ctype := r.Header.Get(headerContentType)
		if !strings.HasPrefix(ctype, contentTypeJSON) {
			http.Error(w, "Content type must be application/json", http.StatusBadRequest)
			return
		}
		//Decode the request body into a users.Credentials struct
		c := &users.Credentials{}
		if err := json.NewDecoder(r.Body).Decode(c); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		//Get the user with the provided email from the UserStore; if not found,
		//respond with a http.StatusUnauthorized error and the message "invalid credentials"
		ru, err := ctx.UserStore.GetByEmail(c.Email)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		//Authenticate the user using the provided password; if that fails, respond
		//with an http.StatusUnauthorized error and the message "invalid credentials"
		if err := ru.Authenticate(c.Password); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		//Begin a new session
		sessionState := &SessionState{time.Now(), ru}
		sessions.BeginSession(ctx.SigningKey, ctx.SessionStore, sessionState, w)
		//Respond to the client with the User encoded as a JSON object
		respond(w, ru)

	default:
		http.Error(w, "Method must be POST", http.StatusMethodNotAllowed)
		return
	}
}

//SessionsMineHandler handles requests for the "current session" resource,
//and allows clients to end that session.
func (ctx *HandlerContext) SessionsMineHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		sessionState := &SessionState{}
		//If there is an error getting the session state, respond with an http.StatusUnauthorized error.
		if _, err := sessions.GetState(r, ctx.SigningKey, ctx.SessionStore, sessionState); err != nil {
			http.Error(w, "Error getting session state", http.StatusUnauthorized)
			return
		}
		//End the current session
		if _, err := sessions.EndSession(r, ctx.SigningKey, ctx.SessionStore); err != nil {
			http.Error(w, "Error signing out", http.StatusInternalServerError)
			return
		}
		//Respond with the string "signed out"
		respond(w, "signed out")

	default:
		http.Error(w, "Method must be DELETE", http.StatusMethodNotAllowed)
		return
	}
}
