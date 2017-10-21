package events

import ()

type PlayerDeadEvent struct {
	id int
}

func NewPlayerDeadEvent(id int) PlayerDeadEvent {
	return PlayerDeadEvent{id}
}
