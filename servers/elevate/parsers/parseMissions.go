package parsers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// ParseMissionCreate handles when a mission is assigned to
// an aircraft
// notifies client and writes new info to db
// assumes Mission_Create topic comes with all information
func ParseMissionCreate(msg *messages.Mission_Create,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	// unmarshal json into correct struct
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	// parse pubsub message for client

	requestor := ""
	receiver := ""
	aircraftStatus := "on a mission" // assume aircraft assigned to mission is on that mission
	// is raw MissionID what we want, or is it mapped?

	if msg.RequestorID != "" {
		reqRow, err := db.Query("SELECT agency_name FROM tblAgency WHERE agency_id=" + msg.RequestorID)
		if err != nil {
			fmt.Printf("Error querying MySQL for requestor: %v", err)
		}
		err = reqRow.Scan(&requestor)
		if err != nil {
			fmt.Printf("Error scanning requestor row: %v", err)
			os.Exit(1)
		}
		msg.RequestorID = requestor
	}
	if msg.ReceiverID != "" {
		recRow, err := db.Query("SELECT agency_name FROM tblAgency WHERE agency_id=" + msg.ReceiverID)
		if err != nil {
			fmt.Printf("Error querying MySQL for receiver: %v", err)
		}
		var receiver string
		err = recRow.Scan(&receiver)
		if err != nil {
			fmt.Printf("Error scanning receiver row: %v", err)
			os.Exit(1)
		}
		msg.ReceiverID = receiver
	}

	// separate crewIDs to build crew members into related groups
	people := []*messages.Person{}

	if len(msg.CrewMemberID) > 0 {
		for _, memberID := range msg.CrewMemberID {
			// retrieve member first and last name
			var fName string
			var lName string
			memRow, err := db.Query("SELECT personnel_F_Name, personnel_L_Name FROM tblPERSONNEL WHERE personnel_id=" + memberID)
			if err != nil {
				fmt.Printf("Error querying MySQL for member: %v", err)
			}
			err = memRow.Scan(&fName, &lName)
			if err != nil {
				fmt.Printf("Error scanning member row: %v", err)
				os.Exit(1)
			}

			// retrieve member role
			roleTitle := ""
			roleRow, err := db.Query("SELECT role_title FROM tblROLES JOIN tblASSIGNED_PERSONNEL_ROLES ON tblASSIGNED_PERSONNEL_ROLES.role_id = tblROLES.role_id JOIN tblPERSONNEL ON tblPERSONNEL.personnel_id = tblASSIGNED_PERSONNEL_ROLES.missionpersonnel_id WHERE tblPERSONNEL.personnel_id = " + memberID)
			if err != nil {
				fmt.Printf("Error querying MySQL for member: %v", err)
			}
			err = roleRow.Scan(&roleTitle)
			if err != nil {
				fmt.Printf("Error scanning role row: %v", err)
				os.Exit(1)
			}

			// fill Person object with crew member info
			person := &messages.Person{
				ID:       memberID,
				FName:    fName,
				LName:    lName,
				Position: roleTitle,
			}
			people = append(people, person)
		}
	}

	var waypoints []*messages.ClientMissionWaypoint
	nextWaypointETE := ""
	if len(msg.Waypoints) > 0 {
		for _, waypoint := range msg.Waypoints {
			wayPtRow, err := db.Query("SELECT waypoint FROM Waypoints WHERE waypointID=" + waypoint.ID)
			if err != nil {
				fmt.Printf("Error querying MySQL for waypoint: %v", err)
			}
			var wayPtName string
			err = wayPtRow.Scan(&wayPtName)
			if err != nil {
				fmt.Printf("Error scanning waypoint row: %v", err)
				os.Exit(1)
			}
			tempWayPt := &messages.ClientMissionWaypoint{
				Name:        wayPtName,
				ETE:         waypoint.ETE,
				ETT:         waypoint.ETT,
				Active:      waypoint.Active,
				FlightRules: waypoint.FlightRules,
			}
			if strings.ToLower(tempWayPt.Active) == "true" {
				nextWaypointETE = tempWayPt.ETE
			}
		}
	}

	mission := &messages.Mission{
		Type:            msg.CallType,
		Vision:          msg.Vision,
		NextWaypointETE: nextWaypointETE,
		Waypoints:       waypoints,
		FlightNum:       msg.TCNum,
	}

	aircraft := &messages.Aircraft{
		Status:   aircraftStatus,
		Callsign: msg.Asset,
		Mission:  mission,
	}

	missionDetail := &messages.MissionDetail{
		Type:            msg.CallType,
		Vision:          msg.Vision,
		NextWaypointETE: nextWaypointETE,
		Waypoints:       waypoints,
		FlightNum:       msg.TCNum,
		RadioReport:     msg.Patient,
		Requestor:       requestor,
		Receiver:        receiver,
	}

	aircraftDetail := &messages.AircraftDetail{
		Status:   aircraftStatus,
		Callsign: msg.Asset,
		Crew:     people,
		Mission:  missionDetail,
	}

	clientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg, notifier)
	clientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg, notifier)
}

