package events

import ()

type PlayerRevealedEvent struct {
	id int
}

func NewPlayerRevealedEvent(id int) PlayerRevealedEvent {
	return PlayerRevealedEvent{id}
}
