package handlers

import (
	"net/http"
)

// Aircraft ...
type Aircraft struct {
	ID            int    `json:"id"`
	Callsign      string `json:"callsign"`
	Crew          string `json:"crew"`
	MissionStatus string `json:"missionStatus"`
	LevelOfCare   string `json:"levelOfCare"`
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
}

/*
// Aircraft ...
type Aircraft struct {
	ID           int    `json:"id"`
	Callsign     string `json:"callsign"`
	NNum         string `json:"nNum"`
	Status       string `json:"status"`
	Lat          string `json:"lat"`
	Long         string `json:"long"`
	Area         string `json:"area"`
	OOSReason    string `json:"OOSReason"`
	OOSRemaining string `json:"OOSRemaining"`
	Type         string `json:"type"`
}
*/

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
