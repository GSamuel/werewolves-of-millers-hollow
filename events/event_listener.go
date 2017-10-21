package events

import ()

type EventListener interface {
	OnGameStarted(GameStartedEvent)
	OnNightStarted(NightStartedEvent)
	OnWerewolfVote(WerewolfVoteEvent)
	OnDailyVote(DailyVoteEvent)
	OnPlayerDead(PlayerDeadEvent)
	OnPlayerRevealed(PlayerRevealedEvent)
}
