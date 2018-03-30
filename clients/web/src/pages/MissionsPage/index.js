// TODO: Replace Box with new invisible style solution
// (Ask Vincent)

import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import ReactMapboxGl, { Layer, Feature } from "react-mapbox-gl";

import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import Toolbar from "../../components/Toolbar";
import ButtonDropdown from "../../components/ButtonDropdown";
import MissionList from "../../components/MissionList";
import FlexFullHeight from "../../components/FlexFullHeight";
import ScrollView from "../../components/ScrollView";

const Map = ReactMapboxGl({
    accessToken: process.env.REACT_APP_MAPBOX
});

class MissionsPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            listView: true
        };
    }

    toggleListView = () => {
        this.setState({
            listView: !this.state.listView
        });
    };

    renderMissions() {
        return this.state.listView ? (
            <MissionList />
        ) : (
            <Map
                style="mapbox://styles/mapbox/streets-v9"
                containerStyle={{
                    flex: 1
                }}
            >
                <Layer
                    type="symbol"
                    id="marker"
                    layout={{ "icon-image": "marker-15" }}
                >
                    <Feature
                        coordinates={[-0.481747846041145, 51.3233379650232]}
                    />
                </Layer>
            </Map>
        );
    }

    render() {
        return (
            <FlexFullHeight flexDirection="column">
                <Helmet>
                    <title>Missions</title>
                </Helmet>

                <TitleBar
                    title="Missions"
                    icon="map"
                    iconOnClick={this.toggleListView}
                />

                <ScrollView>{this.renderMissions()}</ScrollView>
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default MissionsPage;
