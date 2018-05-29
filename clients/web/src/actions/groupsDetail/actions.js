import axios from "axios";

import { API_URI } from "../../constants/constants.js";

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

            const { data } = await axios.get(`${API_URI}/groups/${id}`);

            dispatch({
                type: FETCH_GROUPSDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_GROUPSDETAIL_ERROR,
                error: e
            });
        }
    };
}
