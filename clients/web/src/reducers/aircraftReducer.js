import {
    FETCH_AIRCRAFT_PENDING,
    FETCH_AIRCRAFT_SUCCESS,
    FETCH_AIRCRAFT_ERROR,
    UPDATE_AIRCRAFT_POSITION,
    AIRCRAFT_NEW_MISSION,
    AIRCRAFT_MISSION_COMPLETE
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
        case UPDATE_AIRCRAFT_POSITION:
        case AIRCRAFT_NEW_MISSION:
        case AIRCRAFT_MISSION_COMPLETE:
            return {
                ...state,
                data: state.data.map(item => {
                    if (item.id !== action.payload.id) {
                        return item;
                    }

                    return {
                        ...item,
                        ...action.payload
                    };
                })
            };
        default:
            return state;
    }
};

export default aircraftReducer;
