package messages

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
	ID                int      `json:"ID"`
	Notes             string   `json:"notes"`
	Name              string   `json:"name"`
	Type              string   `json:"type"`
	Address1          string   `json:"address1"`
	Address2          string   `json:"address2"`
	Country           string   `json:"country"`
	State             string   `json:"state"`
	County            string   `json:"county"`
	City              string   `json:"city"`
	Zip               string   `json:"zip"`
	Lat               string   `json:"lat"`
	Long              string   `json:"long"`
	GPSWaypoint       string   `json:"GPSWaypoint"`
	AirportIdentifier string   `json:"AirportIdentifier"`
	Phone             []string `json:"phone"`
	ShortCode         string   `json:"shortCode"`
	PadTime           string   `json:"padTime"`
	RadioChannels     []string `json:"radioChannels"`
	NOTAMS            string   `json:"NOTAMS"`
}

/*
Waypoint Deleted
Topic ID: WP3
Topic Name: waypoint_delete
Description: Waypoint deleted from waypoint catalog
Waypoint types: airport, hospital, landing zone, base
*/
type Waypoint_Delete struct {
	ID int `json:"ID"`
}
