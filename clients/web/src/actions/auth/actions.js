import { push } from "react-router-redux";

// We should think more carefully about naming with regards
// to actions, reducers and pieces of state
export const signIn = values => {
    return (dispatch, getState) => {
        // This shouldn't be missions, but this is just
        // a demonstration
        dispatch(push("/aircraft"));
    };
};
