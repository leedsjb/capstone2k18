import React from "react";

import MasterListItem from "../MasterListItem";

const AircraftListItem = ({ aircraft }) => {
    return <MasterListItem>{aircraft.callsign}</MasterListItem>;
};

export default AircraftListItem;
