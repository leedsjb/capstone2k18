import React from "react";

import MasterListItem from "../MasterListItem";

const AircraftListItem = ({ aircraft, active }) => {
    return <MasterListItem active={active}>{aircraft.callsign}</MasterListItem>;
};

export default AircraftListItem;
