import {
    FETCH_PEOPLE_PENDING,
    FETCH_PEOPLE_SUCCESS,
    FETCH_PEOPLE_ERROR
} from "../actions/people/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const peopleReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_PEOPLE_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_PEOPLE_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        case FETCH_PEOPLE_ERROR:
            return {
                error: action.error
            };
        default:
            return state;
    }
};

export default peopleReducer;
