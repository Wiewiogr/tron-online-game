package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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
	err := w.ws.WriteMessage(1, []byte(`{ "id":`+strconv.Itoa(id)+`}`))
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

		bytes, err := json.Marshal(w.game.players)
		if err != nil {
			log.Println("Error marshalling players", err)
			continue
		}
		err = w.ws.WriteMessage(1, bytes)

		if err != nil {
			log.Println("Returning from writer", err)
			return
		}
		fmt.Println("Tick at", t)
	}
}
