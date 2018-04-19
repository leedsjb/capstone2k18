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

import AircraftProvider from "../../containers/AircraftProvider";

const statusFilters = ["Any status", "On Mission", "OOS"];

class AircraftPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            status: ""
        };
    }

    setStatus = status => {
        this.setState({
            status
        });
    };

    renderMasterView = () => {
        return (
            <div>
                <NavBarItem title="Aircraft" path="/aircraft" />
                <NavBarItem title="People" path="/people" />

                <Flex>
                    <DropdownSelect
                        items={statusFilters}
                        onChange={this.setStatus}
                    />
                </Flex>
                <Box px={3}>
                    <SearchBox />
                </Box>

                <AircraftProvider
                    status={this.state.status}
                    render={test => {
                        console.log(test);
                        return <div />;
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
