import { combineReducers } from "redux";
import { routerReducer } from "react-router-redux";
import { reducer as formReducer } from "redux-form";

import signInReducer from "./signInReducer";

export default function createReducer(injectedReducers) {
    return combineReducers({
        router: routerReducer,
        form: formReducer,
        signIn: signInReducer,
        ...injectedReducers
    });
}
