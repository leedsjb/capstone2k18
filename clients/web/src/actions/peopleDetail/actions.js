import axios from "axios";

import {
    FETCH_PEOPLEDETAIL_PENDING,
    FETCH_PEOPLEDETAIL_SUCCESS,
    FETCH_PEOPLEDETAIL_ERROR
} from "./types";

export function fetchPeopleDetail(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_PEOPLEDETAIL_PENDING
            });

            const { data } = await axios.get(
                `http://api.test.elevate.airliftnw.org/v1/people/${id}`
            );

            dispatch({
                type: FETCH_PEOPLEDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_PEOPLEDETAIL_ERROR,
                error: e
            });
        }
    };
}
