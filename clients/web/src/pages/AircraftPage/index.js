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
import Divider from "../../components/Divider";

import AircraftProvider from "../../containers/AircraftProvider";

const statusFilters = ["Any status", "On Mission", "OOS"];

class AircraftPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isSearching: false
        };
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
                <AircraftProvider
                    render={({ fetchAircraft, aircraft }) => {
                        return (
                            <div>
                                <Box px={3} py={2}>
                                    <Flex
                                        alignItems="center"
                                        justifyContent="space-between"
                                    >
                                        <DropdownSelect
                                            items={statusFilters}
                                            onChange={status => {
                                                if (status === "Any status") {
                                                    status = "";
                                                }
                                                fetchAircraft(status);
                                            }}
                                        />
                                        SEARCH
                                    </Flex>
                                </Box>

                                <Divider />
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
