import { combineReducers } from "redux";
import { routerReducer } from "react-router-redux";

import loginReducer from "./containers/LoginPage/reducer";

export default function createReducer(injectedReducers) {
    return combineReducers({
        router: routerReducer,
        login: loginReducer,
        ...injectedReducers
    });
}
