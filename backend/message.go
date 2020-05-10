package main

type newPlayerIDMessage struct {
	EventType string `json:"type"`
	ID        int    `json:"id"`
}

func createNewPlayerIDMessage(id int) newPlayerIDMessage {
	return newPlayerIDMessage{"newPlayerId", id}
}

type playersPositionsMessage struct {
	EventType string   `json:"type"`
	Players   []player `json:"players"`
}

func createPlayersPositionMessage(players []player) playersPositionsMessage {
	return playersPositionsMessage{"playersPosition", players}
}
