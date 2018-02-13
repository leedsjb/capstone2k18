import React, { Component } from "react";
import { Helmet } from "react-helmet";
import { Switch, Route } from "react-router";

import Footer from "../../components/Footer";

import Header from "../Header";
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
                <Header />
                <Switch>
                    <Route exact path="/" component={SignInPage} />
                    <Route path="/aircrafts" component={AircraftsPage} />
                    <Route component={NotFoundPage} />
                </Switch>
                <Footer />
            </div>
        );
    }
}

export default App;
