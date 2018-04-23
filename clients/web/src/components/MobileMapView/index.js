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

            if (selected) {
                return (
                    <Link
                        to={`/aircraft/${selected.id}?source=map`}
                        key={selected.id}
                    >
                        <AircraftListItem aircraft={selected} />
                    </Link>
                );
            }
        }
    }

    render() {
        return (
            <MasterDetailView>
                <MasterView>
                    <Flex
                        style={{ height: "100%", width: "100%" }}
                        flexDirection="column"
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

function mapStateToProps(state, ownProps) {
    return {
        aircraft: state.aircraft
    };
}

const mapDispatchToProps = {
    fetchAircraft
};

export default connect(mapStateToProps, mapDispatchToProps)(MobileMapView);
