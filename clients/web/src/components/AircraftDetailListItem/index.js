import React from "react";

import Box from "../Box";
import CrewDetailListItem from "../CrewDetailListItem";
import Heading from "../Heading";
import MasterListItem from "../MasterListItem";
import Text from "../Text";

const AircraftDetailListItem = ({ aircraftDetail }) => {
    return (
        <Box>
            {aircraftDetail.data.crew ? (
                <Box>
                    <Heading is="h6" fontSize={3} mt={3}>
                        Assigned Crew
                    </Heading>
                    {aircraftDetail.data.crew.people.map(c => {
                        return <CrewDetailListItem crewDetail={c} key={c.id} />;
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
    );
};

export default AircraftDetailListItem;
