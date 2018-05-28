import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import Media from "react-media";
import { push } from "react-router-redux";
import ReactMapboxGl, { Layer, Feature, ZoomControl } from "react-mapbox-gl";
import { withTheme } from "styled-components";

import Box from "../../components/Box";

import { fetchAircraft } from "../../actions/aircraft/actions";

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
    }

    componentWillReceiveProps(nextProps) {
        if (nextProps.id && nextProps.id !== this.props.id) {
            let aircraft = this.props.aircraft.data.find(air => {
                return air.id === Number(this.props.id);
            });

            this.setState({
                center: [aircraft.long, aircraft.lat]
            });
        }
    }

    componentDidUpdate(prevProps, prevState) {
        if (this.props.id !== prevProps.id && this.state.map) {
            this.state.map.resize();
            this.state.map.flyTo(this.mapCenter());
        }
    }

    fitMapBounds(selected) {
        let minLat = selected.mission.waypoints.reduce(
            (minLat, point) => (point.lat < minLat ? point.lat : minLat),
            selected.lat
        );
        let maxLat = selected.mission.waypoints.reduce(
            (maxLat, point) => (point.lat > maxLat ? point.lat : maxLat),
            selected.lat
        );
        let minLong = selected.mission.waypoints.reduce(
            (minLong, point) => (point.long > minLong ? point.long : minLong),
            selected.long
        );
        let maxLong = selected.mission.waypoints.reduce(
            (maxLong, point) => (point.long < maxLong ? point.long : maxLong),
            selected.long
        );
        this.state.map.fitBounds([[minLong, minLat], [maxLong, maxLat]], {
            padding: 32
        });
    }

    getActive(aircraftID) {
        let selAircraft = this.props.aircraft.data.find(air => {
            return air.id === aircraftID;
        });

        return selAircraft.mission.waypoints.find(point => {
            return point.active;
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

    isAircraftWithWaypoints() {
        let aircraft = this.props.aircraft.data.find(air => {
            return air.id === Number(this.props.id);
        });

        return (
            this.props.aircraft.data &&
            this.props.aircraft.data.length > 0 &&
            aircraft.mission &&
            aircraft.mission.waypoints.length > 0
        );
    }

    mapCenter = () => {
        let aircraft = this.props.aircraft.data.find(air => {
            return air.id === Number(this.props.id);
        });

        if (
            !this.props.aircraft.pending &&
            this.props.aircraft.data &&
            this.props.aircraft.data.length > 0 &&
            (!aircraft.mission ||
                (aircraft.mission && aircraft.mission.waypoints.length === 0))
        ) {
            return [aircraft.long, aircraft.lat];
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
            if (this.isAircraftWithWaypoints() && selected && this.state.map) {
                this.fitMapBounds(selected);
            }
            return (
                <div>
                    {this.props.aircraft.data.map(aircraft => {
                        return (
                            <Box key={aircraft.id}>
                                <Media
                                    query={`(min-width: ${
                                        this.props.theme.breakpoints[1]
                                    }`}
                                >
                                    {matches =>
                                        matches ? (
                                            <Layer
                                                type="symbol"
                                                layout={{
                                                    "icon-image": "airplane",
                                                    "icon-allow-overlap": true,
                                                    "icon-rotate":
                                                        aircraft.mission &&
                                                        aircraft.mission
                                                            .waypoints.length >
                                                            0
                                                            ? Math.atan2(
                                                                  this.getActive(
                                                                      aircraft.id
                                                                  ).long -
                                                                      aircraft.long,
                                                                  this.getActive(
                                                                      aircraft.id
                                                                  ).lat -
                                                                      aircraft.lat
                                                              ) *
                                                              180 /
                                                              Math.PI
                                                            : 0,
                                                    "text-field":
                                                        aircraft.callsign,
                                                    "text-allow-overlap": true,
                                                    "text-anchor": "bottom",
                                                    "text-offset": [0, -1],
                                                    "text-transform":
                                                        "uppercase"
                                                }}
                                                paint={{
                                                    "icon-opacity":
                                                        !selected ||
                                                        (selected &&
                                                            aircraft.id ===
                                                                selected.id)
                                                            ? 1
                                                            : 0.35,
                                                    "text-opacity":
                                                        !selected ||
                                                        (selected &&
                                                            aircraft.id ===
                                                                selected.id)
                                                            ? 1
                                                            : 0.35
                                                }}
                                            >
                                                <Feature
                                                    coordinates={[
                                                        aircraft.long,
                                                        aircraft.lat
                                                    ]}
                                                    onClick={() =>
                                                        this.props.push(
                                                            `/aircraft/${
                                                                aircraft.id
                                                            }`
                                                        )
                                                    }
                                                    onMouseEnter={map => {
                                                        this.state.map.getCanvas().style.cursor =
                                                            "pointer";
                                                    }}
                                                    onMouseLeave={map => {
                                                        this.state.map.getCanvas().style.cursor =
                                                            "";
                                                    }}
                                                />
                                            </Layer>
                                        ) : (
                                            <Layer
                                                type="symbol"
                                                layout={{
                                                    "icon-image": "airplane",
                                                    "icon-allow-overlap": true,
                                                    "text-field":
                                                        aircraft.callsign,
                                                    "text-allow-overlap": true,
                                                    "text-anchor": "bottom",
                                                    "text-offset": [0, -1],
                                                    "text-transform":
                                                        "uppercase"
                                                }}
                                            >
                                                <Feature
                                                    coordinates={[
                                                        aircraft.long,
                                                        aircraft.lat
                                                    ]}
                                                    onClick={() =>
                                                        this.props.push(
                                                            `/aircraft/map/${
                                                                aircraft.id
                                                            }`
                                                        )
                                                    }
                                                    onMouseEnter={map => {
                                                        this.state.map.getCanvas().style.cursor =
                                                            "pointer";
                                                    }}
                                                    onMouseLeave={map => {
                                                        this.state.map.getCanvas().style.cursor =
                                                            "";
                                                    }}
                                                />
                                            </Layer>
                                        )
                                    }
                                </Media>

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
                                {this.isAircraftWithWaypoints() && selected ? (
                                    <Box>
                                        {selected.mission.waypoints.map(
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
                                                    ...selected.mission.waypoints.map(
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
                    zoom={!this.isAircraftWithWaypoints() ? [10] : undefined}
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
        aircraft: state.aircraft
    };
}

const mapDispatchToProps = {
    fetchAircraft,
    push
};

export default connect(mapStateToProps, mapDispatchToProps)(withTheme(MapView));
