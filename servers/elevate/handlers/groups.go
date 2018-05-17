package handlers

import (
	"fmt"
	"net/http"
	"path"

	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

type groupRow struct {
	GroupID   string
	GroupName string
	FName     string
	LName     string
}

type groupDetailRow struct {
	GroupID     string
	GroupName   string
	FName       string
	LName       string
	PersonnelID string
	RoleTitle   string
}

// GroupsHandler ...
func (ctx *HandlerContext) GroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		groups := []*messages.ClientGroup{}

		// type GroupDetail struct {
		// 	ID            string    `json:"id"`
		// 	Name          string    `json:"name"`
		// 	PeoplePreview string    `json:"peoplePreview"`
		// }

		// SELECT group_id, group_name, personnel_F_Name, personnel_L_Name FROM tblPERSONNEL_GROUP
		// JOIN tblPERSONNEL ON tblPERSONNEL_GROUP.personnel_id = tblPERSONNEL.personnel_id
		// JOIN tblGROUP ON tblPERSONNEL_GROUP.group_id = tblGROUP.group_id
		// ORDER BY group_name

		groupRows, err := ctx.GetGroups()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error querying MySQL for groups: %v", err), http.StatusInternalServerError)
			return
		}
		// create variables and fill contents from retrieved rows
		currentRow := &groupRow{}
		currentGroupID := "first"
		currentGroup := &messages.ClientGroup{}
		currentName := ""
		for groupRows.Next() {
			err = groupRows.Scan(currentRow)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error scanning group row: %v", err), http.StatusInternalServerError)
				return
			}
			if currentGroupID != "first" || currentRow.GroupID != currentGroupID {
				groups = append(groups, currentGroup)
			}
			// TODO: maybe optimize to actually check if these already exist
			currentGroup.ID = currentRow.GroupID
			currentGroup.Name = currentRow.GroupName
			currentName = currentRow.FName + currentRow.LName
			currentGroup.PeoplePreview = append(currentGroup.PeoplePreview, currentName)
		}

		respond(w, groups)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}

// GroupDetailHandler ...
func (ctx *HandlerContext) GroupDetailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// TODO: make sure this isn't potentially exposing anything
		id := path.Base(r.URL.Path)

		// type GroupDetail struct {
		// 	ID            string    `json:"id"`
		// 	Name          string    `json:"name"`
		// 	PeoplePreview string    `json:"peoplePreview"`
		// 	People        []*Person `json:"people"`
		// }

		// type Person struct {
		// 	ID       string `json:"id"`
		// 	FName    string `json:"fName"`
		// 	LName    string `json:"lName"`
		// 	Position string `json:"position"`
		// }

		/*
			SELECT group_id, group_name, personnel_F_Name, personnel_L_Name, personnel_id, role_title
			FROM tblPERSONNEL
			JOIN tblASSIGNED_PERSONNEL_ROLES ON tblPERSONNEL.personnel_id = tblASSIGNED_PERSONNEL_ROLES.missionpersonnel_id
			JOIN tblROLES ON tblPERSONNEL.personnel_id = tbl
		*/

		// TODO: Insert stored procedure here
		groupDetailRows, err := ctx.GetGroupDetails(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error querying MySQL for groups: %v", err), http.StatusInternalServerError)
			return
		}

		// create variables and fill contents from retrieved rows
		groupDetail := &messages.GroupDetail{}
		people := []*messages.Person{}
		// var people []*messages.Person

		currentPerson := &messages.Person{}
		row := &groupDetailRow{}
		currentName := ""
		for groupDetailRows.Next() {
			err = groupDetailRows.Scan(row)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error scanning group detail row: %v", err), http.StatusInternalServerError)
				return
			}
			// TODO: maybe optimize to actually check if these already exist
			groupDetail.ID = row.GroupID
			groupDetail.Name = row.GroupName
			currentName = row.FName + row.LName
			groupDetail.PeoplePreview = append(groupDetail.PeoplePreview, currentName)

			currentPerson = &messages.Person{
				ID:       row.PersonnelID,
				FName:    row.FName,
				LName:    row.LName,
				Position: row.RoleTitle,
			}

			people = append(people, currentPerson)
		}

		respond(w, groupDetail)
	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
