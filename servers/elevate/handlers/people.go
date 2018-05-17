package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

type personRow struct {
	PersonID  string
	FName     string
	LName     string
	RoleTitle string
}

type personDetailRow struct {
	PersonnelID    string
	FName          string
	LName          string
	PersonnelTitle string
	Email          string
	// Will we still have UWNetID in the DB if we're using UW Groups?
	// UWNetID        string
	SMS string
	// Infer?
	// SpecialQuals   string
}

// PeopleHandler ...
func (ctx *HandlerContext) PeopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// TODO: replace with SPROC
		/*
			SELECT personnel_id, personnel_F_Name, personnel_L_Name, role_title
		*/
		peopleRows, err := ctx.DB.Query("SELECT group_id, group_name, personnel_F_Name, personnel_L_Name FROM tblPERSONNEL_GROUP JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id ORDER BY group_name")

		if err != nil {
			fmt.Printf("Error querying MySQL for groups: %v", err)
		}

		people := []*messages.Person{}

		currentRow := &personRow{}

		for peopleRows.Next() {
			err = peopleRows.Scan(currentRow)
			if err != nil {
				fmt.Printf("Error scanning person row: %v", err)
				os.Exit(1)
			}
			// TODO: maybe optimize to actually check if these already exist
			currentPerson := &messages.Person{
				ID:       currentRow.PersonID,
				FName:    currentRow.FName,
				LName:    currentRow.LName,
				Position: currentRow.RoleTitle,
			}

			people = append(people, currentPerson)
		}
		respond(w, people)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// PersonDetailHandler ...
func (ctx *HandlerContext) PersonDetailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := path.Base(r.URL.Path)
		if id != "." {
			personDetail := &messages.PersonDetail{}

			personDetailRows, err := ctx.GetPersonDetails(id)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error getting person details from DB: %v", err), http.StatusInternalServerError)
				return
			}

			personDetailRow := &personDetailRow{}
			for personDetailRows.Next() {
				err = personDetailRows.Scan(personDetailRow)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error scanning person details from query: %v", err), http.StatusInternalServerError)
					return
				}
				personDetail = &messages.PersonDetail{
					ID:       personDetailRow.PersonnelID,
					FName:    personDetailRow.FName,
					LName:    personDetailRow.LName,
					Position: personDetailRow.PersonnelTitle,
					Email:    personDetailRow.Email,
					// UWNetID:      personDetailRow.UWNetID,
					Mobile: personDetailRow.SMS,
					// SpecialQuals: personDetailRow.SpecialQuals,
				}

			}
			respond(w, personDetail)
		} else {
			http.Error(w, "No aircraft with that ID", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
