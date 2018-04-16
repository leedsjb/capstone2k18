import { combineReducers } from "redux";
import { routerReducer } from "react-router-redux";
import { reducer as formReducer } from "redux-form";

import signInReducer from "./signInReducer";
import missionsReducer from "./missionsReducer";
import missionDetailReducer from "./missionDetailReducer";
import aircraftReducer from "./aircraftReducer";
import aircraftDetailReducer from "./aircraftDetailReducer";
import peopleReducer from "./peopleReducer";
import peopleDetailReducer from "./peopleDetailReducer";
import groupsReducer from "./groupsReducer";
import groupsDetailReducer from "./groupsDetailReducer";
import profileReducer from "./profileReducer";

export default function createReducer(injectedReducers) {
    return combineReducers({
        router: routerReducer,
        form: formReducer,
        signIn: signInReducer,
        missions: missionsReducer,
        missionDetail: missionDetailReducer,
        aircraft: aircraftReducer,
        aircraftDetail: aircraftDetailReducer,
        people: peopleReducer,
        peopleDetail: peopleDetailReducer,
        groups: groupsReducer,
        groupsDetail: groupsDetailReducer,
        profileReducer: profileReducer,
        ...injectedReducers
    });
}
