// TODO: Replace Box with new invisible style solution
// (Ask Vincent)

import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";
import Media from "react-media";
import { withTheme } from "styled-components";
import ReactMapboxGl, { Layer, Feature } from "react-mapbox-gl";

import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import MissionList from "../../components/MissionList";
import FlexFullHeight from "../../components/FlexFullHeight";
import ScrollView from "../../components/ScrollView";
import Box from "../../components/Box";

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

                <Flex style={{ flex: 1, overflowY: "hidden" }}>
                    <ScrollView w={[1, 1, 1 / 2]} maxWidth={[null, null, 400]}>
                        {this.renderMissions()}
                    </ScrollView>
                    <Media
                        query={`(min-width: ${
                            this.props.theme.breakpoints[1]
                        })`}
                        render={() => {
                            return (
                                <Map
                                    style="mapbox://styles/mapbox/streets-v9"
                                    containerStyle={{
                                        flex: 1
                                    }}
                                >
                                    <Layer
                                        type="symbol"
                                        id="marker"
                                        layout={{
                                            "icon-image": "marker-15"
                                        }}
                                    >
                                        <Feature
                                            coordinates={[
                                                -0.481747846041145,
                                                51.3233379650232
                                            ]}
                                        />
                                    </Layer>
                                </Map>
                            );
                        }}
                    />
                </Flex>
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default withTheme(MissionsPage);
