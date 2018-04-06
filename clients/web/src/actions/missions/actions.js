import axios from "axios";

import { FETCH_MISSIONS_PENDING, FETCH_MISSIONS_SUCCESS } from "./types";

export function fetchMissions() {
    return async dispatch => {
        try {
            const { data } = await axios.get("http://localhost:3004/missions");
            dispatch({
                type: FETCH_MISSIONS_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
