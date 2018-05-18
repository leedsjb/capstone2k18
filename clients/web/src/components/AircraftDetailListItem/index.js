import React from "react";
import { Flex } from "grid-styled";

import Box from "../Box";
import CrewDetailListItem from "../CrewDetailListItem";
import Heading from "../Heading";
import Text from "../Text";
import Divider from "../Divider";
import Badge from "../Badge";
import AicraftIdentifier from "../AircraftIdentifier";
import AircraftIdentifier from "../AircraftIdentifier";

const AircraftDetailListItem = ({ aircraftDetail }) => {
    let { callsign, nNum, status } = aircraftDetail.data;

    return (
        <Box>
            <Flex py={3} justifyContent="center">
                <AircraftIdentifier
                    callsign={callsign}
                    nNum={nNum}
                    status={status}
                />
            </Flex>
            <Divider />
            <Box px={6}>
                <Heading is="h2" fontSize={4}>
                    ETA
                </Heading>
                <Heading is="h2" fontSize={4}>
                    Patient
                </Heading>
                {aircraftDetail.data.crew ? (
                    <Box>
                        <Heading is="h2" fontSize={4}>
                            Assigned Crew
                        </Heading>
                        {aircraftDetail.data.crew.people.map(c => {
                            return (
                                <CrewDetailListItem crewDetail={c} key={c.id} />
                            );
                        })}
                    </Box>
                ) : null}
                {aircraftDetail.data.mission ? (
                    <Box>
                        <Heading is="h6" fontSize={3} my={3}>
                            Radio Report
                        </Heading>
                        <Text>{aircraftDetail.data.mission.radioReport}</Text>
                    </Box>
                ) : null}
                {aircraftDetail.data.mission ? (
                    <Box>
                        <Heading is="h6" fontSize={3} my={3}>
                            Requestor
                        </Heading>
                        <Text>{aircraftDetail.data.mission.requestor}</Text>
                    </Box>
                ) : null}
                {aircraftDetail.data.mission ? (
                    <Box>
                        <Heading is="h6" fontSize={3} my={3}>
                            Receiver
                        </Heading>
                        <Text>{aircraftDetail.data.mission.receiver}</Text>
                    </Box>
                ) : null}
            </Box>
        </Box>
    );
};

export default AircraftDetailListItem;
