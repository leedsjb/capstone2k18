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

class AircraftDetailListItem extends Component {
    render() {
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
                    {this.props.aircraftDetail.data.mission ? (
                        <div>
                            <Heading is="h2" fontSize={4}>
                                ETA
                            </Heading>
                            <Text mt={2}>Route component goes here</Text>
                        </div>
                    ) : null}
                    <Heading is="h2" fontSize={4} mt={4}>
                        Patient
                    </Heading>

                    {this.props.aircraftDetail.data.crew ? (
                        <Box>
                            <Heading is="h2" fontSize={4}>
                                Assigned Crew
                            </Heading>
                            <Flex flexWrap="wrap">
                                {this.props.aircraftDetail.data.crew.people.map(
                                    c => {
                                        return (
                                            <CrewDetailListItem
                                                crewDetail={c}
                                                key={c.id}
                                                w="calc(50% - 8px)"
                                                mx={2}
                                            />
                                        );
                                    }
                                )}
                            </Flex>
                        </Box>
                    ) : null}
                    {this.props.aircraftDetail.data.mission ? (
                        <Box>
                            <Heading is="h6" fontSize={3} my={3}>
                                Radio Report
                            </Heading>
                            <Text />
                        </Box>
                    ) : null}
                    {this.props.aircraftDetail.data.mission ? (
                        <Box>
                            <Heading is="h6" fontSize={3} my={3}>
                                Requestor
                            </Heading>
                            <Text>
                                {
                                    this.props.aircraftDetail.data.mission
                                        .requestor
                                }
                            </Text>
                        </Box>
                    ) : null}
                    {this.props.aircraftDetail.data.mission ? (
                        <Box>
                            <Heading is="h6" fontSize={3} my={3}>
                                Receiver
                            </Heading>
                            <Text>
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
