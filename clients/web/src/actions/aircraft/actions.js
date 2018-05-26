import axios from "axios";

import { PROD_API_URL_V1 } from "../../constants/constants.js";

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

            let url = `${PROD_API_URL_V1}/aircraft`;

            if (term) {
                url = `${url}?q=${term}`;
            } else if (status) {
                url = `${url}?status=${status}`;
            } else if (category) {
                url = `${url}?category=${category}`;
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
