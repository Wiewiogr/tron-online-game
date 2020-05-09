package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\":\"ok\"}")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	var worker = worker{&mainGame, ws}
	worker.initAndDispatch()
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func setupRoutes() {
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/", healthCheck)
	http.HandleFunc("/ws", wsEndpoint)
}
