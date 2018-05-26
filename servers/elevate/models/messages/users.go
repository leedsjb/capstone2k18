package messages

// [PUB/SUB Messages]

/*
User Create
Topic ID: USR1
Topic Name: user_create
Description: Create new user

User Update
Topic ID: USR2
Topic Name: user_update
Description: Update existing user
*/
type User struct {
	ID         string `json:"ID"`
	UserName   string `json:"userName"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Initials   string `json:"initials"`
	Email      string `json:"email"`
	GroupID    int    `json:"groupID"`
	Role       string `json:"role"`
	CellPhone  string `json:"cellPhone"`
	// UWNetID         string `json:"UWNetID"`
	// QualificationID string `json:"qualificationID"`
}

/*
User Delete
Topic ID: USR3
Topic Name: user_delete
Description: Delete existing user
FV Table: Users
*/
type User_Delete struct {
	ID string `json:"id"`
}

// [Client Messages]

type PersonGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Person ...
type Person struct {
	ID       int    `json:"id"`
	FName    string `json:"fName"`
	LName    string `json:"lName"`
	Position string `json:"position"`
}

// PersonDetail ...
type PersonDetail struct {
	ID           int            `json:"id"`
	FName        string         `json:"fName"`
	LName        string         `json:"lName"`
	Position     string         `json:"position"`
	Mobile       string         `json:"mobile"`
	Email        string         `json:"email"`
	MemberGroups []*PersonGroup `json:"memberGroups"`
	// UWNetID      string `json:"uwNetID"`
	// SpecialQuals string `json:"specialQuals"`
}
