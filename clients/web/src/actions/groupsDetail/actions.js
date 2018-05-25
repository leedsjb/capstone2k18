import axios from "axios";

import { PROD_API_URL_V1 } from "../../constants/constants.js";

import {
    FETCH_GROUPSDETAIL_PENDING,
    FETCH_GROUPSDETAIL_SUCCESS,
    FETCH_GROUPSDETAIL_ERROR
} from "./types";

export function fetchGroupsDetail(id) {
    return async dispatch => {
        try {
            dispatch({
                type: FETCH_GROUPSDETAIL_PENDING
            });

            const { data } = await axios.get(
                `${PROD_API_URL_V1}/groups/${id}`
                // `http://api.test.elevate.airliftnw.org/v1/groups/${id}`
            );

            dispatch({
                type: FETCH_GROUPSDETAIL_SUCCESS,
                payload: data
            });
        } catch (e) {
            dispatch({
                type: FETCH_GROUPSDETAIL_ERROR,
                error: e
            });
        }
    };
}
