import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import Footer from "../../components/Footer";
import FlexFullHeight from "../../components/FlexFullHeight";

import Navigation from "../Navigation";
import SignInPage from "../SignInPage";
import AircraftPage from "../AircraftPage";
import MissionsPage from "../MissionsPage";
import PersonnelPage from "../PersonnelPage";
import NotFoundPage from "../NotFoundPage";

class App extends Component {
    render() {
        return (
            <FlexFullHeight flexDirection="column">
                <Helmet
                    titleTemplate="%s - Airlift Northwest"
                    defaultTitle="Airlift Northwest"
                />
                <Navigation />
                <Switch>
                    <Route exact path="/" component={AircraftPage} />
                    <Route path="/missions" component={MissionsPage} />
                    <Route path="/personnel" component={PersonnelPage} />
                    <Route path="/signin" component={SignInPage} />
                    <Route component={NotFoundPage} />
                </Switch>
                <Footer />
            </FlexFullHeight>
        );
    }
}

export default App;
