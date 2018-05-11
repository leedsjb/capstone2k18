package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/leedsjb/capstone2k18/servers/mockserver/indexes"
)

// Mission ...
type Mission struct {
	Type            string `json:"type"`
	Status          string `json:"status"`
	Vision          string `json:"vision"`
	NextWaypointETE string `json:"nextWaypointETE"`
	FlightNum       string `json:"flightNum"`
}

// MissionDetail ...
type MissionDetail struct {
	Type            string `json:"type"`
	Status          string `json:"status"`
	Vision          string `json:"vision"`
	NextWaypointETE string `json:"nextWaypointETE"`
	FlightNum       string `json:"flightNum"`
	RadioReport     string `json:"radioReport"`
	Requestor       string `json:"requestor"`
	Receiver        string `json:"receiver"`
}

// OOS ...
type OOS struct {
	Reason    string `json:"reason"`
	Remaining string `json:"remaining"`
}

// OOSDetail ...
type OOSDetail struct {
	Reason    string `json:"reason"`
	Remaining string `json:"remaining"`
	Duration  string `json:"duration"`
}

// Aircraft ...
type Aircraft struct {
	ID          int      `json:"id"`
	Status      string   `json:"status"`
	Type        string   `json:"type"`
	Callsign    string   `json:"callsign"`
	LevelOfCare string   `json:"levelOfCare"`
	Class       string   `json:"class"`
	Lat         float32  `json:"lat"`
	Long        float32  `json:"long"`
	Area        string   `json:"area"`
	NNum        string   `json:"nNum"`
	Mission     *Mission `json:"mission"`
	OOS         *OOS     `json:"OOS"`
}

// AircraftDetail ...
type AircraftDetail struct {
	ID          int            `json:"id"`
	Status      string         `json:"status"`
	Type        string         `json:"type"`
	Callsign    string         `json:"callsign"`
	Crew        *GroupDetail   `json:"crew"`
	LevelOfCare string         `json:"levelOfCare"`
	Class       string         `json:"class"`
	Lat         float32        `json:"lat"`
	Long        float32        `json:"long"`
	Area        string         `json:"area"`
	NNum        string         `json:"nNum"`
	Mission     *MissionDetail `json:"mission"`
	OOS         *OOSDetail     `json:"OOS"`
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
		Lat:         47.545218,
		Long:        -122.315673,
		Area:        "Inceptos vestibulum",
		NNum:        "N948AL",
		Mission:     nil,
		OOS: &OOSDetail{
			Reason:    "Unscheduled maintenance",
			Remaining: "29 hours",
			Duration:  "7 hours",
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
		Lat:         47.542964,
		Long:        -122.309860,
		Area:        "Ullamcorper fusce",
		NNum:        "N937AL",
		Mission:     nil,
		OOS:         nil,
	},
	{
		ID:          3,
		Status:      "on a mission",
		Type:        "August A109E Power",
		Callsign:    "AL2",
		Crew:        groupDetails[0],
		LevelOfCare: "neonatal",
		Class:       "rotary",
		Lat:         47.528478,
		Long:        -122.291697,
		Area:        "Sem quam Commodo",
		NNum:        "N951AL",
		Mission: &MissionDetail{
			Type:            "RW-SCENE",
			Status:          "ongoing",
			Vision:          "IFR",
			NextWaypointETE: "x min to...",
			FlightNum:       "18-0013",
			RadioReport:     "18-0013, 65, 90, male, GSW to chest. Has chest tube., Yes, 4, Paced externally - bring pacer box, Upper GI Bleed, Less than 5cm - launch without AOC Notification",
			Requestor:       "Lopez Island EMS",
			Receiver:        "Harborview Medical Center",
		},
		OOS: nil,
	},
}

// IndexAircraft ...
func IndexAircraft(trie *indexes.Trie, aircraft *Aircraft) error {
	if err := trie.AddEntity(strings.ToLower(aircraft.Callsign), aircraft.ID); err != nil {
		return fmt.Errorf("Error adding aircraft to trie: %v", err)
	}

	if err := trie.AddEntity(strings.ToLower(aircraft.NNum), aircraft.ID); err != nil {
		return fmt.Errorf("Error adding aircraft to trie: %v", err)
	}

	return nil
}

// LoadAircraftTrie ...
func LoadAircraftTrie(trie *indexes.Trie) error {
	for _, v := range aircraftDetails {
		if err := IndexAircraft(trie, GetAircraftSummary(v)); err != nil {
			return fmt.Errorf("Error loading trie: %v", err)
		}
	}
	return nil
}

// GetTrieAircraft ...
func GetTrieAircraft(aircraftIDS []int) []*Aircraft {
	results := []*Aircraft{}

	for _, aircraftID := range aircraftIDS {
		results = append(results, GetAircraftSummary(aircraftDetails[aircraftID-1]))
	}

	return results
}

// GetAircraftSummary ...
func GetAircraftSummary(v *AircraftDetail) *Aircraft {
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
	}
	if v.Mission != nil {
		a.Mission = &Mission{
			Type:            v.Mission.Type,
			Status:          v.Mission.Status,
			Vision:          v.Mission.Vision,
			NextWaypointETE: v.Mission.NextWaypointETE,
			FlightNum:       v.Mission.FlightNum,
		}
	}
	if v.OOS != nil {
		a.OOS = &OOS{
			Reason:    v.OOS.Reason,
			Remaining: v.OOS.Remaining,
		}
	}
	return a
}

// AircraftHandler ...
func (ctx *HandlerContext) AircraftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()

		term := query.Get("q")

		if len(term) > 0 {
			aircraftIDS := ctx.AircraftTrie.GetEntities(strings.ToLower(term), 20)
			aircraft := GetTrieAircraft(aircraftIDS)
			respond(w, aircraft)
		} else {
			statusFilter := query.Get("status")

			aircraft := []*Aircraft{}

			if len(statusFilter) > 0 {
				for _, v := range aircraftDetails {
					if v.Status == statusFilter {
						aircraft = append(aircraft, GetAircraftSummary(v))
					}
				}
			} else {
				for _, v := range aircraftDetails {
					aircraft = append(aircraft, GetAircraftSummary(v))
				}
			}

			respond(w, aircraft)
		}
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
