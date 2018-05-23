package handlers

import (
	"net/http"
)

//CORSHandler ensures that every handler deals with
//CORS appropriately
type CORSHandler struct {
	Handler http.Handler
}

func (ch *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Respond with the following headers to all requests:
	//Access-Control-Allow-Origin: *
	w.Header().Add("Access-Control-Allow-Origin", "*")
	//Access-Control-Allow-Headers: Content-Type, Authorization
	w.Header().Add(accessControlAllowHeaders, "Content-Type, Authorization")
	//Access-Control-Allow-Methods: GET, PUT, POST, PATCH, DELETE
	w.Header().Add(accessControlAllowMethods, "GET, PUT, POST, PATCH, DELETE")
	//Access-Control-Expose-Headers: Authorization
	w.Header().Add(accessControlExposeHeaders, "Authorization")
	//Access-Control-Max-Age: 600
	w.Header().Add(accessControlMaxAge, "600")

	//if this is preflight request, the method will
	//be OPTIONS, so call the real handler only if
	//the method is something else
	if r.Method != "OPTIONS" {
		ch.Handler.ServeHTTP(w, r)
	}
}

//NewCORSHandler creates a new CORSHandler
func NewCORSHandler(handlerToWrap http.Handler) *CORSHandler {
	return &CORSHandler{handlerToWrap}
}
