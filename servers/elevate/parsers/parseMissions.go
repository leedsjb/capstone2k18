package parsers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// const layout = "2006-01-02 15:04:05.0000000 -07:00"
const layout = "2006-01-02 15:04:05"

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

	requestor := &messages.Agency{}
	receiver := &messages.Agency{}

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
	aircraftStatus := "OAM" // assume aircraft assigned to mission is on that mission

	if requestorID != 0 {
		req, err := ctx.GetAgencyDetailsByID(requestorID)
		if err != nil {
			return fmt.Errorf("Could not retrieve requestor by given ID: %v, error is: %v\n", requestorID, err)
		}
		requestor = req
	}
	if receiverID != 0 {
		rec, err := ctx.GetAgencyDetailsByID(receiverID)
		if err != nil {
			return fmt.Errorf("Could not retrieve receiver by given ID: %v, error is: %v\n", receiverID, err)
		}
		receiver = rec
	}

	// [CREW]
	// use given people IDs to build person objects for clientside
	people := []*messages.Person{}

	intCrewMemberIDs, err := convertToInts(msg.CrewMembers)
	if err != nil {
		return fmt.Errorf("Could not convert crew member IDs to int: %v", err)
	}
	if len(intCrewMemberIDs) > 0 {
		for _, memberID := range intCrewMemberIDs {
			person, err := ctx.GetPersonByID(memberID)
			if err != nil {
				return fmt.Errorf("Could not retrieve person summary: %v", err)
			}
			people = append(people, person)
		}
	}

	// [WAYPOINTS]
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
				FlightRules: waypoint.FlightRules,
				// Completed:   "false",
			}
			if waypoint.Active == "1" {
				tempWayPt.Active = true
				ETA, err := time.Parse(layout, tempWayPt.ETA)
				if err != nil {
					return fmt.Errorf("Couldn't parse string ETA to time: %v", err)
				}
				nextWaypointETE = time.Until(ETA).String()
			} else {
				tempWayPt.Active = false
			}
			waypoints = append(waypoints, tempWayPt)
		}
	}

	aircraftCallsign, err := ctx.GetAircraftCallsignByID(msg.Asset)
	if err != nil {
		return fmt.Errorf("Couldn't get aircraft callsign: %v", err)
	}

	mission := &messages.Mission{
		Type:            msg.CallType,
		NextWaypointETE: nextWaypointETE,
		Waypoints:       waypoints,
		FlightNum:       msg.TCNum,
	}

	aircraft := &messages.Aircraft{
		Status:   aircraftStatus,
		Callsign: aircraftCallsign,
		Mission:  mission,
	}

	report := &messages.ClientPatient{}

	if msg.Patient.ShortReport != "" {
		report.ShortReport = msg.Patient.ShortReport
	}
	if msg.Patient.Intubated == "1" {
		report.Intubated = true
	} else {
		report.Intubated = false
	}
	if msg.Patient.Drips != "" {
		report.Drips, err = strconv.Atoi(msg.Patient.Drips)
		if err != nil {
			return fmt.Errorf("Couldn't convert patient drips to int: %v", err)
		}
	}
	if msg.Patient.Age != "" {
		report.Age, err = strconv.Atoi(msg.Patient.Age)
		if err != nil {
			return fmt.Errorf("Couldn't convert patient age to int: %v", err)
		}
	}
	if msg.Patient.Weight != "" {
		report.Weight, err = strconv.Atoi(msg.Patient.Weight)
		if err != nil {
			return fmt.Errorf("Couldn't convert patient weight to int: %v", err)
		}
	}
	if msg.Patient.Gender != "" {
		report.Gender = msg.Patient.Gender
	}
	if msg.Patient.Cardiac == "1" {
		report.Cardiac = true
	} else {
		report.Cardiac = false
	}
	if msg.Patient.GIBleed == "1" {
		report.GIBleed = true
	} else {
		report.GIBleed = false
	}
	if msg.Patient.OB == "1" {
		report.OB = true
	} else {
		report.OB = false
	}

	missionDetail := &messages.MissionDetail{
		Type:            msg.CallType,
		NextWaypointETE: nextWaypointETE,
		Waypoints:       waypoints,
		FlightNum:       msg.TCNum,
		RadioReport:     report,
		Requestor:       requestor,
		Receiver:        receiver,
	}

	aircraftDetail := &messages.AircraftDetail{
		Status:   aircraftStatus,
		Callsign: aircraftCallsign,
		Crew:     people,
		Mission:  missionDetail,
	}

	log.Printf("[CLIENT NOTIFY] Aircraft: %v", aircraft)
	ctx.ClientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg)
	log.Printf("[CLIENT NOTIFY AircraftDetail: %v", aircraftDetail)
	ctx.ClientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg)

	// Notify crewmembers assigned to mission
	aircraftID, err := strconv.Atoi(msg.Asset)
	if err != nil {
		return fmt.Errorf("Could not convert aircraft ID from string to int: %v", err)
	}

	if len(msg.CrewMembers) > 0 {
		for _, memberID := range intCrewMemberIDs {
			personRows, err := ctx.GetPersonDetailByID(memberID)
			if err != nil {
				return fmt.Errorf("Could not retrieve person details: %v\n", err)
			}
			person := &messages.PersonDetail{}
			var unnecessaryGroupID string
			var unnecessaryGroupName string
			for personRows.Next() {
				err = personRows.Scan(
					&person.ID,
					&person.FName,
					&person.LName,
					&person.Position,
					&person.Mobile,
					&person.Email,
					&unnecessaryGroupID,
					&unnecessaryGroupName,
				)
			}
			mobile := person.Mobile
			if err := MissionNotify(aircraftCallsign, aircraftID, mobile); err != nil {
				return fmt.Errorf("Couldn't send mission notification: %v\n", err)
			}
			close(personRows)
		}
	}

	// [ADD MISSION TO DB]

	// collect info for new mission insert

	if err := ctx.NewMission(msg, aircraftID); err != nil {
		return fmt.Errorf("Error adding new mission to DB: %v", err)
	}

	return nil
}

