package handlers

import (
	"net/http"
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
	id int
}

/*
   "id": 1,
   "nNum": "N951AL",
   "type": "RW-SCENE",
   "status": "ongoing",
   "flightNum": "18-0013",
   "aircraftId": "3",
   "radioReport":
       "18-0013, 65, 90, male, GSW to chest. Has chest tube., Yes, 4, Paced externally - bring pacer box, Upper GI Bleed, Less than 5cm - launch without AOC Notification",
   "crew": "First Last, First Last, First Last",
   "requestor": "First Last"
*/

// MissionsHandler ...
func MissionsHandler(w http.ResponseWriter, r *http.Request) {
	missions := []*Mission{}

	m := &Mission{
		ID:               1,
		NNum:             "N951AL",
		MissionType:      "RW-SCENE",
		Status:           "ongoing",
		FlightNum:        "18-0013",
		AircraftCallsign: "AL42",
	}

	missions = append(missions, m)

	switch r.Method {
	case "GET":
		respond(w, missions)
	default:
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
