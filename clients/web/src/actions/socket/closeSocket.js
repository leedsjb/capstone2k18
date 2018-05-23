import { CLOSE_SOCKET } from "./types";

export default () => {
    return dispatch => {
        dispatch({
            type: CLOSE_SOCKET
        });
    };
};
