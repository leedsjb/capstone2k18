package handlers

import (
	"net/http"
	"path"
)

// Mission ...
type Mission struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	NextWaypointETE  string `json:"nextWaypointETE"`
	FlightNum        string `json:"flightNum"`
	Status           string `json:"status"`
	RadioReport      string `json:"radioReport"`
	Crew             string `json:"crew"`
	Requestor        string `json:"requestor"`
	AircraftID       int    `json:"aircraftID"`
	AircraftNNum     string `json:"aircraftNNum"`
	AircraftCallsign string `json:"aircraftCallsign"`
}

// MissionDetail ...
type MissionDetail struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	NextWaypointETE  string `json:"nextWaypointETE"`
	FlightNum        string `json:"flightNum"`
	Status           string `json:"status"`
	RadioReport      string `json:"radioReport"`
	Crew             string `json:"crew"`
	Requestor        string `json:"requestor"`
	AircraftID       int    `json:"aircraftID"`
	AircraftNNum     string `json:"aircraftNNum"`
	AircraftCallsign string `json:"aircraftCallsign"`
}

var missions = []*Mission{}

// MissionsHandler ...
func MissionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		respond(w, missions)
	default:
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// MissionDetailHandler ...
func MissionDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	print(id)
	switch r.Method {
	case "GET":
		return
	default:
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
