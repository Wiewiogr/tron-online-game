package main

import "math/rand"

type player struct {
	ID int `json:"id"`
	X  int `json:"x"`
	Y  int `json:"y"`
}

type game struct {
	playerIDCounter int
	players         []player
}

func (g *game) addPlayer() int {
	g.playerIDCounter++
	g.players = append(g.players, player{g.playerIDCounter, rand.Int() % 200, rand.Int() % 200})
	return g.playerIDCounter
}
