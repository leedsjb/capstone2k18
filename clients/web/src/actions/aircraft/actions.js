import axios from "axios";

import { API_URI } from "../../constants/constants.js";

import {
    FETCH_AIRCRAFT_PENDING,
    FETCH_AIRCRAFT_SUCCESS,
    FETCH_AIRCRAFT_ERROR
} from "./types";

export function fetchAircraft(term, status, category) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_AIRCRAFT_PENDING
            });

            let url = `${API_URI}/aircraft`;

            if (term) {
                url = `${url}?q=${term}`;
            } else if (status && !category) {
                url = `${url}?status=${status}`;
            } else if (category && !status) {
                url = `${url}?category=${category}`;
            } else if (category && status) {
                url = `${url}?status=${status}&category=${category}`;
            }

            const { data } = await axios.get(url);

            dispatch({
                type: FETCH_AIRCRAFT_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_AIRCRAFT_ERROR,
                error: e
            });
        }
    };
}
