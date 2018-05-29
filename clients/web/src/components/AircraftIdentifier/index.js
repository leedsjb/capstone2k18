import React from "react";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";

import Badge from "../Badge";
import Span from "../Span";
import Box from "../Box";

const statuses = {
    OAM: {
        desc: "On Mission",
        bg: "orange2",
        color: "orange1"
    },
    OOS: {
        desc: "Out of Service",
        bg: "pink2",
        color: "pink1"
    },
    RFM: {
        desc: "Ready",
        bg: "green2",
        color: "green1"
    }
};

const AircraftIdentifier = ({ theme: { colors }, ...props }) => {
    return (
        <Flex alignItems="center">
            <Span fontSize={props.fontSize} fontWeight="bold">
                {props.callsign}
            </Span>
            <Box ml={2}>
                <Span fontSize={props.fontSize}>{props.nNum}</Span>
            </Box>
            <Box ml={2}>
                <Badge
                    bg={statuses[props.status].bg}
                    color={statuses[props.status].color}
                    border={`1px solid ${colors[statuses[props.status].color]}`}
                >
                    {statuses[props.status].desc}
                </Badge>
            </Box>
        </Flex>
    );
};

export default withTheme(AircraftIdentifier);
