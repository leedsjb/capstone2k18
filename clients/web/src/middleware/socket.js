// Courtesy of https://exec64.co.uk/blog/websockets_with_redux/

import { OPEN_SOCKET, CLOSE_SOCKET } from "../actions/socket/types";
import { PROD_WEBSOCKET_URL_V1 } from "../constants/constants.js";

const socketMiddleware = (function() {
    let socket = null;

    const onMessage = (ws, store) => evt => {
        // Parse the JSON message received on the websocket
        store.dispatch(JSON.parse(evt.data));
    };

    return store => next => action => {
        switch (action.type) {
            // The user wants us to connect
            case OPEN_SOCKET:
                // Start a new connection to the server
                if (socket != null) {
                    socket.close();
                }

                socket = new WebSocket(PROD_WEBSOCKET_URL_V1);
                socket.onmessage = onMessage(socket, store);
                socket.addEventListener("error", err => {
                    console.log("WebSocket error");
                });
                socket.addEventListener("open", () => {
                    console.log("WebSocket opened");
                });
                socket.addEventListener("close", () => {
                    console.log("WebSocket closed");
                });

                return next(action);

            // The user wants us to disconnect
            case CLOSE_SOCKET:
                if (socket != null) {
                    socket.close();
                }
                socket = null;

                return next(action);

            default:
                return next(action);
        }
    };
})();

export default socketMiddleware;
