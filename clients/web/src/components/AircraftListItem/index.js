import React from "react";
import { Flex } from "grid-styled";

import Badge from "../Badge";
import Box from "../Box";
import MasterListItem from "../MasterListItem";
import Span from "../Span";
import Icon from "../Icon";

const AircraftListItem = ({ aircraft, active }) => {
    return (
        <MasterListItem active={active}>
            <Box>
                <Flex alignItems="center">
                    <Span fontWeight="bold">{aircraft.callsign}</Span>
                    <Span>{aircraft.nNum}</Span>
                    <Badge>{aircraft.status}</Badge>
                </Flex>
                <Flex alignItems="center">
                    <Icon glyph="earth" size={14} />
                    <Span>{aircraft.area}</Span>
                </Flex>
            </Box>
        </MasterListItem>
    );
};

export default AircraftListItem;
