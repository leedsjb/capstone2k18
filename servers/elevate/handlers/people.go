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
		peopleRows, err := ctx.GetPeople()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error querying MySQL for groups: %v", err), http.StatusInternalServerError)
		}

		people := []*messages.Person{}
		personRow := &personRow{}
		for peopleRows.Next() {
			err = peopleRows.Scan(personRow)
			if err != nil {
				fmt.Printf("Error scanning person row: %v", err)
				os.Exit(1)
			}
			// TODO: maybe optimize to actually check if these already exist
			currentPerson := &messages.Person{
				ID:       personRow.PersonID,
				FName:    personRow.FName,
				LName:    personRow.LName,
				Position: personRow.RoleTitle,
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
