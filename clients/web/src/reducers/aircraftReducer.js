import {
    FETCH_AIRCRAFT_PENDING,
    FETCH_AIRCRAFT_SUCEESS,
    FETCH_AIRCRAFT_ERROR
} from "../actions/aircraft/types";

const intitialState = {
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
        case FETCH_AIRCRAFT_SUCEESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        default:
            return state;
    }
};

export default aircraftReducer;
