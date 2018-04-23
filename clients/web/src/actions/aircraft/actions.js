import axios from "axios";

import { FETCH_AIRCRAFT_PENDING, FETCH_AIRCRAFT_SUCCESS } from "./types";

export function fetchAircraft(term, status) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_AIRCRAFT_PENDING
            });

            let url = "http://localhost:4000/aircraft";

            if (term) {
                url = `${url}?q=${term}`;
            } else if (status) {
                url = `${url}?status=${status}`;
            }

            const { data } = await axios.get(url);

            dispatch({
                type: FETCH_AIRCRAFT_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
