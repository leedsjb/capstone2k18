package parsers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// empty, no-time-out context for sql queries
var sqlCtx = context.Background()

// [AIRCRAFT]

func (ctx *ParserContext) GetAircraftCallsignByID(aircraftID string) (string, error) {
	query := `CALL uspGetAircraftCallsignByID(?)`
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
			return "", fmt.Errorf("Error retrieving aircraft callsign by ID: %v\n", err)
		}
	}
	close(aircraftRow)
	return aircraftCallsign, nil
}

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
		return 0, fmt.Errorf("Error querying MySQL for aircraft ID  by callsign: %v\n", err)
	}
	var aircraftID int
	for aircraftRow.Next() {
		err = aircraftRow.Scan(&aircraftID)
		if err != nil {
			return 0, fmt.Errorf("Error retrieving aircraft ID by callsign: %v\n", err)
		}
	}
	close(aircraftRow)
	return aircraftID, nil
}

func (ctx *ParserContext) GetAircraftIDByMission(missionID string) (int, error) {
	query := `CALL uspGetAircraftIDByMission(?)`
	aircraftRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionID,
	)
	if err != nil {
		return 0, fmt.Errorf("Error querying MySQL for aircraft ID by mission: %v\n", err)
	}
	var aircraftID int
	for aircraftRow.Next() {
		err = aircraftRow.Scan(&aircraftID)
		if err != nil {
			return 0, fmt.Errorf("Error retrieving aircraft ID by mission: %v\n", err)
		}
	}
	close(aircraftRow)
	return aircraftID, nil
}

