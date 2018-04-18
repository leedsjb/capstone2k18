import axios from "axios";

import {
    FETCH_PROFILE_PENDING,
    FETCH_PROFILE_SUCCESS,
    FETCH_PROFILE_ERROR
} from "./types";

export function fetchProfile(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_PROFILE_PENDING
            });

            const { data } = await axios.get(
                `http://localhost:4000/profile/$id`
            );

            dispatch({
                type: FETCH_PROFILE_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
