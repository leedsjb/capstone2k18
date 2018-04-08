import axios from "axios";

import {
    FETCH_AIRCRAFT_PENDING,
    FETCH_AIRCRAFT_SUCCESS,
    FETCH_AIRCRAFT_ERROR
} from "./types";

export function fetchAircaft() {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_AIRCRAFT_PENDING
            });

            const { data } = await axios.get("http://localhost:4000/aircraft");

            dispatch({
                type: FETCH_AIRCRAFT_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
