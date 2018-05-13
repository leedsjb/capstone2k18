package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

type missionRow struct {
	Type        string
	FlightRules string
	FlightNum   string
}

type waypointRow struct {
	Name        string
	ETE         string
	ETT         string
	Active      string
	FlightRules string
}

type oosRow struct {
	Reason  string
	EndTime time.Time
}

const (
	timeFormat = "2006-01-02 15:04 MST"
)

type aircraftRow struct {
	ID           string
	Callsign     string
	Nnum         string
	Manufacturer string // i.e. Augusta, Learjet, etc
	Title        string // i.e. A109E, PC-12, etc
	Class        string // i.e. Rotorcraft, Fixed-wing
	Lat          string
	Long         string
	LocationName string
	Status       string // TODO: double check what exactly this status is
	/* [MISSION]
	MissionType    string
	FlightRules    string
	TCNum      string
	// [WAYPOINT]
	WaypointTitle  string
	WaypointETE    string
	WaypointETT    string
	WaypointActive string
	*/
	// [OOS]
	OOSReason  string
	OOSEndTime string
}

type aircraftDetailRow struct {
	ID           string
	Callsign     string
	Nnum         string
	AircraftType string
	Lat          string
	Long         string
	LocationName string
	// [CREW]
	PersonnelID string
	FName       string
	LName       string
	Role        string
	/* [MISSION DETAIL]
	MissionType string
	FlightRules string
	TCNum string
	// [WAYPOINTS]
	WaypointTitle string
	WaypointETE string
	WaypointETT string
	WaypointActive string
	// [RADIO REPORT]
	ShortReport string
	Intubated string
	Drips string
	Age string
	Weight string
	Sex string
	Cardiac string
	GIBleed string
	OB string
	// [END radio report]
	Requestor string
	Receiverstring
	*/
	// [OOS Detail]
	// OOSReason    string
	// OOSStartTime string
	// OOSEndTime   string
}

// IndexAircraft ...
func IndexAircraft(trie *indexes.Trie, aircraft *messages.Aircraft) error {
	if err := trie.AddEntity(strings.ToLower(aircraft.Callsign), aircraft.ID); err != nil {
		return fmt.Errorf("Error adding aircraft to trie: %v", err)
	}

	if err := trie.AddEntity(strings.ToLower(aircraft.NNum), aircraft.ID); err != nil {
		return fmt.Errorf("Error adding aircraft to trie: %v", err)
	}

	return nil
}

// // LoadAircraftTrie ...
// func LoadAircraftTrie(trie *indexes.Trie) error {
// 	for _, v := range aircraftDetails {
// 		if err := IndexAircraft(trie, GetAircraftSummary(v)); err != nil {
// 			return fmt.Errorf("Error loading trie: %v", err)
// 		}
// 	}
// 	return nil
// }

// // GetTrieAircraft ...
// func GetTrieAircraft(aircraftIDS []int) []*Aircraft {
// 	results := []*Aircraft{}

// 	for _, aircraftID := range aircraftIDS {
// 		results = append(results, GetAircraftSummary(aircraftDetails[aircraftID-1]))
// 	}

// 	return results
// }

