package parsers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

func ParseAircraftCreate(msg *messages.Aircraft_Create,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Aircraft_Create struct {
	// 	ID                string   `json:"ID"`
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

}

func ParseAircraftPropsUpdate(msg *messages.Aircraft_Props_Update,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
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
}

// PENDING: Brian add "id" to message
// ParseAircraftCrewUpdate handles the standard assignment of crews
// to aircraft as shifts change
// does not notify client, writes new info to db
func ParseAircraftCrewUpdate(msg *messages.Aircraft_Crew_Update,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Aircraft_Crew_Update struct {
	// 	ID          string `json:"ID"`
	// 	PIC         string `json:"PIC"`
	// 	AdultRN     string `json:"adultRN"`
	// 	PediatricRN string `json:"pediatricRN"`
	// }
}

// PENDING: Brian add "id" to message
// ParseAircraftServiceSchedule handles when aircraft are scheduled
// to be serviced for maintenance or otherwise
// does not notify client, writes new info to db
func ParseAircraftServiceSchedule(msg *messages.Aircraft_Service_Schedule,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
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
}

// PENDING: Brian add "id" to message
// ParseAircraftPositionUpdate handles anytime the aircraft moves,
// and is highly relevant to missions
// notifies client, writes new info to db
func ParseAircraftPositionUpdate(msg *messages.Aircraft_Pos_Update,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Aircraft_Pos_Update struct {
	// 	ID              string `json:"ID"`
	// 	PosLat          string `json:"posLat"`
	// 	PosLong         string `json:"posLong"`
	// 	PosFriendlyName string `json:"posFriendlyName"`
	// }

	aircraftRow, err := db.Query("SELECT ac_callsign FROM tblAIRCRAFT WHERE ac_id=" + msg.ID)
	if err != nil {
		fmt.Printf("Error querying MySQL for aircraftID: %v", err)
	}
	var aircraftCallsign string
	err = aircraftRow.Scan(&aircraftCallsign)
	if err != nil {
		fmt.Printf("Error getting aircraft callsign: %v", err)
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

	clientNotify(aircraft, "FETCH_AIRCRAFT_SUCCESS", pulledMsg, notifier)
	clientNotify(aircraftDetail, "FETCH_AIRCRAFTDETAIL_SUCCESS", pulledMsg, notifier)
}
