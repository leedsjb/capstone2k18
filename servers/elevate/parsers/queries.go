package parsers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// empty, no-time-out context for sql queries
var sqlCtx = context.Background()

// [AIRCRAFT]

// GetAircraftCallsign retrieves an aircraft's callsign given a missionID
func (ctx *ParserContext) GetAircraftCallsign(missionID string) (string, error) {
	// get mission from db using missionID
	aircraftRow, err := ctx.DB.Query("SELECT ac_callsign FROM tblAIRCRAFT JOIN tblMISSION ON tblMISSION.aircraft_id = tblAIRCRAFT.ac_id WHERE mission_id=" + missionID)
	if err != nil {
		fmt.Printf("Error querying MySQL for aircraftID: %v", err)
	}
	var aircraftCallsign string
	err = aircraftRow.Scan(&aircraftCallsign)
	if err != nil {
		return "", err
	}
	return aircraftCallsign, nil
}

// GetAircraftByID retrieves an aircraft object with the matching ID from the database
func (ctx *ParserContext) GetAircraftByID(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	aircraftRows, err := ctx.DB.Query("CALL uspGetAircraftByID(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

// AddNewAircraft adds a new aircraft object to the database
func (ctx *ParserContext) AddNewAircraft(aircraftInfo *messages.Aircraft_Create) error {
	query := `CALL uspAddNewUser($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.Base,
		aircraftInfo.Callsign,
		aircraftInfo.CallTypes,
		aircraftInfo.CellPhone,
		aircraftInfo.Color,
		aircraftInfo.ID,
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
	query := `CALL uspUpdateAircraftProps($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftInfo.Base,
		aircraftInfo.Callsign,
		aircraftInfo.CellPhone,
		aircraftInfo.ID,
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
	query := `CALL uspUpateAircraftCrew($1, $2, $3, $4)`
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
	query := `CALL uspUpdateAircraftServiceSchedule($1, $2, $3, $4, $5)`
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
	query := `CALL uspUpdateAircraftPosition($1, $2, $3, $4)`
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
	query := `CALL uspAddNewGroup($1, $2, $3)`
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
	query := `CALL uspUpdateGroup($1, $2, $3)`
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
	query := `CALL uspDeleteGroup($1)`
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

// AddNewMission adds a new mission object to the database
func (ctx *ParserContext) AddNewMission(missionInfo *messages.Mission_Create) error {
	query := `CALL uspAddNewMission($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionInfo.Asset,
		missionInfo.CallType,
		missionInfo.CrewMemberID,
		missionInfo.MissionID,
		missionInfo.Patient,
		missionInfo.Priority,
		missionInfo.ReceiverID,
		missionInfo.RequestorID,
		missionInfo.TCNum,
		missionInfo.Vision,
		missionInfo.Waypoints,
	); err != nil {
		return fmt.Errorf("Error adding new mission to DB: %v", err)
	}
	return nil
}

// UpdateMissionWaypoints updates the waypoints for an existing mission
func (ctx *ParserContext) UpdateMissionWaypoints(missionInfo *messages.Mission_Waypoint_Update) error {
	query := `CALL uspUpdateMissionWaypoints($1, $2)`
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
	query := `CALL uspUpdateMissionCrew($1, $2)`
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
	query := `CALL uspAddNewUser($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
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
		userInfo.UWNetID,
		userInfo.GroupID,
		userInfo.Role,
		userInfo.CellPhone,
		userInfo.QualificationID,
	); err != nil {
		return fmt.Errorf("Error adding new user to DB: %v", err)
	}
	return nil
}

// UpdateUser updates an existing user
func (ctx *ParserContext) UpdateUser(userInfo *messages.User) error {
	query := `CALL uspUpdateUser($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
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
		userInfo.UWNetID,
		userInfo.GroupID,
		userInfo.Role,
		userInfo.CellPhone,
		userInfo.QualificationID,
	); err != nil {
		return fmt.Errorf("Error updating user in DB: %v", err)
	}
	return nil
}

// DeleteUser deletes an existing user from the database
func (ctx *ParserContext) DeleteUser(userInfo *messages.User_Delete) error {
	query := `CALL uspDeleteUser($1)`
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
	query := `CALL uspAddNewWaypoint($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`
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
	query := `CALL uspUpdateWaypoint($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`
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
	query := `CALL uspDeleteWaypoint($1)`
	if _, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		waypointInfo.ID,
	); err != nil {
		return fmt.Errorf("Error deleting waypoint from DB: %v", err)
	}
	return nil
}
