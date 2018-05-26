package messages

type ClientMsg struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type ClientMissionWaypoint struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ETA         string `json:"ETA"`
	FlightRules string `json:"flightRules"`
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	Active      string `json:"active"` // denotes active waypoint
	// Completed   string `json:"completed"`
}
