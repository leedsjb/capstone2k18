import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { Provider as RebassProvider } from "rebass";

import "sanitize.css/sanitize.css";

import registerServiceWorker from "./registerServiceWorker";
import "./global-styles";

import App from "./containers/App";

ReactDOM.render(
    <RebassProvider>
        <App />
    </RebassProvider>,
    document.getElementById("root")
);
registerServiceWorker();
