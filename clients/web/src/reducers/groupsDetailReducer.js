import {
    FETCH_GROUPSDETAIL_PENDING,
    FETCH_GROUPSDETAIL_SUCCESS
} from "../actions/groupsDetail/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const groupDetailReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_GROUPSDETAIL_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_GROUPSDETAIL_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        default:
            return state;
    }
};

export default groupDetailReducer;
