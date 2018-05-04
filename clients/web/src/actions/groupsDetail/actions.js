import axios from "axios";

import {
    FETCH_GROUPSDETAIL_PENDING,
    FETCH_GROUPSDETAIL_SUCCESS,
    FETCH_GROUPSDETAIL_ERROR
} from "./types";

export function fetchGroupsDetail(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_GROUPSDETAIL_PENDING
            });

            const { data } = await axios.get(
                `http://localhost:4000/groups/${id}`
            );

            dispatch({
                type: FETCH_GROUPSDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
