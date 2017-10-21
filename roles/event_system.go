package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
)

type EventSystem struct {
	Group
}

func (e *EventSystem) GameStartedEvent(event events.GameStartedEvent) {
	for i := 0; i < e.Count(); i++ {
		e.Role(i).OnGameStarted(event)
	}
}

func (e *EventSystem) NightStartedEvent(event events.NightStartedEvent) {
	for i := 0; i < e.Count(); i++ {
		e.Role(i).OnNightStarted(event)
	}
}

func (e *EventSystem) WerewolfVoteEvent(event events.WerewolfVoteEvent) {
	for i := 0; i < e.Count(); i++ {
		e.Role(i).OnWerewolfVote(event)
	}
}

func (e *EventSystem) DailyVoteEvent(event events.DailyVoteEvent) {
	for i := 0; i < e.Count(); i++ {
		e.Role(i).OnDailyVote(event)
	}
}

func (e *EventSystem) PlayerDeadEvent(event events.PlayerDeadEvent) {
	for i := 0; i < e.Count(); i++ {
		e.Role(i).OnPlayerDead(event)
	}
}

func (e *EventSystem) PlayerRevealedEvent(event events.PlayerRevealedEvent) {
	for i := 0; i < e.Count(); i++ {
		e.Role(i).OnPlayerRevealed(event)
	}
}

func NewEventSystem(group Group) *EventSystem {
	return &EventSystem{group}
}
