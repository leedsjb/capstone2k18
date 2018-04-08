import axios from "axios";

import {
    FETCH_PEOPLE_PENDING,
    FETCH_PEOPLE_SUCCESS,
    FETCH_PEOPLE_ERROR
} from "./types";

export function fetchPeople() {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_PEOPLE_PENDING
            });

            const { data } = await axios.get("http://localhost:4000/people");

            dispatch({
                type: FETCH_PEOPLE_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
