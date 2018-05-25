import {
    FETCH_AIRCRAFT_PENDING,
    FETCH_AIRCRAFT_SUCCESS,
    FETCH_AIRCRAFT_ERROR
} from "../actions/aircraft/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const aircraftReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_AIRCRAFT_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_AIRCRAFT_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        case FETCH_AIRCRAFT_ERROR:
            return {
                error: action.error
            };
        default:
            return state;
    }
};

export default aircraftReducer;
