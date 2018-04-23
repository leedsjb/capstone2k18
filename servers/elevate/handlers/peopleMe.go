package handlers

import "net/http"

// PeopleMeHandler ...
func PeopleMeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		respond(w, personDetails[6])
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
