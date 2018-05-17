package handlers

import "net/http"

// PeopleMeHandler ...
func PeopleMeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// TODO: when auth, store active user
		// respond(w, personDetails[6])
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
