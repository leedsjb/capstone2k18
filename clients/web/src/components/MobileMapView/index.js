import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

import AircraftListItem from "../../components/AircraftListItem";
import Box from "../../components/Box";
import Heading from "../../components/Heading";
import MapView from "../MapView";
import MasterView from "../MasterView";
import MasterDetailView from "../MasterDetailView";
import Text from "../../components/Text";

import { fetchAircraft } from "../../actions/aircraft/actions";

class MobileMapView extends Component {
    constructor(props) {
        super(props);
        this.state = {
            selAircraft: null
        };
    }

    componentDidMount() {
        if (this.props.aircraftID) {
            this.props
                .fetchAircraft()
                .then(this.renderAircraft(this.props.aircraft));
        }
    }

    componentWillReceiveProps(nextProps) {
        if (
            nextProps.aircraftID &&
            nextProps.aircraftID !== this.props.aircraftID
        ) {
            this.renderAircraft(this.props.aircraft);
        }
    }

    renderAircraft(aircraft) {
        if (!aircraft.pending && aircraft.data.length > 0) {
            let selected = aircraft.data.find(air => {
                return air.id == this.props.aircraftID;
            });

            return (
                <Link to={`/aircraft/${selected.id}`} key={selected.id}>
                    <AircraftListItem aircraft={selected} mobile={true} />
                </Link>
            );
        } else if (!aircraft.pending) {
            return (
                <Box mt={4}>
                    <Heading is="h2" textAlign="center" fontSize={4}>
                        No aircraft
                    </Heading>
                    <Text textAlign="center">Empty state text</Text>
                </Box>
            );
        }
    }

    render() {
        return (
            <MasterDetailView>
                <MasterView>
                    <Flex
                        style={{ height: "100%", width: "100%" }}
                        flexDirection={[null, null, "column", "row"]}
                        flexWrap="wrap"
                    >
                        <MapView />
                        {this.props.aircraftID
                            ? this.renderAircraft(this.props.aircraft)
                            : null}
                    </Flex>
                </MasterView>
            </MasterDetailView>
        );
    }
}

function mapStateToProps(state) {
    return {
        aircraft: state.aircraft
    };
}

const mapDispatchToProps = {
    fetchAircraft
};

export default connect(mapStateToProps, mapDispatchToProps)(MobileMapView);
