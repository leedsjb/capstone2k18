package handlers

import (
	"net/http"
)

// Aircraft ...
type Aircraft struct {
	ID                 int    `json:"id"`
	Callsign           string `json:"callsign"`
	Crew               string `json:"crew"`
	FlightStatus       string `json:"flightStatus"`
	LevelOfCare        string `json:"levelOfCare"`
	Location           string `json:"location"`
	NNum               string `json:"nNum"`
	OosReason          string `json:"oosReason"`
	OosDuration        string `json:"oosDuration"`
	ReturnToServiceEta string `json:"returnToService"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	Vision             string `json:"vision"`
}

// AircraftDetail ...
type AircraftDetail struct {
	ID int `json:"id"`
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
