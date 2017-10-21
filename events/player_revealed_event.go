package events

import ()

type PlayerRevealedEvent struct {
	id int
}

func (p *PlayerRevealedEvent) ID() int {
	return p.id
}

func NewPlayerRevealedEvent(id int) PlayerRevealedEvent {
	return PlayerRevealedEvent{id}
}
