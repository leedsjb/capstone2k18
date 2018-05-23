import axios from "axios";

import {
    FETCH_PROFILE_PENDING,
    FETCH_PROFILE_SUCCESS,
    FETCH_PROFILE_ERROR
} from "./types";

export function fetchProfile() {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_PROFILE_PENDING
            });

            const { data } = await axios.get("http://localhost:4000/people/me");

            dispatch({
                type: FETCH_PROFILE_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_PROFILE_ERROR,
                error: e
            });
        }
    };
}
