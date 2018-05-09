package parsers

import (
	"database/sql"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// ParseGroupCreate handles the creation of a new group in
// the Flight Vector DB
// does not notify client, writes new info to db
func ParseGroupCreate(msg *messages.Group,
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
	// type Group struct {
	// 	ID      string   `json:"ID"`
	// 	Name    string   `json:"Name"`
	// 	Members []string `json:"members"`
	// }
}

// ParseGroupUpdate handles modifications to a group in
// the Flight Vector DB
// does not notify client, writes new info to db
func ParseGroupUpdate(msg *messages.Group,
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
	// type Group struct {
	// 	ID      string   `json:"ID"`
	// 	Name    string   `json:"Name"`
	// 	Members []string `json:"members"`
	// }
}

// ParseGroupDelete handles the deletion of a group from
// the Flight Vector DB
// does not notify client, writes new info to db
func ParseGroupDelete(msg *messages.Group_Delete,
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
	// type Group_Delete struct {
	// 	ID string `json:"ID"`
	// }
}
