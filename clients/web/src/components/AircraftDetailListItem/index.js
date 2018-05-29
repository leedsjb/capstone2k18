import React, { Component } from "react";
import { Flex } from "grid-styled";
import { push } from "react-router-redux";
import { connect } from "react-redux";

import AircraftIdentifier from "../AircraftIdentifier";
import Box from "../Box";
import Clickable from "../Clickable";
import CrewDetailListItem from "../CrewDetailListItem";
import Divider from "../Divider";
import Heading from "../Heading";
import Icon from "../Icon";
import OOSInformation from "../OOSInformation";
import RadioReport from "../RadioReport";
import Receiver from "../Receiver";
import Requestor from "../Requestor";
import Route from "../Route";
import ScrollView from "../ScrollView";
import Span from "../Span";

class AircraftDetailListItem extends Component {
    render() {
        const { mission } = this.props.aircraftDetail.data;

        return (
            <Flex flexDirection="column" flex={1}>
                <Flex
                    py={3}
                    justifyContent="space-between"
                    alignItems="center"
                    px={3}
                >
                    <Box w={16} />
                    <AircraftIdentifier
                        callsign={this.props.aircraftDetail.data.callsign}
                        nNum={this.props.aircraftDetail.data.nNum}
                        status={this.props.aircraftDetail.data.status}
                    />
                    <Clickable onClick={() => this.props.push("/aircraft")}>
                        <Icon glyph="close" size={12} />
                    </Clickable>
                </Flex>
                <Divider />
                <ScrollView>
                    <Box mx={6} my={6}>
                        {mission ? (
                            <div>
                                <Heading is="h2" fontSize={4}>
                                    ETA
                                </Heading>
                                <Box mt={6}>
                                    <Route waypoints={mission.waypoints} />
                                </Box>
                                <Heading is="h2" fontSize={4} my={6}>
                                    Patient Summary
                                </Heading>
                                <Box mt={2} mb={4}>
                                    <Span fontWeight="bold">
                                        Transfer Number:{" "}
                                    </Span>
                                    <Span>{mission.flightNum}</Span>
                                </Box>
                                {mission.radioReport &&
                                mission.radioReport.shortReport ? (
                                    <RadioReport
                                        radioReport={mission.radioReport}
                                    />
                                ) : null}
                            </div>
                        ) : null}
                        {this.props.aircraftDetail.data.crew ? (
                            <Box>
                                <Heading
                                    is="h2"
                                    fontSize={4}
                                    mt={
                                        this.props.aircraftDetail.data.mission
                                            ? 6
                                            : 0
                                    }
                                    mb={2}
                                >
                                    {this.props.aircraftDetail.data.mission
                                        ? "Mission Crew"
                                        : "Assigned Crew"}
                                </Heading>
                                <CrewDetailListItem
                                    crew={
                                        this.props.aircraftDetail.data.mission
                                            ? this.props.aircraftDetail.data
                                                  .mission.crew
                                            : this.props.aircraftDetail.data
                                                  .crew
                                    }
                                />
                            </Box>
                        ) : null}
                        {mission ? (
                            <div>
                                <Box>
                                    <Heading is="h2" fontSize={4} my={6}>
                                        Requestor
                                    </Heading>
                                    {mission.requestor ? (
                                        <Requestor
                                            requestor={mission.requestor}
                                        />
                                    ) : null}
                                </Box>
                                <Box>
                                    <Heading is="h2" fontSize={4} my={6}>
                                        Receiver
                                    </Heading>
                                    {mission.receiver ? (
                                        <Receiver receiver={mission.receiver} />
                                    ) : null}
                                </Box>
                            </div>
                        ) : null}
                        {this.props.aircraftDetail.data.OOS ? (
                            <div>
                                <Heading
                                    is="h2"
                                    fontSize={4}
                                    mt={
                                        this.props.aircraftDetail.data.crew
                                            ? 6
                                            : 0
                                    }
                                    mb={4}
                                >
                                    Out of Service Information
                                </Heading>
                                <OOSInformation
                                    OOS={this.props.aircraftDetail.data.OOS}
                                />
                            </div>
                        ) : null}
                    </Box>
                </ScrollView>
            </Flex>
        );
    }
}

const mapDispatchToProps = {
    push
};

export default connect(null, mapDispatchToProps)(AircraftDetailListItem);
