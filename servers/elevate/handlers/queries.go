package handlers

import (
	"database/sql"
	"fmt"
)

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
	aircraftRows, err := ctx.DB.Query("CALL uspGetAircraftByStatus(" + status + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByID(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	aircraftRows, err := ctx.DB.Query("CALL uspGetAircraftByID(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetMissionByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	missionRows, err := ctx.DB.Query("CALL uspGetMissionByAircraft(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for missions: %v", err)
	}
	return missionRows, nil
}

func (ctx *HandlerContext) GetWaypointsByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	waypointRows, err := ctx.DB.Query("CALL uspGetWaypointsByAircraft(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for waypoint: %v", err)
	}
	return waypointRows, nil
}

func (ctx *HandlerContext) GetOOSByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	oosRows, err := ctx.DB.Query("CALL uspGetOOSByAircraft(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for OOS status: %v", err)
	}
	return oosRows, nil
}

func (ctx *HandlerContext) GetAircraftDetailByID(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	acDetailRows, err := ctx.DB.Query("CALL uspGetAircraftDetailByID(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft details: %v", err)
	}
	return acDetailRows, nil
}

func (ctx *HandlerContext) GetCrewByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	crewRows, err := ctx.DB.Query("CALL uspGetCrewByAircraft(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for crew: %v", err)
	}
	return crewRows, nil
}

func (ctx *HandlerContext) GetMissionDetailsByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	mdRows, err := ctx.DB.Query("CALL uspGetMissionDetailsByAircraft(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for mission details: %v", err)
	}
	return mdRows, nil
}

func (ctx *HandlerContext) GetPatientByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	patientRows, err := ctx.DB.Query("CALL uspGetPatientByAircraft(" + aircraftID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for patient info: %v", err)
	}
	return patientRows, nil
}

func (ctx *HandlerContext) GetOOSDetailByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	OOSDetailRows, err := ctx.DB.Query("CALL uspGetOOSDetailByAircraft(" + aircraftID + ")")
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

func (ctx *HandlerContext) GetGroupDetailByID(groupID string) (*sql.Rows, error) {
	// TODO sql sproc
	gdRows, err := ctx.DB.Query("CALL uspGetGroupDetailByID(" + groupID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for groups: %v", err)
	}
	return gdRows, nil
}

func (ctx *HandlerContext) GetGroupByID(groupID string) (*sql.Rows, error) {
	// TODO sql sproc
	groupRows, err := ctx.DB.Query("CALL uspGetGroupByID(" + groupID + ")")
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

func (ctx *HandlerContext) GetPersonDetailsByID(personID string) (*sql.Rows, error) {
	// TODO sql sproc
	pdRows, err := ctx.DB.Query("CALL uspGetPersonDetailsByID(" + personID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	return pdRows, nil
}

func (ctx *HandlerContext) GetPersonByID(personID string) (*sql.Rows, error) {
	// TODO sql sproc
	personRows, err := ctx.DB.Query("CALL uspGetPersonByID(" + personID + ")")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for person details: %v", err)
	}
	return personRows, nil
}
