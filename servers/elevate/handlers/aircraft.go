package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// These structs receive data from SQL queries
// and allow the structure of SQL data to be abstracted
// from the structure of what is sent to the client.

type missionRow struct {
	Type        string
	FlightNum   string
	MissionDate mysql.NullTime
	Completed   string
}

type missionDetailRow struct {
	Type        string
	FlightRules string
	FlightNum   string
	Requestor   string
	Receiver    string
	Completed   string
}

type waypointRow struct {
	ID          int
	Name        string
	ETA         mysql.NullTime
	FlightRules string
	Lat         string
	Long        string
	Active      string
	Completed   string
}

type oosRow struct {
	Reason  string
	EndTime mysql.NullTime
}

type oosDetailRow struct {
	Reason    string
	StartTime mysql.NullTime
	EndTime   mysql.NullTime
}

type crewRow struct {
	PersonnelID int
	FName       string
	LName       string
	Role        string
}

type reportRow struct {
	MissionID   int
	ShortReport sql.NullString
	Intubated   sql.NullBool
	Drips       sql.NullInt64
	Age         sql.NullInt64
	Weight      sql.NullInt64
	Sex         sql.NullString
	Cardiac     sql.NullBool
	GIBleed     sql.NullBool
	OB          sql.NullBool
}

const (
	timeFormat = "2006-01-02 15:04 MST"
)

type aircraftRow struct {
	ID           int
	Callsign     string
	Nnum         string
	Manufacturer string // i.e. Augusta, Learjet, etc
	Title        string // i.e. A109E, PC-12, etc
	Class        string // i.e. Rotorcraft, Fixed-wing
	Lat          string
	Long         string
	LocationName string
	Status       string // TODO: double check what exactly this status is
	/*
		// [MISSION]
		MissionType string
		FlightRules string
		TCNum       string
		// [WAYPOINT]
		WaypointTitle  string
		WaypointETE    string
		WaypointETT    string
		WaypointActive string
		WaypointCompleted string
		// [OOS]
		OOSReason  string
		OOSEndTime string
	*/
}

type aircraftDetailRow struct {
	ID           int
	Callsign     string
	Nnum         string
	Manufacturer string
	Title        string
	Class        string
	Lat          string
	Long         string
	LocationName string
	Status       string
	// [CREW]
	// PersonnelID string
	// FName       string
	// LName       string
	// Role        string
	// // [START MISSION DETAIL]
	// MissionType string
	// FlightRules string
	// TCNum       string
	// Requestor string
	// Receiver  string
	// // [WAYPOINTS]
	// WaypointTitle  string
	// WaypointETE    string
	// WaypointETT    string
	// WaypointActive string
	// WaypointCompleted string
	// // [RADIO REPORT]
	// ShortReport string
	// Intubated   string
	// Drips       string
	// Age         string
	// Weight      string
	// Sex         string
	// Cardiac     string
	// GIBleed     string
	// OB          string
	// // [END radio report]
	// // [END MISSION DETAIL]
	// // [OOS Detail]
	// OOSReason    string
	// OOSStartTime string
	// OOSEndTime   string
}

// // checkNullSql ...
// func checkNullMissionRow(check *missionRow) *missionRow {

// }
// //
// func checkMissionDetailRow  (check *missionDetailRow) *missionDetailRow {

// }
// func checkWaypointRow (check *waypointRow) *waypointRow {

// }
// func checkOosRow (check *oosRow) *oosRow {

// }
// func checkOosDetailRow (check *oosDetailRow) *oosDetailRow {

// }
// func checkCrewRow (check *crewRow) *crewRow {

// }
// func checkReportRow (check *reportRow) *reportRow {

// }
// func checkAircraftRow (check *aircraftRow) *aircraftRow {

// }
// func checkAircraftDetailRow (check *aircraftDetailRow) *aircraftDetailRow {

// }

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

// LoadAircraftTrie ...
func (ctx *HandlerContext) LoadAircraftTrie(trie *indexes.Trie) error {
	aircraftRows, err := ctx.GetAllAircraft()
	if err != nil {
		return fmt.Errorf("Error loading trie: %v", err)
	}
	aircraftRow := &aircraftRow{}
	for aircraftRows.Next() {
		err = aircraftRows.Scan(
			&aircraftRow.ID,
			&aircraftRow.Callsign,
			&aircraftRow.Nnum,
			&aircraftRow.Manufacturer,
			&aircraftRow.Title,
			&aircraftRow.Class,
			&aircraftRow.Lat,
			&aircraftRow.Long,
			&aircraftRow.LocationName,
			&aircraftRow.Status,
		)
		if err != nil {
			return fmt.Errorf("Error scanning aircraft row: %v", err)
		}
		aircraft, err := ctx.GetAircraftSummary(aircraftRow)
		if err != nil {
			return fmt.Errorf("Error populating aircraft: %v", err)
		}
		if err := IndexAircraft(trie, aircraft); err != nil {
			return fmt.Errorf("Error loading trie: %v", err)
		}
	}
	return nil
}

