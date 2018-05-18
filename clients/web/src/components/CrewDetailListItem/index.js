import React from "react";
import { Flex } from "grid-styled";

import Card from "../Card";
import Avatar from "../Avatar";
import Span from "../Span";

const CrewDetailListItem = ({ crewDetail, ...props }) => {
    return (
        <Card {...props}>
            <Flex flexDirection="column" alignItems="center">
                <Avatar />
                <Span fontWeight="bold">{`${crewDetail.lName} ${
                    crewDetail.fName
                }`}</Span>
                <Span>{crewDetail.position}</Span>
            </Flex>
        </Card>
    );
};

export default CrewDetailListItem;
