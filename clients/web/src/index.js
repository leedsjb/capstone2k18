import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { Provider as RebassProvider } from "rebass";
import createHistory from "history/createBrowserHistory";
import { ConnectedRouter } from "react-router-redux";

import registerServiceWorker from "./registerServiceWorker";
import configureStore from "./configureStore";
import "sanitize.css/sanitize.css";
import "./global-styles";

import App from "./containers/App";

const initialState = {};
const history = createHistory();
const store = configureStore(initialState, history);

ReactDOM.render(
    <Provider store={store}>
        <RebassProvider>
            <ConnectedRouter history={history}>
                <App />
            </ConnectedRouter>
        </RebassProvider>
    </Provider>,
    document.getElementById("root")
);
registerServiceWorker();
