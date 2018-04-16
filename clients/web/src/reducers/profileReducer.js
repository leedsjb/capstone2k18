import {
    FETCH_PROFILE_PENDING,
    FETCH_PROFILE_SUCCESS,
    FETCH_PROFILE_ERROR
} from "../actions/profile/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const profileReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_PROFILE_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_PROFILE_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        default:
            return state;
    }
};

export default profileReducer;
