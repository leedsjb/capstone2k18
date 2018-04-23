import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import ReactMapboxGl, { Layer, Feature, Popup } from "react-mapbox-gl";
import { push } from "react-router-redux";

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
            aircraft: null
        };
    }

    componentDidMount() {
        this.props.fetchAircraft();
    }

    getAircraft = (aircraft, e) => {
        this.setState({ aircraft: aircraft });
    };

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
        return [-122.4821475, 47.6129432];
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
                            <Popup
                                coordinates={[aircraft.long, aircraft.lat]}
                                key={aircraft.id}
                                offset={{
                                    bottom: [0, -24]
                                }}
                                style={{ cursor: "pointer" }}
                                onClick={() =>
                                    this.props.push(
                                        `/aircraft/map/${aircraft.id}`
                                    )
                                }
                            >
                                <Span fontWeight="bold">
                                    {aircraft.callsign}
                                </Span>
                            </Popup>
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

export default connect(mapStateToProps, mapDispatchToProps)(MapView);
