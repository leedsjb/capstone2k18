package handlers

import (
	"context"
	"database/sql"
	"fmt"
)

// empty, no-time-out context for sql queries
var sqlCtx = context.Background()

// [AIRCRAFT QUERIES]

func (ctx *HandlerContext) GetAllAircraft() (*sql.Rows, error) {
	aircraftRows, err := ctx.DB.Query("CALL uspGetAllAircraft()")
	// ID           string
	// Callsign     string
	// Nnum         string
	// Manufacturer string // i.e. Augusta, Learjet, etc
	// Title        string // i.e. A109E, PC-12, etc
	// Class        string // i.e. Rotorcraft, Fixed-wing
	// Lat          string
	// Long         string
	// LocationName string
	// Status
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByStatus(status string) (*sql.Rows, error) {
	query := `CALL uspGetAircraftByStatus(?)`
	aircraftRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		status,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft by status: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByCategory(category string) (*sql.Rows, error) {
	query := `CALL uspGetAircraftByCategory(?)`
	aircraftRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		category,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft by category: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByStatusAndCategory(status string, category string) (*sql.Rows, error) {
	query := `CALL uspGetAircraftByStatusAndCategory(?, ?)`
	aircraftRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		status,
		category,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft by status and category: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByID(aircraftID int) (*sql.Rows, error) {
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

func (ctx *HandlerContext) GetMissionByAircraft(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetMissionByAircraft(?)`
	missionRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for missions: %v", err)
	}
	return missionRows, nil
}

func (ctx *HandlerContext) GetWaypointsByAircraft(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetWaypointsByAircraft(?)`
	waypointRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for waypoint: %v", err)
	}
	return waypointRows, nil
}

func (ctx *HandlerContext) GetOOSByAircraft(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetOOSByAircraft(?)`
	oosRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for OOS status: %v", err)
	}
	return oosRows, nil
}

func (ctx *HandlerContext) GetAircraftDetailByID(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetAircraftDetailByID`
	acDetailRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft details: %v", err)
	}
	return acDetailRows, nil
}

func (ctx *HandlerContext) GetCrewByAircraft(aircraftID int) (*sql.Rows, error) {
	query := "CALL uspGetCrewByAircraft(?)"
	crewRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft crew: %v", err)
	}
	return crewRows, nil
}

func (ctx *HandlerContext) GetCrewByMission(missionID int) (*sql.Rows, error) {
	query := "CALL uspGetCrewByMission(?)"
	crewRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		missionID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for mission crew: %v", err)
	}
	return crewRows, nil
}

func (ctx *HandlerContext) GetMissionDetailByAircraft(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetMissionDetailByAircraft(?)`
	mdRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for mission details: %v", err)
	}
	return mdRows, nil
}

func (ctx *HandlerContext) GetPatientByAircraft(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetPatientByAircraft(?)`
	patientRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for patient info: %v", err)
	}
	return patientRows, nil
}

func (ctx *HandlerContext) GetOOSDetailByAircraft(aircraftID int) (*sql.Rows, error) {
	query := `CALL uspGetOOSDetailByAircraft(?)`
	OOSDetailRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		aircraftID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for OOS details: %v", err)
	}
	return OOSDetailRows, nil
}

// [GROUP GUERIES]

func (ctx *HandlerContext) GetAllGroups() (*sql.Rows, error) {
	groupRows, err := ctx.DB.Query("CALL uspGetAllGroups()")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for groups: %v", err)
	}
	return groupRows, nil
}

func (ctx *HandlerContext) GetGroupDetailByID(groupID int) (*sql.Rows, error) {
	query := `CALL uspGetGroupDetailByID(?)`
	gdRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for groups: %v", err)
	}
	return gdRows, nil
}

func (ctx *HandlerContext) GetGroupByID(groupID int) (*sql.Rows, error) {
	query := `CALL uspGetGroupByID(?)`
	groupRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		groupID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for group: %v", err)
	}
	return groupRows, nil
}

// [PEOPLE QUERIES]

/*
	SELECT personnel_id, personnel_F_Name, personnel_L_Name, role_title
*/
func (ctx *HandlerContext) GetAllPeople() (*sql.Rows, error) {
	peopleRows, err := ctx.DB.Query("CALL uspGetAllPeople()")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	return peopleRows, nil
}

func (ctx *HandlerContext) GetPersonDetailByID(personID int) (*sql.Rows, error) {
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

func (ctx *HandlerContext) GetPersonByID(personID int) (*sql.Rows, error) {
	query := `CALL uspGetPersonByID(?)`
	personRows, err := ctx.DB.QueryContext(
		sqlCtx,
		query,
		personID,
	)
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	return personRows, nil
}

// [RESOURCE QUERIES]

func (ctx *HandlerContext) GetAllResources() (*sql.Rows, error) {
	resourceRows, err := ctx.DB.Query("CALL uspGetResources()")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for resource details: %v", err)
	}
	return resourceRows, nil
}
