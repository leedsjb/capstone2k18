package handlers

import (
	"net/http"
)

// Aircraft ...
type Aircraft struct {
	ID            int    `json:"id"`
	Callsign      string `json:"callsign"`
	Crew          string `json:"crew"`
	LevelOfCare   string `json:"levelOfCare"`
	Class         string `json:"class"`
	Lat           string `json:"lat"`
	Long          string `json:"long"`
	Area          string `json:"area"`
	NNum          string `json:"nNum"`
	OOSReason     string `json:"OOSReason"`
	OOSRemaining  string `json:"OOSRemaining"`
	OOSDuration   string `json:"OOSDuration"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Vision        string `json:"vision"`
	MissionStatus string `json:"missionStatus"`
}

// AircraftDetail ...
type AircraftDetail struct {
	ID            int    `json:"id"`
	Callsign      string `json:"callsign"`
	Crew          string `json:"crew"`
	LevelOfCare   string `json:"levelOfCare"`
	Class         string `json:"class"`
	Lat           string `json:"lat"`
	Long          string `json:"long"`
	Area          string `json:"area"`
	NNum          string `json:"nNum"`
	OOSReason     string `json:"OOSReason"`
	OOSRemaining  string `json:"OOSRemaining"`
	OOSDuration   string `json:"OOSDuration"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Vision        string `json:"vision"`
	MissionStatus string `json:"missionStatus"`
}

var aircraft = []*Aircraft{}

// AircraftHandler ...
func AircraftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		respond(w, aircraft)
	default:
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// AircraftDetailHandler ...
func AircraftDetailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
	default:
		http.Error(w, "method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
