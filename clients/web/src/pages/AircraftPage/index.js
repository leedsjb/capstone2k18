import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";
import { connect } from "react-redux";

import FlexFullHeight from "../../components/FlexFullHeight";
import SearchBox from "../../components/SearchBox";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";
import Box from "../../components/Box";
import NavBar from "../../components/NavBar";
import MasterDetailMapView from "../../components/MasterDetailMapView";
import DropdownSelect from "../../components/DropdownSelect";
import NavBarItem from "../../components/NavBarItem";
import AircraftListItem from "../../components/AircraftListItem";
import Divider from "../../components/Divider";
import AircraftDetailListItem from "../../components/AircraftDetailListItem";

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

    componentWillReceiveProps(nextProps) {}

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
        console.log(aircraftDetail);
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
        return <div>{this.renderAircraftDetail(aircraftDetail)}</div>;
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
