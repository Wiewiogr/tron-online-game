package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type playerSession struct {
	game *game
	ws   *websocket.Conn
	id   int
}

func (w *playerSession) initAndDispatch() {
	updatesChannel := make(chan map[int]playerTrace)
	inputChannel := make(chan playerInput)
	w.id = w.game.registerNewPlayer(updatesChannel, inputChannel)

	message := createNewPlayerIDMessage(w.id)
	bytes, _ := json.Marshal(message)

	err := w.ws.WriteMessage(1, bytes)
	if err != nil {
		log.Println(err)
	}

	go w.updatesSender(updatesChannel)
	go w.inputListener(inputChannel)
}

func (w playerSession) updatesSender(updatesChannel chan map[int]playerTrace) {
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

func (w playerSession) inputListener(inputChannel chan playerInput) {
	for {
		_, bytes, err := w.ws.ReadMessage()

		var message playerInputMessage
		json.Unmarshal(bytes, &message)
		input := toPlayerInput(message)
		if input != nil {
			inputChannel <- *input
		}
		if err != nil {
			log.Println("Returning from reader", err)
			w.game.removePlayer(w.id)
			return
		}
		fmt.Println(string(bytes))
	}
}

func toPlayerInput(msg playerInputMessage) *playerInput {
	switch msg.Key {
	case "Right":
		return &playerInput{msg.ID, RIGHT}
	case "Left":
		return &playerInput{msg.ID, LEFT}
	default:
		return nil
	}
}
