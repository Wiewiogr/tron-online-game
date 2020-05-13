package main

type newPlayerIDMessage struct {
	EventType string `json:"type"`
	ID        int    `json:"id"`
}

func createNewPlayerIDMessage(id int) newPlayerIDMessage {
	return newPlayerIDMessage{"newPlayerId", id}
}

type playersPositionsMessage struct {
	EventType string              `json:"type"`
	Board     map[int]playerTrace `json:"board"`
}

func createPlayersPositionMessage(board map[int]playerTrace) playersPositionsMessage {
	return playersPositionsMessage{"playersTrace", board}
}

type playerInputMessage struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}
