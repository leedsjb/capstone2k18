import axios from "axios";

import { FETCH_RESOURCES_PENDING, FETCH_RESOURCES_SUCCESS } from "./types";

export function fetchResources() {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_RESOURCES_PENDING
            });

            const { data } = await axios.get(`http://localhost:4000/resources`);

            dispatch({
                type: FETCH_RESOURCES_SUCCESS,
                payload: data
            });
        } catch (e) {}
    };
}