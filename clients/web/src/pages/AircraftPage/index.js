import React, { Component } from "react";
import { Helmet } from "react-helmet";

import FlexFullHeight from "../../components/FlexFullHeight";
import Toolbar from "../../components/Toolbar";
import SearchBox from "../../components/SearchBox";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import Tab from "../../components/Tab";
import Tabs from "../../components/Tabs";
import ScrollView from "../../components/ScrollView";
import NavBar from "../../components/NavBar";

class AircraftPage extends Component {
    constructor(props) {
        super(props);
        this.state = { active: "availTab" };
    }

    render() {
        const content = {
            availTab: "Available Aircrafts",
            missionTab: "Aircrafts Currently on a Mission",
            oosTab: "Aircrafts Currently Out of Service"
        };
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
                    <Tabs
                        active={this.state.active}
                        onChange={active => this.setState({ active })}
                    >
                        <Tab key="availTab">AVAILABLE</Tab>
                        <Tab key="missionTab">ON MISSION</Tab>
                        <Tab key="oosTab">OOS</Tab>
                    </Tabs>
                    <p>{content[this.state.active]}</p>
                </ScrollView>
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default AircraftPage;
