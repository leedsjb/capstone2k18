package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/info344-a17/challenges-vincentmvdm/servers/gateway/sessions"

	"github.com/gorilla/websocket"
)

//WebSocketsHandler is a handler for WebSocket upgrade requests
type WebSocketsHandler struct {
	notifier   *Notifier
	upgrader   *websocket.Upgrader
	signingKey string
}

//NewWebSocketsHandler constructs a new WebSocketsHandler
func (ctx *HandlerContext) NewWebSocketsHandler(notifier *Notifier) *WebSocketsHandler {
	wsh := &WebSocketsHandler{}
	wsh.notifier = notifier
	wsh.upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	wsh.signingKey = ctx.SigningKey
	return wsh
}

//ServeHTTP implements the http.Handler interface for the WebSocketsHandler
func (wsh *WebSocketsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//https://godoc.org/github.com/gorilla/websocket#hdr-Overview
	//Upgrade the connection to a WebSocket
	_, err := sessions.GetSessionID(r, wsh.signingKey)
	if err != nil {
		//Users must be authenticated to upgrade to a WebSocket
		//If there's an error when retrieving the session state, respond
		//with an http.StatusUnauthorized error.
		http.Error(w, fmt.Sprintf("Must be authenticated to upgrade to a websocket: %v", err), http.StatusUnauthorized)
		return
	}
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error upgrading to a websocket: %v", err), http.StatusInternalServerError)
		return
	}
	//Add the new websock.Conn to the Notifier
	wsh.notifier.AddClient(conn)
}

//Notifier is an object that handles WebSocket notifications
type Notifier struct {
	eventQ  chan []byte
	clients []*websocket.Conn
	//add other channels to protect `clients` from concurrent use
	clientsToAdd    chan *websocket.Conn
	clientsToRemove chan *websocket.Conn
}

//NewNotifier constructs a new Notifier
func NewNotifier() *Notifier {
	//Construct a new Notifier
	n := &Notifier{}
	n.eventQ = make(chan []byte, 100)
	n.clientsToAdd = make(chan *websocket.Conn, 100)
	n.clientsToRemove = make(chan *websocket.Conn, 100)

	//Start the event notification loop
	go n.start()
	return n
}

//AddClient adds a new client to the Notifier
func (n *Notifier) AddClient(client *websocket.Conn) {
	//Add the client to the `clients` slice
	n.clientsToAdd <- client

	//Process incoming control messages from the client
	//https://godoc.org/github.com/gorilla/websocket#hdr-Control_Messages
	for {
		if _, _, err := client.NextReader(); err != nil {
			//Remove the client
			n.clientsToRemove <- client
			break
		}
	}
}

//Notify broadcasts the event to all WebSocket clients
func (n *Notifier) Notify(event []byte) {
	//Add `event` to the `n.eventQ`
	n.eventQ <- event
}

//start starts the notification loop
func (n *Notifier) start() {
	log.Println("starting notifier loop")
	//Start a never-ending loop that reads new events out
	//of the `n.eventQ` and broadcasts them to all WebSocket clients
	for {
		select {
		case clientToAdd := <-n.clientsToAdd:
			n.clients = append(n.clients, clientToAdd)
		case clientToRemove := <-n.clientsToRemove:
			clientToRemove.Close()
			newClients := []*websocket.Conn{}
			for _, client := range n.clients {
				if client != clientToRemove {
					newClients = append(newClients, client)
				}
			}
			n.clients = newClients
		case msg := <-n.eventQ:
			for _, client := range n.clients {
				if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
					n.clientsToRemove <- client
				}
			}
		}
	}
}
