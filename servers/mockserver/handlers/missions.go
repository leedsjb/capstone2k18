package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
)

// Mission ...
type Mission struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	Vision           string `json:"vision"`
	NextWaypointETE  string `json:"nextWaypointETE"`
	FlightNum        string `json:"flightNum"`
	Status           string `json:"status"`
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
	Vision           string `json:"vision"`
	RadioReport      string `json:"radioReport"`
	Crew             string `json:"crew"`
	Requestor        string `json:"requestor"`
	AircraftID       int    `json:"aircraftID"`
	AircraftNNum     string `json:"aircraftNNum"`
	AircraftCallsign string `json:"aircraftCallsign"`
}

var missionDetails = []*MissionDetail{
	{
		ID:               1,
		Type:             "RW-SCENE",
		NextWaypointETE:  "x min to...",
		FlightNum:        "18-0013",
		Status:           "ongoing",
		RadioReport:      "18-0013, 65, 90, male, GSW to chest. Has chest tube., Yes, 4, Paced externally - bring pacer box, Upper GI Bleed, Less than 5cm - launch without AOC Notification",
		Crew:             "First Last, First Last, First Last",
		Requestor:        "First Last",
		Vision:           "IFR",
		AircraftID:       3,
		AircraftNNum:     "N951AL",
		AircraftCallsign: "AL2",
	},
}

// MissionsHandler ...
func MissionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		missions := []*Mission{}
		for _, v := range missionDetails {
			m := &Mission{
				ID:               v.ID,
				Type:             v.Type,
				Vision:           v.Vision,
				NextWaypointETE:  v.NextWaypointETE,
				FlightNum:        v.FlightNum,
				Status:           v.Status,
				AircraftNNum:     v.AircraftNNum,
				AircraftCallsign: v.AircraftCallsign,
			}
			missions = append(missions, m)
		}
		respond(w, missions)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// MissionDetailHandler ...
func MissionDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding ID: %v", err), http.StatusBadRequest)
		return
	}
	var md *MissionDetail
	for _, v := range missionDetails {
		if v.ID == id {
			md = v
			break
		}
	}
	if md == nil {
		http.Error(w, "No mission with that ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		respond(w, md)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
