package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
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

type peopleAndGroups struct {
	People []*messages.Person
	Groups []*messages.ClientGroup
}

// IndexPerson
func IndexPerson(trie *indexes.Trie, person *messages.Person) error {
	personID, err := strconv.Atoi(person.ID)
	if err != nil {
		fmt.Printf("Error changing person ID from string to int")
	}
	personName := person.FName + " " + person.LName
	if err := trie.AddEntity(strings.ToLower(personName), personID); err != nil {
		return fmt.Errorf("Error adding person name to trie: %v", err)
	}

	if err := trie.AddEntity(strings.ToLower(person.Position), personID); err != nil {
		return fmt.Errorf("Errod adding person role to trie: %v", err)
	}
	return nil
}

// LoadPeopleTrie
// *** Pass in the same trie as for handlers/groups.go LoadGroupsTrie
// to allow both to be mutually searchable ***
func (ctx *HandlerContext) LoadPeopleTrie(trie *indexes.Trie) error {
	peopleRows, err := ctx.GetAllPeople()
	if err != nil {
		return fmt.Errorf("Error querying MySQL for people: %v", err)
	}

	personRow := &personRow{}
	for peopleRows.Next() {
		err = peopleRows.Scan(
			&personRow.PersonID,
			&personRow.FName,
			&personRow.LName,
			&personRow.RoleTitle,
		)
		if err != nil {
			return fmt.Errorf("Error scanning person row: %v", err)
		}
		// TODO: maybe optimize to actually check if these already exist
		currentPerson := &messages.Person{
			ID:       personRow.PersonID,
			FName:    personRow.FName,
			LName:    personRow.LName,
			Position: personRow.RoleTitle,
		}

		if err := IndexPerson(trie, currentPerson); err != nil {
			return fmt.Errorf("Error loading people trie: %v", err)
		}
	}
	return nil
}

// GetTriePeople retrieves information on people who match
// the search term based on IDs found in the personnel trie
func (ctx *HandlerContext) GetTriePeople(peopleIDS []int) ([]*messages.Person, error) {
	people := []*messages.Person{}

	// get each group whose prefix matches the search term
	for _, personID := range peopleIDS {
		person := &messages.Person{}
		ID := strconv.Itoa(personID)
		fmt.Printf("Query ID: %v\n", ID)
		peopleRows, err := ctx.GetPersonByID(ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving people from the DB: %v", err)
		}
		fmt.Println("got people rows")
		personRow := &personRow{}
		for peopleRows.Next() {
			err = peopleRows.Scan(
				&personRow.PersonID,
				&personRow.FName,
				&personRow.LName,
				&personRow.RoleTitle,
			)
			fmt.Printf("personRow.PersonID: %v\n", personRow.PersonID)
			fmt.Printf("personRow.FName: %v\n", personRow.FName)
			fmt.Printf("personRow.LName: %v\n", personRow.LName)
			fmt.Printf("personRow.RoleTitle: %v\n", personRow.RoleTitle)
			if err != nil {
				return nil, fmt.Errorf("Error scanning person row: %v", err)
			}
			// TODO: maybe optimize to actually check if these already exist
			person = &messages.Person{
				ID:       personRow.PersonID,
				FName:    personRow.FName,
				LName:    personRow.LName,
				Position: personRow.RoleTitle,
			}
			fmt.Printf("person.ID: %v\n", person.ID)
			fmt.Printf("person.FName: %v\n", person.FName)
			fmt.Printf("person.LName: %v\n", person.LName)
			fmt.Printf("person.Position: %v\n", person.Position)
		}
		// TODO: append outside the people scan loop to ensure you only get
		// one person per ID?
		// Or some sort of logic to append multiple positions to a person?
		people = append(people, person)
	}
	return people, nil
}

// PeopleHandler ...
func (ctx *HandlerContext) PeopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()

		term := query.Get("q")

		people := []*messages.Person{}

		if len(term) > 0 {
			// search term non-empty, filter which people are returned
			peopleIDS := ctx.PeopleTrie.GetEntities(strings.ToLower(term), 20)
			// retrieve the actual people information
			people, err := ctx.GetTriePeople(peopleIDS)
			if err != nil {
				fmt.Printf("Error pulling people from trie: %v", err)
				return
			}
			respond(w, people)
		} else {
			// search term empty, return all people
			peopleRows, err := ctx.GetAllPeople()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error retrieving people from DB: %v", err), http.StatusInternalServerError)
			}

			personRow := &personRow{}
			for peopleRows.Next() {
				err = peopleRows.Scan(
					&personRow.PersonID,
					&personRow.FName,
					&personRow.LName,
					&personRow.RoleTitle,
				)
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
		}
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
		fmt.Printf("path.Base of URL is: %v", id)
		if id != "." && id != "people" {
			personDetail := &messages.PersonDetail{}

			personDetailRows, err := ctx.GetPersonDetailByID(id)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error getting person details from DB: %v", err), http.StatusInternalServerError)
				return
			}

			personDetailRow := &personDetailRow{}
			for personDetailRows.Next() {
				err = personDetailRows.Scan(
					&personDetailRow.PersonnelID,
					&personDetailRow.FName,
					&personDetailRow.LName,
					&personDetailRow.PersonnelTitle,
					&personDetailRow.Email,
					&personDetailRow.SMS,
				)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error scanning person details: %v", err), http.StatusInternalServerError)
					return
				}
				personDetail = &messages.PersonDetail{
					ID:       personDetailRow.PersonnelID,
					FName:    personDetailRow.FName,
					LName:    personDetailRow.LName,
					Position: personDetailRow.PersonnelTitle,
					Email:    personDetailRow.Email,
					Mobile:   personDetailRow.SMS,
				}

			}
			respond(w, personDetail)
		} else if id == "people" {
			ctx.PeopleHandler(w, r)
		} else {
			http.Error(w, "No person with that ID", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
