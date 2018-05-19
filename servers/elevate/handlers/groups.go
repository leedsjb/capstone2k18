package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/leedsjb/capstone2k18/servers/elevate/indexes"
	"github.com/leedsjb/capstone2k18/servers/elevate/models/messages"
)

// These structs receive data from SQL queries
// and allow the structure of SQL data to be abstracted
// from the structure of what is sent to the client.

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

// IndexGroup ...
func IndexGroup(trie *indexes.Trie, group *messages.ClientGroup) error {
	fmt.Println("group.ID is: " + group.ID)
	groupID, err := strconv.Atoi(group.ID)
	if err != nil {
		fmt.Printf("Error changing group ID from string to int")
	}
	if err := trie.AddEntity(strings.ToLower(group.Name), groupID); err != nil {
		return fmt.Errorf("Error adding group to trie: %v", err)
	}

	for _, member := range group.PeoplePreview {
		nameParts := strings.Fields(member)
		for _, namePart := range nameParts {
			if err := trie.AddEntity(strings.ToLower(namePart), groupID); err != nil {
				return fmt.Errorf("Error adding group to trie: %v", err)
			}
		}
	}
	return nil
}

// LoadGroupsTrie ...
// *** Pass in the same trie as for handlers/people.go LoadPeopleTrie
// to allow both to be mutually searchable ***
func (ctx *HandlerContext) LoadGroupsTrie(trie *indexes.Trie) error {
	groupRows, err := ctx.GetAllGroups()
	if err != nil {
		return fmt.Errorf("Error retrieving groups for personnel trie: %v", err)
	}
	// create variables and fill contents from retrieved rows

	currentRow := &groupRow{}
	currentRow.GroupID = "first"
	currentGroupID := "first"
	currentGroup := &messages.ClientGroup{}
	var rowID string
	var rowName string
	var rowFName string
	var rowLName string

	for groupRows.Next() {
		err = groupRows.Scan(&rowID, &rowName, &rowFName, &rowLName)
		if err != nil {
			return fmt.Errorf("Error scanning group row: %v", err)
		}
		// if this isn't the first group or if it is a different group than before,
		// append the previous group

		if currentGroupID == "first" {
			fmt.Println("[GROUP HANDLER] currentGroupID was first")
			currentGroupID = rowID
		}
		fmt.Printf("[GROUP HANDLER] Current RowID is: [%v], GroupID is: [%v]\n", rowID, currentGroupID)
		if rowID != currentGroupID {
			fmt.Printf("[GROUP HANDLER] Non-matching rowID [%v] and groupID [%v]\n", rowID, currentGroupID)
			if err := IndexGroup(trie, currentGroup); err != nil {
				return fmt.Errorf("Error loading trie: %v", err)
			}
			// empty out the current group being built
			currentGroup = &messages.ClientGroup{}
			// update current groupID
			currentGroupID = rowID
		}
		// otherwise, continue as usual and assign values to the receiving struct
		currentRow = &groupRow{
			GroupID:   rowID,
			GroupName: rowName,
			FName:     rowFName,
			LName:     rowLName,
		}
		// TODO: maybe optimize to actually check if these already exist
		currentGroup, err = ctx.GetGroupSummary(currentRow, currentGroup)
		if err != nil {
			return fmt.Errorf("Error populating group for trie: %v", err)
		}
	}
	// Index last group
	if err := IndexGroup(trie, currentGroup); err != nil {
		return fmt.Errorf("Error loading trie: %v", err)
	}
	return nil
}

