import React from "react";
import ContentLoader, { Facebook } from "react-content-loader";

import MissionListItem from "../MissionListItem";
import Box from "../Box";

const MissionList = ({ data }) => {
    let content;

    if (data.length === 0) {
        content = <Facebook />;
    } else {
        content = data.map(mission => {
            return <div>{mission.nNum}</div>;
        });
    }

    return <Box p={3}>{content}</Box>;
};

export default MissionList;
