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

// ParseGroupCreate handles the creation of a new group in
// the Flight Vector DB
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseGroupCreate(msg *messages.Group,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) error {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Group-Create: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Group struct {
	// 	ID      string   `json:"ID"`
	// 	Name    string   `json:"Name"`
	// 	Members []string `json:"members"`
	// }

	// [ADD GROUP TO DB]
	// if err := ctx.AddNewGroup(stuff); err != nil {
	// 	return fmt.Errorf("Error adding new group to DB: %v", err)
	// }
	return nil
}

// ParseGroupUpdate handles modifications to a group in
// the Flight Vector DB
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseGroupUpdate(msg *messages.Group,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) error {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Group-Update: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Group struct {
	// 	ID      string   `json:"ID"`
	// 	Name    string   `json:"Name"`
	// 	Members []string `json:"members"`
	// }

	// [ADD GROUP UPDATE TO DB]
	// if err := ctx.UpdateGroup(stuff); err != nil {
	// 	return fmt.Errorf("Error adding group updates to DB: %v", err)
	// }
	return nil
}

// ParseGroupDelete handles the deletion of a group from
// the Flight Vector DB
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseGroupDelete(msg *messages.Group_Delete,
	pulledMsg *pubsub.Message, msgType string, db *sql.DB, notifier *handlers.Notifier) error {
	log.Printf("before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in Group-Delete: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type Group_Delete struct {
	// 	ID string `json:"ID"`
	// }

	// [DELETE GROUP FROM DB]
	// if err := ctx.DeleteGroup(stuff); err != nil {
	// 	return fmt.Errorf("Error deleting group from DB: %v", err)
	// }
	return nil
}
