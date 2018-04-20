import axios from "axios";

import {
    FETCH_AIRCRAFTDETAIL_PENDING,
    FETCH_AIRCRAFTDETAIL_SUCCESS,
    FETCH_AIRCRAFTDETAILS_ERROR
} from "./types";

export function fetchAircraftDetail(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_AIRCRAFTDETAIL_PENDING
            });

            const { data } = await axios.get(
                `http://localhost:4000/aircraft/${id}`
            );

            dispatch({
                type: FETCH_AIRCRAFTDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}
