import React, { Component } from "react";
import { Helmet } from "react-helmet";

class App extends Component {
    render() {
        return (
            <div>
                <Helmet
                    titleTemplate="%s - Airlift Northwest App"
                    defaultTitle="Airlift Northwest App"
                />
                <div>This is the App wrapper</div>
            </div>
        );
    }
}

export default App;
