package events

import ()

type GameStartedEvent struct {
}

func NewGameStartedEvent() GameStartedEvent {
	return GameStartedEvent{}
}
