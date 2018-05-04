import React from "react";

import MasterListItem from "../MasterListItem";

const AircraftDetailListItem = ({ aircraftDetail }) => {
    return <MasterListItem>{aircraftDetail.data.nNum}</MasterListItem>;
};

export default AircraftDetailListItem;
