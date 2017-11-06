package game

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
)

type EventSystem struct {
}

func (e EventSystem) GameStartedEvent(g *Game, event *events.GameStartedEvent) {
	for i := 0; i < g.Count(); i++ {
		g.Player(i).OnGameStarted(g, event)
	}
}

func (e EventSystem) NightStartedEvent(g *Game, event *events.NightStartedEvent) {
	for i := 0; i < g.Count(); i++ {
		g.Player(i).OnNightStarted(g, event)
	}
}

func (e EventSystem) WerewolfVoteEvent(g *Game, event *events.WerewolfVoteEvent) {
	for i := 0; i < g.Count(); i++ {
		g.Player(i).OnWerewolfVote(g, event)
	}
}

func (e EventSystem) DailyVoteEvent(g *Game, event *events.DailyVoteEvent) {
	for i := 0; i < g.Count(); i++ {
		g.Player(i).OnDailyVote(g, event)
	}
}

func (e EventSystem) PlayerDeadEvent(g *Game, event *events.PlayerDeadEvent) {
	for i := 0; i < g.Count(); i++ {
		g.Player(i).OnPlayerDead(g, event)
	}
}

func (e EventSystem) PlayerRevealedEvent(g *Game, event *events.PlayerRevealedEvent) {
	for i := 0; i < g.Count(); i++ {
		g.Player(i).OnPlayerRevealed(g, event)
	}
}

func NewEventSystem() *EventSystem {
	return &EventSystem{}
}
