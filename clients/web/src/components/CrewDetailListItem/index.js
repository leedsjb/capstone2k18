import React from "react";

import MasterListItem from "../MasterListItem";

const CrewDetailListItem = ({ crewDetail }) => {
    return (
        <MasterListItem>
            {crewDetail.lName}, {crewDetail.fName} ({crewDetail.position})
        </MasterListItem>
    );
};

export default CrewDetailListItem;
