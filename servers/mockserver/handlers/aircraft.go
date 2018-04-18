package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Mission struct {
	Type            string `json:"type"`
	Status          string `json:"status"`
	Vision          string `json:"vision"`
	NextWaypointETE string `json:"nextWaypointETE"`
	FlightNum       string `json:"flightNum"`
}

type MissionDetail struct {
	Type            string `json:"type"`
	Status          string `json:"status"`
	Vision          string `json:"vision"`
	NextWaypointETE string `json:"nextWaypointETE"`
	FlightNum       string `json:"flightNum"`
	RadioReport     string `json:"radioReport"`
	Requestor       string `json:"requestor"`
}

type OOS struct {
	reason    string `json:"reason"`
	remaining string `json:"remaining"`
}

type OOSDetail struct {
	reason    string `json:"reason"`
	remaining string `json:"remaining"`
	duration  string `json:"duration"`
}

// Aircraft ...
type Aircraft struct {
	ID          int     `json:"id"`
	Status      string  `json:"status"`
	Type        string  `json:"type"`
	Callsign    string  `json:"callsign"`
	LevelOfCare string  `json:"levelOfCare"`
	Class       string  `json:"class"`
	Lat         string  `json:"lat"`
	Long        string  `json:"long"`
	Area        string  `json:"area"`
	NNum        string  `json:"nNum"`
	Mission     Mission `json:"mission"`
	OOS         OOS     `json:"OOS"`
}

// AircraftDetail ...
type AircraftDetail struct {
	ID          int           `json:"id"`
	Status      string        `json:"status"`
	Type        string        `json:"type"`
	Callsign    string        `json:"callsign"`
	Crew        GroupDetail   `json:"crew"`
	LevelOfCare string        `json:"levelOfCare"`
	Class       string        `json:"class"`
	Lat         string        `json:"lat"`
	Long        string        `json:"long"`
	Area        string        `json:"area"`
	NNum        string        `json:"nNum"`
	Mission     MissionDetail `json:"mission"`
	OOS         OOSDetail     `json:"OOS"`
}

var aircraftDetails = []*AircraftDetail{
	{
		ID:          1,
		Status:      "OOS",
		Type:        "August A109E Power",
		Callsign:    "AL5",
		Crew:        groupDetails[2],
		LevelOfCare: "neonatal",
		Class:       "rotary",
		Lat:         "47.543265",
		Long:        "-122.309759",
		Area:        "Inceptos vestibulum",
		NNum:        "N948AL",
		Mission:     {},
		OOS: {
			reason:    "Unscheduled maintenance",
			remaining: "29 hours",
			duration:  "7 hours",
		},
	},
	{
		ID:          2,
		Status:      "standby",
		Type:        "Pilatus PC-12",
		Callsign:    "AL3",
		Crew:        groupDetails[1],
		LevelOfCare: "pediatric",
		Class:       "fixed",
		Lat:         "47.543265",
		Long:        "-122.309759",
		Area:        "Ullamcorper fusce",
		NNum:        "N937AL",
		Mission:     {},
		OOS:         {},
	},
	{
		ID:          3,
		Status:      "on a mission",
		Type:        "August A109E Power",
		Callsign:    "AL2",
		Crew:        groupDetails[0],
		LevelOfCare: "neonatal",
		Class:       "rotary",
		Lat:         "47.543265",
		Long:        "-122.309759",
		Area:        "Sem quam Commodo",
		NNum:        "N951AL",
		Mission: {
			Type:            "RW-SCENE",
			Status:          "ongoing",
			Vision:          "IFR",
			NextWaypointETE: "x min to...",
			FlightNum:       "18-0013",
			RadioReport:     "18-0013, 65, 90, male, GSW to chest. Has chest tube., Yes, 4, Paced externally - bring pacer box, Upper GI Bleed, Less than 5cm - launch without AOC Notification",
			Requestor:       "First Last",
		},
		OOS: {},
	},
}

// AircraftHandler ...
func AircraftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		aircraft := []*Aircraft{}
		for _, v := range aircraftDetails {
			a := &Aircraft{
				ID:          v.ID,
				Status:      v.Status,
				Type:        v.Type,
				Callsign:    v.Callsign,
				LevelOfCare: v.LevelOfCare,
				Class:       v.Class,
				Lat:         v.Lat,
				Long:        v.Long,
				Area:        v.Area,
				NNum:        v.NNum,
				Mission: {
					Type:            v.Mission.Type,
					Status:          v.Mission.Status,
					Vision:          v.Mission.Vision,
					NextWaypointETE: v.Mission.NextWaypointETE,
					FlightNum:       v.Mission.FlightNum,
				},
				OOS: {
					reason:    v.OOS.reason,
					remaining: v.OOS.remaining,
				},
			}
			aircraft = append(aircraft, a)
		}
		respond(w, aircraft)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// AircraftDetailHandler ...
func AircraftDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding ID: %v", err), http.StatusBadRequest)
		return
	}
	var ad *AircraftDetail
	for _, v := range aircraftDetails {
		if v.ID == id {
			ad = v
			break
		}
	}
	if ad == nil {
		http.Error(w, "No aircraft with that ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		respond(w, ad)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
