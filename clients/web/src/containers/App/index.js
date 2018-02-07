import React from "react";
import { Router, Route } from "react-router";

const App = () => {
    return (
        <Router>
            <div>
                <Route
                    path="/vincentsucks"
                    component={() => {
                        return <p>Jessica sucks</p>;
                    }}
                />
            </div>
        </Router>
    );
};

export default App;
