import React from "react";
import { Helmet } from "react-helmet";

import FlexFullHeight from "../../components/FlexFullHeight";
import Toolbar from "../../components/Toolbar";
import SearchBox from "../../components/SearchBox";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";
import NavBar from "../../components/NavBar";

const AircraftPage = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet>
                <title>Aircraft</title>
            </Helmet>

            <TitleBar title="Aircraft" />
            <NavBar />

            <Toolbar>Test</Toolbar>
            <ScrollView>
                <SearchBox />
                Aircraft page
            </ScrollView>
            <TabBar />
        </FlexFullHeight>
    );
};

export default AircraftPage;
