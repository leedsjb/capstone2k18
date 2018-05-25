package parsers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// const layout = "2018-05-26 14:00:00"
const layout = "2006-01-02 15:04:05.0000000 -07:00"

// ParseMissionCreate handles when a mission is assigned to
// an aircraft
// notifies client and writes new info to db
// assumes Mission_Create topic comes with all information
func (ctx *ParserContext) ParseMissionCreate(msg *messages.Mission_Create,
	pulledMsg *pubsub.Message, msgType string) error {
	// unmarshal json into correct struct
	log.Printf("[MISSION CREATE] before unmarshaling: %v\n", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Mission-Create: %v", err)
	}

	log.Printf("[MISSION CREATE] after unmarshaling: %+v\n", msg)

	// parse pubsub message for client

	requestor := ""
	receiver := ""

	fmt.Printf("[MISSION CREATE] requestor ID: %v\n", msg.RequestorID)
	fmt.Printf("[MISSION CREATE] receiver ID: %v\n", msg.ReceiverID)
	requestorID, err := strconv.Atoi(msg.RequestorID)
	if err != nil {
		return fmt.Errorf("could not convert string to int: %v", err)
	}
	receiverID, err := strconv.Atoi(msg.ReceiverID)
	if err != nil {
		return fmt.Errorf("could not convert string to int: %v", err)
	}
	aircraftStatus := "on a mission" // assume aircraft assigned to mission is on that mission

	log.Printf("[MISSION CREATE] requestor ID: %v\n", requestorID)
	if requestorID != 0 {
		req, err := ctx.GetRequestorByID(requestorID)
		if err != nil {
			return fmt.Errorf("Could not retrieve requestor by given ID: %v, error is: %v\n", requestorID, err)
		}
		requestor = req
	}
	log.Printf("[MISSION CREATE] requestor ID: %v\n", requestorID)
	if receiverID != 0 {
		rec, err := ctx.GetReceiverByID(receiverID)
		if err != nil {
			return fmt.Errorf("Could not retrieve receiver by given ID: %v, error is: %v\n", receiverID, err)
		}
		receiver = rec
	}

	// separate crewIDs to build crew members into related groups
	people := []*messages.Person{}

	if len(msg.CrewMemberID) > 0 {
		intCrewMemberIDs, err := convertToInts(msg.CrewMemberID)
		if err != nil {
			return fmt.Errorf("Could not convert crew member IDs to int: %v", err)
		}
		for _, memberID := range intCrewMemberIDs {
			person, err := ctx.GetPersonByID(memberID)
			if err != nil {
				return fmt.Errorf("Could not retrieve person summary: %v", err)
			}
			people = append(people, person)
		}
	}

	var waypoints []*messages.ClientMissionWaypoint
	nextWaypointETE := ""
	if len(msg.Waypoints) > 0 {
		for _, waypoint := range msg.Waypoints {
			waypointID, err := strconv.Atoi(waypoint.ID)
			if err != nil {
				return fmt.Errorf("Could not convert waypoint ID from string to int: %v", err)
			}
			waypointName, err := ctx.GetWaypointNameByID(waypointID)
			if err != nil {
				return fmt.Errorf("Couldn't get waypoint name with given ID: %v, %v", waypointID, err)
			}
			tempWayPt := &messages.ClientMissionWaypoint{
				Name:        waypointName,
				ETA:         waypoint.ETA,
				Active:      waypoint.Active,
				FlightRules: waypoint.FlightRules,
				// Completed:   "false",
			}
			if tempWayPt.Active == "1" {
				// nextWaypointETA = tempWayPt.ETA
				ETA, err := time.Parse(layout, tempWayPt.ETA)
				if err != nil {
					return fmt.Errorf("Couldn't parse string ETA to time: %v", err)
				}
				nextWaypointETE = time.Until(ETA).String()
			}
			// TODO: calculate ETE/ETT
			waypoints = append(waypoints, tempWayPt)
		}
	}

	mission := &messages.Mission{
		Type: msg.CallType,
		// Vision:          msg.Vision,
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
		Type: msg.CallType,
		// Vision:          msg.Vision,
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

	log.Printf("[CLIENT NOTIFY] Aircraft: %v", aircraft)
	ctx.ClientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg)
	log.Printf("[CLIENT NOTIFY AircraftDetail: %v", aircraftDetail)
	ctx.ClientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg)

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
	nextWaypointETE := ""
	aircraftStatus := "available" // assume if waypoints are updated and none are active, mission complete

	if len(msg.Waypoints) > 0 {
		for _, waypoint := range msg.Waypoints {
			waypointID, err := strconv.Atoi(waypoint.ID)
			if err != nil {
				return fmt.Errorf("Could not convert waypoint ID from string to int: %v", err)
			}
			waypointName, err := ctx.GetWaypointNameByID(waypointID)
			if err != nil {
				return fmt.Errorf("Couldn't get waypoint name with given ID: %v, %v", waypointID, err)
			}
			tempWayPt := &messages.ClientMissionWaypoint{
				Name: waypointName,
				ETA:  waypoint.ETA,
				// ETT:         waypoint.ETT,
				Active:      waypoint.Active,
				FlightRules: waypoint.FlightRules,
			}
			if strings.ToLower(tempWayPt.Active) == "true" {
				ETA, err := time.Parse(layout, tempWayPt.ETA)
				if err != nil {
					return fmt.Errorf("Couldn't parse string ETA to time: %v", err)
				}
				nextWaypointETE = time.Until(ETA).String()
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
	missionID, err := strconv.Atoi(msg.MissionID)
	if err != nil {
		return fmt.Errorf("Could not convert mission ID from string to int: %v", err)
	}
	aircraftCallsign, err := ctx.GetAircraftCallsign(missionID)
	if err != nil {
		return fmt.Errorf("Error getting aircraftCallsign: %v", err)
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
	tcNum, err := ctx.GetTCNumByMissionID(missionID)
	if err != nil {
		return fmt.Errorf("Couldnt get tcNum with given mission ID: %v, %v", msg.MissionID, err)
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

	ctx.ClientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg)
	ctx.ClientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg)

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
	log.Printf("[MISSION CREW UPDATE] before unmarshaling: %v\n", string(pulledMsg.Data))
	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v\n", msg)
		log.Printf("Could not decode message data: %#v\n", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Mission-Crew-Update: %v\n", err)
	}

	// parse pubsub message for client
	// type Mission_Crew_Update struct {
	// 	MissionID    string   `json:"missionID"`
	// 	CrewMemberID []string `json:"crewMemberID"`
	// }

	// separate crewIDs to build crew members into related groups
	people := []*messages.Person{}

	if len(msg.CrewMemberID) > 0 {
		intCrewMemberIDs, err := convertToInts(msg.CrewMemberID)
		if err != nil {
			return fmt.Errorf("Could not convert crew member IDs to int: %v\n", err)
		}
		for _, memberID := range intCrewMemberIDs {
			person, err := ctx.GetPersonByID(memberID)
			if err != nil {
				return fmt.Errorf("Could not retrieve person summary: %v", err)
			}
			people = append(people, person)
		}
	}

	missionID, err := strconv.Atoi(msg.MissionID)
	if err != nil {
		return fmt.Errorf("Could not convert mission ID from string to int: %v\n", err)
	}
	aircraftCallsign, err := ctx.GetAircraftCallsign(missionID)
	if err != nil {
		fmt.Printf("Error getting aircraft callsign: %v\n", err)
	}

	aircraftDetail := &messages.AircraftDetail{
		Callsign: aircraftCallsign,
		Crew:     people,
	}
	ctx.ClientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg)

	// [ADD CREW UPDATES TO DB]
	// if err := ctx.UpdateMissionCrew(msg); err != nil {
	// 	return fmt.Errorf("Error adding mission crew updates to DB: %v", err)
	// }

	return nil
}

func convertToInts(toInt []string) ([]int, error) {
	converted := []int{}
	for _, str := range toInt {
		tempInt, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		converted = append(converted, tempInt)
	}
	return converted, nil
}
