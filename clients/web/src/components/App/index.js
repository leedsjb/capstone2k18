import React from "react";
import { Helmet } from "react-helmet";
import { Switch, Route, Redirect } from "react-router";
import Media from "react-media";
import { withTheme } from "styled-components";

import FlexFullHeight from "../FlexFullHeight";

import SignInPage from "../../pages/SignInPage";
import AircraftPage from "../../pages/AircraftPage";
import AircraftDetailPage from "../../pages/AircraftDetailPage";
import PeoplePage from "../../pages/PeoplePage";
import NotFoundPage from "../../pages/NotFoundPage";
import HomePage from "../../pages/HomePage";
import ProfilePage from "../../pages/ProfilePage";
import ResourcesPage from "../../pages/ResourcesPage";

const App = ({ theme: { breakpoints } }) => {
    return (
        <FlexFullHeight flexDirection="column">
            <Helmet
                titleTemplate="%s - AirliftNW Elevate"
                defaultTitle="AirliftNW Elevate"
            />

            <Media query={`(min-width: ${breakpoints[1]})`}>
                {matches =>
                    matches ? (
                        <Switch>
                            <Route exact path="/" component={HomePage} />
                            <Route path="/signin" component={SignInPage} />
                            <Route
                                path="/aircraft/:id"
                                component={AircraftPage}
                            />
                            <Route path="/aircraft" component={AircraftPage} />
                            <Route path="/people/:id" component={PeoplePage} />
                            <Route path="/people" component={PeoplePage} />
                            <Route path="/profile" component={ProfilePage} />
                            <Redirect from="/resources" to="/aircraft" />
                            <Route component={NotFoundPage} />
                        </Switch>
                    ) : (
                        <Switch>
                            <Route exact path="/" component={HomePage} />
                            <Route path="/signin" component={SignInPage} />
                            <Route
                                path="/aircraft/:id"
                                component={AircraftDetailPage}
                            />
                            <Route path="/aircraft" component={AircraftPage} />
                            <Route path="/people" component={PeoplePage} />
                            <Route path="/profile" component={ProfilePage} />
                            <Route
                                path="/resources"
                                component={ResourcesPage}
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
