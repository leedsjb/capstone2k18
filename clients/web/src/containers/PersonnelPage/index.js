import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Page from "../../components/Page";

class PersonnelPage extends Component {
    render() {
        return (
            <Page>
                <Helmet>
                    <title>Personnel</title>
                </Helmet>
                <div>Personnel Page</div>
            </Page>
        );
    }
}

export default PersonnelPage;
