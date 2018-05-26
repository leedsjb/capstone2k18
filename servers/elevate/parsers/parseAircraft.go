package parsers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// For each action, consider:
// Send to client?
// Write to db?
// Write to trie?

func (ctx *ParserContext) ParseAircraftCreate(msg *messages.Aircraft_Create,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("[AIRCRAFT CREATE] before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Aircraft-Create: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc

	// type Aircraft_Create struct {
	// 	ID                int   `json:"ID"`
	// 	NNum              string   `json:"nNum"`
	// 	SatPhone          string   `json:"satPhone"`
	// 	CellPhone         string   `json:"cellPhone"`
	// 	Base              string   `json:"baseID"`
	// 	Callsign          string   `json:"callsign"`
	// 	MaxPatientWeight  string   `json:"maxPatientWeight"`
	// 	PadTimeDay        string   `json:"padTimeDay"`
	// 	PadTimeNight      string   `json:"padTimeNight"`
	// 	Vendor            string   `json:"vendor"`
	// 	Status            string   `json:"status"`
	// 	SpecialEquipment  string   `json:"specialEquipment"`
	// 	Color             string   `json:"color"`
	// 	LastKnownLocation string   `json:"lastKnownLocation"`
	// 	Model             string   `json:"model"`
	// 	CallTypes         []string `json:"callTypes"`
	// }

	// [ADD AIRCRAFT TO DB]
	// if err := ctx.AddNewAircraft(msg); err != nil {
	// 	return fmt.Errorf("Error adding new aircraft to DB: %v", err)
	// }

	// [ADD AIRCRAFT TO TRIE]
	// aircraftID, err := strconv.Atoi(msg.ID)
	// if err != nil {
	// 	return fmt.Errorf("Error converting aircraft ID to int: %v", err)
	// }

	// if err = ctx.AircraftTrie.AddEntity(strings.ToLower(msg.Callsign), aircraftID); err != nil {
	// 	return fmt.Errorf("Error adding aircraft to trie: %v", err)
	// }

	// if err = ctx.AircraftTrie.AddEntity(strings.ToLower(msg.NNum), aircraftID); err != nil {
	// 	return fmt.Errorf("Error adding aircraft to trie: %v", err)
	// }

	return nil
}

func (ctx *ParserContext) ParseAircraftPropsUpdate(msg *messages.Aircraft_Props_Update,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Aircraft-Props-Update: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Aircraft_Props_Update struct {
	// 	ID               string `json:"ID"`
	// 	SatPhone         string `json:"satPhone"`
	// 	CellPhone        string `json:"cellPhone"`
	// 	Base             string `json:"base"`
	// 	Callsign         string `json:"callsign"`
	// 	MaxPatientWeight string `json:"maxPatientWeight"`
	// 	PadTimeDay       string `json:"padTimeDay"`
	// 	PadTimeNight     string `json:"padTimeNight"`
	// 	Vendor           string `json:"vendor"`
	// 	SpecialEquipment string `json:"specialEquipment"`
	// }

	// [ADD PROPS UPDATE TO DB]
	// if err := ctx.UpdateAircraftProps(msg); err != nil {
	// 	return fmt.Errorf("Error adding aircraft updates to DB: %v", err)
	// }

	return nil
}

// PENDING: Brian add "id" to message
// ParseAircraftCrewUpdate handles the standard assignment of crews
// to aircraft as shifts change
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseAircraftCrewUpdate(msg *messages.Aircraft_Crew_Update,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("[AIRCRAFT CREW UPDATE] before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Aircraft-Crew-Update: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Aircraft_Crew_Update struct {
	// 	ID          string `json:"ID"`
	// 	PIC         string `json:"PIC"`
	// 	AdultRN     string `json:"adultRN"`
	// 	PediatricRN string `json:"pediatricRN"`
	// }

	// [ADD CREW UPDATE TO DB]
	// if err := ctx.UpdateAircraftCrew(msg); err != nil {
	// 	return fmt.Errorf("Error adding aircraft crew updates to DB: %v", err)
	// }

	return nil
}

// PENDING: Brian add "id" to message
// ParseAircraftServiceSchedule handles when aircraft are scheduled
// to be serviced for maintenance or otherwise
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseAircraftServiceSchedule(msg *messages.Aircraft_Service_Schedule,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Aircraft-Service-Schedule: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Aircraft_Service_Schedule struct {
	// 	ID        string `json:"ID"`
	// 	OosReason string `json:"oosReason"`
	// 	StartTime string `json:"startTime"`
	// 	EndTime   string `json:"endTime"`
	// 	Status    string `json:"status"`
	// }

	// TODO: Check Aircraft_Service_Schedule.Status -  is it a status we want?

	// [ADD SERVICE SCHEDULE UPDATE TO DB]
	// if err := ctx.UpdateAircraftServiceSchedule(msg); err != nil {
	// 	return fmt.Errorf("Error adding aircraft service schedule updates to DB: %v", err)
	// }

	return nil
}

// PENDING: Brian add "id" to message
// ParseAircraftPositionUpdate handles anytime the aircraft moves,
// and is highly relevant to missions
// notifies client, writes new info to db
func (ctx *ParserContext) ParseAircraftPositionUpdate(msg *messages.Aircraft_Pos_Update,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("[AIRCRAFT POSITION UPDATE] before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Aircraft-Position-Update: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Aircraft_Pos_Update struct {
	// 	ID              string `json:"ID"`
	// 	PosLat          string `json:"posLat"`
	// 	PosLong         string `json:"posLong"`
	// 	PosFriendlyName string `json:"posFriendlyName"`
	// }

	msgID, err := strconv.Atoi(msg.ID)
	if err != nil {
		return fmt.Errorf("Could not convert message ID from string to int: %v", err)
	}

	aircraftCallsign, err := ctx.GetAircraftCallsign(msgID)
	if err != nil {
		return fmt.Errorf("Error getting aircraft callsign: %v", err)
	}

	aircraft := &messages.Aircraft{
		Callsign: aircraftCallsign,
		Lat:      msg.PosLat,
		Long:     msg.PosLong,
		Area:     msg.PosFriendlyName,
	}

	aircraftDetail := &messages.AircraftDetail{
		Callsign: aircraftCallsign,
		Lat:      msg.PosLat,
		Long:     msg.PosLong,
		Area:     msg.PosFriendlyName,
	}

	ctx.ClientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg)
	ctx.ClientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg)

	// ADD POSITION UPDATE TO DB
	// if err := ctx.UpdateAircraftPosition(msg); err != nil {
	// 	return fmt.Errorf("Error adding aircraft position updates to DB: %v", err)
	// }

	return nil
}
