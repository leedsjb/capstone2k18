// TODO: Replace Box with new invisible style solution
// (Ask Vincent)

import React, { Component } from "react";
import { Flex } from "grid-styled";
import { Helmet } from "react-helmet";

import Page from "../../components/Page";
import Box from "../../components/Box";
import Toolbar from "../../components/Toolbar";
import ButtonDropdown from "../../components/ButtonDropdown";
import MissionList from "../../components/MissionList";

class MissionsPage extends Component {
    render() {
        return (
            <Page>
                <Helmet>
                    <title>Missions</title>
                </Helmet>
                <Toolbar>
                    <Flex justifyContent="center">
                        <ButtonDropdown mr={1}>Ongoing</ButtonDropdown>
                        <ButtonDropdown ml={1}>Any aircraft</ButtonDropdown>
                    </Flex>
                </Toolbar>
                <MissionList />
            </Page>
        );
    }
}

export default MissionsPage;
