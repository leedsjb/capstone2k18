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
import Heading from "../../components/Heading";
import MasterDetailMapView from "../../components/MasterDetailMapView";
import NavBar from "../../components/NavBar";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";
import Text from "../../components/Text";
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
            nextProps.match.params.id !== this.props.match.params.id
        ) {
            this.props.fetchAircraftDetail(nextProps.match.params.id);
        }
    }

    renderAircraft(aircraft) {
        if (!aircraft.pending && aircraft.data.length > 0) {
            return aircraft.data.map(a => {
                return (
                    <Link to={`/aircraft/${a.id}`} key={a.id}>
                        <AircraftListItem
                            aircraft={a}
                            active={this.props.match.params.id == a.id}
                        />
                    </Link>
                );
            });
        } else if (!aircraft.pending) {
            return (
                <Box mt={4}>
                    <Heading is="h2" textAlign="center" fontSize={4}>
                        No aircraft
                    </Heading>
                    <Text textAlign="center">Empty state text</Text>
                </Box>
            );
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
                                    status = null;
                                }
                                this.props.fetchAircraft(status);
                            }}
                        />
                    </Flex>
                </Box>

                <Divider />
                {this.renderAircraft(this.props.aircraft)}
            </div>
        );
    };

    renderDetailView = () => {
        return (
            <div>
                <Span onClick={() => this.props.push("/aircraft")}>CLOSE</Span>
                {this.renderAircraftDetail(this.props.aircraftDetail)}
            </div>
        );
    };

    render() {
        return (
            <FlexFillVH flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>
                <TitleBar
                    title="Aircraft"
                    showMap={true}
                    link="/aircraft/map"
                />
                <NavBar />
                <MasterDetailMapView
                    renderMasterView={this.renderMasterView}
                    renderDetailView={this.renderDetailView}
                    showDetail={this.props.match.params.id}
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
