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
                    <Box ml={2}>
                        <Span>{aircraft.nNum}</Span>
                    </Box>
                    <Box ml={2}>
                        <Badge>{aircraft.status}</Badge>
                    </Box>
                </Flex>
                <Flex alignItems="center" mt={2}>
                    <Icon glyph="earth" size={14} />
                    <Box ml={2}>
                        <Span>{aircraft.area}</Span>
                    </Box>
                </Flex>
            </Box>
        </MasterListItem>
    );
};

export default AircraftListItem;
