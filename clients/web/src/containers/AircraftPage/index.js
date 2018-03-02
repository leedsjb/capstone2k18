import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Toolbar from "../../components/Toolbar";
import Border from "../../components/Border";

class AircraftPage extends Component {
    render() {
        return (
            <div>
                <Helmet>
                    <title>Aircraft</title>
                </Helmet>
                <Toolbar>Test</Toolbar>
            </div>
        );
    }
}

export default AircraftPage;
