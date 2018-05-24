package parsers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// ParseMissionCreate handles when a mission is assigned to
// an aircraft
// notifies client and writes new info to db
// assumes Mission_Create topic comes with all information
func (ctx *ParserContext) ParseMissionCreate(msg *messages.Mission_Create,
	pulledMsg *pubsub.Message, msgType string) error {
	// unmarshal json into correct struct
	log.Printf("[MISSION CREATE] before unmarshaling: %v", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Mission-Create: %v", err)
	}

	// parse pubsub message for client

	requestor := ""
	receiver := ""
	aircraftStatus := "on a mission" // assume aircraft assigned to mission is on that mission

	if msg.RequestorID != 0 {
		req, err := ctx.GetRequestorByID(msg.RequestorID)
		if err != nil {
			return fmt.Errorf("Could not retrieve requestor by given ID: %v", msg.RequestorID)
		}
		requestor = req
	}
	if msg.ReceiverID != 0 {
		rec, err := ctx.GetReceiverByID(msg.ReceiverID)
		if err != nil {
			return fmt.Errorf("Could not retrieve requestor by given ID: %v", msg.ReceiverID)
		}
		receiver = rec
	}

	// separate crewIDs to build crew members into related groups
	people := []*messages.Person{}

	if len(msg.CrewMemberID) > 0 {
		for _, memberID := range msg.CrewMemberID {
			person, err := ctx.getPersonSummary(memberID)
			if err != nil {
				return fmt.Errorf("Couldn't retrieve person summary: %v", err)
			}
			people = append(people, person)
		}
	}

	var waypoints []*messages.ClientMissionWaypoint
	nextWaypointETA := ""
	if len(msg.Waypoints) > 0 {
		for _, waypoint := range msg.Waypoints {
			waypointName, err := ctx.GetWaypointNameByID(waypoint.ID)
			if err != nil {
				fmt.Printf("Couldn't get waypoint name with given ID: %v, %v", waypoint.ID, err)
			}
			tempWayPt := &messages.ClientMissionWaypoint{
				Name:        waypointName,
				ETA:         waypoint.ETA,
				Active:      waypoint.Active,
				FlightRules: waypoint.FlightRules,
				// Completed:   "false",
			}
			if strings.ToLower(tempWayPt.Active) == "true" {
				nextWaypointETA = tempWayPt.ETA
			}
			// TODO: calculate ETE/ETT
			waypoints = append(waypoints, tempWayPt)
		}
	}

	mission := &messages.Mission{
		Type: msg.CallType,
		// Vision:          msg.Vision,
		NextWaypointETA: nextWaypointETA,
		Waypoints:       waypoints,
		FlightNum:       msg.TCNum,
	}

	aircraft := &messages.Aircraft{
		Status:   aircraftStatus,
		Callsign: msg.Asset,
		Mission:  mission,
	}

	missionDetail := &messages.MissionDetail{
		Type: msg.CallType,
		// Vision:          msg.Vision,
		NextWaypointETA: nextWaypointETA,
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

	clientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg, ctx.Notifier)
	clientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg, ctx.Notifier)

	// [ADD MISSION TO DB]
	// if err := ctx.AddNewMission(msg); err != nil {
	// 	return fmt.Errorf("Error adding new mission to DB: %v", err)
	// }

	return nil
}

func (ctx *ParserContext) ParseMissionComplete(msg *messages.Mission_Complete,
	pulledMsg *pubsub.Message, msgType string) error {
	fmt.Printf("[MISSION COMPLETE]\n")
	return nil
}

