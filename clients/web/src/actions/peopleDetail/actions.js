import axios from "axios";

import { API_URI } from "../../constants/constants.js";

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

            const { data } = await axios.get(`${API_URI}/people/${id}`);

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
