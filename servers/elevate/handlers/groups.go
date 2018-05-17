package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

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

func (ctx *HandlerContext) GetTrieGroups(groupIDS []int) ([]*messages.ClientGroup, error) {
	groups := []*messages.ClientGroup{}

	// get each group with matching id
	for _, groupID := range groupIDS {
		group := &messages.ClientGroup{}
		ID := strconv.Itoa(groupID)
		groupRows, err := ctx.GetGroupByID(ID)
		groupRow := &groupRow{}
		// for each group with matching id,
		// get all matching rows
		for groupRows.Next() {
			err = groupRows.Scan(groupRow)
			if err != nil {
				return nil, fmt.Errorf("Error scanning group row: %v", err)
			}
			// for each matching row, re-define
			// groupID and groupName, which should stay the same
			// and append members until there are no more
			group, err = ctx.GetGroupSummary(groupRow, group)
			if err != nil {
				return nil, fmt.Errorf("Error populating group for trie: %v", err)
			}
		}
		// after getting all the members, add the group
		// to the list of returned groups
		groups = append(groups, group)
	}
	return groups, nil
}

// GetGroupSummary populates a passed-in group with ID, Name, and
// appends the current given row's member name to the group's list of members
func (ctx *HandlerContext) GetGroupSummary(currentRow *groupRow, group *messages.ClientGroup) (*messages.ClientGroup, error) {
	person := currentRow.FName + " " + currentRow.LName
	people := append(group.PeoplePreview, person)

	group = &messages.ClientGroup{
		ID:            currentRow.GroupID,
		Name:          currentRow.GroupName,
		PeoplePreview: people,
	}
	return group, nil
}

// GroupsHandler ...
func (ctx *HandlerContext) GroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()

		term := query.Get("q")

		groups := []*messages.ClientGroup{}

		if len(term) > 0 {
			// search query non-empty
			// get matching group prefixes
			groupIDS := ctx.AircraftTrie.GetEntities(strings.ToLower(term), 20)
			groups, err := ctx.GetTrieGroups(groupIDS)
			if err != nil {
				fmt.Printf("Error pulling groups from trie: %v", err)
				return
			}
			respond(w, groups)
		} else {
			// no filters, return all groups

			// type Group struct {
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
		}
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