// GetTrieAircraft ...
func (ctx *HandlerContext) GetTrieAircraft(aircraftIDS []int) ([]*messages.Aircraft, error) {
	results := []*messages.Aircraft{}

	for _, aircraftID := range aircraftIDS {
		aircraftRows, err := ctx.GetAircraftByID(aircraftID)
		if err != nil {
			return nil, fmt.Errorf("Error getting trie aircraft: %v", err)
		}
		aircraftRow := &aircraftRow{}
		for aircraftRows.Next() {
			err = aircraftRows.Scan(
				&aircraftRow.ID,
				&aircraftRow.Callsign,
				&aircraftRow.Nnum,
				&aircraftRow.Manufacturer,
				&aircraftRow.Title,
				&aircraftRow.Class,
				&aircraftRow.Lat,
				&aircraftRow.Long,
				&aircraftRow.LocationName,
				&aircraftRow.Status,
			)
			if err != nil {
				return nil, fmt.Errorf("Error scanning aircraft row: %v", err)
			}
		}
		result, err := ctx.GetAircraftSummary(aircraftRow)
		if err != nil {
			return nil, fmt.Errorf("Error populating aircraft for trie: %v", err)
		}
		results = append(results, result)
	}

	return results, nil
}

// GetAircraftSummary ...
func (ctx *HandlerContext) GetAircraftSummary(currentRow *aircraftRow) (*messages.Aircraft, error) {
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
	if strings.ToLower(aircraft.Status) == "on mission" {
		mission := &messages.Mission{}
		missionRows, err := ctx.GetMissionByAircraft(currentRow.ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving missions for aircraft [%v]: %v", currentRow.Callsign, err)
		}
		missionRow := &missionRow{}
		for missionRows.Next() {
			err = missionRows.Scan(
				&missionRow.Type,
				&missionRow.FlightNum,
				&missionRow.MissionDate,
				&missionRow.Completed,
			)
			if err != nil {
				return nil, fmt.Errorf("Error scanning mission row: %v", err)
			}
			mission = &messages.Mission{
				Type:      missionRow.Type,
				FlightNum: missionRow.FlightNum,
				Completed: missionRow.Completed,
			}
		}
		nextETE := ""

		// [Waypoint]
		waypoints := []*messages.ClientMissionWaypoint{}
		waypointRows, err := ctx.GetWaypointsByAircraft(currentRow.ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving waypoints for aircraft [%v]: %v", currentRow.Callsign, err)
		}
		waypointRow := &waypointRow{}
		for waypointRows.Next() {
			err = waypointRows.Scan(
				&waypointRow.ID,
				&waypointRow.Name,
				&waypointRow.ETA,
				&waypointRow.Active,
				&waypointRow.Completed,
				&waypointRow.FlightRules,
				&waypointRow.Lat,
				&waypointRow.Long,
			)
			if err != nil {
				fmt.Printf("Error scanning waypoint row: %v", err)
			}

			waypoint := &messages.ClientMissionWaypoint{
				ID:          waypointRow.ID,
				Name:        waypointRow.Name,
				FlightRules: waypointRow.FlightRules,
				Lat:         waypointRow.Lat,
				Long:        waypointRow.Long,
			}

			if waypointRow.Active == "0" {
				waypoint.Active = false
			} else {
				waypoint.Active = true
			}

			if strings.ToLower(waypointRow.Completed) == "0" {
				waypoint.Completed = false
			} else {
				waypoint.Completed = true
			}

			if waypointRow.ETA.Valid {
				waypoint.ETA = waypointRow.ETA.Time.String()

				if strings.ToLower(waypointRow.Active) == "true" {
					nextETE = time.Until(waypointRow.ETA.Time).String()
				}
			}

			waypoints = append(waypoints, waypoint)
		}
		// add waypoints to mission
		mission.Waypoints = waypoints
		mission.NextWaypointETE = nextETE

		// add mission to aircraft
		aircraft.Mission = mission
	}

	// [OOS]
	if strings.ToLower(aircraft.Status) == "out of service" {
		oos := &messages.OOS{}
		oosRows, err := ctx.GetOOSByAircraft(currentRow.ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving OOS for aircraft [%v]: %v", currentRow.Callsign, err)
		}
		oosRow := &oosRow{}
		for oosRows.Next() {
			err = oosRows.Scan(
				&oosRow.Reason,
				&oosRow.EndTime,
			)
			if err != nil {
				fmt.Printf("Error scanning OOS row: %v", err)
			}

			remaining := ""
			if oosRow.EndTime.Valid {
				oosFinishTime := time.Until(oosRow.EndTime.Time)
				remaining = oosFinishTime.String()
			}

			oos = &messages.OOS{
				Reason:    oosRow.Reason,
				Remaining: remaining,
			}
		}

		// add OOS to aircraft
		aircraft.OOS = oos
	}
	return aircraft, nil
}

