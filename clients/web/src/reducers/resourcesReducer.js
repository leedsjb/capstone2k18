import {
    FETCH_RESOURCES_PENDING,
    FETCH_RESOURCES_SUCCESS,
    FETCH_RESOURCES_ERROR
} from "../actions/resources/types";

const initialState = {
    pending: false,
    error: null,
    data: []
};

const resourcesReducer = (state = initialState, action) => {
    switch (action.type) {
        case FETCH_RESOURCES_PENDING:
            return {
                ...state,
                pending: true
            };
        case FETCH_RESOURCES_SUCCESS:
            return {
                ...state,
                pending: false,
                data: action.payload
            };
        case FETCH_RESOURCES_ERROR:
            return {
                error: action.error
            };
        default:
            return state;
    }
};

export default resourcesReducer;
