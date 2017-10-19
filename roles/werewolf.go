package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

//Basic Werewolf role
type Werewolf struct {
	BaseRole
}

func (w *Werewolf) Name() string {
	return WEREWOLF
}

func (w *Werewolf) OnWerewolfVote(e events.WerewolfVoteEvent) {
	i := utils.ReadInput()
	e.Vote(i, 1)
}
