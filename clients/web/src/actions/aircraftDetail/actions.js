import axios from "axios";

import { PROD_API_URL_V1 } from "../../constants/constants.js";

import {
    FETCH_AIRCRAFTDETAIL_PENDING,
    FETCH_AIRCRAFTDETAIL_SUCCESS,
    FETCH_AIRCRAFTDETAIL_ERROR
} from "./types";

export function fetchAircraftDetail(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_AIRCRAFTDETAIL_PENDING
            });

            const { data } = await axios.get(
                `${PROD_API_URL_V1}/aircraft/${id}`
            );

            dispatch({
                type: FETCH_AIRCRAFTDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_AIRCRAFTDETAIL_ERROR,
                error: e
            });
        }
    };
}