// GetAircraftDetailSummary ...
func (ctx *HandlerContext) GetAircraftDetailSummary(currentRow *aircraftDetailRow) (*messages.AircraftDetail, error) {
	// [GENERAL AIRCRAFT INFO]
	aircraftType := currentRow.Manufacturer + " " + currentRow.Title

	aircraftDetail := &messages.AircraftDetail{
		ID:       currentRow.ID,
		Status:   currentRow.Status,
		Type:     aircraftType,
		Callsign: currentRow.Callsign,
		// Crew
		Class: currentRow.Class,
		Lat:   currentRow.Lat,
		Long:  currentRow.Long,
		Area:  currentRow.LocationName,
		NNum:  currentRow.Nnum,
		// Mission
		// Waypoints
		// OOS &message.OOSDetail{}
	}

	// [CREW]
	crew := []*messages.Person{}
	crewRows, err := ctx.GetCrewByAircraft(currentRow.ID)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving crew for aircraft [%v]: %v", currentRow.Callsign, err)
	}
	crewRow := &crewRow{}
	for crewRows.Next() {
		err = crewRows.Scan(
			&crewRow.PersonnelID,
			&crewRow.FName,
			&crewRow.LName,
			&crewRow.Role,
		)
		if err != nil {
			return nil, fmt.Errorf("Error scanning crew row: %v", err)
		}
		crewMember := &messages.Person{
			ID:       crewRow.PersonnelID,
			FName:    crewRow.FName,
			LName:    crewRow.LName,
			Position: crewRow.Role,
		}
		crew = append(crew, crewMember)
	}
	aircraftDetail.Crew = crew

	// [MISSION]
	if strings.ToLower(aircraftDetail.Status) == "on mission" {
		missionDetail := &messages.MissionDetail{}
		missionDetailRows, err := ctx.GetMissionDetailByAircraft(currentRow.ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving MissionDetails for aircraft [%v]: %v", currentRow.Callsign, err)
		}
		missionDetailRow := &missionDetailRow{}
		for missionDetailRows.Next() {
			err = missionDetailRows.Scan(
				&missionDetailRow.Type,
				// &missionDetailRow.FlightRules,
				&missionDetailRow.FlightNum,
				&missionDetailRow.Requestor,
				&missionDetailRow.Receiver,
				&missionDetailRow.Completed,
			)
			if err != nil {
				return nil, fmt.Errorf("Error scanning mission detail row: %v", err)
			}
			missionDetail = &messages.MissionDetail{
				Type: missionDetailRow.Type,
				// Vision: missionDetailRow.FlightRules,
				// NextWaypointETE
				// Waypoints
				FlightNum: missionDetailRow.FlightNum,
				// RadioReport
				Requestor: missionDetailRow.Requestor,
				Receiver:  missionDetailRow.Receiver,
				Completed: missionDetailRow.Completed,
			}
		}
		// [Waypoint]
		nextETE := ""
		waypoints := []*messages.ClientMissionWaypoint{}
		waypointRows, err := ctx.GetWaypointsByAircraft(currentRow.ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving waypoints for aircraft [%v]: %v", currentRow.Callsign, err)
		}
		waypointRow := &waypointRow{}
		for waypointRows.Next() {
			err = waypointRows.Scan(
				&waypointRow.ID,
				&waypointRow.Name,
				&waypointRow.ETA,
				&waypointRow.Active,
				&waypointRow.Completed,
				&waypointRow.FlightRules,
				&waypointRow.Lat,
				&waypointRow.Long,
			)
			if err != nil {
				return nil, fmt.Errorf("Error scanning waypoint row: %v", err)
			}
			waypoint := &messages.ClientMissionWaypoint{
				ID:          waypointRow.ID,
				Name:        waypointRow.Name,
				FlightRules: waypointRow.FlightRules,
				Lat:         waypointRow.Lat,
				Long:        waypointRow.Long,
			}

			if waypointRow.Active == "0" {
				waypoint.Active = false
			} else {
				waypoint.Active = true
			}

			if strings.ToLower(waypointRow.Completed) == "0" {
				waypoint.Completed = false
			} else {
				waypoint.Completed = true
			}

			if waypointRow.ETA.Valid {
				waypoint.ETA = waypointRow.ETA.Time.String()

				if strings.ToLower(waypointRow.Active) == "1" {
					nextETE = time.Until(waypointRow.ETA.Time).String()
				}
			}

			waypoints = append(waypoints, waypoint)
		}
		// add waypoints to mission
		missionDetail.Waypoints = waypoints
		missionDetail.NextWaypointETE = nextETE
		// [RADIO REPORT]
		report := &messages.Patient{}
		reportRows, err := ctx.GetPatientByAircraft(currentRow.ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving patient info for aircraft [%v]: %v", currentRow.Callsign, err)
		}
		reportRow := &reportRow{}
		for reportRows.Next() {
			err = reportRows.Scan(
				&reportRow.MissionID,
				&reportRow.ShortReport,
				&reportRow.Intubated,
				&reportRow.Drips,
				&reportRow.Age,
				&reportRow.Weight,
				&reportRow.Sex,
				&reportRow.Cardiac,
				&reportRow.GIBleed,
				&reportRow.OB,
			)
			if err != nil {
				return nil, fmt.Errorf("Error scanning report row: %v", err)
			}

			report = &messages.Patient{}

			if reportRow.ShortReport.Valid {
				report.ShortReport = reportRow.ShortReport.String
			}
			if reportRow.Intubated.Valid {
				report.Intubated = reportRow.Intubated.Bool
			}
			if reportRow.Drips.Valid {
				report.Drips = int(reportRow.Drips.Int64)
			}
			if reportRow.Age.Valid {
				report.Age = int(reportRow.Age.Int64)
			}
			if reportRow.Weight.Valid {
				report.Weight = int(reportRow.Weight.Int64)
			}
			if reportRow.Sex.Valid {
				report.Gender = reportRow.Sex.String
			}
			if reportRow.Cardiac.Valid {
				if reportRow.Cardiac.Bool {
					report.Cardiac = true
				} else {
					report.Cardiac = false
				}
			}
			if reportRow.GIBleed.Valid {
				if reportRow.GIBleed.Bool {
					report.GIBleed = true
				} else {
					report.GIBleed = false
				}
			}
			if reportRow.OB.Valid {
				if reportRow.OB.Bool {
					report.OB = true
				} else {
					report.OB = false
				}
			}
		}
		// if a patient report exists
		if reportRow.MissionID != 0 {
			// add patient information to mission
			missionDetail.RadioReport = report
		}
		// add mission, waypoints, and radio report to aircraft detail
		aircraftDetail.Mission = missionDetail
	}

	// [OOS]
	if strings.ToLower(aircraftDetail.Status) == "out of service" {
		oosDetail := &messages.OOSDetail{}
		oosDetailRows, err := ctx.GetOOSDetailByAircraft(currentRow.ID)
		if err != nil {
			return nil, fmt.Errorf("Error returning OOS details: %v", err)
		}
		oosDetailRow := &oosDetailRow{}
		for oosDetailRows.Next() {
			err = oosDetailRows.Scan(
				&oosDetailRow.Reason,
				&oosDetailRow.StartTime,
				&oosDetailRow.EndTime,
			)
			if err != nil {
				fmt.Printf("Error scanning OOS row: %v", err)
			}

			remaining := ""
			if oosDetailRow.EndTime.Valid {
				oosFinishTime := time.Until(oosDetailRow.EndTime.Time)
				remaining = oosFinishTime.String()
			}

			duration := ""
			if oosDetailRow.StartTime.Valid {
				oosElapsedTime := time.Since(oosDetailRow.StartTime.Time)
				duration = oosElapsedTime.String()
			}

			oosDetail = &messages.OOSDetail{
				Reason:    oosDetailRow.Reason,
				Remaining: remaining,
				Duration:  duration,
			}
		}
		// add OOS to aircraft
		aircraftDetail.OOS = oosDetail
	}

	return aircraftDetail, nil
}

