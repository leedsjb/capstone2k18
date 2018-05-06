import axios from "axios";

import {
    FETCH_PEOPLEDETAIL_PENDING,
    FETCH_PEOPLEDETAIL_SUCCESS
} from "./types";

export function fetchPeopleDetail(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_PEOPLEDETAIL_PENDING
            });

            const { data } = await axios.get(
                `http://localhost:4000/people/${id}`
            );

            dispatch({
                type: FETCH_PEOPLEDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
