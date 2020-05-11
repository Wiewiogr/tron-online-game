package main

type newPlayerIDMessage struct {
	EventType string `json:"type"`
	ID        int    `json:"id"`
}

func createNewPlayerIDMessage(id int) newPlayerIDMessage {
	return newPlayerIDMessage{"newPlayerId", id}
}

type playersPositionsMessage struct {
	EventType string                 `json:"type"`
	Board     map[int]playerPosition `json:"board"`
}

func createPlayersPositionMessage(board map[int]playerPosition) playersPositionsMessage {
	return playersPositionsMessage{"playersPosition", board}
}
