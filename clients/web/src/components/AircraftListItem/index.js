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
                    <Box mr={2}>
                        <Span fontWeight="bold">{aircraft.callsign}</Span>
                        <Icon glyph="hourglass" />
                    </Box>
                    <Box mr={1}>
                        <Span>{aircraft.nNum}</Span>
                    </Box>
                    <Badge>{aircraft.status}</Badge>
                </Flex>
            </Box>
        </MasterListItem>
    );
};

export default AircraftListItem;
