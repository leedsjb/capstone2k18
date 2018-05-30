import { createStore, applyMiddleware, compose } from "redux";
import reduxThunkMiddleware from "redux-thunk";
import { routerMiddleware } from "react-router-redux";
// import logger from "redux-logger";

import socketMiddleware from "./middleware/socket";

import createReducer from "./reducers";

export default function configureStore(initialState = {}, history) {
    const middleware = [
        reduxThunkMiddleware,
        routerMiddleware(history),
        // logger,
        socketMiddleware
    ];

    const enhancers = [applyMiddleware(...middleware)];

    const store = createStore(
        createReducer(),
        initialState,
        compose(...enhancers)
    );

    return store;
}
