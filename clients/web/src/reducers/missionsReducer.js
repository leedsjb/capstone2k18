import {
    FETCH_MISSIONS_PENDING,
    FETCH_MISSIONS_SUCCESS,
    FETCH_MISSIONS_ERROR
} from "../actions/missions/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const missionsReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_MISSIONS_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_MISSIONS_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        default:
            return state;
    }
};

export default missionsReducer;
