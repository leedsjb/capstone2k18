package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
	ETE         time.Time
	ETT         time.Time
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
	aircraftID, err := strconv.Atoi(aircraft.ID)
	if err != nil {
		fmt.Printf("Error changing aircraft ID from string to int")
	}
	if err := trie.AddEntity(strings.ToLower(aircraft.Callsign), aircraftID); err != nil {
		return fmt.Errorf("Error adding aircraft to trie: %v", err)
	}

	if err := trie.AddEntity(strings.ToLower(aircraft.NNum), aircraftID); err != nil {
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

// GetTrieAircraft ...
func (ctx *HandlerContext) GetTrieAircraft(aircraftIDS []int) []*messages.Aircraft {
	results := []*messages.Aircraft{}

	for _, aircraftID := range aircraftIDS {
		ID := strconv.Itoa(aircraftID)
		aircraftRows, err := ctx.GetAircraftByID(ID)
		aircraftRow := &aircraftRow{}
		for aircraftRows.Next() {
			err = aircraftRows.Scan(aircraftRow)
			if err != nil {
				fmt.Printf("Error scanning aircraft row: %v", err)
			}
		}
		results = append(results, ctx.GetAircraftSummary(aircraftRow))
	}

	return results
}

// GetAircraftSummary ...
func (ctx *HandlerContext) GetAircraftSummary(currentRow *aircraftRow) *messages.Aircraft {
	// [GENERAL AIRCRAFT INFO]
	aircraftType := currentRow.Manufacturer + " " + currentRow.Title

	aircraft := &messages.Aircraft{
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
	missionRows, err := ctx.GetMissionByAircraft(currentRow.ID)
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
	nextETE := ""

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
			ETE:         waypointRow.ETE.String(),
			ETT:         waypointRow.ETT.String(),
			Active:      waypointRow.Active,
			FlightRules: waypointRow.FlightRules,
		}

		if strings.ToLower(waypointRow.Active) == "true" {
			nextETE = waypointRow.ETE.String()
		}
		waypoints = append(waypoints, waypoint)
	}
	// add waypoints to mission
	mission.Waypoints = waypoints
	mission.NextWaypointETE = nextETE

	// [OOS]
	// TODO: SQL sproc for finding OOS status by aircraftID
	oos := &messages.OOS{}
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

		oos = &messages.OOS{
			Reason:    oosRow.Reason,
			Remaining: remaining,
		}
	}

	// add mission to aircraft
	aircraft.Mission = mission
	// add OOS to aircraft
	aircraft.OOS = oos
	return aircraft
}

// AircraftHandler ...
func (ctx *HandlerContext) AircraftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()

		term := query.Get("q")

		// TODO: refactor to be cleaner
		if len(term) > 0 {
			// search query non-empty
			aircraftIDS := ctx.AircraftTrie.GetEntities(strings.ToLower(term), 20)
			aircraft := ctx.GetTrieAircraft(aircraftIDS)
			respond(w, aircraft)
		} else {
			// search query empty
			statusFilter := query.Get("status")

			aircraft := []*messages.Aircraft{}

			if len(statusFilter) > 0 {
				// filter by status

				// TODO: SQL sproc for aircraft by status
				// map status to statusID
				aircraftRows, err := ctx.GetAircraftByStatus(statusFilter)

				currentRow := &aircraftRow{}
				for aircraftRows.Next() {
					err = aircraftRows.Scan(currentRow)
					if err != nil {
						fmt.Printf("Error scanning aircraft row: %v", err)
						os.Exit(1)
					}
					aircraft = append(aircraft, ctx.GetAircraftSummary(currentRow))
				}
			} else {
				// no filter, return all
				aircraftRows, err := ctx.GetAllAircraft()
				currentRow := &aircraftRow{}
				for aircraftRows.Next() {
					err = aircraftRows.Scan(currentRow)
					if err != nil {
						fmt.Printf("Error scanning aircraft row: %v", err)
						os.Exit(1)
					}
					aircraft = append(aircraft, ctx.GetAircraftSummary(currentRow))
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
