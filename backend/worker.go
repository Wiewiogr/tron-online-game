package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type worker struct {
	game *game
	ws   *websocket.Conn
	id   int
}

func (w *worker) initAndDispatch() {
	log.Println("Client Connected")
	updatesChannel := make(chan map[int]playerPosition)
	inputChannel := make(chan input)
	w.id = w.game.registerNewPlayer(updatesChannel, inputChannel)

	message := createNewPlayerIDMessage(w.id)
	bytes, _ := json.Marshal(message)

	err := w.ws.WriteMessage(1, bytes)
	if err != nil {
		log.Println(err)
	}

	go w.writer(updatesChannel)
	go w.reader(inputChannel)
}

func (w worker) writer(updatesChannel chan map[int]playerPosition) {
	for {
		newBoard := <-updatesChannel

		message := createPlayersPositionMessage(newBoard)
		bytes, _ := json.Marshal(message)
		err := w.ws.WriteMessage(1, bytes)

		if err != nil {
			log.Println("Returning from writer", err)
			w.game.removePlayer(w.id)
			return
		}
	}
}

func (w worker) reader(inputChannel chan input) {
	for {
		_, p, err := w.ws.ReadMessage()

		if err != nil {
			log.Println("Returning from reader", err)
			w.game.removePlayer(w.id)
			return
		}
		fmt.Println(string(p))
	}
}
