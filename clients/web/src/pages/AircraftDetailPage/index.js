import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { withTheme } from "styled-components";
import { Helmet } from "react-helmet";

import Accordion from "../../components/Accordion";
import AccordionSection from "../../components/AccordionSection";
import AircraftIdentifier from "../../components/AircraftIdentifier";
import Container from "../../components/Container";
import CrewDetailListItem from "../../components/CrewDetailListItem";
import Box from "../../components/Box";
import Divider from "../../components/Divider";
import Error from "../../components/Error";
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
                <Box borderTop={`1px solid ${this.props.theme.colors.gray4}`}>
                    {mission && mission.radioReport ? (
                        <AccordionSection title="Radio Report">
                            {mission.radioReport &&
                            mission.radioReport.shortReport ? (
                                <div>
                                    <Span fontWeight="bold">
                                        Patient Summary
                                    </Span>
                                    <Box my={3}>
                                        <Span>{mission.flightNum}</Span>
                                    </Box>
                                    <RadioReport
                                        radioReport={mission.radioReport}
                                    />
                                </div>
                            ) : null}
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.crew ? (
                        <AccordionSection title="Assigned Crew">
                            <CrewDetailListItem
                                crew={this.props.aircraftDetail.data.crew}
                            />
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.mission ? (
                        <AccordionSection title="Requestor">
                            <Requestor requestor={mission.requestor} />
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.mission ? (
                        <AccordionSection title="Receiver">
                            <Receiver receiver={mission.receiver} />
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {this.props.aircraftDetail.data.OOS ? (
                        <AccordionSection title="OOS Information">
                            <OOSInformation
                                OOS={this.props.aircraftDetail.data.OOS}
                            />
                        </AccordionSection>
                    ) : null}
                </Box>
            </Accordion>
        );
    }

    renderContent() {
        if (this.props.aircraftDetail.error) {
            return (
                <Flex
                    flexDirection="column"
                    flex={1}
                    alignItems="center"
                    justifyContent="center"
                >
                    <Error
                        title="An error has occurred"
                        content={this.props.aircraftDetail.error.toString()}
                    />
                </Flex>
            );
        } else if (
            this.props.aircraftDetail &&
            this.props.aircraftDetail.data.id === Number(this.props.id) &&
            !Array.isArray(this.props.aircraftDetail.data)
        ) {
            return (
                <ScrollView>
                    <Container my={3}>
                        <AircraftIdentifier
                            callsign={this.props.aircraftDetail.data.callsign}
                            nNum={this.props.aircraftDetail.data.nNum}
                            status={this.props.aircraftDetail.data.status}
                            fontSize={3}
                        />
                    </Container>
                    <Divider />
                    <Container mt={3}>
                        {this.props.aircraftDetail.data.mission ? (
                            <div>
                                <Heading is="h2" fontSize={3} mt={1} mb={4}>
                                    Route
                                </Heading>
                                <Route
                                    mb={4}
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
                        <Flex alignItems="center">
                            <Icon glyph="earth" size={16} />
                            <Box ml={2}>
                                <Span>
                                    {this.props.aircraftDetail.data.area}
                                </Span>
                            </Box>
                        </Flex>
                    </Container>
                    <Box py={4}>
                        {this.renderAircraftDetail(this.props.aircraftDetail)}
                    </Box>
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
        if (
            this.props.aircraftDetail.data &&
            !Array.isArray(this.props.aircraftDetail.data)
        ) {
            if (this.props.aircraftDetail.data.mission) {
                title = `Flight #${
                    this.props.aircraftDetail.data.mission.flightNum
                }`;
            } else {
                title = `Aircraft ${this.props.aircraftDetail.data.nNum}`;
            }
        } else if (this.props.aircraftDetail.error) {
            title = "Error";
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

export default connect(mapStateToProps, mapDispatchToProps)(
    withTheme(AircraftDetailPage)
);
