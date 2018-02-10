import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import SignUpPage from "../SignUpPage";
import NotFoundPage from "../NotFoundPage";

class App extends Component {
    render() {
        return (
            <div>
                <Helmet
                    titleTemplate="%s - Airlift Northwest App"
                    defaultTitle="Airlift Northwest App"
                />
                <Switch>
                    <Route exact path="/" component={SignUpPage} />
                    <Route component={NotFoundPage} />
                </Switch>
            </div>
        );
    }
}

export default App;
