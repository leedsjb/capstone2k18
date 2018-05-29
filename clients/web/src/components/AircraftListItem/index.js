import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import MasterListItem from "../MasterListItem";
import Span from "../Span";
import Icon from "../Icon";
import AircraftIdentifier from "../AircraftIdentifier";

const AircraftListItem = ({ active, aircraft }) => {
    return (
        <MasterListItem active={active}>
            <Box>
                <AircraftIdentifier
                    callsign={aircraft.callsign}
                    nNum={aircraft.nNum}
                    status={aircraft.status}
                />
                <Flex alignItems="center" mt={1}>
                    <Icon glyph="earth" size={14} color="black2" />
                    <Box ml={2}>
                        <Span color="black2">{aircraft.area}</Span>
                    </Box>
                </Flex>
            </Box>
        </MasterListItem>
    );
};

export default AircraftListItem;
