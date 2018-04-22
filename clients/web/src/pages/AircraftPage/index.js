import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import { Link } from "react-router-dom";
import { Layer, Feature, Popup } from "react-mapbox-gl";
import { push } from "react-router-redux";

import AircraftListItem from "../../components/AircraftListItem";
import AircraftDetailListItem from "../../components/AircraftDetailListItem";
import Box from "../../components/Box";
import Divider from "../../components/Divider";
import DropdownSelect from "../../components/DropdownSelect";
import FlexFillVH from "../../components/FlexFillVH";
import MasterDetailMapView from "../../components/MasterDetailMapView";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";
import SearchBox from "../../components/SearchBox";
import Span from "../../components/Span";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

import airplane from "../../images/airplane.svg";

const statusFilters = ["Any status", "On Mission", "OOS"];

const image = new Image(32, 32);
image.src = airplane;

class AircraftPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isSearching: false
        };
    }

    componentDidMount() {
        this.props.fetchAircraft();
        if (this.props.match.params.id) {
            this.props.fetchAircraftDetail(this.props.match.params.id);
        }
    }

    componentWillReceiveProps(nextProps) {
        if (
            nextProps.match.params.id &&
            !(nextProps.match.params.id === this.props.match.params.id)
        ) {
            this.props.fetchAircraftDetail(nextProps.match.params.id);
        }
    }

    renderAircraft(aircraft) {
        if (!aircraft.pending) {
            return aircraft.data.map(a => {
                return (
                    <Link to={`/aircraft/${a.id}`} key={a.id}>
                        <AircraftListItem aircraft={a} />
                    </Link>
                );
            });
        }
    }

    renderAircraftDetail(aircraftDetail) {
        if (!aircraftDetail.pending) {
            return (
                <AircraftDetailListItem
                    aircraftDetail={aircraftDetail}
                    key={aircraftDetail.data.id}
                />
            );
        }
    }

    renderMasterView = () => {
        return (
            <div>
                <Box px={3} py={2}>
                    <SearchBox />
                    <Flex alignItems="center" mt={2}>
                        <DropdownSelect
                            items={statusFilters}
                            onChange={status => {
                                if (status === "Any status") {
                                    status = "";
                                }
                            }}
                        />
                        <Box ml={2}>
                            <DropdownSelect
                                items={statusFilters}
                                onChange={status => {
                                    if (status === "Any status") {
                                        status = "";
                                    }
                                }}
                            />
                        </Box>
                    </Flex>
                </Box>

                <Divider />
                {this.renderAircraft(this.props.aircraft)}
            </div>
        );
    };

    renderDetailView = () => {
        return (
            <div>{this.renderAircraftDetail(this.props.aircraftDetail)}</div>
        );
    };

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>

                <TitleBar title="Aircraft" />
                <NavBar />

                <MasterDetailMapView
                    renderMasterView={this.renderMasterView}
                    renderDetailView={this.renderDetailView}
                    renderMapView={() => {
                        if (
                            !this.props.aircraft.pending &&
                            this.props.aircraft.data.length > 0
                        ) {
                            return (
                                <div>
                                    {this.props.aircraft.data.map(aircraft => {
                                        const images = [
                                            aircraft.callsign,
                                            image
                                        ];

                                        return (
                                            <Layer
                                                type="symbol"
                                                layout={{
                                                    "icon-image":
                                                        aircraft.callsign
                                                }}
                                                images={images}
                                                key={aircraft.id}
                                            >
                                                <Feature
                                                    coordinates={[
                                                        aircraft.long,
                                                        aircraft.lat
                                                    ]}
                                                />
                                            </Layer>
                                        );
                                    })}
                                    {this.props.aircraft.data.map(aircraft => {
                                        return (
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
                                                        `/aircraft/${
                                                            aircraft.id
                                                        }`
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
                        return;
                    }}
                    showDetail={this.props.match.params.id}
                    mapCenter={() => {
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
                    }}
                />
                <TabBar />
            </FlexFillVH>
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

export default connect(mapStateToProps, mapDispatchToProps)(AircraftPage);
