import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import Media from "react-media";
import { push } from "react-router-redux";
import ReactMapboxGl, { Layer, Feature, Popup } from "react-mapbox-gl";
import { withTheme } from "styled-components";

import MasterView from "../MasterView";
import MasterDetailView from "../MasterDetailView";
import Span from "../../components/Span";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

import airplane from "../../images/airplane.svg";

const image = new Image(32, 32);
image.src = airplane;

const Map = ReactMapboxGl({
    accessToken: process.env.REACT_APP_MAPBOX
});

class MapView extends Component {
    constructor(props) {
        super(props);
        this.state = {
            center: [-122.4821475, 47.6129432],
            map: null
        };
    }

    componentDidMount() {
        this.props.fetchAircraft();
    }

    componentDidUpdate(prevProps, prevState) {
        if (
            (!prevProps.id || !this.props.id) &&
            this.props.id !== prevProps.id &&
            this.state.map
        ) {
            this.setState({
                center: [
                    this.props.aircraftDetail.data.long,
                    this.props.aircraftDetail.data.lat
                ]
            });
            this.state.map.resize();
            this.state.map.setCenter(this.mapCenter());
        }
    }

    mapCenter = () => {
        if (
            !this.props.aircraftDetail.pending &&
            !Array.isArray(this.props.aircraftDetail.data)
        ) {
            return [
                this.props.aircraftDetail.data.long,
                this.props.aircraftDetail.data.lat
            ];
        }
        return this.state.center;
    };

    renderMapView = () => {
        if (
            !this.props.aircraft.pending &&
            this.props.aircraft.data.length > 0
        ) {
            return (
                <div>
                    {this.props.aircraft.data.map(aircraft => {
                        const images = [aircraft.callsign, image];

                        return (
                            <Layer
                                type="symbol"
                                layout={{
                                    "icon-image": "airplane"
                                }}
                                key={aircraft.id}
                            >
                                <Feature
                                    coordinates={[aircraft.long, aircraft.lat]}
                                />
                            </Layer>
                        );
                    })}
                    {this.props.aircraft.data.map(aircraft => {
                        return (
                            <Media
                                query={`(min-width: ${
                                    this.props.theme.breakpoints[1]
                                }`}
                            >
                                {matches =>
                                    matches ? (
                                        <Popup
                                            coordinates={[
                                                aircraft.long,
                                                aircraft.lat
                                            ]}
                                            key={aircraft.id}
                                            offset={{
                                                bottom: [0, -24]
                                            }}
                                            style={{ cursor: "pointer" }}
                                            onClick={() =>
                                                this.props.push(
                                                    `/aircraft/${aircraft.id}`
                                                )
                                            }
                                        >
                                            <Span fontWeight="bold">
                                                {aircraft.callsign}
                                            </Span>
                                        </Popup>
                                    ) : (
                                        <Popup
                                            coordinates={[
                                                aircraft.long,
                                                aircraft.lat
                                            ]}
                                            key={aircraft.id}
                                            offset={{
                                                bottom: [0, -24]
                                            }}
                                            style={{ cursor: "pointer" }}
                                            onClick={() =>
                                                this.props.push(
                                                    `/aircraft/map/${
                                                        aircraft.id
                                                    }`
                                                )
                                            }
                                        >
                                            <Span fontWeight="bold">
                                                {aircraft.callsign}
                                            </Span>
                                        </Popup>
                                    )
                                }
                            </Media>
                        );
                    })}
                </div>
            );
        }
        return <div />;
    };

    render() {
        return (
            <Flex flex={1}>
                <Map
                    onStyleLoad={map => this.setState({ map })}
                    style="mapbox://styles/vincentmvdm/cjga7b9nz28b82st2j6jhwu91"
                    containerStyle={{
                        width: "100%",
                        height: "100%"
                    }}
                    center={this.mapCenter()}
                >
                    {this.renderMapView()}
                </Map>
            </Flex>
        );
    }
}

function mapStateToProps(state) {
    return {
        aircraft: state.aircraft,
        aircraftDetail: state.aircraftDetail
    };
}

const mapDispatchToProps = {
    fetchAircraft,
    fetchAircraftDetail,
    push
};

export default connect(mapStateToProps, mapDispatchToProps)(withTheme(MapView));
