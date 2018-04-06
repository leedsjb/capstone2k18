import { combineReducers } from "redux";
import { routerReducer } from "react-router-redux";
import { reducer as formReducer } from "redux-form";

import signInReducer from "./signInReducer";
import missionsReducer from "./missionsReducer";

export default function createReducer(injectedReducers) {
    return combineReducers({
        router: routerReducer,
        form: formReducer,
        signIn: signInReducer,
        missions: missionsReducer,
        ...injectedReducers
    });
}
