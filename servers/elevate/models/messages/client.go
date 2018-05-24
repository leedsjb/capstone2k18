package messages

type ClientMsg struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type ClientMissionWaypoint struct {
	Name        string `json:"name"`
	ETA         string `json:"ETA"`         // time to next point
	Active      string `json:"active"`      // cumulative mission time
	FlightRules string `json:"flightRules"` // denotes active waypoint
	// Completed   string `json:"completed"`
}
