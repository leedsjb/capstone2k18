import React from "react";
import ContentLoader, { Facebook } from "react-content-loader";
import { Flex } from "grid-styled";

import MissionListItem from "../MissionListItem";
import Box from "../Box";
import Badge from "../Badge";

const MissionList = ({ data }) => {
    let content;
    console.log(data);

    if (data.length === 0) {
        content = <Facebook />;
    } else {
        content = data.map(mission => {
            return (
                <div>
                    <Flex alignItems="center">
                        <Badge>{mission.type}</Badge>
                        {mission.nNum}
                    </Flex>
                </div>
            );
        });
    }

    return <Box p={3}>{content}</Box>;
};

export default MissionList;
