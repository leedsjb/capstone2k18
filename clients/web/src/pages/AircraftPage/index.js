import React, { Component } from "react";
import { Helmet } from "react-helmet";

import FlexFullHeight from "../../components/FlexFullHeight";
import Toolbar from "../../components/Toolbar";
import Border from "../../components/Border";
import SearchBox from "../../components/SearchBox";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import ScrollView from "../../components/ScrollView";

class AircraftPage extends Component {
    render() {
        return (
            <FlexFullHeight flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>

                <TitleBar />
                <Toolbar>Test</Toolbar>
                <ScrollView>
                    <SearchBox />
                    Aircraft page
                </ScrollView>
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default AircraftPage;
