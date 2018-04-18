package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
)

// Person ...
type Person struct {
	ID       int    `json:"id"`
	FName    string `json:"fName"`
	LName    string `json:"lName"`
	Position string `json:"position"`
}

// PersonDetail ...
type PersonDetail struct {
	ID           int    `json:"id"`
	FName        string `json:"fName"`
	LName        string `json:"lName"`
	Position     string `json:"position"`
	Email        string `json:"email"`
	UWNetID      string `json:"uwNetID"`
	Mobile       string `json:"mobile"`
	SpecialQuals string `json:"specialQuals"`
}

var personDetails = []*PersonDetail{
	{
		ID:           1,
		FName:        "Andrew",
		LName:        "Wiles",
		Position:     "Adult RN",
		Email:        "andrewwiles@uw.edu",
		UWNetID:      "andrewwiles",
		Mobile:       "(811) 328 3218",
		SpecialQuals: "",
	},
	{
		ID:           2,
		FName:        "Austin",
		LName:        "Bailey",
		Position:     "Pilot",
		Email:        "austinbailey@airliftnw.org",
		UWNetID:      "austinbailey",
		Mobile:       "(206) 456 7890",
		SpecialQuals: "",
	},
	{
		ID:           3,
		FName:        "Brenda",
		LName:        "Nelson",
		Position:     "Chief Flight Nurse",
		Email:        "brendanelson@airliftnw.org",
		UWNetID:      "brendanelson",
		Mobile:       "(123) 456 7890",
		SpecialQuals: "",
	},
	{
		ID:           4,
		FName:        "Donald",
		LName:        "Lynch",
		Position:     "Flight Nurse",
		Email:        "donaldlynch@uw.edu",
		UWNetID:      "donaldlynch",
		Mobile:       "(456) 789 0123",
		SpecialQuals: "",
	},
	{
		ID:           5,
		FName:        "Christine",
		LName:        "Engle",
		Position:     "Pilot",
		Email:        "chrisengle@airliftnw.org",
		UWNetID:      "chrisengle",
		Mobile:       "(425) 598 6442",
		SpecialQuals: "",
	},
	{
		ID:           6,
		FName:        "Atif",
		LName:        "Mack",
		Position:     "Flight Nurse",
		Email:        "atifmack@uw.edu",
		UWNetID:      "atifmack",
		Mobile:       "(206) 280 7212",
		SpecialQuals: "",
	},
}

// PeopleHandler ...
func PeopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		people := []*Person{}
		for _, v := range personDetails {
			p := &Person{
				ID:       v.ID,
				FName:    v.FName,
				LName:    v.LName,
				Position: v.Position,
			}
			people = append(people, p)
		}
		respond(w, people)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// PersonDetailHandler ...
func PersonDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding ID: %v", err), http.StatusBadRequest)
		return
	}
	var pd *PersonDetail
	for _, v := range personDetails {
		if v.ID == id {
			pd = v
			break
		}
	}
	if pd == nil {
		http.Error(w, "No person with that ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		respond(w, pd)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
