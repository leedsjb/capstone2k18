import React, { Component } from "react";
import { connect } from "react-redux";
import { Helmet } from "react-helmet";

import Accordion from "../../components/Accordion";
import AccordionSection from "../../components/AccordionSection";
import CrewDetailListItem from "../../components/CrewDetailListItem";
import FlexFillVH from "../../components/FlexFillVH";
import InsetMapView from "../../components/InsetMapView";
import ScrollView from "../../components/ScrollView";
import TabBar from "../../components/TabBar";
import Text from "../../components/Text";
import TitleBar from "../../components/TitleBar";
import Icon from "../../components/Icon";

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
            return (
                <Accordion>
                    {aircraftDetail.data.crew ? (
                        <AccordionSection title="Assigned Crew">
                            {aircraftDetail.data.crew.people.map(c => {
                                return (
                                    <CrewDetailListItem
                                        crewDetail={c}
                                        key={c.id}
                                    />
                                );
                            })}
                        </AccordionSection>
                    ) : (
                        <div />
                    )}

                    {aircraftDetail.data.mission ? (
                        <AccordionSection title="Requestor">
                            <Text>{aircraftDetail.data.mission.requestor}</Text>
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
                    {aircraftDetail.data.mission ? (
                        <AccordionSection title="Receiver">
                            <Text>{aircraftDetail.data.mission.receiver}</Text>
                        </AccordionSection>
                    ) : (
                        <div />
                    )}
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
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Missions</title>
                </Helmet>
                <TitleBar back backPath={backPath} />
                {this.props.aircraftDetail.error ? (
                    <FlexFillVH flexDirection="column">
                        An error has occurred:{" "}
                        {this.props.aircraftDetail.error.toString()}
                    </FlexFillVH>
                ) : (
                    <ScrollView>
                        <InsetMapView id={this.props.id} />
                        {this.renderAircraftDetail(this.props.aircraftDetail)}
                    </ScrollView>
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
