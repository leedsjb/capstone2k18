import { OPEN_SOCKET } from "./types";

export default () => {
    return dispatch => {
        dispatch({
            type: OPEN_SOCKET
        });
    };
};
