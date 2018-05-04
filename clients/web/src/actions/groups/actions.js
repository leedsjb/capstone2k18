import axios from "axios";

import {
    FETCH_GROUPS_PENDING,
    FETCH_GROUPS_SUCCESS,
    FETCH_GROUPS_ERROR
} from "./types";

export function fetchGroups() {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_GROUPS_PENDING
            });

            const { data } = await axios.get("http://localhost:4000/groups");

            dispatch({
                type: FETCH_GROUPS_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