// ParseMissionWaypointsUpdate handles changes to a mission's
// waypoints, including ETA, ETT, friendly names,
// and modifications to the route
// notifies client and writes new info to db
func (ctx *ParserContext) ParseMissionWaypointsUpdate(msg *messages.Mission_Waypoint_Update,
	pulledMsg *pubsub.Message, msgType string) error {
	// unmarshal json into correct struct
	log.Printf("[MISSION WAYPOINT UPDATE] before unmarshaling: %v", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Mission-Waypoints-Update: %v", err)
	}

	// parse pubsub message for client

	waypoints := []*messages.ClientMissionWaypoint{}
	nextWaypointETA := ""
	aircraftStatus := "available" // assume if waypoints are updated and none are active, mission complete

	if len(msg.Waypoints) > 0 {
		for _, waypoint := range msg.Waypoints {
			waypointName, err := ctx.GetWaypointNameByID(waypoint.ID)
			if err != nil {
				return fmt.Errorf("Couldn't get waypoint name with given ID: %v, %v", waypoint.ID, err)
			}
			tempWayPt := &messages.ClientMissionWaypoint{
				Name: waypointName,
				ETA:  waypoint.ETA,
				// ETT:         waypoint.ETT,
				Active:      waypoint.Active,
				FlightRules: waypoint.FlightRules,
			}
			if strings.ToLower(tempWayPt.Active) == "true" {
				nextWaypointETA = tempWayPt.ETA
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
	aircraftCallsign, err := ctx.GetAircraftCallsign(msg.MissionID)
	if err != nil {
		return fmt.Errorf("Error getting aircraftCallsign: %v", err)
	}

	mission := &messages.Mission{
		NextWaypointETA: nextWaypointETA,
		Waypoints:       waypoints,
	}

	aircraft := &messages.Aircraft{
		Status:   aircraftStatus, // assume aircraft assigned to mission is on that mission
		Callsign: aircraftCallsign,
		Mission:  mission,
	}
	// [END format aircraft]

	// [START format aircraftDetail]
	tcNum, err := ctx.GetTCNumByMissionID(msg.MissionID)
	if err != nil {
		return fmt.Errorf("Couldnt get tcNum with given mission ID: %v, %v", msg.MissionID, err)
	}

	missionDetail := &messages.MissionDetail{
		NextWaypointETA: nextWaypointETA,
		Waypoints:       waypoints,
		FlightNum:       tcNum,
	}

	aircraftDetail := &messages.AircraftDetail{
		Status:   aircraftStatus,
		Callsign: aircraftCallsign,
		Mission:  missionDetail,
	}
	// [END format aircraftDetail]

	clientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg, ctx.Notifier)
	clientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg, ctx.Notifier)

	// [ADD WAYPOINT UPDATE TO DB]
	// if err := ctx.UpdateMissionWaypoints(msg); err != nil {
	// 	return fmt.Errorf("Error adding mission waypoint updates to DB: %v", err)
	// }

	return nil
}

// ParseMissionCrewUpdate handles when an aircraft has crew adjusted
// with respect to an assigned mission
// notifies client, writes new info to db
func (ctx *ParserContext) ParseMissionCrewUpdate(msg *messages.Mission_Crew_Update,
	pulledMsg *pubsub.Message, msgType string) error {
	// unmarshal json into correct struct
	log.Printf("[MISSION CREW UPDATE] before unmarshaling: %v", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Mission-Crew-Update: %v", err)
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
			person, err := ctx.getPersonSummary(memberID)
			if err != nil {
				return fmt.Errorf("Couldn't retrieve person summary: %v", err)
			}
			people = append(people, person)
		}
	}

	aircraftCallsign, err := ctx.GetAircraftCallsign(msg.MissionID)
	if err != nil {
		fmt.Printf("Error getting aircraft callsign: %v", err)
	}

	aircraftDetail := &messages.AircraftDetail{
		Callsign: aircraftCallsign,
		Crew:     people,
	}
	clientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg, ctx.Notifier)

	// [ADD CREW UPDATES TO DB]
	// if err := ctx.UpdateMissionCrew(msg); err != nil {
	// 	return fmt.Errorf("Error adding mission crew updates to DB: %v", err)
	// }

	return nil
}

func (ctx *ParserContext) getPersonSummary(memberID int) (*messages.Person, error) {
	// retrieve member first and last name
	fName, lName, err := ctx.GetCrewMemberByID(memberID)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve crew member by given ID: %v", memberID)
	}

	// retrieve member role
	roleTitle, err := ctx.GetRoleByMemberID(memberID)
	if err != nil {
		return nil, fmt.Errorf("Could not retrieve roleTitle with given ID: %v, %v", memberID, err)
	}

	// fill Person object with crew member info
	person := &messages.Person{
		ID:       memberID,
		FName:    fName,
		LName:    lName,
		Position: roleTitle,
	}
	return person, nil
}
