import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";

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
import LoadingSpinner from "../../components/LoadingSpinner";
import OOSInformation from "../../components/OOSInformation";
import ScrollView from "../../components/ScrollView";
import Span from "../../components/Span";
import RadioReport from "../../components/RadioReport";
import Receiver from "../../components/Receiver";
import Requestor from "../../components/Requestor";
import Route from "../../components/Route";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";

import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

class AircraftDetailPage extends Component {
    componentDidMount() {
        if (this.props.id) {
            this.props.fetchAircraftDetail(this.props.id);
        }
    }

    renderAircraftDetail(aircraftDetail) {
        const { mission } = this.props.aircraftDetail.data;
        return (
            <Accordion>
                <Box borderTop="1px solid black">
                    {mission && mission.radioReport ? (
                        <AccordionSection title="Radio Report">
                            {mission.radioReport &&
                            mission.radioReport.shortReport ? (
                                <Box mb={3}>
                                    <Box my={6}>
                                        <Span fontWeight="bold">
                                            Patient Summary
                                        </Span>
                                        <Box mt={2} mb={4}>
                                            <Span>{mission.flightNum}</Span>
                                        </Box>
                                        <RadioReport
                                            radioReport={mission.radioReport}
                                        />
                                    </Box>
                                </Box>
                            ) : null}
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.crew ? (
                        <AccordionSection title="Assigned Crew">
                            <Box mb={8}>
                                <CrewDetailListItem
                                    crew={this.props.aircraftDetail.data.crew}
                                />
                            </Box>
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.mission ? (
                        <AccordionSection title="Requestor">
                            <Box my={3}>
                                <Requestor requestor={mission.requestor} />
                            </Box>
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.mission ? (
                        <AccordionSection title="Receiver">
                            <Box my={3}>
                                <Receiver receiver={mission.receiver} />
                            </Box>
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.OOS ? (
                        <AccordionSection title="OOS Information">
                            <Box my={3}>
                                <OOSInformation
                                    OOS={this.props.aircraftDetail.data.OOS}
                                />
                            </Box>
                        </AccordionSection>
                    ) : null}
                </Box>
            </Accordion>
        );
    }

    renderContent() {
        if (this.props.aircraftDetail.error) {
            return (
                <FlexFillVH>
                    An error has occurred:{" "}
                    {this.props.aircraftDetail.error.toString()}
                </FlexFillVH>
            );
        } else if (!Array.isArray(this.props.aircraftDetail.data)) {
            return (
                <ScrollView>
                    <Container my={3}>
                        <AircraftIdentifier
                            callsign={this.props.aircraftDetail.data.callsign}
                            nNum={this.props.aircraftDetail.data.nNum}
                            status={this.props.aircraftDetail.data.status}
                            fontSize={4}
                            ml={2}
                        />
                    </Container>
                    <Divider />
                    <Container mt={3}>
                        {this.props.aircraftDetail.data.mission ? (
                            <div>
                                <Heading is="h2" fontSize={4}>
                                    Route
                                </Heading>
                                <Route
                                    waypoints={
                                        this.props.aircraftDetail.data.mission
                                            .waypoints
                                    }
                                />
                            </div>
                        ) : null}

                        <Box height={160} my={4}>
                            <InsetMapView id={this.props.id} />
                        </Box>
                        <Flex alignItems="center" mb={4}>
                            <Icon glyph="earth" size={16} />
                            <Box ml={2}>
                                <Span>
                                    {this.props.aircraftDetail.data.area}
                                </Span>
                            </Box>
                        </Flex>
                    </Container>

                    <Divider />
                    {this.renderAircraftDetail(this.props.aircraftDetail)}
                </ScrollView>
            );
        }
        return <LoadingSpinner />;
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
                {this.renderContent()}
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
