import axios from "axios";

import {
    FETCH_MISSIONS_PENDING,
    FETCH_MISSIONS_SUCCESS,
    FETCH_MISSIONS_ERROR
} from "./types";

export function fetchMissions() {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_MISSIONS_PENDING
            });

            const { data } = await axios.get("http://localhost:4000/missions");

            dispatch({
                type: FETCH_MISSIONS_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
