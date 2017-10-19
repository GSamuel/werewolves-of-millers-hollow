package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type BaseRole struct {
	playerId int
}

func (b *BaseRole) Name() string {
	return UNDEFINED
}

func (b *BaseRole) SetPlayerId(id int) {
	b.playerId = id
}

func (b *BaseRole) OnNightStarted(e events.NightStartedEvent) {
}

func (b *BaseRole) OnGameStarted(e events.GameStartedEvent) {
}

func (b *BaseRole) OnWerewolfVote(e events.WerewolfVoteEvent) {
}

func (b *BaseRole) OnDailyVote(e events.DailyVoteEvent) {
	i := utils.ReadInput()
	e.Vote(i, 1)
}
