import React, { Component } from "react";
import { Helmet } from "react-helmet";

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
