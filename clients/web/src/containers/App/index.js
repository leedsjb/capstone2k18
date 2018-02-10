import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import LoginPage from "../LoginPage";
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
                    <Route exact path="/" component={LoginPage} />
                    <Route component={NotFoundPage} />
                </Switch>
            </div>
        );
    }
}

export default App;
