package parsers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// TODO: may not work with strongly typed messages
// parse data for delivery
func parse(msg interface{}, pulledMsg *pubsub.Message, msgType string) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	log.Printf("Message contents: %#v", msg)
}

// send message to client
func clientNotify(msg interface{}, msgType string, pulledMsg *pubsub.Message, notifier *handlers.Notifier) {
	// TODO: parse pubsub message into client struct
	toClient := &messages.ClientMsg{
		Type:    msgType,
		Payload: msg,
	}

	// send msg contents to websockets
	send, err := json.Marshal(toClient)
	if err != nil {
		log.Printf("PROBLEM marshaling json: %v", err)
		pulledMsg.Ack()
		return
	}
	notifier.Notify(send)
}

func getAircraftCallsign(ID int, db *sql.DB) (string, error) {
	missionID := strconv.Itoa(ID)
	// get mission from db using missionID
	aircraftRow, err := db.Query("SELECT ac_callsign FROM tblAIRCRAFT JOIN tblMISSION ON tblMISSION.aircraft_id = tblAIRCRAFT.ac_id WHERE mission_id=" + missionID)
	if err != nil {
		fmt.Printf("Error querying MySQL for aircraftID: %v", err)
	}
	var aircraftCallsign string
	err = aircraftRow.Scan(&aircraftCallsign)
	if err != nil {
		return "", err
	}
	return aircraftCallsign, nil
}
