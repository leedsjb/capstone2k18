package handlers

import (
	"net/http"
)

// Aircraft ...
type Aircraft struct {
	id int
}

// AircraftDetail ...
type AircraftDetail struct {
	id int
}

// AircraftHandler ...
func AircraftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	default:
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
