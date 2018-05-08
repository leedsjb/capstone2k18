package messages

// [PUB/SUB Messages]

/*
Group Create
Topic ID: GR1
Topic Name: group_create
Description: Create a new group

Group Update
Topic ID: GR2
Topic Name: group_update
Description: Add or delete member, change group name
FV Table: UserGroup/GroupOfUsers/PagableGroups?
*/
type Group struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Members []string `json:"members"`
}

/*
Group Delete
Topic ID: GR3
Topic Name: group_delete
Description: Delete existing group
FV Table: UserGroup/GroupOfUsers/PagableGroups?
*/
type Group_Delete struct {
	ID string `json:"ID"`
}

// [Client Messages]

// Group ...
type ClientGroup struct {
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
