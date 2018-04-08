import {
    FETCH_PEOPLEDETAIL_PENDING,
    FETCH_PEOPLEDETAIL_SUCCESS,
    FETCH_PEOPLEDETAIL_ERROR
} from "../actions/peopleDetail/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const peopleDetailReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_PEOPLEDETAIL_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_PEOPLEDETAIL_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        default:
            return state;
    }
};

export default peopleDetailReducer;
