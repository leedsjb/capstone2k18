import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import FlexFullHeight from "../../components/FlexFullHeight";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";

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
                <TitleBar />
                <Switch>
                    <Route exact path="/" component={AircraftPage} />
                    <Route path="/missions" component={MissionsPage} />
                    <Route path="/personnel" component={PersonnelPage} />
                    <Route path="/signin" component={SignInPage} />
                    <Route component={NotFoundPage} />
                </Switch>
                <TabBar />
            </FlexFullHeight>
        );
    }
}

export default App;
