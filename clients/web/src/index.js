import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { Provider as RebassProvider } from "rebass";

import store from "./configureStore";
import "sanitize.css/sanitize.css";

import registerServiceWorker from "./registerServiceWorker";
import "./global-styles";

import App from "./containers/App";

ReactDOM.render(
    <Provider store={store}>
        <RebassProvider>
            <App />
        </RebassProvider>
    </Provider>,
    document.getElementById("root")
);
registerServiceWorker();
