package messages

// [PUB/SUB Messages]
/*
Missions
Topic ID: NM1
Topic Name: mission_create
Description: New mission created
*/
type Patient struct {
	ShortReport string `json:"shortReport"`
	Intubated   string `json:"intubated"`
	Drips       string `json:"drips"`
	Age         string `json:"age"`
	Weight      string `json:"weight"`
	Gender      string `json:"gender"`
	Cardiac     string `json:"cardiac"`
	GIBleed     string `json:"GIBleed"`
	OB          string `json:"OB"`
}

type MissionWaypoint struct {
	ID          string `json:"id"`
	ETA         string `json:"ETA"`         // time to next point
	Active      string `json:"active"`      // cumulative mission time
	FlightRules string `json:"flightRules"` // denotes active waypoint
	Completed   string `json:"completed"`
}

type Mission_Crew struct {
	ID     string `json:"id"`
	RoleID string `json:"roleID"`
}

type Mission_Create struct {
	MissionID   string             `json:"missionID"`
	TCNum       string             `json:"TCNum"`
	Asset       string             `json:"asset"`
	RequestorID string             `json:"reqID"`
	ReceiverID  string             `json:"recID"`
	CallType    string             `json:"callType"`
	Patient     *Patient           `json:"patient"`
	CrewMembers []*Mission_Crew    `json:"crewMembers"`
	Waypoints   []*MissionWaypoint `json:"waypoints"`
}

type Mission_Complete struct {
	MissionID string `json:"missionID"`
}

/*
Mission Waypoint Update
Topic ID: MU1
Topic Name: mission_waypoints_update
Description: Changes to waypoint after mission creation (waypoints, ETE, ETT, active leg)
*/
type Mission_Waypoint_Update struct {
	MissionID string             `json:"missionID"`
	Waypoints []*MissionWaypoint `json:"waypoints"`
}

/*
Mission Crew Update
Topic ID: MU2
Topic Name: mission_crew_update
Description: Changes to crew after mission creation
*/
type Mission_Crew_Update struct {
	MissionID   string          `json:"missionID"`
	CrewMembers []*Mission_Crew `json:"crewMembers"`
}

// [Client Messages]

// Mission ...
type Mission struct {
	Type            string                   `json:"type"`
	NextWaypointETE string                   `json:"nextWaypointETE"`
	Waypoints       []*ClientMissionWaypoint `json:"waypoints"`
	FlightNum       string                   `json:"flightNum"`
	Completed       bool                     `json:"completed"`
}

type Agency struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Type    string `json:"type"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

// MissionDetail ...
type MissionDetail struct {
	Type            string                   `json:"type"`
	NextWaypointETE string                   `json:"nextWaypointETE"`
	Waypoints       []*ClientMissionWaypoint `json:"waypoints"`
	FlightNum       string                   `json:"flightNum"`
	RadioReport     *ClientPatient           `json:"radioReport"`
	Requestor       *Agency                  `json:"requestor"`
	Receiver        *Agency                  `json:"receiver"`
	Completed       bool                     `json:"completed"`
	Crew            []*Person                `json:"crew"`
}

type ClientPatient struct {
	ShortReport string `json:"shortReport"`
	Intubated   bool   `json:"intubated"`
	Drips       int    `json:"drips"`
	Age         int    `json:"age"`
	Weight      int    `json:"weight"`
	Gender      string `json:"gender"`
	Cardiac     bool   `json:"cardiac"`
	GIBleed     bool   `json:"GIBleed"`
	OB          bool   `json:"OB"`
}
