package game

import (
	"github.com/GSamuel/werewolves-of-millers-hollow/events"
)

type EventListener interface {
	OnGameStarted(*Game, *events.GameStartedEvent)
	OnNightStarted(*Game, *events.NightStartedEvent)
	OnWerewolfVote(*Game, *events.WerewolfVoteEvent)
	OnDailyVote(*Game, *events.DailyVoteEvent)
	OnPlayerDead(*Game, *events.PlayerDeadEvent)
	OnPlayerRevealed(*Game, *events.PlayerRevealedEvent)
}
