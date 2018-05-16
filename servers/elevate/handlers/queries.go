package handlers

import (
	"database/sql"
	"fmt"
)

func (ctx *HandlerContext) GetAllAircraft() (*sql.Rows, error) {
	// TODO sql sproc
	aircraftRows, err := ctx.DB.Query("SELECT group_id, group_name, personnel_F_Name, personnel_L_Name, personnel_id,  FROM tblPERSONNEL_GROUP JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id WHERE group_id = ORDER BY group_name")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByStatus(status string) (*sql.Rows, error) {
	// TODO sql sproc
	aircraftRows, err := ctx.DB.Query("SELECT group_id, group_name, personnel_F_Name, personnel_L_Name, personnel_id,  FROM tblPERSONNEL_GROUP JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id WHERE group_id = ORDER BY group_name")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetAircraftByID(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	aircraftRows, err := ctx.DB.Query("SELECT group_id, group_name, personnel_F_Name, personnel_L_Name, personnel_id,  FROM tblPERSONNEL_GROUP JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id WHERE group_id = ORDER BY group_name")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft: %v", err)
	}
	return aircraftRows, nil
}

func (ctx *HandlerContext) GetMissionByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	missionRows, err := ctx.DB.Query("SELECT group_id, group_name, personnel_F_Name, personnel_L_Name, personnel_id,  FROM tblPERSONNEL_GROUP JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id WHERE group_id = ORDER BY group_name")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for missions: %v", err)
	}
	return missionRows, nil
}

func (ctx *HandlerContext) GetWaypointsByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	waypointRows, err := ctx.DB.Query("SELECT things")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for waypoint: %v", err)
	}
	return waypointRows, nil
}

func (ctx *HandlerContext) GetOOSByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	oosRows, err := ctx.DB.Query("SELECT things")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for OOS status: %v", err)
	}
	return oosRows, nil
}

func (ctx *HandlerContext) GetAircraftDetailById(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	acDetailRows, err := ctx.DB.Query("SELECT things")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for aircraft details: %v", err)
	}
	return acDetailRows, nil
}

func (ctx *HandlerContext) GetCrewByAircraft(aircraftID string) (*sql.Rows, error) {
	// TODO sql sproc
	crewRows, err := ctx.DB.Query("SELECT things")
	if err != nil {
		return nil, fmt.Errorf("Error querying MySQL for crew: %v", err)
	}
	return crewRows, nil
}