// ParseMissionWaypointsUpdate handles changes to a mission's
// waypoints, including ETE, ETT, friendly names,
// and modifications to the route
// notifies client and writes new info to db
func ParseMissionWaypointsUpdate(msg *messages.Mission_Waypoint_Update,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	// unmarshal json into correct struct
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	// parse pubsub message for client

	waypoints := []*messages.ClientMissionWaypoint{}
	nextWaypointETE := ""
	aircraftStatus := "available" // assume if waypoints are updated and none are active, mission complete

	if len(msg.Waypoints) > 0 {
		for _, waypoint := range msg.Waypoints {
			// TODO: Fix SQL query
			wayPtRow, err := db.Query("SELECT waypoint_title FROM tblWAYPOINT WHERE waypoint_id=" + waypoint.ID)
			if err != nil {
				fmt.Printf("Error querying MySQL for waypoint: %v", err)
			}
			var wayPtName string
			err = wayPtRow.Scan(&wayPtName)
			if err != nil {
				fmt.Printf("Error scanning waypoint row: %v", err)
				os.Exit(1)
			}
			tempWayPt := &messages.ClientMissionWaypoint{
				Name:        wayPtName,
				ETE:         waypoint.ETE,
				ETT:         waypoint.ETT,
				Active:      waypoint.Active,
				FlightRules: waypoint.FlightRules,
			}
			if strings.ToLower(tempWayPt.Active) == "true" {
				nextWaypointETE = tempWayPt.ETE
				aircraftStatus = "on a mission" // if any waypoints active, mission must be active
			}
			waypoints = append(waypoints, tempWayPt)
		}
	}

	// type Mission_Waypoint_Update struct {
	// 	MissionID		string 		`json:"missionID"`
	// 	Waypoints		[]*Waypoint `json:"waypoints"`
	// }

	// [START format aircraft]
	// get mission from db using missionID
	aircraftCallsign, err := getAircraftCallsign(msg.MissionID, db)
	if err != nil {
		fmt.Printf("Error getting aircraftCallsign: %v", err)
		// TODO: continue with empty aircraft callsign?
	}

	mission := &messages.Mission{
		NextWaypointETE: nextWaypointETE,
		Waypoints:       waypoints,
	}

	aircraft := &messages.Aircraft{
		Status:   aircraftStatus, // assume aircraft assigned to mission is on that mission
		Callsign: aircraftCallsign,
		Mission:  mission,
	}
	// [END format aircraft]

	// [START format aircraftDetail]
	missionRow, err := db.Query("SELECT tc_number FROM tblMISSION WHERE mission_id=" + msg.MissionID)
	if err != nil {
		fmt.Printf("Error querying MySQL for mission: %v", err)
	}
	var tcNum string
	err = missionRow.Scan(&tcNum)
	if err != nil {
		fmt.Printf("Error scanning mission row: %v", err)
		os.Exit(1)
	}

	missionDetail := &messages.MissionDetail{
		NextWaypointETE: nextWaypointETE,
		Waypoints:       waypoints,
		FlightNum:       tcNum,
	}

	aircraftDetail := &messages.AircraftDetail{
		Status:   aircraftStatus,
		Callsign: aircraftCallsign,
		Mission:  missionDetail,
	}
	// [END format aircraftDetail]

	clientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg, notifier)
	clientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg, notifier)
}

// ParseMissionCrewUpdate handles when an aircraft has crew adjusted
// with respect to an assigned mission
// notifies client, writes new info to db
func ParseMissionCrewUpdate(msg *messages.Mission_Crew_Update,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	// unmarshal json into correct struct
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	// parse pubsub message for client
	// type Mission_Crew_Update struct {
	// 	MissionID    string   `json:"missionID"`
	// 	CrewMemberID []string `json:"crewMemberID"`
	// }

	// separate crewIDs to build crew members into related groups
	people := []*messages.Person{}

	if len(msg.CrewMemberID) > 0 {
		for _, memberID := range msg.CrewMemberID {
			// retrieve member first and last name
			var fName string
			var lName string
			memRow, err := db.Query("SELECT personnel_F_Name, personnel_L_Name FROM tblPERSONNEL WHERE personnel_id=" + memberID)
			if err != nil {
				fmt.Printf("Error querying MySQL for member: %v", err)
			}
			err = memRow.Scan(&fName, &lName)
			if err != nil {
				fmt.Printf("Error scanning member row: %v", err)
				os.Exit(1)
			}

			// retrieve member role
			roleTitle := ""
			roleRow, err := db.Query("SELECT role_title FROM tblROLES JOIN tblASSIGNED_PERSONNEL_ROLES ON tblASSIGNED_PERSONNEL_ROLES.role_id = tblROLES.role_id JOIN tblPERSONNEL ON tblPERSONNEL.personnel_id = tblASSIGNED_PERSONNEL_ROLES.missionpersonnel_id WHERE tblPERSONNEL.personnel_id = " + memberID)
			if err != nil {
				fmt.Printf("Error querying MySQL for member: %v", err)
			}
			err = roleRow.Scan(&roleTitle)
			if err != nil {
				fmt.Printf("Error scanning role row: %v", err)
				os.Exit(1)
			}

			// fill Person object with crew member info
			person := &messages.Person{
				ID:       memberID,
				FName:    fName,
				LName:    lName,
				Position: roleTitle,
			}
			people = append(people, person)
		}
	}

	aircraftCallsign, err := getAircraftCallsign(msg.MissionID, db)
	if err != nil {
		fmt.Printf("Error getting aircraft callsign: %v", err)
	}

	aircraftDetail := &messages.AircraftDetail{
		Callsign: aircraftCallsign,
		Crew:     people,
	}
	clientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg, notifier)
}
