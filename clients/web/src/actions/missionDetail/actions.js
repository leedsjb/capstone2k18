import axios from "axios";

import {
    FETCH_MISSIONDETAIL_PENDING,
    FETCH_MISSIONDETAIL_SUCCESS,
    FETCH_MISSIONDETAIL_ERROR
} from "./types";

export function fetchMissionDetail(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_MISSIONDETAIL_PENDING
            });

            const { data } = await axios.get(
                `http://localhost:4000/missions/$id`
            );

            dispatch({
                type: FETCH_MISSIONDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
