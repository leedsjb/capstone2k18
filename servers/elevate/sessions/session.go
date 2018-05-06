package sessions

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"
const paramAuthorization = "auth"
const schemeBearer = "Bearer "

//ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + headerAuthorization + " header")

//ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

//BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
//Authorization header to the response with the SessionID, and returns the new SessionID
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {
	//create a new SessionID
	si, err := NewSessionID(signingKey)
	if err != nil {
		return InvalidSessionID, fmt.Errorf("Error creating a new Session ID: %v", err)
	}

	//save the sessionState to the store
	if err := store.Save(si, sessionState); err != nil {
		return InvalidSessionID, fmt.Errorf("Error saving state: %v", err)
	}

	//add a header to the ResponseWriter that looks like this:
	//"Authorization: Bearer <sessionID>"
	//where "<sessionID>" is replaced with the newly-created SessionID
	w.Header().Add(headerAuthorization, schemeBearer+si.String())
	return si, nil
}

//GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {
	//get the value of the Authorization header,
	id := r.Header.Get(headerAuthorization)
	//or the "auth" query string parameter if no Authorization header is present
	if len(id) == 0 {
		id = r.URL.Query().Get(paramAuthorization)
	}
	//check scheme
	if !strings.HasPrefix(id, schemeBearer) {
		return InvalidSessionID, ErrInvalidScheme
	}
	id = strings.TrimPrefix(id, schemeBearer)
	//and validate it.
	si, err := ValidateID(id, signingKey)
	//If it is not valid, return the validation error.
	if err != nil {
		return InvalidSessionID, fmt.Errorf("Error validating session ID: %v", err)
	}
	//If it's valid, return the SessionID.
	return si, nil
}

//GetState extracts the SessionID from the request,
//gets the associated state from the provided store into
//the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {
	//get the SessionID from the request
	si, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, fmt.Errorf("Error getting SessionID from request: %v", err)
	}
	//get the data associated with that SessionID from the store.
	if err = store.Get(si, sessionState); err != nil {
		return InvalidSessionID, ErrStateNotFound
	}
	return si, nil
}

//EndSession extracts the SessionID from the request,
//and deletes the associated data in the provided store, returning
//the extracted SessionID.
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	//get the SessionID from the request
	si, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, fmt.Errorf("Error getting SessionID from the request: %v", err)
	}
	//and delete the data associated with it in the store.
	if err = store.Delete(si); err != nil {
		return InvalidSessionID, fmt.Errorf("Error deleting session data: %v", err)
	}
	return si, nil
}
