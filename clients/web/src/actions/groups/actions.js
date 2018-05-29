import axios from "axios";

import { API_URI } from "../../constants/constants.js";

import {
    FETCH_GROUPS_PENDING,
    FETCH_GROUPS_SUCCESS,
    FETCH_GROUPS_ERROR
} from "./types";

export function fetchGroups(term) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_GROUPS_PENDING
            });

            let url = `${API_URI}/groups`;

            if (term) {
                url = `${url}?q=${term}`;
            }

            const { data } = await axios.get(url);

            dispatch({
                type: FETCH_GROUPS_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_GROUPS_ERROR,
                error: e
            });
        }
    };
}
