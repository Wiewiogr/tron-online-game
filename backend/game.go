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

type playerTrace struct {
	Position coordinates   `json:"position"`
	Traces   []coordinates `json:"traces"`
}

type coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type playerInfo struct {
	updatesChannel chan map[int]playerTrace
	inputChannel   chan playerInput
	direction      int
}

type game struct {
	playerIDCounter int
	board           map[int]playerTrace
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
	return game{0, make(map[int]playerTrace), make(map[int]playerInfo), 800, 600}
}

func (g *game) registerNewPlayer(updatesChannel chan map[int]playerTrace, inputChannel chan playerInput) int {
	g.playerIDCounter++
	id := g.playerIDCounter
	startingX := rand.Int() % g.boardWidth
	startingY := rand.Int() % g.boardHeight

	g.board[id] = playerTrace{coordinates{startingX, startingY}, []coordinates{coordinates{startingX, startingY}}}
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

	/// adding old position as trace
	playerTrace := g.board[playerInput.id]
	playerTrace.Traces = append(playerTrace.Traces, playerTrace.Position)
	g.board[playerInput.id] = playerTrace
}

func (g *game) updatePosition() {
	for id, info := range g.players {
		position := g.board[id]
		currentDirection := directions[info.direction]
		position.Position.X = computeNewPosition(position.Position.X, currentDirection[0], g.boardWidth)
		position.Position.Y = computeNewPosition(position.Position.Y, currentDirection[1], g.boardHeight)
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
