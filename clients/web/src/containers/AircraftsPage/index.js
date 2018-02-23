import React, { Component } from "react";
import { Helmet } from "react-helmet";

import StyledMeasure from "../../components/StyledMeasure";

class AircraftsPage extends Component {
    render() {
        return (
            <div>
                <Helmet>
                    <title>Aircrafts</title>
                </Helmet>
                <div>Aircrafts page</div>
            </div>
        );
    }
}

export default AircraftsPage;
