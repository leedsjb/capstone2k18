package messages

type ClientMsg struct {
	Type	string `json:"type"`
	Payload interface{} `json:"payload"`
}

type ClientMissionWaypoint struct {
	Name	string `json:"name"`
	ETE		string `json:"ETE"`		// time to next point
	ETT		string `json:"ETT"`		// cumulative mission time
	Active	string `json:"active"`	// denotes active waypoint
}

type Client_Mission_Create struct {
	MissionID			string 				`json:"missionID"`
	TCNum				string   			`json:"TCNum"`
	Asset				string   			`json:"asset"`
	Requestor			string   			`json:"requestor"`
	Receiver			string 			 	`json:"receiver"`
	Priority			string 			 	`json:"priority"`
	CallType			string 	 			`json:"callType"`
	Patient				*Patient 			`json:"patient"`
	CrewMembers		    string	 			`json:"crewMemberID"`
	Waypoints			[]*ClientMissionWaypoint  `json:"waypoints"`
}