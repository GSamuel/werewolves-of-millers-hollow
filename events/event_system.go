package events

import ()

type EventSystem struct {
	listeners []EventListener
}

func (e *EventSystem) Add(listener EventListener) {
	e.listeners = append(e.listeners, listener)
}

func (e *EventSystem) GameStartedEvent(event GameStartedEvent) {
	for i := 0; i < len(e.listeners); i++ {
		e.listeners[i].OnGameStarted(event)
	}
}
func (e *EventSystem) NightStartedEvent(event NightStartedEvent) {
	for i := 0; i < len(e.listeners); i++ {
		e.listeners[i].OnNightStarted(event)
	}
}
func (e *EventSystem) WerewolfVoteEvent(event WerewolfVoteEvent) {
	for i := 0; i < len(e.listeners); i++ {
		e.listeners[i].OnWerewolfVote(event)
	}
}
func (e *EventSystem) DailyVoteEvent(event DailyVoteEvent) {
	for i := 0; i < len(e.listeners); i++ {
		e.listeners[i].OnDailyVote(event)
	}
}
func (e *EventSystem) PlayerDeadEvent(event PlayerDeadEvent) {
	for i := 0; i < len(e.listeners); i++ {
		e.listeners[i].OnPlayerDead(event)
	}
}
func (e *EventSystem) PlayerRevealedEvent(event PlayerRevealedEvent) {
	for i := 0; i < len(e.listeners); i++ {
		e.listeners[i].OnPlayerRevealed(event)
	}
}

func NewEventSystem() *EventSystem {
	return &EventSystem{make([]EventListener, 0)}
}
