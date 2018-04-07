package handlers

import (
	"net/http"
	"path"
)

// Mission ...
type Mission struct {
	ID               int    `json:"id"`
	NNum             string `json:"nNum"`
	MissionType      string `json:"missionType"`
	Status           string `json:"status"`
	FlightNum        string `json:"flightNum"`
	AircraftCallsign string `json:"aircraftCallsign"`
}

// MissionDetail ...
type MissionDetail struct {
	ID               int    `json:"id"`
	NNum             string `json:"nNum"`
	MissionType      string `json:"missionType"`
	Status           string `json:"status"`
	FlightNum        string `json:"flightNum"`
	AircraftCallsign string `json:"aircraftCallsign"`
	RadioReport      string `json:"radioReport"`
	Crew             string `json:"crew"`
	Requestor        string `json:"requestor"`
}

var missions = []*Mission{
	{
		ID:               1,
		NNum:             "N951AL",
		MissionType:      "RW-SCENE",
		Status:           "ongoing",
		FlightNum:        "18-0013",
		AircraftCallsign: "AL42",
	},
}

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
