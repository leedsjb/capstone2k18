import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import { Link } from "react-router-dom";

import Accordion from "../../components/Accordion";
import AccordionSection from "../../components/AccordionSection";
import AircraftIdentifier from "../../components/AircraftIdentifier";
import Container from "../../components/Container";
import CrewDetailListItem from "../../components/CrewDetailListItem";
import Box from "../../components/Box";
import Divider from "../../components/Divider";
import FlexFillVH from "../../components/FlexFillVH";
import Heading from "../../components/Heading";
import Icon from "../../components/Icon";
import InsetMapView from "../../components/InsetMapView";
import ScrollView from "../../components/ScrollView";
import Span from "../../components/Span";
import TabBar from "../../components/TabBar";
import Text from "../../components/Text";
import TitleBar from "../../components/TitleBar";

import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

class AircraftDetailPage extends Component {
    componentDidMount() {
        if (this.props.id) {
            this.props.fetchAircraftDetail(this.props.id);
        }
    }

    renderAircraftDetail(aircraftDetail) {
        if (
            !this.props.aircraftDetail.pending &&
            !Array.isArray(this.props.aircraftDetail.data)
        ) {
            const { mission } = this.props.aircraftDetail.data;

            return (
                <Accordion>
                    <Box borderTop="1px solid black">
                        {mission ? (
                            <AccordionSection title="Radio Report">
                                {mission.radioReport &&
                                mission.radioReport.shortReport ? (
                                    <Box mb={3}>
                                        <Box mt={4}>
                                            <Span fontWeight="bold">
                                                Patient Summary
                                            </Span>
                                            <Box mt={2} mb={6}>
                                                <Span>{mission.flightNum}</Span>
                                            </Box>
                                            <Span fontWeight="bold">
                                                Short report
                                            </Span>
                                            <Text mt={1}>
                                                {
                                                    mission.radioReport
                                                        .shortReport
                                                }
                                            </Text>
                                        </Box>
                                        {mission.radioReport.gender ? (
                                            <Box mt={4}>
                                                <Span fontWeight="bold">
                                                    Sex:{" "}
                                                </Span>
                                                <Span>
                                                    {mission.radioReport.gender}
                                                </Span>
                                            </Box>
                                        ) : null}
                                        {mission.radioReport.age ? (
                                            <Box mt={4}>
                                                <Span fontWeight="bold">
                                                    Age:{" "}
                                                </Span>
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
                                                        mission.radioReport
                                                            .weight
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
                                                        mission.radioReport
                                                            .GIBleed
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
                                                        mission.radioReport
                                                            .cardiac
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
                                                    {
                                                        mission.radioReport
                                                            .intubated
                                                    }
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
                                    </Box>
                                ) : null}
                            </AccordionSection>
                        ) : (
                            <div />
                        )}
                        <AccordionSection title="Assigned Crew">
                            {this.props.aircraftDetail.data.crew ? (
                                <Box mb={3}>
                                    <Flex
                                        flexWrap="wrap"
                                        justifyContent="space-between"
                                        mt={3}
                                    >
                                        {this.props.aircraftDetail.data.crew.people.map(
                                            c => {
                                                return (
                                                    <Box
                                                        key={c.id}
                                                        w="calc(50% - 8px)"
                                                    >
                                                        <Link
                                                            to={`/people/${
                                                                c.id
                                                            }`}
                                                        >
                                                            <CrewDetailListItem
                                                                crewDetail={c}
                                                            />
                                                        </Link>
                                                    </Box>
                                                );
                                            }
                                        )}
                                    </Flex>
                                </Box>
                            ) : (
                                <div />
                            )}
                        </AccordionSection>
                        {this.props.aircraftDetail.data.mission ? (
                            <AccordionSection title="Requestor">
                                <Box my={3}>
                                    <Text>
                                        {
                                            this.props.aircraftDetail.data
                                                .mission.requestor
                                        }
                                    </Text>
                                </Box>
                            </AccordionSection>
                        ) : (
                            <div />
                        )}
                        {this.props.aircraftDetail.data.mission ? (
                            <AccordionSection title="Receiver">
                                <Box my={3}>
                                    <Text>
                                        {
                                            this.props.aircraftDetail.data
                                                .mission.receiver
                                        }
                                    </Text>
                                </Box>
                            </AccordionSection>
                        ) : (
                            <div />
                        )}
                    </Box>
                </Accordion>
            );
        }
        return <div>Loading...</div>;
    }

    render() {
        let backPath =
            new URLSearchParams(window.location.search).get("source") === "map"
                ? `/aircraft/map/${this.props.id}`
                : "/aircraft";
        let title = "Loading...";
        if (this.props.aircraftDetail.data.mission) {
            title = `Flight #${
                this.props.aircraftDetail.data.mission.flightNum
            }`;
        } else {
            title = `Aircraft ${this.props.aircraftDetail.data.nNum}`;
        }
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Missions</title>
                </Helmet>
                <TitleBar back backPath={backPath} title={title} />
                {this.props.aircraftDetail.error ? (
                    <FlexFillVH>
                        An error has occurred:{" "}
                        {this.props.aircraftDetail.error.toString()}
                    </FlexFillVH>
                ) : (
                    <FlexFillVH flexDirection="column">
                        <Container my={3}>
                            <AircraftIdentifier
                                callsign={
                                    this.props.aircraftDetail.data.callsign
                                }
                                nNum={this.props.aircraftDetail.data.nNum}
                                status={this.props.aircraftDetail.data.status}
                                fontSize={4}
                                ml={2}
                            />
                        </Container>
                        <Divider />
                        <Container mt={3}>
                            <Heading is="h2" fontSize={4}>
                                Route
                            </Heading>
                            <Text mt={2}>Route component goes here</Text>
                            <Box height={160} my={4}>
                                <InsetMapView id={this.props.id} />
                            </Box>
                            <Icon glyph="earth" />
                        </Container>

                        <Divider />
                        <ScrollView>
                            {this.renderAircraftDetail(
                                this.props.aircraftDetail
                            )}
                        </ScrollView>
                    </FlexFillVH>
                )}
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        aircraftDetail: state.aircraftDetail,
        id: ownProps.match.params.id
    };
}

const mapDispatchToProps = {
    fetchAircraftDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(AircraftDetailPage);
