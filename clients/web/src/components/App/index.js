import React from "react";
import { Helmet } from "react-helmet";
import { Switch, Route, Redirect } from "react-router";
import Media from "react-media";
import { withTheme } from "styled-components";

import FlexFillVH from "../FlexFillVH";

import AircraftPage from "../../pages/AircraftPage";
import AircraftDetailPage from "../../pages/AircraftDetailPage";
import AircraftMapPage from "../../pages/AircraftMapPage";
import GroupsDetailPage from "../../pages/GroupsDetailPage";
import HomePage from "../../pages/HomePage";
import NotFoundPage from "../../pages/NotFoundPage";
import PeoplePage from "../../pages/PeoplePage";
import PeopleDetailPage from "../../pages/PeopleDetailPage";
import ProfilePage from "../../pages/ProfilePage";
import ResourcesPage from "../../pages/ResourcesPage";
import SignInPage from "../../pages/SignInPage";

const App = ({ theme: { breakpoints } }) => {
    return (
        <FlexFillVH flexDirection="column">
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
                                path="/aircraft/map/:id"
                                render={({ match }) => (
                                    <Redirect
                                        to={`/aircraft/${match.params.id}`}
                                    />
                                )}
                            />
                            <Redirect from="/aircraft/map" to="/aircraft" />
                            <Route
                                path="/aircraft/:id"
                                component={AircraftPage}
                            />
                            <Route path="/aircraft" component={AircraftPage} />
                            <Route path="/people/:id" component={PeoplePage} />
                            <Route path="/people" component={PeoplePage} />
                            <Route
                                path="/groups/:groupID/:id"
                                component={PeoplePage}
                            />
                            <Route
                                path="/groups/:groupID"
                                component={PeoplePage}
                            />
                            <Route path="/groups" component={PeoplePage} />
                            <Route path="/profile" component={ProfilePage} />
                            <Redirect from="/resources" to="/aircraft" />
                            <Route component={NotFoundPage} />
                        </Switch>
                    ) : (
                        <Switch>
                            <Route exact path="/" component={HomePage} />
                            <Route path="/signin" component={SignInPage} />
                            <Route
                                path="/aircraft/map/:id"
                                component={AircraftMapPage}
                            />
                            <Route
                                path="/aircraft/map"
                                component={AircraftMapPage}
                            />
                            <Route
                                path="/aircraft/:id"
                                component={AircraftDetailPage}
                            />
                            <Route path="/aircraft" component={AircraftPage} />
                            <Route
                                path="/groups/:groupID"
                                component={GroupsDetailPage}
                            />
                            <Route path="/groups" component={PeoplePage} />
                            <Route
                                path="/people/:id"
                                component={PeopleDetailPage}
                            />
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
        </FlexFillVH>
    );
};

export default withTheme(App);
