import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";
import { Link } from "react-router-dom";

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
import AircraftDetailListItem from "../../components/AircraftDetailListItem";

import AircraftProvider from "../../containers/AircraftProvider";
import AircraftDetailProvider from "../../containers/AircraftDetailProvider";

const statusFilters = ["Any status", "On Mission", "OOS"];

class AircraftPage extends Component {
    constructor(props) {
        super(props);
    }

    componentWillReceiveProps(nextProps) {
        if (this.props.match.params.id) {
            this.renderDetailView;
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
                <NavBarItem title="Aircraft" path="/aircraft" />
                <NavBarItem title="People" path="/people" />

                <AircraftProvider
                    render={({ fetchAircraft, aircraft }) => {
                        return (
                            <div>
                                <Flex>
                                    <DropdownSelect
                                        items={statusFilters}
                                        onChange={status =>
                                            fetchAircraft(status)
                                        }
                                    />
                                </Flex>
                                <Box px={3}>
                                    <SearchBox />
                                </Box>
                                {this.renderAircraft(aircraft)}
                            </div>
                        );
                    }}
                />
            </div>
        );
    };

    renderDetailView = () => {
        return (
            <AircraftDetailProvider
                id={this.props.match.params.id}
                render={({ aircraftDetail }) => {
                    console.log(aircraftDetail);
                    return (
                        <div>{this.renderAircraftDetail(aircraftDetail)}</div>
                    );
                }}
            />
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
                />
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default AircraftPage;
