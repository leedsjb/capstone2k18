import { createStore, applyMiddleware, compose } from "redux";
import reduxThunkMiddleware from "redux-thunk";
import { routerMiddleware } from "react-router-redux";
import logger from "redux-logger";

import createReducer from "./reducers";

// TODO: Change to createReducer
export default (initialState = {}, history) => {
    const middleware = [reduxThunkMiddleware, routerMiddleware(history)];

    const enhancers = [applyMiddleware(...middleware)];

    const store = createStore(
        createrReducer(),
        initialState,
        composeEnhancers(...enhancers)
    );

    return store;
};
