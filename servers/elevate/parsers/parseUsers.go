package parsers

import (
	"encoding/json"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// ParseUserCreate handles the creation of a new user in
// the Flight Vector DB
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseUserCreate(msg *messages.User,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("[USER CREATE] before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in User-Create: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type User struct {
	// 	ID              int `json:"ID"`
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

	// [ADD USER TO DB]
	// if err := ctx.AddNewUser(msg); err != nil {
	// 	return fmt.Errorf("Error adding new user to DB: %v", err)
	// }

	return nil
}

// ParseUserUpdate handles modifications to a user in
// the Flight Vector DB
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseUserUpdate(msg *messages.User,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("[USER UPDATE] before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in User-Update: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type User struct {
	// 	ID              int `json:"ID"`
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

	// [ADD USER UPDATE TO DB]
	// if err := ctx.UpdateUser(msg); err != nil {
	// 	return fmt.Errorf("Error adding user updates to DB: %v", err)
	// }

	return nil
}

// ParseUserDelete handles the deletion of a user from
// the Flight Vector DB
// does not notify client, writes new info to db
func (ctx *ParserContext) ParseUserDelete(msg *messages.User_Delete,
	pulledMsg *pubsub.Message, msgType string) error {
	log.Printf("[USER DELETE] before unmarshaling: %v", string(pulledMsg.Data))

	if err := json.Unmarshal(pulledMsg.Data, &msg); err != nil {
		log.Printf("PROBLEM contents of decoded json: %#v", msg)
		log.Printf("Could not decode message data: %#v", pulledMsg)
		pulledMsg.Ack()
		return fmt.Errorf("Error unmarshaling message data in User-Delete: %v", err)
	}

	log.Printf("Message contents: %#v", msg)

	// TODO: parse and write to sql sproc
	// type User_Delete struct {
	// 	ID string `json:"ID"`
	// }

	// [DELETE USER FROM DB]
	// if err := ctx.DeleteUser(msg); err != nil {
	// 	return fmt.Errorf("Error deleting user from DB: %v", err)
	// }

	return nil
}
