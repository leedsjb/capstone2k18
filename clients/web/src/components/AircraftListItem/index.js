import React from "react";
import { Flex } from "grid-styled";

import Badge from "../Badge";
import Box from "../Box";
import MasterListItem from "../MasterListItem";
import Span from "../Span";
import Icon from "../Icon";
import AircraftIdentifier from "../AircraftIdentifier";

const AircraftListItem = ({ aircraft, active }) => {
    return (
        <MasterListItem active={active}>
            <Box>
                <AircraftIdentifier
                    callsign={aircraft.callsign}
                    nNum={aircraft.nNum}
                    status={aircraft.status}
                />
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
