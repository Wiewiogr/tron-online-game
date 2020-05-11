package main

import (
	"log"
	"net/http"
)

var mainGame = NewGame()

func main() {
	setupRoutes()
	mainGame.start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
