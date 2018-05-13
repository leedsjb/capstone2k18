import {
    FETCH_AIRCRAFTDETAIL_PENDING,
    FETCH_AIRCRAFTDETAIL_SUCCESS,
    FETCH_AIRCRAFTDETAIL_ERROR
} from "../actions/aircraftDetail/types";

const intitialState = {
    pending: false,
    error: null,
    data: []
};

const aircraftDetailReducer = (state = intitialState, action) => {
    switch (action.type) {
        case FETCH_AIRCRAFTDETAIL_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_AIRCRAFTDETAIL_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        case FETCH_AIRCRAFTDETAIL_ERROR:
            return {
                error: action.error
            };
        default:
            return state;
    }
};

export default aircraftDetailReducer;
