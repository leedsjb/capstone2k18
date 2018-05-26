package parsers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// empty, no-time-out context for sql queries
var sqlCtx = context.Background()

// [AIRCRAFT]

// GetAircraftIDByCallsign retrieves an aircraft's ID given its callsign
// ctx.GetAircraftIDByCallsign("AL3")
func (ctx *ParserContext) GetAircraftIDByCallsign(aircraftCallsign string) (int, error) {
	fmt.Printf("[GET AIRCRAFT ID] callsign: %v\n", aircraftCallsign)
	query := `CALL uspGetAircraftIDByCallsign(?)`
	aircraftRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftCallsign,
	)
	if err != nil {
		return -1, fmt.Errorf("Error querying MySQL for aircraft ID: %v\n", err)
	}
	var aircraftID int
	for aircraftRow.Next() {
		err = aircraftRow.Scan(&aircraftID)
		if err != nil {
			return -1, fmt.Errorf("Error retrieving aircraft ID by callsign: %v", err)
		}
	}
	return aircraftID, nil
}

// GetAircraftCallsign retrieves an aircraft's callsign given the aircraft's ID
func (ctx *ParserContext) GetAircraftCallsign(aircraftID int) (string, error) {
	// get callsign from db using aircraftID
	// aircraftID := strconv.Itoa(ID)

	// "SELECT ac_callsign FROM tblAIRCRAFT JOIN tblMISSION ON tblMISSION.aircraft_id = tblAIRCRAFT.ac_id WHERE mission_id=" + aircraftID

	query := `CALL uspGetAircraftCallsign(?)`
	aircraftRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return "", fmt.Errorf("Error querying MySQL for aircraft callsign: %v\n", err)
	}
	var aircraftCallsign string
	for aircraftRow.Next() {
		err = aircraftRow.Scan(&aircraftCallsign)
		if err != nil {
			return "", err
		}
	}
	return aircraftCallsign, nil
}

// GetAircraftByID retrieves an aircraft object with the matching ID from the database
func (ctx *ParserContext) GetAircraftByID(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetAircraftByID(?)`
	aircraftRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

// AddNewAircraft adds a new aircraft object to the database
func (ctx *ParserContext) AddNewAircraft(aircraftInfo *messages.Aircraft_Create) error {
	query := `CALL uspAddNewAircraft(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.ID,
		aircraftInfo.Base,
		aircraftInfo.Callsign,
		aircraftInfo.CallTypes,
		aircraftInfo.CellPhone,
		aircraftInfo.Color,
		aircraftInfo.LastKnownLocation,
		aircraftInfo.MaxPatientWeight,
		aircraftInfo.Model,
		aircraftInfo.NNum,
		aircraftInfo.PadTimeDay,
		aircraftInfo.PadTimeNight,
	); err != nil {
		return fmt.Errorf("Error adding aircraft to DB: %v", err)
	}
	return nil
}

// UpdateAircraftProps updates an existing aircraft object
func (ctx *ParserContext) UpdateAircraftProps(aircraftInfo *messages.Aircraft_Props_Update) error {
	query := `CALL uspUpdateAircraftProps(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.ID,
		aircraftInfo.Base,
		aircraftInfo.Callsign,
		aircraftInfo.CellPhone,
		aircraftInfo.MaxPatientWeight,
		aircraftInfo.PadTimeDay,
		aircraftInfo.PadTimeNight,
		aircraftInfo.SatPhone,
		aircraftInfo.SpecialEquipment,
		aircraftInfo.Vendor,
	); err != nil {
		return fmt.Errorf("Error updating aircraft props: %v", err)
	}
	return nil
}

// UpdateAircraftCrew updates the crewmembers that are assigned to an aircraft
func (ctx *ParserContext) UpdateAircraftCrew(aircraftInfo *messages.Aircraft_Crew_Update) error {
	query := `CALL uspUpateAircraftCrew(?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.AdultRN,
		aircraftInfo.ID,
		aircraftInfo.PediatricRN,
		aircraftInfo.PIC,
	); err != nil {
		return fmt.Errorf("Error updating aircraft crew in DB: %v", err)
	}
	return nil
}

