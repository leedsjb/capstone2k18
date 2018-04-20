import React, { Component } from "react";
import { connect } from "react-redux";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";
import { Link } from "react-router-dom";

import AircraftListItem from "../../components/AircraftListItem";
import AircraftDetailListItem from "../../components/AircraftDetailListItem";
import Box from "../../components/Box";
import Divider from "../../components/Divider";
import DropdownSelect from "../../components/DropdownSelect";
import FlexFullHeight from "../../components/FlexFullHeight";
import SearchBox from "../../components/SearchBox";
import MasterDetailMapView from "../../components/MasterDetailMapView";
import NavBar from "../../components/NavBar";
import NavBarItem from "../../components/NavBarItem";
import TabBar from "../../components/TabBar";
import TitleBar from "../../components/TitleBar";

import { fetchAircraft } from "../../actions/aircraft/actions";
import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

const statusFilters = ["Any status", "On Mission", "OOS"];

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
                    <Link to={`/aircraft/${a.id}`}>
                        <AircraftListItem aircraft={a} key={a.id} />
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
                    <Flex alignItems="center" justifyContent="space-between">
                        <DropdownSelect
                            items={statusFilters}
                            onChange={status => {
                                if (status === "Any status") {
                                    status = "";
                                }
                            }}
                        />
                        SEARCH
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
            <FlexFullHeight flexDirection="column">
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>

                <TitleBar title="Aircraft" />
                <NavBar />

                <MasterDetailMapView
                    renderMasterView={this.renderMasterView}
                    renderDetailView={this.renderDetailView}
                    renderMapView={() => {}}
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
            </FlexFullHeight>
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
    fetchAircraftDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(AircraftPage);
