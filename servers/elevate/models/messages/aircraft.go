package messages

// [PUB/SUB Messages]

/*
Aircraft Created
Topic ID: AC1
Topic Name: aircraft_create
Description: New aircraft added
FV Table: Vendor
*/
type Aircraft_Create struct {
	ID                string   `json:"ID"`
	NNum              string   `json:"nNum"`
	SatPhone          string   `json:"satPhone"`
	CellPhone         string   `json:"cellPhone"`
	Base              string   `json:"baseID"`
	Callsign          string   `json:"callsign"`
	MaxPatientWeight  string   `json:"maxPatientWeight"`
	PadTimeDay        string   `json:"padTimeDay"`
	PadTimeNight      string   `json:"padTimeNight"`
	Vendor            string   `json:"vendor"`
	Status            string   `json:"status"`
	SpecialEquipment  string   `json:"specialEquipment"`
	Color             string   `json:"color"`
	LastKnownLocation string   `json:"lastKnownLocation"`
	Model             string   `json:"model"`
	CallTypes         []string `json:"callTypes"`
}

/*
Aircraft Properties Update
Topic ID: AC2
Topic Name: ac_properties_update
Description: Aircraft persistent properties updated
*/
type Aircraft_Props_Update struct {
	ID               string `json:"ID"`
	SatPhone         string `json:"satPhone"`
	CellPhone        string `json:"cellPhone"`
	Base             string `json:"base"`
	Callsign         string `json:"callsign"`
	MaxPatientWeight string `json:"maxPatientWeight"`
	PadTimeDay       string `json:"padTimeDay"`
	PadTimeNight     string `json:"padTimeNight"`
	Vendor           string `json:"vendor"`
	SpecialEquipment string `json:"specialEquipment"`
}

/*
Aircraft Crew Update
Topic ID: AC3
Topic Name: ac_crew_update
Description: Aircraft crew reassigned
*/
type Aircraft_Crew_Update struct {
	PIC         string `json:"PIC"`
	AdultRN     string `json:"adultRN"`
	PediatricRN string `json:"pediatricRN"`
}

/*
Aircraft Scheduled Service
Topic ID: AC4
Topic Name: ac_service_schedule
Description: Aircraft service scheduled
*/
type Aircraft_Service_Schedule struct {
	OosReason string `json:"oosReason"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Status    string `json:"status"`
}

/*
Aircraft Postion Update
Topic ID: AC5
Topic Name: ac_position_update
Description: Aircraft position updated
*/
type Aircraft_Pos_Update struct {
	PosLat          string `json:"posLat"`
	PosLong         string `json:"posLong"`
	PosFriendlyName string `json:"posFriendlyName"`
}

// [Client Messages]

// OOS ...
type OOS struct {
	Key       string `json:"key"`
	Reason    string `json:"reason"`
	Remaining string `json:"remaining"`
}

// OOSDetail ...
type OOSDetail struct {
	Key       string `json:"key"`
	Reason    string `json:"reason"`
	Remaining string `json:"remaining"`
	Duration  string `json:"duration"`
}

// Aircraft ...
type Aircraft struct {
	Key         string   `json:"key"`
	ID          int      `json:"id"`
	Status      string   `json:"status"`
	Type        string   `json:"type"`
	Callsign    string   `json:"callsign"`
	LevelOfCare string   `json:"levelOfCare"`
	Class       string   `json:"class"`
	Lat         float32  `json:"lat"`
	Long        float32  `json:"long"`
	Area        string   `json:"area"`
	NNum        string   `json:"nNum"`
	Mission     *Mission `json:"mission"`
	OOS         *OOS     `json:"OOS"`
}

// AircraftDetail ...
type AircraftDetail struct {
	Key         string         `json:"key"`
	ID          int            `json:"id"`
	Status      string         `json:"status"`
	Type        string         `json:"type"`
	Callsign    string         `json:"callsign"`
	Crew        *GroupDetail   `json:"crew"`
	LevelOfCare string         `json:"levelOfCare"`
	Class       string         `json:"class"`
	Lat         float32        `json:"lat"`
	Long        float32        `json:"long"`
	Area        string         `json:"area"`
	NNum        string         `json:"nNum"`
	Mission     *MissionDetail `json:"mission"`
	OOS         *OOSDetail     `json:"OOS"`
}
