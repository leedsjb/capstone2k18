import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import SignInPage from "../SignInPage";
import AircraftsPage from "../AircraftsPage";
import NotFoundPage from "../NotFoundPage";

class App extends Component {
    render() {
        return (
            <div>
                <Helmet
                    titleTemplate="%s - Airlift Northwest"
                    defaultTitle="Airlift Northwest"
                />
                <Switch>
                    <Route exact path="/" component={SignInPage} />
                    <Route path="/aircrafts" component={AircraftsPage} />
                    <Route component={NotFoundPage} />
                </Switch>
            </div>
        );
    }
}

export default App;