// GetAircraftCallsign retrieves an aircraft's callsign given mission's ID
func (ctx *ParserContext) GetAircraftCallsignByMission(missionID int) (string, error) {
	// get callsign from db using aircraftID
	// aircraftID := strconv.Itoa(ID)

	// "SELECT ac_callsign FROM tblAIRCRAFT JOIN tblMISSION ON tblMISSION.aircraft_id = tblAIRCRAFT.ac_id WHERE mission_id=" + aircraftID

	query := `CALL uspGetAircraftCallsign(?)`
	aircraftRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionID,
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
	close(aircraftRow)
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
	connect, err := ctx.DB.QueryContext(
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
	)
	if err != nil {
		return fmt.Errorf("Error adding aircraft to DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateAircraftProps updates an existing aircraft object
func (ctx *ParserContext) UpdateAircraftProps(aircraftInfo *messages.Aircraft_Props_Update) error {
	query := `CALL uspUpdateAircraftProps(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
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
	)
	if err != nil {
		return fmt.Errorf("Error updating aircraft props: %v", err)
	}
	close(connect)
	return nil
}

// UpdateAircraftCrew updates the crewmembers that are assigned to an aircraft
func (ctx *ParserContext) UpdateAircraftCrew(aircraftInfo *messages.Aircraft_Crew_Update) error {
	query := `CALL uspUpateAircraftCrew(?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.AdultRN,
		aircraftInfo.ID,
		aircraftInfo.PediatricRN,
		aircraftInfo.PIC,
	)
	if err != nil {
		return fmt.Errorf("Error updating aircraft crew in DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateAircraftServiceSchedule updates the OOS status of an existing aircraft
func (ctx *ParserContext) UpdateAircraftServiceSchedule(aircraftInfo *messages.Aircraft_Service_Schedule) error {
	query := `CALL uspUpdateAircraftServiceSchedule(?, ?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.ID,
		aircraftInfo.OosReason,
		aircraftInfo.Status,
		aircraftInfo.StartTime,
		aircraftInfo.EndTime,
	)
	if err != nil {
		return fmt.Errorf("Error updating aircraft service schedule in DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateAircraftPosition updates the lat/long and a human-friendly area name
// for an existing aircraft's position
func (ctx *ParserContext) UpdateAircraftPosition(aircraftInfo *messages.Aircraft_Pos_Update) error {
	query := `CALL uspUpdateACLocation(?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.ID,
		aircraftInfo.PosLat,
		aircraftInfo.PosLong,
		aircraftInfo.PosFriendlyName,
	)
	if err != nil {
		return fmt.Errorf("Error updating aircraft position in DB: %v", err)
	}
	fmt.Printf("[AIRCRAFT POS QUERY] Updated aircraft!!!")
	close(connect)
	return nil
}

// [GROUPS]

// AddNewGroup adds a new group object to the database
func (ctx *ParserContext) AddNewGroup(groupInfo *messages.Group) error {
	query := `CALL uspAddNewGroup(?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupInfo.ID,
		groupInfo.Members,
		groupInfo.Name,
	)
	if err != nil {
		return fmt.Errorf("Error adding new group to DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateGroup updates an existing group object
func (ctx *ParserContext) UpdateGroup(groupInfo *messages.Group) error {
	query := `CALL uspUpdateGroup(?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupInfo.ID,
		groupInfo.Members,
		groupInfo.Name,
	)
	if err != nil {
		return fmt.Errorf("Error updating group in DB: %v", err)
	}
	close(connect)
	return nil
}

// DeleteGroup deletes an existing group object from the database
func (ctx *ParserContext) DeleteGroup(groupInfo *messages.Group_Delete) error {
	query := `CALL uspDeleteGroup(?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupInfo.ID,
	)
	if err != nil {
		return fmt.Errorf("Error deleting group from DB: %v", err)
	}
	close(connect)
	return nil
}

// [MISSIONS]

func (ctx *ParserContext) GetAgencyDetailsByID(agencyID int) (*messages.Agency, error) {
	// query := `CALL uspGetAgencyByID(?)`
	// reqRow, err := ctx.DB.QueryContext(
	// 	sqlCtx,
	// 	query,
	// 	requestorID,
	// )
	query := `CALL uspGetAgencyDetailsByID(?)`
	agencyRow, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		agencyID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for agency: %v", err)
	}
	agency := &messages.Agency{}
	var name string
	var areaCode string
	var phone string
	var agencyType string
	var street1 sql.NullString
	var street2 sql.NullString
	var city string
	var state string
	var zip string
	for agencyRow.Next() {
		err = agencyRow.Scan(
			&name,
			&areaCode,
			&phone,
			&agencyType,
			&street1,
			&street2,
			&city,
			&state,
			&zip,
		)
		if err != nil {
			return nil, fmt.Errorf("Error scanning agency row: %v", err)
		}

		var address string
		if street1.Valid {
			address += street1.String
		}
		if street2.Valid {
			address += street2.String
		}

		phone = areaCode + phone

		agency = &messages.Agency{
			Name:    name,
			Phone:   phone,
			Type:    agencyType,
			Address: address,
			City:    city,
			State:   state,
			Zip:     zip,
		}
	}
	close(agencyRow)
	return agency, nil
}

// func (ctx *ParserContext) GetReceiverByID(receiverID int) (string, error) {
// 	// query := `CALL uspGetReceiverByID(?)`
// 	// reqRow, err := ctx.DB.QueryContext(
// 	// 	sqlCtx,
// 	// 	query,
// 	// 	receiverID,
// 	// )
// 	query := `SELECT agency_name FROM tblAGENCY WHERE agency_id=(?)`
// 	recRow, err := ctx.DB.QueryContext(
// 		sqlCtx,
// 		query,
// 		receiverID,
// 	)
// 	if err != nil {
// 		return "", fmt.Errorf("Error querying MySQL for requestor: %v", err)
// 	}
// 	var receiver string
// 	for recRow.Next() {
// 		err = recRow.Scan(&receiver)
// 		if err != nil {
// 			return "", fmt.Errorf("Error scanning requestor row: %v", err)
// 		}
// 	}
// 	return receiver, nil
// }

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
	close(memRow)
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
	close(roleRow)
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
	close(wayPtRows)
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
	close(missionRow)
	return tcNum, nil
}

// AddNewMission adds a new mission object to the database
func (ctx *ParserContext) NewMission(missionInfo *messages.Mission_Create, aircraftID int) error {

	marshaledMembers, err := json.Marshal(missionInfo.CrewMembers)
	if err != nil {
		return fmt.Errorf("Error marshaling crew members: %v", err)
	}
	crew := string(marshaledMembers)

	fmt.Printf("[STRINGIFIED] crew: %v\n", crew)

	marshaledWaypoints, err := json.Marshal(missionInfo.Waypoints)
	if err != nil {
		return fmt.Errorf("Error marshaling waypoints: %v", err)
	}
	waypoints := string(marshaledWaypoints)

	fmt.Printf("[STRINGIFIED] waypoints: %v\n", waypoints)

	query := `CALL uspNewMission(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionInfo.MissionID,
		missionInfo.TCNum,
		aircraftID,
		missionInfo.RequestorID,
		missionInfo.ReceiverID,
		missionInfo.CallType,
		missionInfo.Patient.ShortReport,
		missionInfo.Patient.Intubated,
		missionInfo.Patient.Drips,
		missionInfo.Patient.Age,
		missionInfo.Patient.Weight,
		missionInfo.Patient.Gender,
		missionInfo.Patient.Cardiac,
		missionInfo.Patient.GIBleed,
		missionInfo.Patient.OB,
		crew,
		waypoints,
	)
	if err != nil {
		return fmt.Errorf("Error adding new mission to DB: %v", err)
	}
	fmt.Printf("[NEW MISSON] SUCCESFULLY ADDED TO DB?!?!?!?!!!")
	close(connect)
	return nil
}

func (ctx *ParserContext) CompleteMission(missionID string) error {
	query := `CALL uspCompleteMission(?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionID,
	)
	if err != nil {
		return fmt.Errorf("Error completing mission in DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateMissionWaypoints updates the waypoints for an existing mission
func (ctx *ParserContext) UpdateMissionWaypoints(missionInfo *messages.Mission_Waypoint_Update) error {
	query := `CALL uspUpdateMissionWaypoints(?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionInfo.MissionID,
		missionInfo.Waypoints,
	)
	if err != nil {
		return fmt.Errorf("Error updating mission waypoints in DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateMissionCrew updates the crew assigned to a mission
// Note: Although crew are usually tied to an aircraft, this handles
// the case when crew is reassigned for a particular mission
func (ctx *ParserContext) UpdateMissionCrew(missionInfo *messages.Mission_Crew_Update) error {
	query := `CALL uspUpdateMissionCrew(?, ?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionInfo.MissionID,
		missionInfo.CrewMembers,
	)
	if err != nil {
		return fmt.Errorf("Error updating mission crew in DB: %v", err)
	}
	close(connect)
	return nil
}

// [USERS]

// AddNewUser adds a new user object to the database
func (ctx *ParserContext) AddNewUser(userInfo *messages.User) error {
	query := `CALL uspAddNewUser(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
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
	)
	if err != nil {
		return fmt.Errorf("Error adding new user to DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateUser updates an existing user
func (ctx *ParserContext) UpdateUser(userInfo *messages.User) error {
	query := `CALL uspUpdateUser(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
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
	)
	if err != nil {
		return fmt.Errorf("Error updating user in DB: %v", err)
	}
	close(connect)
	return nil
}

// DeleteUser deletes an existing user from the database
func (ctx *ParserContext) DeleteUser(userInfo *messages.User_Delete) error {
	query := `CALL uspDeleteUser(?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		userInfo.ID,
	)
	if err != nil {
		return fmt.Errorf("Error deleting user from DB: %v", err)
	}
	close(connect)
	return nil
}

// [WAYPOINTS]

// AddNewWaypoint adds a new waypoint object to the database
func (ctx *ParserContext) AddNewWaypoint(waypointInfo *messages.Waypoint) error {
	query := `CALL uspAddNewWaypoint(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
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
	)
	if err != nil {
		return fmt.Errorf("Error adding new waypoint to DB: %v", err)
	}
	close(connect)
	return nil
}

// UpdateWaypoint updates an existing waypoint's information
func (ctx *ParserContext) UpdateWaypoint(waypointInfo *messages.Waypoint) error {
	query := `CALL uspUpdateWaypoint(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	connect, err := ctx.DB.QueryContext(
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
	)
	if err != nil {
		return fmt.Errorf("Error updating waypoint in DB: %v", err)
	}
	close(connect)
	return nil
}

// DeleteWaypoint deletes an existing waypoint from the database
func (ctx *ParserContext) DeleteWaypoint(waypointInfo *messages.Waypoint_Delete) error {
	query := `CALL uspDeleteWaypoint(?)`
	connect, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		waypointInfo.ID,
	)
	if err != nil {
		return fmt.Errorf("Error deleting waypoint from DB: %v", err)
	}
	close(connect)
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
	close(personRows)
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
