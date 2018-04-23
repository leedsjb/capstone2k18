import React, { Component } from "react";
import { Helmet } from "react-helmet";

import FlexFillVH from "../../components/FlexFillVH";
import MobileMapView from "../../components/MobileMapView";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";

const AircraftMapPage = props => {
    return (
        <FlexFillVH flexDirection="column">
            <Helmet>
                <title>Aircraft</title>
            </Helmet>
            <TitleBar title="Aircraft" showMap={true} link="/aircraft" />
            <NavBar />
            <MobileMapView aircraftID={props.match.params.id} />
            <TabBar />
        </FlexFillVH>
    );
};

export default AircraftMapPage;
