package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
)

const (
	UNDEFINED = "Undefined"
	VILLAGER  = "Villager"
	WEREWOLF  = "Werewolf"
	HUNTER    = "Hunter"
)

type Role interface {
	Name() string
	ID() int
	SetID(int)
	Alive() bool
	SetAlive(bool)
	OnGameStarted(events.GameStartedEvent)
	OnNightStarted(events.NightStartedEvent)
	OnWerewolfVote(events.WerewolfVoteEvent)
	OnDailyVote(events.DailyVoteEvent)
}
