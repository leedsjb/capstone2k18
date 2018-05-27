import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import Media from "react-media";
import { push } from "react-router-redux";
import ReactMapboxGl, {
    Layer,
    Feature,
    Popup,
    ZoomControl
} from "react-mapbox-gl";
import { withTheme } from "styled-components";

import Box from "../../components/Box";
import Span from "../../components/Span";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

// import mapStyle from "../../utils/mapbox/style.json";

const Map = ReactMapboxGl({
    accessToken: process.env.REACT_APP_MAPBOX
});

class MapView extends Component {
    constructor(props) {
        super(props);
        this.state = {
            center: [-122.4821475, 47.6129432],
            userPos: null,
            map: null
        };
    }

    componentDidMount() {
        this.props.fetchAircraft();
        if (this.props.id) {
            this.props.fetchAircraftDetail(this.props.id);
        }
    }

    componentWillReceiveProps(nextProps) {
        if (nextProps.id && nextProps.id !== this.props.id) {
            this.props.fetchAircraftDetail(nextProps.id);
            if (
                !this.props.aircraftDetail.error &&
                !this.props.aircraftDetail.pending &&
                !Array.isArray(this.props.aircraftDetail.data)
            ) {
                this.setState({
                    center: [
                        nextProps.aircraftDetail.data.long,
                        nextProps.aircraftDetail.data.lat
                    ]
                });
            }
        }
    }

    componentDidUpdate(prevProps, prevState) {
        if (this.props.id !== prevProps.id && this.state.map) {
            this.state.map.resize();
            this.state.map.flyTo(this.mapCenter());
        }
    }

    fitMapBounds(selected) {
        let minLat = this.props.aircraftDetail.data.mission.waypoints.reduce(
            (minLat, point) => (point.lat < minLat ? point.lat : minLat),
            selected.lat
        );
        let maxLat = this.props.aircraftDetail.data.mission.waypoints.reduce(
            (maxLat, point) => (point.lat > maxLat ? point.lat : maxLat),
            selected.lat
        );
        let minLong = this.props.aircraftDetail.data.mission.waypoints.reduce(
            (minLong, point) => (point.long > minLong ? point.long : minLong),
            selected.long
        );
        let maxLong = this.props.aircraftDetail.data.mission.waypoints.reduce(
            (maxLong, point) => (point.long < maxLong ? point.long : maxLong),
            selected.long
        );
        this.state.map.fitBounds([[minLong, minLat], [maxLong, maxLat]], {
            padding: 32
        });
    }

    getUserLocation() {
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(
                position => {
                    if (this.refs.aircraftPage) {
                        this.setState({
                            userPos: [
                                position.coords.longitude,
                                position.coords.latitude
                            ]
                        });
                        if (!this.props.id) {
                            this.setState({
                                center: [
                                    position.coords.longitude,
                                    position.coords.latitude
                                ]
                            });
                            this.state.map.setCenter([
                                position.coords.longitude,
                                position.coords.latitude
                            ]);
                        }
                    }
                },
                null,
                { enableHighAccuracy: true }
            );
        }
    }

    isSelAirWithWaypoints() {
        return (
            !this.props.aircraftDetail.error &&
            !this.props.aircraftDetail.pending &&
            !Array.isArray(this.props.aircraftDetail.data) &&
            this.props.aircraftDetail.data.mission &&
            this.props.aircraftDetail.data.mission.waypoints.length > 0
        );
    }

    mapCenter = () => {
        if (
            !this.props.aircraftDetail.error &&
            !this.props.aircraftDetail.pending &&
            !Array.isArray(this.props.aircraftDetail.data) &&
            (!this.props.aircraftDetail.data.mission ||
                (this.props.aircraftDetail.data.mission &&
                    this.props.aircraftDetail.data.mission.waypoints.length ===
                        0))
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
            this.props.aircraft.data &&
            this.props.aircraft.data.length > 0
        ) {
            let selected = this.props.aircraft.data.find(aircraft => {
                return aircraft.id === Number(this.props.id);
            });
            if (this.isSelAirWithWaypoints() && selected && this.state.map) {
                this.fitMapBounds(selected);
            }
            return (
                <div>
                    {this.props.aircraft.data.map(aircraft => {
                        return (
                            <Box key={aircraft.id}>
                                <Layer
                                    type="symbol"
                                    layout={{
                                        "icon-image": "airplane",
                                        "icon-allow-overlap": true
                                    }}
                                    paint={{
                                        "icon-opacity":
                                            selected &&
                                            aircraft.id === selected.id
                                                ? 1
                                                : 0.5
                                    }}
                                >
                                    <Feature
                                        coordinates={[
                                            aircraft.long,
                                            aircraft.lat
                                        ]}
                                    />
                                </Layer>
                                {this.state.userPos ? (
                                    <Layer
                                        type="symbol"
                                        layout={{
                                            "icon-image": "circle-15",
                                            "icon-allow-overlap": true
                                        }}
                                    >
                                        <Feature
                                            coordinates={this.state.userPos}
                                        />
                                    </Layer>
                                ) : null}
                                {this.isSelAirWithWaypoints() && selected ? (
                                    <Box>
                                        {this.props.aircraftDetail.data.mission.waypoints.map(
                                            point => {
                                                return (
                                                    <Layer
                                                        key={point.id}
                                                        type="symbol"
                                                        layout={{
                                                            "icon-image":
                                                                "circle-15",
                                                            "icon-allow-overlap": true,
                                                            "text-field":
                                                                point.name,
                                                            "text-allow-overlap": true,
                                                            "text-anchor":
                                                                "top",
                                                            "text-offset": [
                                                                0,
                                                                0.5
                                                            ],
                                                            "text-transform":
                                                                "uppercase"
                                                        }}
                                                    >
                                                        <Feature
                                                            coordinates={[
                                                                point.long,
                                                                point.lat
                                                            ]}
                                                        />
                                                    </Layer>
                                                );
                                            }
                                        )}
                                        <Layer type="line">
                                            <Feature
                                                coordinates={[
                                                    [
                                                        selected.long,
                                                        selected.lat
                                                    ],
                                                    ...this.props.aircraftDetail.data.mission.waypoints.map(
                                                        point => {
                                                            return [
                                                                point.long,
                                                                point.lat
                                                            ];
                                                        }
                                                    )
                                                ]}
                                            />
                                        </Layer>
                                    </Box>
                                ) : null}
                            </Box>
                        );
                    })}
                    {this.props.aircraft.data.map(aircraft => {
                        return (
                            <Media
                                query={`(min-width: ${
                                    this.props.theme.breakpoints[1]
                                }`}
                                key={aircraft.id}
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
                                            style={{
                                                cursor: "pointer",
                                                zIndex:
                                                    selected &&
                                                    aircraft.id === selected.id
                                                        ? 1
                                                        : 0
                                            }}
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
            <Flex flex={1} ref="aircraftPage">
                <Map
                    onStyleLoad={map =>
                        this.setState({ map }, () => {
                            this.getUserLocation();
                        })
                    }
                    // style={mapStyle}
                    style="mapbox://styles/tzchen/cjhl4cawj17o92rlazjfvmmmg"
                    containerStyle={{
                        width: "100%",
                        height: "100%"
                    }}
                    center={this.mapCenter()}
                    zoom={!this.isSelAirWithWaypoints() ? [10] : undefined}
                >
                    {this.renderMapView()}

                    <ZoomControl
                        onClick={() => {
                            this.state.map.zoomIn();
                        }}
                    />
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
