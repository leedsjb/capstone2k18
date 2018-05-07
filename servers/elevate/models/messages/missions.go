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
	ID     string `json:"ID"`
	ETE    string `json:"ETE"`    // time to next point
	ETT    string `json:"ETT"`    // cumulative mission time
	Active string `json:"active"` // denotes active waypoint
}

type Mission_Create struct {
	MissionID    string             `json:"missionID"`
	TCNum        string             `json:"TCNum"`
	Asset        string             `json:"asset"`
	RequestorID  string             `json:"requestorID"`
	ReceiverID   string             `json:"receiverID"`
	Priority     string             `json:"priority"`
	CallType     string             `json:"callType"`
	Patient      *Patient           `json:"patient"`
	CrewMemberID []string           `json:"crewMemberID"`
	Waypoints    []*MissionWaypoint `json:"waypoints"`
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
	MissionID    string   `json:"missionID"`
	CrewMemberID []string `json:"crewMemberID"`
}

// [Client Messages]

// Mission ...
type Mission struct {
	Key             string             `json:"key"`
	Type            string             `json:"type"`
	Status          string             `json:"status"`
	Vision          string             `json:"vision"`
	NextWaypointETE string             `json:"nextWaypointETE"`
	Waypoints       []*MissionWaypoint `json:"waypoints"`
	FlightNum       string             `json:"flightNum"`
}

// MissionDetail ...
type MissionDetail struct {
	Key             string             `json:"key"`
	Type            string             `json:"type"`
	Status          string             `json:"status"`
	Vision          string             `json:"vision"`
	NextWaypointETE string             `json:"nextWaypointETE"`
	Waypoints       []*MissionWaypoint `json:"waypoints"`
	FlightNum       string             `json:"flightNum"`
	RadioReport     *Patient           `json:"radioReport"`
	Requestor       string             `json:"requestor"`
	Receiver        string             `json:"receiver"`
}