func (ctx *HandlerContext) getAircraftList(aircraftRows *sql.Rows) ([]*messages.Aircraft, error) {
	aircraftList := []*messages.Aircraft{}
	currentRow := &aircraftRow{}
	for aircraftRows.Next() {
		err := aircraftRows.Scan(
			&currentRow.ID,
			&currentRow.Callsign,
			&currentRow.Nnum,
			&currentRow.Manufacturer,
			&currentRow.Title,
			&currentRow.Class,
			&currentRow.Lat,
			&currentRow.Long,
			&currentRow.LocationName,
			&currentRow.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("Error scanning aircraft row: %v", err)
		}
		aircraft, err := ctx.GetAircraftSummary(currentRow)
		if err != nil {
			return nil, fmt.Errorf("Error populating aircraft: %v", err)
		}
		aircraftList = append(aircraftList, aircraft)
	}
	return aircraftList, nil
}

// AircraftHandler ...
func (ctx *HandlerContext) AircraftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()

		term := query.Get("q")

		statusFilter := query.Get("status")

		category := query.Get("category")

		if len(term) > 0 {
			// filter by user query
			aircraftIDS := ctx.AircraftTrie.GetEntities(strings.ToLower(term), 20)
			aircraftList, err := ctx.GetTrieAircraft(aircraftIDS)
			if err != nil {
				fmt.Printf("Error pulling aircrafts from trie: %v", err)
				return
			}
			respond(w, aircraftList)
		} else if len(statusFilter) > 0 {
			// filter by aircraft status: OAM (On Mission), RFM (Ready for Mission), OOS (Out of Service)
			aircraftRows, err := ctx.GetAircraftByStatus(statusFilter)
			if err != nil {
				http.Error(w, fmt.Sprintf("Couldn't get aircraft by status: %v", err), http.StatusInternalServerError)
				return
			}

			aircraftList, err := ctx.getAircraftList(aircraftRows)
			if err != nil {
				http.Error(w, fmt.Sprintf("Couldn't get aircraft list by status: %v", err), http.StatusInternalServerError)
				return
			}

			respond(w, aircraftList)

		} else if len(category) > 0 {
			// filter by aircraft category: rotorcraft, fixed-wing
			aircraftRows, err := ctx.GetAircraftByCategory(category)
			if err != nil {
				fmt.Printf("Couldn't get aircraft by category: %v", err)
			}

			aircraftList, err := ctx.getAircraftList(aircraftRows)
			if err != nil {
				http.Error(w, fmt.Sprintf("Couldn't get aircraft list by status: %v", err), http.StatusInternalServerError)
				return
			}

			respond(w, aircraftList)
		} else {
			// no filter, return all
			aircraftRows, err := ctx.GetAllAircraft()
			if err != nil {
				http.Error(w, fmt.Sprintf("Couldn't get all aircraft: %v", err), http.StatusInternalServerError)
				return
			}

			aircraftList, err := ctx.getAircraftList(aircraftRows)
			if err != nil {
				http.Error(w, fmt.Sprintf("Couldn't get aircraft list by status: %v", err), http.StatusInternalServerError)
				return
			}

			respond(w, aircraftList)
		}

	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// AircraftDetailHandler ...
