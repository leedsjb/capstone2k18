import React from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import FlexFullHeight from "../../components/FlexFullHeight";

import SignInPage from "../../pages/SignInPage";
import MissionsPage from "../../pages/MissionsPage";
import MissionDetailPage from "../../pages/MissionDetailPage";
import AircraftPage from "../../pages/AircraftPage";
import AircraftDetailPage from "../../pages/AircraftDetailPage";
import PersonnelPage from "../../pages/PersonnelPage";
import NotFoundPage from "../../pages/NotFoundPage";

const App = () => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet
                titleTemplate="%s - Airlift Northwest"
                defaultTitle="Airlift Northwest"
            />
            <Switch>
                <Route exact path="/" component={SignInPage} />
                <Route path="/missions/:id" component={MissionDetailPage} />
                <Route path="/missions" component={MissionsPage} />
                <Route path="/aircraft/:id" component={AircraftDetailPage} />
                <Route path="/aircraft" component={AircraftPage} />
                <Route path="/personnel" component={PersonnelPage} />
                <Route component={NotFoundPage} />
            </Switch>
        </FlexFullHeight>
    );
};

export default App;
