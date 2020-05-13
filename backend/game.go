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

type playerInput struct {
	id  int
	key input
}

type playerPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type playerInfo struct {
	updatesChannel chan map[int]playerPosition
	inputChannel   chan playerInput
	direction      int
}

type game struct {
	playerIDCounter int
	board           map[int]playerPosition
	players         map[int]playerInfo
	boardWidth      int
	boardHeight     int
}

var directions = [4][2]int{
	[2]int{1, 0},
	[2]int{0, 1},
	[2]int{-1, 0},
	[2]int{0, -1},
}

func NewGame() game {
	return game{0, make(map[int]playerPosition), make(map[int]playerInfo), 800, 600}
}

func (g *game) registerNewPlayer(updatesChannel chan map[int]playerPosition, inputChannel chan playerInput) int {
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
	g.readInput()
	g.updatePosition()
	g.broadcastNewBoard()
}

func (g game) readInput() {
	for _, info := range g.players {
		select {
		case playerInput := <-info.inputChannel:
			g.handlePlayerInput(playerInput)
		default:
		}
	}
}

func (g *game) handlePlayerInput(playerInput playerInput) {
	playerInfo := g.players[playerInput.id]
	if playerInput.key == RIGHT {
		playerInfo.direction = (playerInfo.direction + 1) % 4
	} else if playerInput.key == LEFT {
		if playerInfo.direction == 0 {
			playerInfo.direction = 3
		} else {
			playerInfo.direction--
		}
	}
	g.players[playerInput.id] = playerInfo
}

func (g *game) updatePosition() {
	for id, info := range g.players {
		position := g.board[id]
		currentDirection := directions[info.direction]
		position.X = computeNewPosition(position.X, currentDirection[0], g.boardWidth)
		position.Y = computeNewPosition(position.Y, currentDirection[1], g.boardHeight)
		g.board[id] = position
	}
}

func computeNewPosition(position, direction, limit int) int {
	newPosition := position + direction
	if newPosition < 0 {
		newPosition = limit - 1
	} else if newPosition >= limit {
		newPosition = 0
	}
	return newPosition
}

func (g game) broadcastNewBoard() {
	for _, info := range g.players {
		info.updatesChannel <- g.board
	}
}