// UpdateAircraftServiceSchedule updates the OOS status of an existing aircraft
func (ctx *ParserContext) UpdateAircraftServiceSchedule(aircraftInfo *messages.Aircraft_Service_Schedule) error {
	query := `CALL uspUpdateAircraftServiceSchedule(?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.ID,
		aircraftInfo.OosReason,
		aircraftInfo.Status,
		aircraftInfo.StartTime,
		aircraftInfo.EndTime,
	); err != nil {
		return fmt.Errorf("Error updating aircraft service schedule in DB: %v", err)
	}
	return nil
}

// UpdateAircraftPosition updates the lat/long and a human-friendly area name
// for an existing aircraft's position
func (ctx *ParserContext) UpdateAircraftPosition(aircraftInfo *messages.Aircraft_Pos_Update) error {
	query := `CALL uspUpdateAircraftPosition(?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.ID,
		aircraftInfo.PosFriendlyName,
		aircraftInfo.PosLat,
		aircraftInfo.PosLong,
	); err != nil {
		return fmt.Errorf("Error updating aircraft position in DB: %v", err)
	}
	return nil
}

// [GROUPS]

// AddNewGroup adds a new group object to the database
func (ctx *ParserContext) AddNewGroup(groupInfo *messages.Group) error {
	query := `CALL uspAddNewGroup(?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupInfo.ID,
		groupInfo.Members,
		groupInfo.Name,
	); err != nil {
		return fmt.Errorf("Error adding new group to DB: %v", err)
	}
	return nil
}

// UpdateGroup updates an existing group object
func (ctx *ParserContext) UpdateGroup(groupInfo *messages.Group) error {
	query := `CALL uspUpdateGroup(?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupInfo.ID,
		groupInfo.Members,
		groupInfo.Name,
	); err != nil {
		return fmt.Errorf("Error updating group in DB: %v", err)
	}
	return nil
}

