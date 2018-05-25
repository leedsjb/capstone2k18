import axios from "axios";

import { PROD_API_URL_V1 } from "../../constants/constants.js";

import {
    FETCH_PEOPLE_PENDING,
    FETCH_PEOPLE_SUCCESS,
    FETCH_PEOPLE_ERROR
} from "./types";

export function fetchPeople(term) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_PEOPLE_PENDING
            });

            let url = `${PROD_API_URL_V1}/people`;

            if (term) {
                url = `${url}?q=${term}`;
            }

            const { data } = await axios.get(url);

            dispatch({
                type: FETCH_PEOPLE_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_PEOPLE_ERROR,
                error: e
            });
        }
    };
}
