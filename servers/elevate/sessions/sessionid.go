package sessions

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
)

//InvalidSessionID represents an empty, invalid session ID
const InvalidSessionID SessionID = ""

//idLength is the length of the ID portion
const idLength = 32

//signedLength is the full length of the signed session ID
//(ID portion plus signature)
const signedLength = idLength + sha256.Size

//SessionID represents a valid, digitally-signed session ID.
//This is a base64 URL encoded string created from a byte slice
//where the first `idLength` bytes are crytographically random
//bytes representing the unique session ID, and the remaining bytes
//are an HMAC hash of those ID bytes (i.e., a digital signature).
//The byte slice layout is like so:
//+-----------------------------------------------------+
//|...32 crypto random bytes...|HMAC hash of those bytes|
//+-----------------------------------------------------+
type SessionID string

//ErrInvalidID is returned when an invalid session id is passed to ValidateID()
var ErrInvalidID = errors.New("Invalid Session ID")

//NewSessionID creates and returns a new digitally-signed session ID,
//using `signingKey` as the HMAC signing key. An error is returned only
//if there was an error generating random bytes for the session ID
func NewSessionID(signingKey string) (SessionID, error) {
	//if `signingKey` is zero-length, return InvalidSessionID
	//and an error indicating that it may not be empty
	if len(signingKey) == 0 {
		return InvalidSessionID, fmt.Errorf("You might be missing a signing key")
	}

	//Generate a new digitally-signed SessionID by doing the following:
	//create a byte slice where the first `idLength` of bytes
	//are cryptographically random bytes for the new session ID,
	//and the remaining bytes are an HMAC hash of those ID bytes,
	//using the provided `signingKey` as the HMAC key.
	si := make([]byte, signedLength)
	if _, err := rand.Read(si[:idLength]); err != nil {
		return InvalidSessionID, fmt.Errorf("Error generating random bytes: %v", err)
	}

	h := hmac.New(sha256.New, []byte(signingKey))
	if _, err := h.Write(si[:idLength]); err != nil {
		return InvalidSessionID, fmt.Errorf("Error hashing ID bytes: %v", err)
	}
	copy(si[idLength:], h.Sum(nil))

	//encode that byte slice using base64 URL Encoding and return
	//the result as a SessionID type
	return SessionID(base64.URLEncoding.EncodeToString(si)), nil
}

//ValidateID validates the string in the `id` parameter
//using the `signingKey` as the HMAC signing key
//and returns an error if invalid, or a SessionID if valid
func ValidateID(id string, signingKey string) (SessionID, error) {
	//validate the `id` parameter using the provided `signingKey`.
	//base64 decode the `id` parameter,
	si, err := base64.URLEncoding.DecodeString(id)
	if err != nil {
		return InvalidSessionID, fmt.Errorf("Error base64-decoding: %v", err)
	}

	//HMAC hash the ID portion of the byte slice, and compare that to the
	//HMAC hash stored in the remaining bytes.
	h := hmac.New(sha256.New, []byte(signingKey))
	if _, err := h.Write(si[:idLength]); err != nil {
		return InvalidSessionID, fmt.Errorf("Error hashing ID bytes: %v", err)
	}
	newHash := h.Sum(nil)

	//If they match, return the entire `id` parameter as a SessionID type.
	if subtle.ConstantTimeCompare(newHash, si[idLength:]) == 1 {
		return SessionID(id), nil
	}
	//If not, return InvalidSessionID and ErrInvalidID
	return InvalidSessionID, ErrInvalidID
}

//String returns a string representation of the sessionID
func (sid SessionID) String() string {
	return string(sid)
}
