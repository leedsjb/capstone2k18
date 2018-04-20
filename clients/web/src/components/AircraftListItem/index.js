import React from "react";
import { Flex } from "grid-styled";

import Badge from "../Badge";
import Box from "../Box";
import MasterListItem from "../MasterListItem";
import Span from "../Span";

const AircraftListItem = ({ aircraft }) => {
    return (
        <MasterListItem>
            <Box>
                <Flex alignItems="center">
                    <Box mr={2}>
                        <Span fontWeight="bold">{aircraft.callsign}</Span>
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
