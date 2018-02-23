import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { ThemeProvider } from "styled-components";
import { Provider as RebassProvider } from "rebass";
import createHistory from "history/createBrowserHistory";
import { ConnectedRouter } from "react-router-redux";
import "sanitize.css/sanitize.css";

import App from "./containers/App";

import registerServiceWorker from "./registerServiceWorker";
import configureStore from "./configureStore";
import "./global-styles";

const initialState = {};
const history = createHistory();
const store = configureStore(initialState, history);

ReactDOM.render(
    <Provider store={store}>
        <ThemeProvider>
            <ConnectedRouter history={history}>
                <App />
            </ConnectedRouter>
        </ThemeProvider>
    </Provider>,
    document.getElementById("root")
);
registerServiceWorker();
