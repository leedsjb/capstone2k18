import axios from "axios";

import {
    FETCH_PEOPLE_PENDING,
    FETCH_PEOPLE_SUCCESS,
    FETCH_PEOPLE_ERROR
} from "./types";

export function fetchPeople(term) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_PEOPLE_PENDING
            });

            let url = "http://api.test.elevate.airliftnw.org/v1/people";

            if (term) {
                url = `${url}?q=${term}`;
            }

            const { data } = await axios.get(url);

            dispatch({
                type: FETCH_PEOPLE_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_PEOPLE_ERROR,
                error: e
            });
        }
    };
}
