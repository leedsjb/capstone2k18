import {
    FETCH_GROUPS_PENDING,
    FETCH_GROUPS_SUCCESS
} from "../actions/groups/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const groupsReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_GROUPS_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_GROUPS_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        default:
            return state;
    }
};

export default groupsReducer;