import {
    FETCH_PEOPLE_PENDING,
    FETCH_PEOPLE_SUCCESS,
    FETCH_PEOPLE_ERROR
} from "../actions/people/types";

const initialState = {
    pending: false,
    erro: null,
    data: []
};

const peopleReducer = (state = intialState, action) => {
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
        default:
            return state;
    }
};

export default peopleReducer;
