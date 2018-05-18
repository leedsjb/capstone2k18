import React, { Component } from "react";
import { Flex } from "grid-styled";
import { push } from "react-router-redux";
import { connect } from "react-redux";

import Box from "../Box";
import CrewDetailListItem from "../CrewDetailListItem";
import Heading from "../Heading";
import Text from "../Text";
import Divider from "../Divider";
import Badge from "../Badge";
import AicraftIdentifier from "../AircraftIdentifier";
import AircraftIdentifier from "../AircraftIdentifier";
import Icon from "../Icon";
import Clickable from "../Clickable";
import Span from "../Span";

class AircraftDetailListItem extends Component {
    render() {
        const { mission } = this.props.aircraftDetail.data;

        return (
            <Box>
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
                <Box mx={6} my={6}>
                    {mission ? (
                        <div>
                            <Heading is="h2" fontSize={4}>
                                ETA
                            </Heading>
                            <Text mt={2}>Route component goes here</Text>
                            <Heading is="h2" fontSize={4} mt={8}>
                                Patient
                            </Heading>
                            <Box mt={2} mb={6}>
                                <Span>{mission.flightNum}</Span>
                            </Box>
                            {mission.radioReport &&
                            mission.radioReport.shortReport ? (
                                <div>
                                    <Box mt={4}>
                                        <Span fontWeight="bold">
                                            Short report
                                        </Span>
                                        <Text mt={1}>
                                            {mission.radioReport.shortReport}
                                        </Text>
                                    </Box>
                                    {mission.radioReport.gender ? (
                                        <Box mt={4}>
                                            <Span fontWeight="bold">Sex: </Span>
                                            <Span>
                                                {mission.radioReport.gender}
                                            </Span>
                                        </Box>
                                    ) : null}
                                    {mission.radioReport.age ? (
                                        <Box mt={4}>
                                            <Span fontWeight="bold">Age: </Span>
                                            <Span>
                                                {mission.radioReport.age}
                                            </Span>
                                        </Box>
                                    ) : null}
                                    {mission.radioReport.weight ? (
                                        <Box mt={4}>
                                            <Span fontWeight="bold">
                                                Weight:{" "}
                                            </Span>
                                            <Span>
                                                {`${
                                                    mission.radioReport.weight
                                                } kg`}
                                            </Span>
                                        </Box>
                                    ) : null}
                                    {mission.radioReport.GIBleed ? (
                                        <Box mt={4}>
                                            <Span fontWeight="bold">
                                                Placeholder
                                            </Span>
                                            <Text mt={1}>
                                                {`${
                                                    mission.radioReport.GIBleed
                                                } kg`}
                                            </Text>
                                        </Box>
                                    ) : null}
                                    {mission.radioReport.cardiac ? (
                                        <Box mt={4}>
                                            <Span fontWeight="bold">
                                                Placeholder
                                            </Span>
                                            <Text mt={1}>
                                                {`${
                                                    mission.radioReport.cardiac
                                                } kg`}
                                            </Text>
                                        </Box>
                                    ) : null}
                                    {mission.radioReport.intubated ? (
                                        <Box mt={4}>
                                            <Span fontWeight="bold">
                                                Intubated:{" "}
                                            </Span>
                                            <Span>
                                                {mission.radioReport.intubated}
                                            </Span>
                                        </Box>
                                    ) : null}
                                    {mission.radioReport.drips ? (
                                        <Box mt={4}>
                                            <Span fontWeight="bold">
                                                Drips:{" "}
                                            </Span>
                                            <Span>
                                                {mission.radioReport.drips}
                                            </Span>
                                        </Box>
                                    ) : null}
                                </div>
                            ) : null}
                        </div>
                    ) : null}

                    {this.props.aircraftDetail.data.crew ? (
                        <Box>
                            <Heading is="h2" fontSize={4} mt={8}>
                                Assigned Crew
                            </Heading>
                            <Flex
                                flexWrap="wrap"
                                justifyContent="space-between"
                                mt={6}
                            >
                                {this.props.aircraftDetail.data.crew.people.map(
                                    c => {
                                        return (
                                            <CrewDetailListItem
                                                crewDetail={c}
                                                key={c.id}
                                                w="calc(50% - 8px)"
                                            />
                                        );
                                    }
                                )}
                            </Flex>
                        </Box>
                    ) : null}
                    {this.props.aircraftDetail.data.mission ? (
                        <Box>
                            <Heading is="h6" fontSize={3} mt={8}>
                                Requestor
                            </Heading>
                            <Text mt={2}>
                                {
                                    this.props.aircraftDetail.data.mission
                                        .requestor
                                }
                            </Text>
                        </Box>
                    ) : null}
                    {this.props.aircraftDetail.data.mission ? (
                        <Box>
                            <Heading is="h6" fontSize={3} mt={8}>
                                Receiver
                            </Heading>
                            <Text mt={2}>
                                {
                                    this.props.aircraftDetail.data.mission
                                        .receiver
                                }
                            </Text>
                        </Box>
                    ) : null}
                </Box>
            </Box>
        );
    }
}

const mapDispatchToProps = {
    push
};

export default connect(null, mapDispatchToProps)(AircraftDetailListItem);
