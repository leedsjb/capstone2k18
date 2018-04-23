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
                <TitleBar title="Aircraft" showMap={true} link="/aircraft" />
                <NavBar />
                <MobileMapView id={this.props.id} />
                <TabBar />
            </FlexFillVH>
        );
    }
}

function mapStateToProps(state, ownProps) {
    console.log("state", ownProps);
    return {
        id: ownProps.match.params.id
    };
}

export default connect(mapStateToProps, null)(AircraftMapPage);
