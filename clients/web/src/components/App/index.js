import React from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";
import Media from "react-media";
import { withTheme } from "styled-components";

import FlexFullHeight from "../FlexFullHeight";
import NavBar from "../NavBar";
import TitleBar from "../TitleBar";

import SignInPage from "../../pages/SignInPage";
import MissionsPage from "../../pages/MissionsPage";
import MissionDetailPage from "../../pages/MissionDetailPage";
import AircraftPage from "../../pages/AircraftPage";
import AircraftDetailPage from "../../pages/AircraftDetailPage";
import PeoplePage from "../../pages/PeoplePage";
import NotFoundPage from "../../pages/NotFoundPage";
import HomePage from "../../pages/HomePage";

const App = ({ theme: { breakpoints } }) => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet
                titleTemplate="%s - Airlift Northwest"
                defaultTitle="Airlift Northwest"
            />

            <Media query={`(min-width: ${breakpoints[1]})`}>
                {matches =>
                    matches ? (
                        <Switch>
                            <Route exact path="/" component={HomePage} />
                            <Route path="/signin" component={SignInPage} />
                            <Route
                                path="/missions/:id"
                                component={MissionsPage}
                            />
                            <Route path="/missions" component={MissionsPage} />
                            <Route path="/aircraft" component={AircraftPage} />
                            <Route path="/people/:id" component={PeoplePage} />
                            <Route path="/people" component={PeoplePage} />
                            <Route component={NotFoundPage} />
                        </Switch>
                    ) : (
                        <Switch>
                            <Route exact path="/" component={HomePage} />
                            <Route path="/signin" component={SignInPage} />
                            <Route
                                path="/missions/:id"
                                component={MissionDetailPage}
                            />
                            <Route path="/missions" component={MissionsPage} />
                            <Route
                                path="/aircraft/:id"
                                component={AircraftDetailPage}
                            />
                            <Route path="/aircraft" component={AircraftPage} />
                            <Route path="/people" component={PeoplePage} />
                            <Route component={NotFoundPage} />
                        </Switch>
                    )
                }
            </Media>
        </FlexFullHeight>
    );
};

export default withTheme(App);
