package events

import ()

type PlayerDeadEvent struct {
	id        int
	prevented bool
}

func (p *PlayerDeadEvent) ID() int {
	return p.id
}

func (p *PlayerDeadEvent) DeathPrevented() bool {
	return p.prevented
}

func (p *PlayerDeadEvent) PreventDeath() {
	p.prevented = true
}

func NewPlayerDeadEvent(id int) PlayerDeadEvent {
	return PlayerDeadEvent{id, false}
}
