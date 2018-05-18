import React from "react";
import { Flex } from "grid-styled";

import Badge from "../Badge";
import Span from "../Span";

const AircraftIdentifier = ({ callsign, nNum, status }) => {
    return (
        <Flex alignItems="center">
            <Span fontWeight="bold">{callsign}</Span>
            <Span>{nNum}</Span>
            <Badge>{status}</Badge>
        </Flex>
    );
};

export default AircraftIdentifier;
