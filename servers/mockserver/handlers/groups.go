package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
)

// Group ...
type Group struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PeoplePreview string `json:"peoplePreview"`
}

// GroupDetail ...
type GroupDetail struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	PeoplePreview string    `json:"peoplePreview"`
	People        []*Person `json:"people"`
}

var groupDetails = []*GroupDetail{
	{
		ID:            1,
		Name:          "AL2",
		PeoplePreview: fmt.Sprintf("%v and %v", personDetails[0].FName, personDetails[1].FName),
		People: []*Person{
			{
				ID:       personDetails[0].ID,
				FName:    personDetails[0].FName,
				LName:    personDetails[0].LName,
				Position: personDetails[0].Position,
			},
			{
				ID:       personDetails[1].ID,
				FName:    personDetails[1].FName,
				LName:    personDetails[1].LName,
				Position: personDetails[1].Position,
			},
		},
	},
	{
		ID:            2,
		Name:          "AL3",
		PeoplePreview: fmt.Sprintf("%v and %v", personDetails[2].FName, personDetails[3].FName),
		People: []*Person{
			{
				ID:       personDetails[2].ID,
				FName:    personDetails[2].FName,
				LName:    personDetails[2].LName,
				Position: personDetails[2].Position,
			},
			{
				ID:       personDetails[3].ID,
				FName:    personDetails[3].FName,
				LName:    personDetails[3].LName,
				Position: personDetails[3].Position,
			},
		},
	},
}

// GroupsHandler ...
func GroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		groups := []*Group{}
		for _, v := range groupDetails {
			g := &Group{
				ID:            v.ID,
				Name:          v.Name,
				PeoplePreview: v.PeoplePreview,
			}
			groups = append(groups, g)
		}
		respond(w, groups)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// GroupDetailHandler ...
func GroupDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding ID: %v", err), http.StatusBadRequest)
		return
	}
	var gd *GroupDetail
	for _, v := range groupDetails {
		if v.ID == id {
			gd = v
			break
		}
	}
	if gd == nil {
		http.Error(w, "No group with that ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		respond(w, gd)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
