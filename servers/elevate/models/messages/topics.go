package messages

/* 
Missions
Topic ID: NM1
Topic Name: mission_create
Description: New mission created
*/
type Patient struct {
	ShortReport		string `json:"shortReport"`
	Intubated		string `json:"intubated"`
	Drips			string `json:"drips"`
	Age				string `json:"age"`
	Weight			string `json:"weight"`
	Gender			string `json:"gender"`
	Cardiac			string `json:"cardiac"`
	GIBleed			string `json:"GIBleed"`
	OB				string `json:"OB"`
}

type MissionWaypoint struct {
	ID		string `json:"ID"`
	ETE		string `json:"ETE"`		// time to next point
	ETT		string `json:"ETT"`		// cumulative mission time
	Active	string `json:"active"`	// denotes active waypoint
}

type Mission_Create struct {
	MissionID			string 				`json:"missionID"`
	TCNum				string   			`json:"TCNum"`
	Asset				string   			`json:"asset"`
	RequestorID			string   			`json:"requestorID"`
	ReceiverID			string 			 	`json:"receiverID"`
	Priority			string 			 	`json:"priority"`
	CallType			string 	 			`json:"callType"`
	Patient				*Patient 			`json:"patient"`
	CrewMemberID		[]string 			`json:"crewMemberID"`
	Waypoints			[]*MissionWaypoint  `json:"waypoints"`
}

/* 
Mission Waypoint Update
Topic ID: MU1
Topic Name: mission_waypoints_update
Description: Changes to waypoint after mission creation (waypoints, ETE, ETT, active leg)
*/
type Mission_Waypoint_Update struct {
	MissionID		string 		`json:"missionID"`
	Waypoints		[]*Waypoint `json:"waypoints"`
}

/* 
Mission Crew Update
Topic ID: MU2
Topic Name: mission_crew_update
Description: Changes to crew after mission creation
*/
type Mission_Crew_Update struct {
	MissionID		string `json:"missionID"`
    CrewMemberID	[]string `json:"crewMemberID"`
}

/* 
Waypoint Created
Topic ID: WP1
Topic Name: waypoint_create
Description: New waypoint added to Flight Vector db by dispatcher
Waypoint types: agency, hospital, airport, landing zone

Waypoint Updated
Topic ID: WP2
Topic Name: waypoint_update
Description: Waypoint updated in waypoint catalog
Waypoint types: airport, hospital, landing zone, base
*/
type Waypoint struct {
	ID					string 	 `json:"ID"`
	Notes				string 	 `json:"notes"`
    Name				string 	 `json:"name"`
    Type				string 	 `json:"type"`
    Address1			string 	 `json:"address1"`
    Address2			string 	 `json:"address2"`
    Country				string 	 `json:"country"`
    State				string 	 `json:"state"`
    County				string 	 `json:"county"`
    City				string 	 `json:"city"`
    Zip					string 	 `json:"zip"`
    Lat					string 	 `json:"lat"`
    Long				string 	 `json:"long"`
    GPSWaypoint			string 	 `json:"GPSWaypoint"`
    AirportIdentifier 	string 	 `json:"AirportIdentifier"`
    Phone				[]string `json:"phone"`
    ShortCode			string	 `json:"shortCode"`
    PadTime				string	 `json:"padTime"`
	RadioChannels		[]string `json:"radioChannels"`
    NOTAMS				string   `json:"NOTAMS"`
}

/* 
Waypoint Deleted
Topic ID: WP3
Topic Name: waypoint_delete
Description: Waypoint deleted from waypoint catalog
Waypoint types: airport, hospital, landing zone, base
*/
type Waypoint_Delete struct {
    ID		string `json:"ID"`
}

/* 
Aircraft Created
Topic ID: AC1
Topic Name: aircraft_create
Description: New aircraft added
FV Table: Vendor
*/
type Aircraft_Create struct {
	ID 						string   `json:"ID"`
    NNum					string   `json:"nNum"`
    SatPhone				string   `json:"satPhone"`
    CellPhone				string   `json:"cellPhone"`
    Base					string   `json:"baseID"`
    Callsign				string   `json:"callsign"`
    MaxPatientWeight		string   `json:"maxPatientWeight"`
    PadTimeDay				string   `json:"padTimeDay"`
    PadTimeNight			string   `json:"padTimeNight"`
    Vendor					string   `json:"vendor"`
    Status					string   `json:"status"`
    SpecialEquipment		string   `json:"specialEquipment"`
    Color					string   `json:"color"`
    LastKnownLocation		string   `json:"lastKnownLocation"`
    Model					string   `json:"model"`
    CallTypes				[]string `json:"callTypes"`
}

/* 
Aircraft Properties Update
Topic ID: AC2
Topic Name: ac_properties_update
Description: Aircraft persistent properties updated
*/
type Aircraft_Props_Update struct {
	ID					string `json:"ID"`
	SatPhone			string `json:"satPhone"`
    CellPhone			string `json:"cellPhone"`
    Base				string `json:"base"`
    Callsign			string `json:"callsign"`
    MaxPatientWeight 	string `json:"maxPatientWeight"`
    PadTimeDay			string `json:"padTimeDay"`
    PadTimeNight		string `json:"padTimeNight"`
    Vendor				string `json:"vendor"`
    SpecialEquipment	string `json:"specialEquipment"`
}

/* 
Aircraft Crew Update
Topic ID: AC3
Topic Name: ac_crew_update
Description: Aircraft crew reassigned
*/
type Aircraft_Crew_Update struct { 
	PIC				string `json:"PIC"`
	AdultRN			string `json:"adultRN"`
    PediatricRN		string `json:"pediatricRN"`
}

/* 
Aircraft Scheduled Service
Topic ID: AC4
Topic Name: ac_service_schedule
Description: Aircraft service scheduled
*/
type Aircraft_Service_Schedule struct {
	OosReadon		string `json:"oosReason"`
	StartTime		string `json:"startTime"`
    EndTime			string `json:"endTime"`
    Status			string `json:"status"`
}

/* 
Aircraft Postion Update
Topic ID: AC5
Topic Name: ac_position_update
Description: Aircraft position updated
*/
type Aircraft_Pos_Update struct {
	PosLat				string `json:"posLat"`
	PosLong				string `json:"posLong"`
    PosFriendlyName		string `json:"posFriendlyName"`
}

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
	ID			string   `json:"ID"`
    Name		string   `json:"Name"`
    Members		[]string `json:"members"`
}

/*
Group Delete
Topic ID: GR3
Topic Name: group_delete
Description: Delete existing group
FV Table: UserGroup/GroupOfUsers/PagableGroups?
*/
type Group_Delete struct {
    ID		string `json:"ID"`
}