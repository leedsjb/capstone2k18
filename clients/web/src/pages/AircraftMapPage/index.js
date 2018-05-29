import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { connect } from "react-redux";

import FlexFillVH from "../../components/FlexFillVH";
import MobileMapView from "../../components/MobileMapView";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";

class AircraftMapPage extends Component {
    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>
                <TitleBar title="Aircraft" showMap link="/aircraft" />
                <NavBar />
                <MobileMapView aircraftID={this.props.aircraftID} />
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    return {
        aircraftID: ownProps.match.params.id
    };
}

export default connect(mapStateToProps, null)(AircraftMapPage);
