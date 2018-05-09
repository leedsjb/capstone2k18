package parsers

import (
	"database/sql"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/handlers"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// ParseUserCreate handles the creation of a new user in
// the Flight Vector DB
// does not notify client, writes new info to db
func ParseUserCreate(msg *messages.User,
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
	// type User struct {
	// 	ID              string `json:"ID"`
	// 	UserName        string `json:"userName"`
	// 	FirstName       string `json:"firstName"`
	// 	MiddleName      string `json:"middleName"`
	// 	LastName        string `json:"lastName"`
	// 	Initials        string `json:"initials"`
	// 	Email           string `json:"email"`
	// 	UWNetID         string `json:"UWNetID"`
	// 	GroupID         string `json:"groupID"`
	// 	Role            string `json:"role"`
	// 	CellPhone       string `json:"cellPhone"`
	// 	QualificationID string `json:"qualificationID"`
	// }
}

// ParseUserUpdate handles modifications to a user in
// the Flight Vector DB
// does not notify client, writes new info to db
func ParseUserUpdate(msg *messages.User,
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
	// type User struct {
	// 	ID              string `json:"ID"`
	// 	UserName        string `json:"userName"`
	// 	FirstName       string `json:"firstName"`
	// 	MiddleName      string `json:"middleName"`
	// 	LastName        string `json:"lastName"`
	// 	Initials        string `json:"initials"`
	// 	Email           string `json:"email"`
	// 	UWNetID         string `json:"UWNetID"`
	// 	GroupID         string `json:"groupID"`
	// 	Role            string `json:"role"`
	// 	CellPhone       string `json:"cellPhone"`
	// 	QualificationID string `json:"qualificationID"`
	// }
}

// ParseUserDelete handles the deletion of a user from
// the Flight Vector DB
// does not notify client, writes new info to db
func ParseUserDelete(msg *messages.User_Delete,
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
	// type User_Delete struct {
	// 	ID string `json:"ID"`
	// }
}
