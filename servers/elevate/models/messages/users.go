package messages

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
	ID					string `json:"ID"`
    UserName			string `json:"userName"`
    FirstName			string `json:"firstName"`
    MiddleName			string `json:"middleName"`
    LastName			string `json:"lastName"`
    Initials			string `json:"initials"`
    Email				string `json:"email"`
    UWNetID				string `json:"UWNetID"`
    GroupID				string `json:"groupID"`
    Role				string `json:"role"`
    CellPhone			string `json:"cellPhone"`
    QualificationID		string `json:"qualificationID"`     
}
/* 
User Delete
Topic ID: USR3
Topic Name: user_delete
Description: Delete existing user
FV Table: Users
*/
type User_Delete struct {
	ID		string `json:"ID"`
}