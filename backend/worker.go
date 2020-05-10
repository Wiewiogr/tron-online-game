package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type worker struct {
	game *game
	ws   *websocket.Conn
}

func (w worker) initAndDispatch() {
	log.Println("Client Connected")
	id := w.game.addPlayer()
	message := createNewPlayerIDMessage(id)
	bytes, _ := json.Marshal(message)

	err := w.ws.WriteMessage(1, bytes)
	if err != nil {
		log.Println(err)
	}

	ticker := time.NewTicker(100 * time.Millisecond)

	go w.writer(ticker)
	go w.reader()
}

func (w worker) reader() {
	for {
		_, p, err := w.ws.ReadMessage()

		if err != nil {
			log.Println("Returning from reader", err)
			return
		}
		fmt.Println(string(p))
	}
}

func (w worker) writer(ticker *time.Ticker) {
	for {
		t := <-ticker.C

		message := createPlayersPositionMessage(w.game.players)
		bytes, _ := json.Marshal(message)
		err := w.ws.WriteMessage(1, bytes)

		if err != nil {
			log.Println("Returning from writer", err)
			return
		}
		fmt.Println("Tick at", t)
	}
}
