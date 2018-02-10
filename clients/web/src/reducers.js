import { combineReducers } from "react-redux";

import { signUpReducer } from "./containers/SignUpPage/reducer";

const reducers = combineReducers({
    signUp: signUpReducer
});

export default reducers;
