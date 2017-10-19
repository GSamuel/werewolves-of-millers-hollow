package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type BaseRole struct {
	*State
}

func (b *BaseRole) Name() string {
	return UNDEFINED
}

func (b *BaseRole) OnNightStarted(e events.NightStartedEvent) {
}

func (b *BaseRole) OnGameStarted(e events.GameStartedEvent) {
}

func (b *BaseRole) OnWerewolfVote(e events.WerewolfVoteEvent) {
}

func (b *BaseRole) OnPlayerDead(e events.PlayerDeadEvent) {

}
func (b *BaseRole) OnPlayerRevealed(e events.PlayerRevealedEvent) {

}

func (b *BaseRole) OnDailyVote(e events.DailyVoteEvent) {
	if !b.Alive() {
		return
	}

	fmt.Printf("Player %v:", b.ID())
	i := utils.ReadInput()
	e.Vote(i, 1)

}