// DeleteGroup deletes an existing group object from the database
func (ctx *ParserContext) DeleteGroup(groupInfo *messages.Group_Delete) error {
	query := `CALL uspDeleteGroup(?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupInfo.ID,
	); err != nil {
		return fmt.Errorf("Error deleting group from DB: %v", err)
	}
	return nil
}

// [MISSIONS]

func (ctx *ParserContext) GetRequestorByID(requestorID int) (string, error) {
	// query := `CALL uspGetRequestorByID(?)`
	// reqRow, err := ctx.DB.QueryContext(
	// 	sqlCtx,
	// 	query,
	// 	requestorID,
	// )
	query := `SELECT agency_name FROM tblAGENCY WHERE agency_id=?`
	reqRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		requestorID,
	)
	if err != nil {
		return "", fmt.Errorf("Error querying MySQL for requestor: %v", err)
	}
	var requestor string
	for reqRow.Next() {
		err = reqRow.Scan(&requestor)
		if err != nil {
			return "", fmt.Errorf("Error scanning requestor row: %v", err)
		}
	}
	return requestor, nil
}

func (ctx *ParserContext) GetReceiverByID(receiverID int) (string, error) {
	// query := `CALL uspGetReceiverByID(?)`
	// reqRow, err := ctx.DB.QueryContext(
	// 	sqlCtx,
	// 	query,
	// 	receiverID,
	// )
	query := `SELECT agency_name FROM tblAGENCY WHERE agency_id=(?)`
	recRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		receiverID,
	)
	if err != nil {
		return "", fmt.Errorf("Error querying MySQL for requestor: %v", err)
	}
	var receiver string
	for recRow.Next() {
		err = recRow.Scan(&receiver)
		if err != nil {
			return "", fmt.Errorf("Error scanning requestor row: %v", err)
		}
	}
	return receiver, nil
}

func (ctx *ParserContext) GetCrewMemberByID(memberID int) (string, string, error) {
	// query := `CALL uspGetCrewMemberByID(?)`
	// memRow, err := ctx.DB.QueryContext(
	// 	sqlCtx,
	// 	query,
	// 	memberID,
	// )
	query := `SELECT personnel_F_Name, personnel_L_Name FROM tblPERSONNEL WHERE personnel_id=?`
	memRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		memberID,
	)
	// ctx.DB.Query("SELECT personnel_F_Name, personnel_L_Name FROM tblPERSONNEL WHERE personnel_id=" + memberID)
	if err != nil {
		return "", "", fmt.Errorf("Error querying MySQL for member: %v", err)
	}
	var fName string
	var lName string
	for memRow.Next() {
		err = memRow.Scan(&fName, &lName)
		if err != nil {
			return "", "", fmt.Errorf("Error scanning member row: %v", err)
		}
	}
	return fName, lName, nil
}

// GetRoleByMemberID gets a member's assigned role from the database
func (ctx *ParserContext) GetRoleByMemberID(memberID int) (string, error) {
	// SELECT role_title FROM tblROLES JOIN tblASSIGNED_PERSONNEL_ROLES ON tblASSIGNED_PERSONNEL_ROLES.role_id = tblROLES.role_id JOIN tblPERSONNEL ON tblPERSONNEL.personnel_id = tblASSIGNED_PERSONNEL_ROLES.missionpersonnel_id WHERE tblPERSONNEL.personnel_id = " + strconv.Itoa(memberID)
	// query := `CALL uspGetRoleByMemberID(?)`
	// roleRow, err := ctx.DB.QueryContext(
	// 	sqlCtx,
	// 	query,
	// 	memberID,
	// )
	query := `SELECT role_title FROM tblROLES JOIN tblASSIGNED_PERSONNEL_ROLES ON tblASSIGNED_PERSONNEL_ROLES.role_id = tblROLES.role_id JOIN tblPERSONNEL ON tblPERSONNEL.personnel_id = tblASSIGNED_PERSONNEL_ROLES.missionpersonnel_id WHERE tblPERSONNEL.personnel_id=?`
	roleRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		memberID,
	)
	if err != nil {
		return "", fmt.Errorf("Error querying MySQL for member: %v", err)
	}
	var roleTitle string
	for roleRow.Next() {
		err = roleRow.Scan(&roleTitle)
		if err != nil {
			return "", fmt.Errorf("Error scanning role row: %v", err)
		}
	}
	return roleTitle, nil
}

func (ctx *ParserContext) GetWaypointNameByID(waypointID int) (string, error) {
	// query := `CALL uspGetWaypointNameByID(?)`
	query := `SELECT waypoint_title FROM tblWAYPOINT WHERE waypoint_id=?`
	wayPtRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		waypointID,
	)
	if err != nil {
		return "", fmt.Errorf("Error querying MySQL for waypoint: %v", err)
	}
	var wayPtName string
	for wayPtRows.Next() {
		err = wayPtRows.Scan(&wayPtName)
		if err != nil {
			return "", fmt.Errorf("Error scanning waypoint row: %v", err)
		}
	}
	return wayPtName, nil
}

func (ctx *ParserContext) GetTCNumByMissionID(missionID int) (string, error) {
	query := `CALL uspGetTCNumByMissionID(?)`
	missionRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionID,
	)
	if err != nil {
		return "", fmt.Errorf("Error querying MySQL for tcnum: %v", err)
	}
	var tcNum string
	for missionRow.Next() {
		err = missionRow.Scan(&tcNum)
		if err != nil {
			return "", fmt.Errorf("Error scanning tcnum row: %v", err)
		}
	}
	return tcNum, nil
}

// AddNewMission adds a new mission object to the database
func (ctx *ParserContext) AddNewMission(missionInfo *messages.Mission_Create) error {
	query := `CALL uspAddNewMission(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionInfo.Asset,
		missionInfo.CallType,
		missionInfo.CrewMemberID,
		missionInfo.MissionID,
		missionInfo.Patient,
		// missionInfo.Priority,
		missionInfo.ReceiverID,
		missionInfo.RequestorID,
		missionInfo.TCNum,
		// missionInfo.Vision,
		missionInfo.Waypoints,
	); err != nil {
		return fmt.Errorf("Error adding new mission to DB: %v", err)
	}
	return nil
}

// UpdateMissionWaypoints updates the waypoints for an existing mission
func (ctx *ParserContext) UpdateMissionWaypoints(missionInfo *messages.Mission_Waypoint_Update) error {
	query := `CALL uspUpdateMissionWaypoints(?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionInfo.MissionID,
		missionInfo.Waypoints,
	); err != nil {
		return fmt.Errorf("Error updating mission waypoints in DB: %v", err)
	}
	return nil
}

// UpdateMissionCrew updates the crew assigned to a mission
// Note: Although crew are usually tied to an aircraft, this handles
// the case when crew is reassigned for a particular mission
func (ctx *ParserContext) UpdateMissionCrew(missionInfo *messages.Mission_Crew_Update) error {
	query := `CALL uspUpdateMissionCrew(?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionInfo.MissionID,
		missionInfo.CrewMemberID,
	); err != nil {
		return fmt.Errorf("Error updating mission crew in DB: %v", err)
	}
	return nil
}

// [USERS]

