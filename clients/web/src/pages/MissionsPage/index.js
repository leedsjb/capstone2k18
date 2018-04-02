// TODO: Replace Box with new invisible style solution
// (Ask Vincent)

import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import { connect } from "react-redux";
import { withTheme } from "styled-components";
import ReactMapboxGl, { Layer, Feature } from "react-mapbox-gl";

import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import Box from "../../components/Box";
import MissionList from "../../components/MissionList";
import NavBar from "../../components/NavBar";
import FlexFullHeight from "../../components/FlexFullHeight";
import MasterDetailView from "../../components/MasterDetailView";
import MasterView from "../../components/MasterView";
import DetailView from "../../components/DetailView";
import ButtonDropdown from "../../components/ButtonDropdown";
import Relative from "../../components/Relative";
import Absolute from "../../components/Absolute";
import ScrollView from "../../components/ScrollView";
import Heading from "../../components/Heading";
import Measure from "../../components/Measure";

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

    componentWillReceiveProps(nextProps) {
        if (nextProps !== this.props) {
            console.log(this.props.match.params.id);
        }
    }

    toggleListView = () => {
        this.setState({
            listView: !this.state.listView
        });
    };

    renderMissions() {
        return this.state.listView ? (
            <div>
                <Flex justifyContent="center">
                    <ButtonDropdown>Ongoing</ButtonDropdown>
                    <ButtonDropdown>Any aircraft</ButtonDropdown>
                </Flex>
                <MissionList />
            </div>
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

                <NavBar />

                <MasterDetailView>
                    <MasterView>{this.renderMissions()}</MasterView>
                    <DetailView>
                        <Flex
                            style={{ height: "100%", width: "100%" }}
                            flexDirection={[null, null, "column", "row"]}
                        >
                            <ScrollView
                                maxWidth={[null, null, null, 320]}
                                height={[null, null, "50%", "100%"]}
                            >
                                <Box p={3}>
                                    <Heading is="h3">Information</Heading>
                                    <Measure>
                                        Integer posuere erat a ante venenatis
                                        dapibus posuere velit aliquet. Donec
                                        ullamcorper nulla non metus auctor
                                        fringilla. Cum sociis natoque penatibus
                                        et magnis dis parturient montes,
                                        nascetur ridiculus mus. Duis mollis, est
                                        non commodo luctus, nisi erat porttitor
                                        ligula, eget lacinia odio sem nec elit.
                                        Lorem ipsum dolor sit amet, consectetur
                                        adipiscing elit. Morbi leo risus, porta
                                        ac consectetur ac, vestibulum at eros.
                                        Fusce dapibus, tellus ac cursus commodo,
                                        tortor mauris condimentum nibh, ut
                                        fermentum massa justo sit amet risus.
                                    </Measure>
                                    <Heading is="h3">Information</Heading>
                                    <Measure>
                                        Integer posuere erat a ante venenatis
                                        dapibus posuere velit aliquet. Donec
                                        ullamcorper nulla non metus auctor
                                        fringilla. Cum sociis natoque penatibus
                                        et magnis dis parturient montes,
                                        nascetur ridiculus mus. Duis mollis, est
                                        non commodo luctus, nisi erat porttitor
                                        ligula, eget lacinia odio sem nec elit.
                                        Lorem ipsum dolor sit amet, consectetur
                                        adipiscing elit. Morbi leo risus, porta
                                        ac consectetur ac, vestibulum at eros.
                                        Fusce dapibus, tellus ac cursus commodo,
                                        tortor mauris condimentum nibh, ut
                                        fermentum massa justo sit amet risus.
                                    </Measure>
                                    <Heading is="h3">Information</Heading>
                                    <Measure>
                                        Integer posuere erat a ante venenatis
                                        dapibus posuere velit aliquet. Donec
                                        ullamcorper nulla non metus auctor
                                        fringilla. Cum sociis natoque penatibus
                                        et magnis dis parturient montes,
                                        nascetur ridiculus mus. Duis mollis, est
                                        non commodo luctus, nisi erat porttitor
                                        ligula, eget lacinia odio sem nec elit.
                                        Lorem ipsum dolor sit amet, consectetur
                                        adipiscing elit. Morbi leo risus, porta
                                        ac consectetur ac, vestibulum at eros.
                                        Fusce dapibus, tellus ac cursus commodo,
                                        tortor mauris condimentum nibh, ut
                                        fermentum massa justo sit amet risus.
                                    </Measure>
                                </Box>
                            </ScrollView>

                            <Flex flex={1}>
                                <Map
                                    style="mapbox://styles/mapbox/streets-v9"
                                    containerStyle={{
                                        width: "100%",
                                        height: "100%"
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
                            </Flex>
                        </Flex>
                    </DetailView>
                </MasterDetailView>

                <TabBar />
            </FlexFullHeight>
        );
    }
}

const mapStateToProps = ({ router }) => {
    return {
        router
    };
};

export default withTheme(connect(mapStateToProps)(MissionsPage));
