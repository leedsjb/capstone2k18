import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import FlexFullHeight from "../../components/FlexFullHeight";
import TitleBar from "../../components/TitleBar";
import TabBar from "../../components/TabBar";

import SignInPage from "../SignInPage";
import MissionsPage from "../MissionsPage";
import AircraftPage from "../AircraftPage";
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
                    <Route exact path="/" component={MissionsPage} />
                    <Route path="/aircraft" component={AircraftPage} />
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