func (ctx *ParserContext) ParseMissionComplete(msg *messages.Mission_Complete,
	pulledMsg *pubsub.Message, msgType string) error {
	fmt.Printf("[MISSION COMPLETE]\n")

	aircraftID, err := ctx.GetAircraftIDByMission(msg.MissionID)
	if err != nil {
		return fmt.Errorf("Couldn't get aircraft ID given mission ID: %v", err)
	}

	mission := &messages.Mission{
		Completed: "1",
	}

	aircraft := &messages.Aircraft{
		ID:      aircraftID,
		Mission: mission,
	}

	missionDetail := &messages.MissionDetail{
		Completed: "1",
	}

	aircraftDetail := &messages.AircraftDetail{
		ID:      aircraftID,
		Mission: missionDetail,
	}

	ctx.ClientNotify(aircraft, "AIRCRAFT_MISSION_COMPLETE", pulledMsg)
	ctx.ClientNotify(aircraftDetail, "AIRCRAFTDETAIL_MISSION_COMPLETE", pulledMsg)

	if err := ctx.CompleteMission(msg.MissionID); err != nil {
		return fmt.Errorf("Couldn't complete mission: %v", err)
	}
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
	aircraftStatus := "Ready" // assume if waypoints are updated and none are active, mission complete

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
				FlightRules: waypoint.FlightRules,
			}
			if waypoint.Active == "1" {
				tempWayPt.Active = true
				ETA, err := time.Parse(layout, tempWayPt.ETA)
				if err != nil {
					return fmt.Errorf("Couldn't parse string ETA to time: %v", err)
				}
				nextWaypointETE = time.Until(ETA).String()
				aircraftStatus = "On Mission"
			} else {
				tempWayPt.Active = false
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
	aircraftCallsign, err := ctx.GetAircraftCallsignByMission(missionID)
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

	if len(msg.CrewMembers) > 0 {
		intCrewMemberIDs, err := convertToInts(msg.CrewMembers)
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
	aircraftCallsign, err := ctx.GetAircraftCallsignByMission(missionID)
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

func convertToInts(toInt []*messages.Mission_Crew) ([]int, error) {
	converted := []int{}
	for _, member := range toInt {
		tempInt, err := strconv.Atoi(member.ID)
		if err != nil {
			return nil, err
		}
		converted = append(converted, tempInt)
	}
	return converted, nil
}
