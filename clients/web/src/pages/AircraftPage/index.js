import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Flex } from "grid-styled";

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

import AircraftProvider from "../../containers/AircraftProvider";

const statusFilters = ["Any status", "On Mission", "OOS"];

class AircraftPage extends Component {
    constructor(props) {
        super(props);
    }

    renderAircraft(aircraft) {
        if (!aircraft.pending) {
            return aircraft.data.map(a => {
                return <AircraftListItem aircraft={a} key={a.id} />;
            });
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
                                        onChange={status => {
                                            if (status === "Any status") {
                                                status = "";
                                            }
                                            fetchAircraft(status);
                                        }}
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
                    renderDetailView={() => {
                        return <div />;
                    }}
                    renderMapView={() => {}}
                />
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default AircraftPage;