// AddNewUser adds a new user object to the database
func (ctx *ParserContext) AddNewUser(userInfo *messages.User) error {
	query := `CALL uspAddNewUser(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		userInfo.ID,
		userInfo.UserName,
		userInfo.FirstName,
		userInfo.MiddleName,
		userInfo.LastName,
		userInfo.Initials,
		userInfo.Email,
		// userInfo.UWNetID,
		strconv.Itoa(userInfo.GroupID),
		userInfo.Role,
		userInfo.CellPhone,
		// userInfo.QualificationID,
	); err != nil {
		return fmt.Errorf("Error adding new user to DB: %v", err)
	}
	return nil
}

// UpdateUser updates an existing user
func (ctx *ParserContext) UpdateUser(userInfo *messages.User) error {
	query := `CALL uspUpdateUser(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		userInfo.ID,
		userInfo.UserName,
		userInfo.FirstName,
		userInfo.MiddleName,
		userInfo.LastName,
		userInfo.Initials,
		userInfo.Email,
		// userInfo.UWNetID,
		userInfo.GroupID,
		userInfo.Role,
		userInfo.CellPhone,
		// userInfo.QualificationID,
	); err != nil {
		return fmt.Errorf("Error updating user in DB: %v", err)
	}
	return nil
}

// DeleteUser deletes an existing user from the database
func (ctx *ParserContext) DeleteUser(userInfo *messages.User_Delete) error {
	query := `CALL uspDeleteUser(?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		userInfo.ID,
	); err != nil {
		return fmt.Errorf("Error deleting user from DB: %v", err)
	}
	return nil
}

// [WAYPOINTS]

// AddNewWaypoint adds a new waypoint object to the database
func (ctx *ParserContext) AddNewWaypoint(waypointInfo *messages.Waypoint) error {
	query := `CALL uspAddNewWaypoint(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		waypointInfo.ID,
		waypointInfo.Notes,
		waypointInfo.Name,
		waypointInfo.Type,
		waypointInfo.Address1,
		waypointInfo.Address2,
		waypointInfo.Country,
		waypointInfo.State,
		waypointInfo.County,
		waypointInfo.City,
		waypointInfo.Zip,
		waypointInfo.Lat,
		waypointInfo.Long,
		waypointInfo.GPSWaypoint,
		waypointInfo.AirportIdentifier,
		waypointInfo.Phone,
		waypointInfo.ShortCode,
		waypointInfo.PadTime,
		waypointInfo.RadioChannels,
		waypointInfo.NOTAMS,
	); err != nil {
		return fmt.Errorf("Error adding new waypoint to DB: %v", err)
	}
	return nil
}

// UpdateWaypoint updates an existing waypoint's information
func (ctx *ParserContext) UpdateWaypoint(waypointInfo *messages.Waypoint) error {
	query := `CALL uspUpdateWaypoint(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		waypointInfo.ID,
		waypointInfo.Notes,
		waypointInfo.Name,
		waypointInfo.Type,
		waypointInfo.Address1,
		waypointInfo.Address2,
		waypointInfo.Country,
		waypointInfo.State,
		waypointInfo.County,
		waypointInfo.City,
		waypointInfo.Zip,
		waypointInfo.Lat,
		waypointInfo.Long,
		waypointInfo.GPSWaypoint,
		waypointInfo.AirportIdentifier,
		waypointInfo.Phone,
		waypointInfo.ShortCode,
		waypointInfo.PadTime,
		waypointInfo.RadioChannels,
		waypointInfo.NOTAMS,
	); err != nil {
		return fmt.Errorf("Error updating waypoint in DB: %v", err)
	}
	return nil
}

// DeleteWaypoint deletes an existing waypoint from the database
func (ctx *ParserContext) DeleteWaypoint(waypointInfo *messages.Waypoint_Delete) error {
	query := `CALL uspDeleteWaypoint(?)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		waypointInfo.ID,
	); err != nil {
		return fmt.Errorf("Error deleting waypoint from DB: %v", err)
	}
	return nil
}

// [DUPLICATE IN handlers/queries.go]

func (ctx *ParserContext) GetPersonByID(personID int) (*messages.Person, error) {
	query := `CALL uspGetPersonByID(?)`
	personRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		personID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	person := &messages.Person{}
	for personRows.Next() {
		if err := personRows.Scan(&person.ID, &person.FName, &person.LName, &person.Position); err != nil {
			return nil, fmt.Errorf("Error populating person: %v", err)
		}
	}
	return person, nil
}

func (ctx *ParserContext) GetPersonDetailByID(personID int) (*sql.Rows, error) {
	query := `CALL uspGetPersonDetailByID(?)`
	pdRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		personID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	return pdRows, nil
}
