import React, { Component } from "react";
import { Helmet } from "react-helmet";

import FlexFullHeight from "../../components/FlexFullHeight";
import SearchBox from "../../components/SearchBox";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import Tab from "../../components/Tab";
import Tabs from "../../components/Tabs";
import NavBar from "../../components/NavBar";
import MasterDetailMapView from "../../components/MasterDetailMapView";

class AircraftPage extends Component {
    constructor(props) {
        super(props);
        this.state = { active: "availTab" };
    }

    renderMasterView = () => {
        const content = {
            availTab: "Available Aircrafts",
            missionTab: "Aircrafts Currently on a Mission",
            oosTab: "Aircrafts Currently Out of Service"
        };

        return (
            <div>
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
            </div>
        );
    };

    render() {
        return (
            <FlexFullHeight flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>

                <TitleBar title="Aircraft" />
                <NavBar />

                <MasterDetailMapView
                    renderMasterView={this.renderMasterView}
                    renderDetailView={() => {
                        <div />;
                    }}
                    renderMapView={() => {}}
                />
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default AircraftPage;