func (ctx *HandlerContext) AircraftDetailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := path.Base(r.URL.Path)
		if id != "." && id != "aircraft" {
			aircraftDetail := &messages.AircraftDetail{}

			aircraftID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Printf("Error changing aircraft ID from string to int")
			}

			aircraftDetailRows, err := ctx.GetAircraftByID(aircraftID)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error getting aircraft details from DB: %v", err), http.StatusInternalServerError)
				return
			}

			aircraftDetailRow := &aircraftDetailRow{}
			for aircraftDetailRows.Next() {
				err = aircraftDetailRows.Scan(
					&aircraftDetailRow.ID,
					&aircraftDetailRow.Callsign,
					&aircraftDetailRow.Nnum,
					&aircraftDetailRow.Manufacturer,
					&aircraftDetailRow.Title,
					&aircraftDetailRow.Class,
					&aircraftDetailRow.Lat,
					&aircraftDetailRow.Long,
					&aircraftDetailRow.LocationName,
					&aircraftDetailRow.Status,
				)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error scanning aircraft details from query: %v", err), http.StatusInternalServerError)
					return
				}
				aircraftDetail, err = ctx.GetAircraftDetailSummary(aircraftDetailRow)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error populating aircraft detail summary: %v", err), http.StatusInternalServerError)
					return
				}
			}
			respond(w, aircraftDetail)
		} else if id == "aircraft" {
			ctx.AircraftHandler(w, r)
		} else {
			http.Error(w, "No aircraft with that ID", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
