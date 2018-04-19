import axios from "axios";

import {
    FETCH_AIRCRAFT_PENDING,
    FETCH_AIRCRAFT_SUCCESS,
    FETCH_AIRCRAFT_ERROR
} from "./types";

export function fetchAircraft(status) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_AIRCRAFT_PENDING
            });

            let url = "http://localhost:4000/aircraft";

            if (status) {
                url = `${url}?status=${status}`;
            }

            console.log(url);
            const { data } = await axios.get(url);

            dispatch({
                type: FETCH_AIRCRAFT_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
