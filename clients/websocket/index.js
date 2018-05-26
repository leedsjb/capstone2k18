"use strict";

//change this if you run your server on a different port number
const host = "localhost:8000";
const status = document.querySelector("#status")
// const notifications = document.querySelector("#notifications");
// const dummyAnnouncementBtn = document.querySelector("#DummyAnnouncementBtn");
const errors = document.querySelector("#errors");

// let auth = getAuth();

// refreshLocalUser();

const websocket = new WebSocket("wss://" + host + "/v1/ws");
websocket.addEventListener("error", function (err) {
    errors.textContent = err.message;
});
websocket.addEventListener("open", function () {
    status.textContent = "Connected";
});
websocket.addEventListener("close", function () {
    status.textContent = "Closed";
});
// websocket.addEventListener("message", function (event) {

//     let notification = JSON.parse(event.data);
//     let data = notification.data;
//     let messageWrapper = document.createElement("div");
//     messageWrapper.classList.add("message-wrapper");

//     let name = document.createElement("p");
//     name.classList.add("message-top-row")
//     name.classList.add("message-name");
//     name.textContent = data.createdBy.firstName + " " + data.createdBy.lastName;

//     let timestamp = document.createElement("p");
//     timestamp.classList.add("message-top-row");
//     timestamp.classList.add("message-timestamp");
//     timestamp.textContent = moment(data.postDate).format("lll");

//     let body = document.createElement("p");
//     body.classList.add("message-body");
//     body.textContent = data.message;

//     messageWrapper.appendChild(name);
//     messageWrapper.appendChild(timestamp);
//     messageWrapper.appendChild(body);
//     notifications.appendChild(messageWrapper);
// });