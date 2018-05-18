import React from "react";
import { Flex } from "grid-styled";

import Badge from "../Badge";
import Span from "../Span";
import Box from "../Box";

const AircraftIdentifier = ({ callsign, nNum, status }) => {
    return (
        <Flex alignItems="center">
            <Span fontWeight="bold">{callsign}</Span>
            <Box ml={2}>
                <Span>{nNum}</Span>
            </Box>
            <Box ml={2}>
                <Badge>{status}</Badge>
            </Box>
        </Flex>
    );
};

export default AircraftIdentifier;
