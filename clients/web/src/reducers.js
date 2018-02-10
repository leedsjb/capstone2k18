import { combineReducers } from "redux";
import { routerReducer } from "react-router-redux";
import { reducer as formReducer } from "redux-form";

import loginReducer from "./containers/LoginPage/reducer";

export default function createReducer(injectedReducers) {
    return combineReducers({
        router: routerReducer,
        form: formReducer,
        login: loginReducer,
        ...injectedReducers
    });
}
