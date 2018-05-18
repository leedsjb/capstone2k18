package parsers

import "github.com/leedsjb/capstone2k18/servers/elevate/models/messages"

// [AIRCRAFT]

func (ctx *ParserContext) AddNewAircraft(aircraftInfo *messages.Aircraft_Create) error {
	return nil
}

func (ctx *ParserContext) UpdateAircraftProps(aircraftInfo *messages.Aircraft_Props_Update) error {
	return nil
}

func (ctx *ParserContext) UpdateAircraftCrew(aircraftInfo *messages.Aircraft_Crew_Update) error {
	return nil
}

func (ctx *ParserContext) UpdateAircraftServiceSchedule(aircraftInfo *messages.Aircraft_Service_Schedule) error {
	return nil
}

func (ctx *ParserContext) UpdateAircraftPosition(aircraftInfo *messages.Aircraft_Pos_Update) error {
	return nil
}

// [GROUPS]

func (ctx *ParserContext) AddNewGroup(groupInfo *messages.Group) error {
	return nil
}

func (ctx *ParserContext) UpdateGroup(groupInfo *messages.Group) error {
	return nil
}

func (ctx *ParserContext) DeleteGroup(groupInfo *messages.Group) error {
	return nil
}
