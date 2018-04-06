package handlers

import (
	"net/http"
)

// Mission ...
type Mission struct {
	id               int
	nNum             string
	missionType      string
	status           string
	flightNum        string
	aircraftCallsign string
}

// MissionsHandler ...
func MissionsHandler(w http.ResponseWriter, r *http.Request) {
	missions := []*Mission{}

	switch r.Method {
	case "GET":

	default:
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
