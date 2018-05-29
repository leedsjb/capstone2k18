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
	GroupID   int
	GroupName string
	FName     string
	LName     string
}

type groupDetailRow struct {
	GroupID     int
	GroupName   string
	PersonnelID int
	FName       string
	LName       string
	RoleTitle   string
	SMS         string
	Email       string
}

// IndexGroup ...
func IndexGroup(trie *indexes.Trie, group *messages.Group) error {

	groupID, err := strconv.Atoi(group.ID)
	if err != nil {
		return fmt.Errorf("Could not convert group ID to int: %v", err)
	}

	if err := trie.AddEntity(strings.ToLower(group.Name), groupID); err != nil {
		return fmt.Errorf("Error adding group to trie: %v", err)
	}

	for _, member := range group.Members {
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
	currentRow.GroupID = -1
	currentGroupID := -1
	currentGroup := &messages.Group{}

	for groupRows.Next() {
		err = groupRows.Scan(&currentRow.GroupID, &currentRow.GroupName, &currentRow.FName, &currentRow.LName)
		if err != nil {
			return fmt.Errorf("Error scanning group row: %v", err)
		}
		// if this isn't the first group or if it is a different group than before,
		// append the previous group

		if currentGroupID == -1 {
			fmt.Println("[GROUP HANDLER] currentGroupID was first")
			currentGroupID = currentRow.GroupID
		}
		fmt.Printf("[GROUP HANDLER] Current RowID is: [%v], GroupID is: [%v]\n", currentRow.GroupID, currentGroupID)
		if currentRow.GroupID != currentGroupID {
			fmt.Printf("[GROUP HANDLER] Non-matching rowID [%v] and groupID [%v]\n", currentRow.GroupID, currentGroupID)
			if err := IndexGroup(trie, currentGroup); err != nil {
				return fmt.Errorf("Error loading trie: %v", err)
			}
			// empty out the current group being built
			currentGroup = &messages.Group{}
			// update current groupID
			currentGroupID = currentRow.GroupID
		}
		// otherwise, continue as usual and assign values to the receiving struct
		// TODO: maybe optimize to actually check if these already exist
		currentGroup, err = ctx.GetGroupSummary(currentRow, currentGroup)
		if err != nil {
			return fmt.Errorf("Error populating group for trie: %v", err)
		}
	}
	close(groupRows)
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
		currentGroup := &messages.Group{}
		groupRows, err := ctx.GetGroupByID(groupID)
		if err != nil {
			return nil, fmt.Errorf("Error retrieving groups by ID: %v", err)
		}
		currentRow := &groupRow{}
		for groupRows.Next() {
			err = groupRows.Scan(&currentRow.GroupID, &currentRow.GroupName, &currentRow.FName, &currentRow.LName)
			if err != nil {
				return nil, fmt.Errorf("Error scanning group row: %v", err)
			}
			// for each matching row, re-define
			// groupID and groupName, which should stay the same
			// and append members until there are no more
			currentGroup, err = ctx.GetGroupSummary(currentRow, currentGroup)
			if err != nil {
				return nil, fmt.Errorf("Error populating group for trie: %v", err)
			}
		}
		// after getting all the members, add the group
		// to the list of returned groups
		fmt.Printf("[GROUP HANDLER TRIE] Add group to return groups: %#v\n", currentGroup)
		clientGroup, err := GroupToClientGroup(currentGroup)
		if err != nil {
			return nil, fmt.Errorf("Could not parse Group to ClientGroup: %v", err)
		}
		groups = append(groups, clientGroup)
		close(groupRows)
	}
	return groups, nil
}

func GroupToClientGroup(group *messages.Group) (*messages.ClientGroup, error) {
	var preview string
	var firstMember string
	var count int
	for i, member := range group.Members {
		nameParts := strings.Fields(member)
		for j, namePart := range nameParts {
			if j == 0 {
				if i == 0 {
					preview = preview + namePart
					firstMember = namePart
				}
				if i == 1 {
					preview = preview + " and " + namePart
				}
			}
			break
		}
		count = i
	}
	if count > 1 {
		otherMembers := strconv.Itoa(count)
		preview = firstMember + " and " + otherMembers + " others"
	}

	client := &messages.ClientGroup{
		Name:          group.Name,
		PeoplePreview: preview,
	}
	if len(group.ID) > 0 {
		groupID, err := strconv.Atoi(group.ID)
		if err != nil {
			return nil, fmt.Errorf("Could not convert group ID to int: %v", err)
		}
		client.ID = groupID
	}
	return client, nil
}

