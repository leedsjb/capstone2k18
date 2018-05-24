package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

	// aircraftRows, err := ctx.DB.Query("CALL uspGetAircraftByStatus(\"" + status + "\")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByID(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	aircraftRows, err := ctx.DB.Query("CALL uspGetAircraftByID(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetMissionByAircraft(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	missionRows, err := ctx.DB.Query("CALL uspGetMissionByAircraft(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for missions: %v", err)
	}
	return missionRows, nil
}

func (ctx *HandlerContext) GetWaypointsByAircraft(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	waypointRows, err := ctx.DB.Query("CALL uspGetWaypointsByAircraft(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for waypoint: %v", err)
	}
	return waypointRows, nil
}

func (ctx *HandlerContext) GetOOSByAircraft(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	oosRows, err := ctx.DB.Query("CALL uspGetOOSByAircraft(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for OOS status: %v", err)
	}
	return oosRows, nil
}

func (ctx *HandlerContext) GetAircraftDetailByID(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	acDetailRows, err := ctx.DB.Query("CALL uspGetAircraftDetailByID(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft details: %v", err)
	}
	return acDetailRows, nil
}

func (ctx *HandlerContext) GetCrewByAircraft(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	crewRows, err := ctx.DB.Query("CALL uspGetCrewByAircraft(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for crew: %v", err)
	}
	return crewRows, nil
}

func (ctx *HandlerContext) GetMissionDetailByAircraft(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	mdRows, err := ctx.DB.Query("CALL uspGetMissionDetailByAircraft(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for mission details: %v", err)
	}
	return mdRows, nil
}

func (ctx *HandlerContext) GetPatientByAircraft(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	patientRows, err := ctx.DB.Query("CALL uspGetPatientByAircraft(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for patient info: %v", err)
	}
	return patientRows, nil
}

func (ctx *HandlerContext) GetOOSDetailByAircraft(aircraftID int) (*sql.Rows, error) {
	// TODO sql sproc
	OOSDetailRows, err := ctx.DB.Query("CALL uspGetOOSDetailByAircraft(" + strconv.Itoa(aircraftID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for OOS details: %v", err)
	}
	return OOSDetailRows, nil
}

// [GROUP GUERIES]

func (ctx *HandlerContext) GetAllGroups() (*sql.Rows, error) {
	// TODO sql sproc
	groupRows, err := ctx.DB.Query("CALL uspGetAllGroups()")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for groups: %v", err)
	}
	return groupRows, nil
}

func (ctx *HandlerContext) GetGroupDetailByID(groupID int) (*sql.Rows, error) {
	// TODO sql sproc
	gdRows, err := ctx.DB.Query("CALL uspGetGroupDetailByID(" + strconv.Itoa(groupID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for groups: %v", err)
	}
	return gdRows, nil
}

func (ctx *HandlerContext) GetGroupByID(groupID int) (*sql.Rows, error) {
	// TODO sql sproc
	groupRows, err := ctx.DB.Query("CALL uspGetGroupByID(" + strconv.Itoa(groupID) + ")")
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
	// TODO sql sproc
	peopleRows, err := ctx.DB.Query("CALL uspGetAllPeople()")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	return peopleRows, nil
}

func (ctx *HandlerContext) GetPersonDetailByID(personID int) (*sql.Rows, error) {
	// TODO sql sproc
	pdRows, err := ctx.DB.Query("CALL uspGetPersonDetailByID(" + strconv.Itoa(personID) + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	return pdRows, nil
}

func (ctx *HandlerContext) GetPersonByID(personID int) (*sql.Rows, error) {
	// TODO sql sproc
	personRows, err := ctx.DB.Query("CALL uspGetPersonByID(" + strconv.Itoa(personID) + ")")
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
