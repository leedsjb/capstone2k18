import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Page from "../../components/Page";
import Toolbar from "../../components/Toolbar";
import Border from "../../components/Border";
import SearchBox from "../../components/SearchBox";

class AircraftPage extends Component {
    render() {
        return (
            <Page>
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>
                <SearchBox />
                <Toolbar>Test</Toolbar>
                Aircraft page
            </Page>
        );
    }
}

export default AircraftPage;
