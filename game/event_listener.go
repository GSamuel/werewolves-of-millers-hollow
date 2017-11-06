package game

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
)

type EventListener interface {
	OnGameStarted(*events.GameStartedEvent)
	OnNightStarted(*events.NightStartedEvent)
	OnWerewolfVote(*events.WerewolfVoteEvent)
	OnDailyVote(*events.DailyVoteEvent)
	OnPlayerDead(*events.PlayerDeadEvent)
	OnPlayerRevealed(*events.PlayerRevealedEvent)
}
