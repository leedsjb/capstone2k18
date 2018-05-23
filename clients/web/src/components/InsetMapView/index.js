import React, { Component } from "react";
import { connect } from "react-redux";
import { push } from "react-router-redux";
import ReactMapboxGl, { Layer, Feature } from "react-mapbox-gl";
import { withTheme } from "styled-components";

import Box from "../../components/Box";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

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
        if (this.state.map) {
            this.state.map.fitBounds(
                [
                    [
                        Math.min(
                            this.props.aircraftDetail.data.long,
                            -122.28567
                        ),
                        Math.min(this.props.aircraftDetail.data.lat, 47.552965)
                    ],
                    [
                        Math.max(
                            this.props.aircraftDetail.data.long,
                            -122.28567
                        ),
                        Math.max(this.props.aircraftDetail.data.lat, 47.552965)
                    ]
                ],
                { padding: { top: 20, bottom: 20, left: 15, right: 15 } }
            );
        }
    }

    mapCenter = () => {
        if (
            !this.props.aircraftDetail.pending &&
            !Array.isArray(this.props.aircraftDetail.data)
        ) {
            // return [
            //     this.props.aircraftDetail.data.long,
            //     this.props.aircraftDetail.data.lat
            // ];
            return [-122.28567, 47.552965];
        }
        return [-122.4821475, 47.6129432];
    };

    renderMapView = () => {
        if (
            !this.props.aircraft.pending &&
            this.props.aircraft.data &&
            this.props.aircraft.data.length > 0
        ) {
            let selected = this.props.aircraft.data.find(air => {
                return air.id == this.props.id;
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
                            <Layer
                                type="symbol"
                                layout={{
                                    "icon-image": "circle-15",
                                    "text-field": "Squaxin Ballfields",
                                    "text-anchor": "top",
                                    "text-offset": [0, 0.5],
                                    "text-size": 10,
                                    "text-transform": "uppercase"
                                }}
                            >
                                <Feature
                                    coordinates={[-122.28567, 47.552965]}
                                />
                            </Layer>
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
                    width: "80%",
                    height: "25%",
                    borderRadius: "8px",
                    margin: "0 auto"
                }}
                center={this.mapCenter()}
                onClick={() =>
                    this.props.push(`/aircraft/map/${this.props.id}`)
                }
                onStyleLoad={map => this.setState({ map })}
                style="mapbox://styles/vincentmvdm/cjga7b9nz28b82st2j6jhwu91"
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
