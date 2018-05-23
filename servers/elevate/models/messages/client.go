package messages

type ClientMsg struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type ClientMissionWaypoint struct {
	Name        string `json:"name"`
	ETE         string `json:"ETE"`    // time to next point
	ETT         string `json:"ETT"`    // cumulative mission time
	Active      string `json:"active"` // denotes active waypoint
	FlightRules string `json:"flightRules"`
	Completed   string `json:"completed"`
}
