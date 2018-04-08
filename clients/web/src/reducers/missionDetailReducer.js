import {
    FETCH_MISSIONDETAIL_PENDING,
    FETCH_MISSIONDETAIL_SUCCESS,
    FETCH_MISSIONDETAIL_ERROR
} from "../actions/missionDetail/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const missionDetailReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_MISSIONDETAIL_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_MISSIONDETAIL_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        default:
            return state;
    }
};

export default missionDetailReducer;
