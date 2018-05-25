import axios from "axios";

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

            let url = "http://api.test.elevate.airliftnw.org/v1/groups";

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

// http://api.test.elevate.airliftnw.org/v1/groups?q=dEr