func PeopleToPreview(people []*messages.GroupMember) string {
	var preview string
	var firstMember string
	var count int
	for i, person := range people {
		if i == 0 {
			preview = preview + person.FName
			firstMember = person.FName
		}
		if i == 1 {
			preview = preview + " and " + person.FName
		}
		count = i
	}
	if count > 1 {
		otherMembers := strconv.Itoa(count)
		preview = firstMember + " and " + otherMembers + " others"
	}
	return preview
}

// GetGroupSummary populates a passed-in group with ID, Name, and
// appends the current given row's member name to the group's list of members

// TODO: does this need an error??
func (ctx *HandlerContext) GetGroupSummary(currentRow *groupRow, group *messages.Group) (*messages.Group, error) {
	person := currentRow.FName + " " + currentRow.LName
	people := append(group.Members, person)

	groupID := strconv.Itoa(currentRow.GroupID)

	group = &messages.Group{
		ID:      groupID,
		Name:    currentRow.GroupName,
		Members: people,
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
			groupIDS := ctx.GroupsTrie.GetEntities(strings.ToLower(term), 20)
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
			currentRow.GroupID = -1
			currentGroupID := -1
			currentGroup := &messages.Group{}
			for groupRows.Next() {
				err = groupRows.Scan(&currentRow.GroupID, &currentRow.GroupName, &currentRow.FName, &currentRow.LName)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error scanning group row: %v", err), http.StatusInternalServerError)
					return
				}
				// if this isn't the first group or if it is a different group than before,
				// append the previous group

				if currentGroupID == -1 {
					fmt.Println("[GROUP HANDLER] currentGroupID was first")
					currentGroupID = currentRow.GroupID
				}
				if currentRow.GroupID != currentGroupID {
					clientGroup, err := GroupToClientGroup(currentGroup)
					if err != nil {
						http.Error(w, fmt.Sprintf("Error parsing group to client group: %v", err), http.StatusInternalServerError)
					}
					groups = append(groups, clientGroup)
					// empty out the current group being built
					currentGroup = &messages.Group{}
					// update current groupID
					currentGroupID = currentRow.GroupID
				}
				// otherwise, continue as usual and assign values to the receiving struct
				// TODO: maybe optimize to actually check if these already exist
				currentGroup, err = ctx.GetGroupSummary(currentRow, currentGroup)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error populating group for trie: %v", err), http.StatusInternalServerError)
					return
				}
			}
			// add last group to the list of groups
			clientGroup, err := GroupToClientGroup(currentGroup)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing group to client group: %v", err), http.StatusInternalServerError)
			}
			groups = append(groups, clientGroup)

			close(groupRows)

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

		if id != "." && id != "groups" {
			groupID, err := strconv.Atoi(id)
			if err != nil {
				fmt.Printf("Error changing group ID from string to int")
			}
			groupDetailRows, err := ctx.GetGroupDetailByID(groupID)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error retrieving groups details for group [%v]: %v", id, err), http.StatusInternalServerError)
				return
			}

			// create variables and fill contents from retrieved rows
			groupDetail := &messages.GroupDetail{}
			members := []*messages.GroupMember{}

			currentMember := &messages.GroupMember{}
			row := &groupDetailRow{}
			// currentName := ""
			for groupDetailRows.Next() {
				err = groupDetailRows.Scan(
					&row.GroupID,
					&row.GroupName,
					&row.PersonnelID,
					&row.FName,
					&row.LName,
					&row.RoleTitle,
					&row.SMS,
					&row.Email,
				)
				if err != nil {
					http.Error(w, fmt.Sprintf("Error scanning group detail row for group [%v]: %v", id, err), http.StatusInternalServerError)
					return
				}
				// TODO: maybe optimize to actually check if these already exist
				groupDetail.ID = row.GroupID
				groupDetail.Name = row.GroupName

				currentMember = &messages.GroupMember{
					ID:       row.PersonnelID,
					FName:    row.FName,
					LName:    row.LName,
					Position: row.RoleTitle,
					SMS:      row.SMS,
					Email:    row.Email,
				}

				members = append(members, currentMember)
			}

			close(groupDetailRows)

			// change array of members to client-friendly string
			preview := PeopleToPreview(members)
			groupDetail.PeoplePreview = preview
			// attach list of people to the groupDetail
			groupDetail.Members = members
			respond(w, groupDetail)
		} else if id == "groups" {
			ctx.GroupsHandler(w, r)
		} else {
			http.Error(w, "No group with that ID", http.StatusBadRequest)
		}

	default:
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
}