// GetAircraftSummary ...
func (ctx *HandlerContext) GetAircraftSummary(v *messages.AircraftDetail) *messages.Aircraft {
	aircrafts := []*messages.Aircraft{}
	// TODO: SQL sproc for getting aircraft info
	aircraftRows, err := ctx.DB.Query("SELECT group_id, group_name, personnel_F_Name, personnel_L_Name, personnel_id,  FROM tblPERSONNEL_GROUP JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id ORDER BY group_name")
	if err != nil {
		fmt.Printf("Error querying MySQL for groups: %v", err)
	}

	// create variables and fill contents from retrieved rows

	currentRow := &aircraftRow{}
	for aircraftRows.Next() {
		err = aircraftRows.Scan(currentRow)
		if err != nil {
			fmt.Printf("Error scanning aircraft row: %v", err)
			os.Exit(1)
		}

		// [GENERAL AIRCRAFT INFO]
		aircraftType := currentRow.Manufacturer + " " + currentRow.Title

		currentAircraft := &messages.Aircraft{
			ID:       currentRow.ID,
			Status:   currentRow.Status,
			Type:     aircraftType,
			Callsign: currentRow.Callsign,
			Class:    currentRow.Class,
			Lat:      currentRow.Lat,
			Long:     currentRow.Long,
			Area:     currentRow.LocationName,
			NNum:     currentRow.Nnum,
		}

		// [MISSION]
		mission := &messages.Mission{}
		// TODO: SQL sproc for finding mission by aircraftID
		missionRows, err := ctx.DB.Query("SELECT things")
		if err != nil {
			fmt.Printf("Error querying MySQL for mission: %v", err)
		}
		missionRow := &missionRow{}
		for missionRows.Next() {
			err = missionRows.Scan(missionRow)
			if err != nil {
				fmt.Printf("Error scanning mission row: %v", err)
			}
			mission = &messages.Mission{
				Type:      missionRow.Type,
				Vision:    missionRow.FlightRules,
				FlightNum: missionRow.FlightRules,
			}
		}

		// [Waypoint]
		waypoints := []*messages.ClientMissionWaypoint{}
		// TODO: SQL sproc for finding waypoints by missionID
		waypointRows, err := ctx.DB.Query("SELECT things")
		if err != nil {
			fmt.Printf("Error querying MySQL for waypoint: %v", err)
		}
		waypointRow := &waypointRow{}
		for waypointRows.Next() {
			err = waypointRows.Scan(waypointRow)
			if err != nil {
				fmt.Printf("Error scanning waypoint row: %v", err)
			}
			waypoint := &messages.ClientMissionWaypoint{
				Name:        waypointRow.Name,
				ETE:         waypointRow.ETE,
				ETT:         waypointRow.ETT,
				Active:      waypointRow.Active,
				FlightRules: waypointRow.FlightRules,
			}

			waypoints = append(waypoints, waypoint)
		}
		// add waypoints to mission
		mission.Waypoints = waypoints

		// [OOS]
		// TODO: SQL sproc for finding OOS status by aircraftID
		oosRows, err := ctx.DB.Query("SELECT things")
		if err != nil {
			fmt.Printf("Error querying MySQL for OOS status: %v", err)
		}
		oosRow := &oosRow{}
		for oosRows.Next() {
			err = oosRows.Scan(oosRow)
			if err != nil {
				fmt.Printf("Error scanning OOS row: %v", err)
			}

			oosFinishTime := time.Until(oosRow.EndTime)
			remaining := oosFinishTime.String()

			oos := &messages.OOS{
				Reason:    oosRow.Reason,
				Remaining: remaining,
			}
		}

		aircrafts = append(aircrafts, currentAircraft)
	}

	a := &messages.Aircraft{
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
		a.Mission = &messages.Mission{
			Type:            v.Mission.Type,
			Vision:          v.Mission.Vision,
			NextWaypointETE: v.Mission.NextWaypointETE,
			FlightNum:       v.Mission.FlightNum,
		}
	}
	if v.OOS != nil {
		a.OOS = &messages.OOS{
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

			aircraft := []*messages.Aircraft{}

			if len(statusFilter) > 0 {
				for _, aircraftDetail := range aircraftDetails {
					if aircraftDetail.Status == statusFilter {
						aircraft = append(aircraft, GetAircraftSummary(aircraftDetail))
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

// // AircraftDetailHandler ...
// func AircraftDetailHandler(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(path.Base(r.URL.Path))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error decoding ID: %v", err), http.StatusBadRequest)
// 		return
// 	}
// 	var ad *AircraftDetail
// 	for _, v := range aircraftDetails {
// 		if v.ID == id {
// 			ad = v
// 			break
// 		}
// 	}
// 	if ad == nil {
// 		http.Error(w, "No aircraft with that ID", http.StatusBadRequest)
// 		return
// 	}
// 	switch r.Method {
// 	case "GET":
// 		respond(w, ad)
// 	default:
// 		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
// 		return
// 	}
// }
