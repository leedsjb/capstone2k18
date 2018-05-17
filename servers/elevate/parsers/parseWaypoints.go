package parsers

import (
	"database/sql"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// ParseWaypointCreate handles the creation of new persistent waypoints
// in the Flight Vector DB
// does not notify client, writes new info to db
func ParseWaypointCreate(msg *messages.Waypoint,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write for sql sproc
	// type Waypoint struct {
	// 	ID					string 	 `json:"ID"`
	// 	Notes				string 	 `json:"notes"`
	// 	Name				string 	 `json:"name"`
	// 	Type				string 	 `json:"type"`
	// 	Address1			string 	 `json:"address1"`
	// 	Address2			string 	 `json:"address2"`
	// 	Country				string 	 `json:"country"`
	// 	State				string 	 `json:"state"`
	// 	County				string 	 `json:"county"`
	// 	City				string 	 `json:"city"`
	// 	Zip					string 	 `json:"zip"`
	// 	Lat					string 	 `json:"lat"`
	// 	Long				string 	 `json:"long"`
	// 	GPSWaypoint			string 	 `json:"GPSWaypoint"`
	// 	AirportIdentifier 	string 	 `json:"AirportIdentifier"`
	// 	Phone				[]string `json:"phone"`
	// 	ShortCode			string	 `json:"shortCode"`
	// 	PadTime				string	 `json:"padTime"`
	// 	RadioChannels		[]string `json:"radioChannels"`
	// 	NOTAMS				string   `json:"NOTAMS"`
	// }

	// [START WAYPOINT]
	/* TODO: replace with real sproc name
	typeRow, err := db.Query("SELECT waypointtype_id FROM tblWAYPOINT_TYPE WHERE waypointtype_name=" + msg.Type)
	var typeID string
	err = typeRow.Scan(&typeID)
	if err != nil {
		fmt.Printf("Error scanning type: %v", err)
		// TODO: insert if new type?
	}
	msg.ID
	msg.Name
	msg.Lat
	msg.Long
	typeID

	_, err := db.Query("CALL insertWaypoint(" + info + ")")
	if err != nil {
		fmt.Printf("Error sproc-ing MySQL: %v", err)
		os.Exit(1)
	}
	*/
	// [END WAYPOINT]

	// [START HOSPITAL]
	/* TODO: replace with real sproc name
	msg.Notes
	msg.PadTime
	msg.RadioChannels

	waypointRow, err := db.Query("SELECT waypoint_id FROM tblWAYPOINT WHERE waypoint_title=" + msg.Name)
	var typeID string
	err = typeRow.Scan(&typeID)
	if err != nil {
		fmt.Printf("Error scanning type: %v", err)
		// TODO: insert if new type?
	}
	*/
	// [END HOSPITAL]

	// [START AIRCRAFT]

	// [END AIRCRAFT]
}

// ParseWaypointUpdate handles the modification of persistent waypoints
// in the Flight Vector DB
// does not notify client, writes new info to db
func ParseWaypointUpdate(msg *messages.Waypoint,
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
	// type Waypoint struct {
	// 	ID					string 	 `json:"ID"`
	// 	Notes				string 	 `json:"notes"`
	// 	Name				string 	 `json:"name"`
	// 	Type				string 	 `json:"type"`
	// 	Address1			string 	 `json:"address1"`
	// 	Address2			string 	 `json:"address2"`
	// 	Country				string 	 `json:"country"`
	// 	State				string 	 `json:"state"`
	// 	County				string 	 `json:"county"`
	// 	City				string 	 `json:"city"`
	// 	Zip					string 	 `json:"zip"`
	// 	Lat					string 	 `json:"lat"`
	// 	Long				string 	 `json:"long"`
	// 	GPSWaypoint			string 	 `json:"GPSWaypoint"`
	// 	AirportIdentifier 	string 	 `json:"AirportIdentifier"`
	// 	Phone				[]string `json:"phone"`
	// 	ShortCode			string	 `json:"shortCode"`
	// 	PadTime				string	 `json:"padTime"`
	// 	RadioChannels		[]string `json:"radioChannels"`
	// 	NOTAMS				string   `json:"NOTAMS"`
	// }
}

// ParseWaypointDelete handles the deletion of persistent waypoints
// from the Flight Vector DB
// does not notify client, writes new info to db
func ParseWaypointDelete(msg *messages.Waypoint_Delete,
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
	// type Waypoint_Delete struct {
	// 	ID		string `json:"ID"`
	// }
}