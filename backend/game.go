package main

import (
	"math/rand"
	"time"
)

type input int

const (
	LEFT input = iota
	RIGHT
)

type playerPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type playerInfo struct {
	updatesChannel chan map[int]playerPosition
	inputChannel   chan input
	direction      int
}

type game struct {
	playerIDCounter int
	board           map[int]playerPosition
	players         map[int]playerInfo
}

func NewGame() game {
	return game{0, make(map[int]playerPosition), make(map[int]playerInfo)}
}

func (g *game) registerNewPlayer(updatesChannel chan map[int]playerPosition, inputChannel chan input) int {
	g.playerIDCounter++
	id := g.playerIDCounter
	g.board[id] = playerPosition{rand.Int() % 200, rand.Int() % 200}
	g.players[id] = playerInfo{updatesChannel, inputChannel, 0}
	return id
}

func (g *game) removePlayer(id int) {
	delete(g.board, id)
	delete(g.players, id)
}

func (g *game) start() {
	go g.run()
}

func (g *game) run() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		<-ticker.C
		g.tick()
	}
}

func (g *game) tick() {
	// update players positions
	g.broadcastNewBoard()
	// schedule next tick
}

func (g game) broadcastNewBoard() {
	for _, info := range g.players {
		info.updatesChannel <- g.board
	}
}
