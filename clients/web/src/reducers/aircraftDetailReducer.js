import {
    FETCH_AIRCRAFTDETAIL_PENDING,
    FETCH_AIRCRAFTDETAIL_SUCCESS,
    FETCH_AIRCRAFTDETAIL_ERROR,
    UPDATE_AIRCRAFTDETAIL_POSITION,
    AIRCRAFTDETAIL_NEW_MISSION,
    AIRCRAFTDETAIL_MISSION_COMPLETE
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
        case UPDATE_AIRCRAFTDETAIL_POSITION:
        case AIRCRAFTDETAIL_NEW_MISSION:
        case AIRCRAFTDETAIL_MISSION_COMPLETE:
            return {
                ...state,
                data: {
                    ...state.data,
                    ...action.payload
                }
            };
        default:
            return state;
    }
};

export default aircraftDetailReducer;
