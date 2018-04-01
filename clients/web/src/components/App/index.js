import React from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";
import Media from "react-media";
import { withTheme } from "styled-components";

import FlexFullHeight from "../FlexFullHeight";
import NavBar from "../NavBar";

import SignInPage from "../../pages/SignInPage";
import MissionsPage from "../../pages/MissionsPage";
import MissionDetailPage from "../../pages/MissionDetailPage";
import AircraftPage from "../../pages/AircraftPage";
import AircraftDetailPage from "../../pages/AircraftDetailPage";
import PersonnelPage from "../../pages/PersonnelPage";
import NotFoundPage from "../../pages/NotFoundPage";

const App = ({ theme: { breakpoints } }) => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet
                titleTemplate="%s - Airlift Northwest"
                defaultTitle="Airlift Northwest"
            />
            <NavBar />
            <Media query={`(min-width: ${breakpoints[1]})`}>
                {matches =>
                    matches ? (
                        <Switch>
                            <Route exact path="/" component={SignInPage} />
                            <Route path="/missions" component={MissionsPage} />
                            <Route path="/aircraft" component={AircraftPage} />
                            <Route
                                path="/personnel"
                                component={PersonnelPage}
                            />
                            <Route component={NotFoundPage} />
                        </Switch>
                    ) : (
                        <Switch>
                            <Route exact path="/" component={SignInPage} />
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
                            <Route
                                path="/personnel"
                                component={PersonnelPage}
                            />
                            <Route component={NotFoundPage} />
                        </Switch>
                    )
                }
            </Media>
        </FlexFullHeight>
    );
};

export default withTheme(App);
