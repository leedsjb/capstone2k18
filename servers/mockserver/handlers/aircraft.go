package handlers

import (
	"net/http"
)

// Aircraft ...
type Aircraft struct {
	ID            int    `json:"id"`
	Callsign      string `json:"callsign"`
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
	Status        string `json:"status"`
	Type          string `json:"type"`
	MissionStatus string `json:"missionStatus"`
}

var aircraft = []*Aircraft{
	{
		ID:            1,
		Callsign:      "AL5",
		Crew:          "First Last, First Last, First Last",
		LevelOfCare:   "neonatal",
		Class:         "rotary",
		Lat:           "47.543265",
		Long:          "-122.309759",
		Area:          "Inceptos vestibulum",
		NNum:          "N948AL",
		OOSReason:     "Unscheduled maintenance",
		OOSRemaining:  "29 hours",
		OOSDuration:   "7 hours",
		Status:        "OOS",
		Type:          "August A109E Power",
		MissionStatus: "",
	},
	{
		ID:            2,
		Callsign:      "AL3",
		Crew:          "First Last, First Last, First Last",
		LevelOfCare:   "pediatric",
		Class:         "fixed",
		Lat:           "47.543265",
		Long:          "-122.309759",
		Area:          "Ullamcorper fusce",
		NNum:          "N937AL",
		OOSReason:     "",
		OOSRemaining:  "",
		OOSDuration:   "",
		Status:        "standby",
		Type:          "Pilatus PC-12",
		MissionStatus: "",
	},
	{
		ID:            3,
		Callsign:      "AL2",
		Crew:          "First Last, First Last, First Last",
		LevelOfCare:   "neonatal",
		Class:         "rotary",
		Lat:           "47.543265",
		Long:          "-122.309759",
		Area:          "Sem quam Commodo",
		NNum:          "N951AL",
		OOSReason:     "",
		OOSRemaining:  "",
		OOSDuration:   "",
		Status:        "on a mission",
		Type:          "August A109E Power",
		MissionStatus: "Enroute to Squaxin Ballfields",
	},
}

// AircraftHandler ...
func AircraftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		respond(w, aircraft)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// AircraftDetailHandler ...
func AircraftDetailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		return
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
