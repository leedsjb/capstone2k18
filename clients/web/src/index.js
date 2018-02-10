import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { Provider as RebassProvider } from "rebass";
import createHistory from "history/createBrowserHistory";

import registerServiceWorker from "./registerServiceWorker";
import store from "./configureStore";
import "sanitize.css/sanitize.css";
import "./global-styles";

const history = createHistory();

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
