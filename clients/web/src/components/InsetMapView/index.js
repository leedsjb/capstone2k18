import React, { Component } from "react";
import { connect } from "react-redux";
import { push } from "react-router-redux";
import ReactMapboxGl, { Layer, Feature } from "react-mapbox-gl";
import { withTheme } from "styled-components";

import Box from "../../components/Box";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

// import mapStyle from "../../utils/mapbox/style.json";

const Map = ReactMapboxGl({
    accessToken: process.env.REACT_APP_MAPBOX,
    interactive: false
});

class InsetMapView extends Component {
    constructor(props) {
        super(props);
        this.state = {
            map: null
        };
    }

    componentDidMount() {
        this.props.fetchAircraft();
        if (this.props.id) {
            this.props.fetchAircraftDetail(this.props.id);
        }
    }

    componentDidUpdate() {
        if (this.isSelAirWithWaypoints()) {
            let active = this.props.aircraftDetail.data.mission.waypoints.find(
                point => {
                    return point.active;
                }
            );
            this.state.map.fitBounds(
                [
                    [
                        Math.min(
                            this.props.aircraftDetail.data.long,
                            active.long
                        ),
                        Math.min(this.props.aircraftDetail.data.lat, active.lat)
                    ],
                    [
                        Math.max(
                            this.props.aircraftDetail.data.long,
                            active.long
                        ),
                        Math.max(this.props.aircraftDetail.data.lat, active.lat)
                    ]
                ],
                { padding: { top: 20, bottom: 20, left: 15, right: 15 } }
            );
        }
    }

    isSelAirWithWaypoints() {
        return (
            !this.props.aircraftDetail.error &&
            !this.props.aircraftDetail.pending &&
            !Array.isArray(this.props.aircraftDetail.data) &&
            this.props.aircraftDetail.data.mission &&
            this.props.aircraftDetail.data.mission.waypoints.length > 0 &&
            this.state.map
        );
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
        return [-122.4821475, 47.6129432]; // Center of Seattle
    };

    renderMapView = () => {
        if (
            !this.props.aircraft.pending &&
            this.props.aircraft.data &&
            this.props.aircraft.data.length > 0
        ) {
            let active;
            if (this.isSelAirWithWaypoints()) {
                active = this.props.aircraftDetail.data.mission.waypoints.find(
                    point => {
                        return point.active;
                    }
                );
            }
            let selected = this.props.aircraft.data.find(aircraft => {
                return aircraft.id === Number(this.props.id);
            });
            if (selected) {
                return (
                    <div>
                        <Box key={selected.id}>
                            <Layer
                                type="symbol"
                                layout={{
                                    "icon-image": "airplane",
                                    "icon-allow-overlap": true
                                }}
                            >
                                <Feature
                                    coordinates={[selected.long, selected.lat]}
                                />
                            </Layer>
                            {active ? (
                                <Layer
                                    type="symbol"
                                    layout={{
                                        "icon-image": "circle-15",
                                        "text-field": active.name,
                                        "text-anchor": "top",
                                        "text-offset": [0, 0.5],
                                        "text-size": 10,
                                        "text-transform": "uppercase"
                                    }}
                                >
                                    <Feature
                                        coordinates={[active.long, active.lat]}
                                    />
                                </Layer>
                            ) : null}
                        </Box>
                    </div>
                );
            }
        }
        return <div />;
    };

    render() {
        return (
            <Map
                animationOptions={{ animate: false }}
                containerStyle={{
                    width: "100%",
                    height: "100%",
                    borderRadius: "8px"
                }}
                center={this.mapCenter()}
                onClick={() =>
                    this.props.push(`/aircraft/map/${this.props.id}`)
                }
                onStyleLoad={map => this.setState({ map })}
                // style={mapStyle}
                style="mapbox://styles/tzchen/cjhl4cawj17o92rlazjfvmmmg"
            >
                {this.renderMapView()}
            </Map>
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

export default connect(mapStateToProps, mapDispatchToProps)(
    withTheme(InsetMapView)
);
