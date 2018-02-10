import { combineReducers } from "redux";
import { routerReducer } from "react-router-redux";

import signUpReducer from "./containers/SignUpPage/reducer";

export default function createReducer(injectedReducers) {
    return combineReducers({
        router: routerReducer,
        signUp: signUpReducer,
        ...injectedReducers
    });
}
