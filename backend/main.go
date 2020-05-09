package main

import (
	"log"
	"net/http"
)

var mainGame = game{}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
