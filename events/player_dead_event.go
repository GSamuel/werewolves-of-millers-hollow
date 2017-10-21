package events

import ()

type PlayerDeadEvent struct {
	id         int
	werewolves bool
	prevented  bool
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

func (p *PlayerDeadEvent) WerewolfAttack() bool {
	return p.werewolves
}

func (p *PlayerDeadEvent) SetWerewolfAttack(werewolves bool) {
	p.werewolves = werewolves
}

func NewPlayerDeadEvent(id int) *PlayerDeadEvent {
	return &PlayerDeadEvent{id, false, false}
}