func (ctx *HandlerContext) GetTrieGroups(groupIDS []int) ([]*messages.ClientGroup, error) {
	groups := []*messages.ClientGroup{}

	// get each group whose prefix matches the search term
	for _, groupID := range groupIDS {
		currentGroup := &messages.ClientGroup{}
		ID := strconv.Itoa(groupID)
		groupRows, err := ctx.GetGroupByID(ID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving groups by ID: %v", err)
		}
		currentRow := &groupRow{}
		var rowID string
		var rowName string
		var rowFName string
		var rowLName string
		for groupRows.Next() {
			err = groupRows.Scan(&rowID, &rowName, &rowFName, &rowLName)
			if err != nil {
				return nil, fmt.Errorf("Error scanning group row: %v", err)
			}
			// for each matching row, re-define
			// groupID and groupName, which should stay the same
			// and append members until there are no more
			currentRow = &groupRow{
				GroupID:   rowID,
				GroupName: rowName,
				FName:     rowFName,
				LName:     rowLName,
			}
			currentGroup, err = ctx.GetGroupSummary(currentRow, currentGroup)
			if err != nil {
				return nil, fmt.Errorf("Error populating group for trie: %v", err)
			}
		}
		// after getting all the members, add the group
		// to the list of returned groups
		fmt.Printf("[GROUP HANDLER TRIE] Add group to return groups: %#v\n", currentGroup)
		groups = append(groups, currentGroup)
	}
	return groups, nil
}

// GetGroupSummary populates a passed-in group with ID, Name, and
// appends the current given row's member name to the group's list of members

// TODO: does this need an error??
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
		fmt.Printf("[GROUP TRIE] query is: %v\n", query)

		term := query.Get("q")
		fmt.Printf("[GROUP TRIE] term is: %v\n", term)

		groups := []*messages.ClientGroup{}

		if len(term) > 0 {
			fmt.Println("[GROUP TRIE] in filtering")
			// search query non-empty
			// find groupIDs that match the search term
			groupIDS := ctx.PersonnelTrie.GetEntities(strings.ToLower(term), 20)
			fmt.Printf("[GROUP TRIE] groupIDS: %v", groupIDS)
			// retrieve the actual group information
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

			groupRows, err := ctx.GetAllGroups()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error getting groups: %v", err), http.StatusInternalServerError)
				return
			}
			// create variables and fill contents from retrieved rows
			currentRow := &groupRow{}
			var rowID string
			var rowName string
			var rowFName string
			var rowLName string
			currentRow.GroupID = "first"
			currentGroupID := "first"
			currentGroup := &messages.ClientGroup{}
			for groupRows.Next() {
				err = groupRows.Scan(&rowID, &rowName, &rowFName, &rowLName)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error scanning group row: %v", err), http.StatusInternalServerError)
					return
				}
				// if this isn't the first group or if it is a different group than before,
				// append the previous group

				if currentGroupID == "first" {
					fmt.Println("[GROUP HANDLER] currentGroupID was first")
					currentGroupID = rowID
				}
				fmt.Printf("[GROUP HANDLER] Current RowID is: [%v], GroupID is: [%v]\n", rowID, currentGroupID)
				if rowID != currentGroupID {
					fmt.Printf("[GROUP HANDLER] Non-matching rowID [%v] and groupID [%v]\n", rowID, currentGroupID)
					groups = append(groups, currentGroup)
					// empty out the current group being built
					currentGroup = &messages.ClientGroup{}
					// update current groupID
					currentGroupID = rowID
				}
				// otherwise, continue as usual and assign values to the receiving struct
				currentRow = &groupRow{
					GroupID:   rowID,
					GroupName: rowName,
					FName:     rowFName,
					LName:     rowLName,
				}
				// TODO: maybe optimize to actually check if these already exist
				currentGroup, err = ctx.GetGroupSummary(currentRow, currentGroup)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error populating group for trie: %v", err), http.StatusInternalServerError)
					return
				}
			}
			// add last group to the list of groups
			groups = append(groups, currentGroup)

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

		// TODO: make sure this isn't potentially exposing anything
		id := path.Base(r.URL.Path)

		if id != "." {
			// TODO: Insert stored procedure here
			groupDetailRows, err := ctx.GetGroupDetailsByID(id)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error retrieving groups details for group [%v]: %v", id, err), http.StatusInternalServerError)
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
					http.Error(w, fmt.Sprintf("Error scanning group detail row for group [%v]: %v", id, err), http.StatusInternalServerError)
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
		} else {
			http.Error(w, "No group with that ID", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
