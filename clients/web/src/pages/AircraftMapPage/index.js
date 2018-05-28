import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { connect } from "react-redux";
import { Link } from "react-router-dom";

import AircraftListItem from "../../components/AircraftListItem";
import Box from "../../components/Box";
import FlexFillVH from "../../components/FlexFillVH";
import LoadingSpinner from "../../components/LoadingSpinner";
import MapView from "../../components/MapView";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";

import { fetchAircraft } from "../../actions/aircraft/actions";

class AircraftMapPage extends Component {
    componentDidMount() {
        if (this.props.id) {
            this.props
                .fetchAircraft()
                .then(this.renderAircraft(this.props.aircraft));
        }
    }

    renderAircraft(aircraft) {
        if (!aircraft.pending && aircraft.data.length > 0) {
            let selected = aircraft.data.find(air => {
                return air.id === Number(this.props.id);
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
        return <LoadingSpinner />;
    }

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>
                <TitleBar title="Aircraft" showMap link="/aircraft" />
                <NavBar />
                <MapView id={this.props.id} />
                {this.props.id ? (
                    <Box>{this.renderAircraft(this.props.aircraft)}</Box>
                ) : null}
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        aircraft: state.aircraft,
        id: ownProps.match.params.id
    };
}

const mapDispatchToProps = {
    fetchAircraft
};

export default connect(mapStateToProps, mapDispatchToProps)(AircraftMapPage);
