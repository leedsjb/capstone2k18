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

type Mission_Create struct {
	MissionID    string             `json:"missionID"`
	TCNum        string             `json:"TCNum"`
	Asset        string             `json:"asset"`
	RequestorID  string             `json:"reqID"`
	ReceiverID   string             `json:"recID"`
	CallType     string             `json:"callType"`
	Patient      *Patient           `json:"patient"`
	CrewMemberID []string           `json:"crewMemberID"`
	Waypoints    []*MissionWaypoint `json:"waypoints"`
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
	MissionID    string   `json:"missionID"`
	CrewMemberID []string `json:"crewMemberID"`
}

// [Client Messages]

// Mission ...
type Mission struct {
	Type            string                   `json:"type"`
	NextWaypointETE string                   `json:"nextWaypointETE"`
	Waypoints       []*ClientMissionWaypoint `json:"waypoints"`
	FlightNum       string                   `json:"flightNum"`
	Completed       string                   `json:"completed"`
}

// MissionDetail ...
type MissionDetail struct {
	Type            string                   `json:"type"`
	NextWaypointETE string                   `json:"nextWaypointETE"`
	Waypoints       []*ClientMissionWaypoint `json:"waypoints"`
	FlightNum       string                   `json:"flightNum"`
	RadioReport     *Patient                 `json:"radioReport"`
	Requestor       string                   `json:"requestor"`
	Receiver        string                   `json:"receiver"`
	Completed       string                   `json:"completed"`
}
