import React, { Component } from "react";
import { Helmet } from "react-helmet";

import StyledMeasure from "../../components/StyledMeasure";

class NotFoundPage extends Component {
    render() {
        return (
            <div>
                <Helmet>
                    <title>Page not found</title>
                </Helmet>
                <div>Page not found</div>
            </div>
        );
    }
}

export default NotFoundPage;
