package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

//Basic Werewolf role
type Werewolf struct {
	*BaseRole
}

func (w *Werewolf) Name() string {
	return WEREWOLF
}

func (w *Werewolf) OnWerewolfVote(e events.WerewolfVoteEvent) {
	if !w.Alive() {
		return
	}

	fmt.Printf("Werewolf %v:", w.ID())
	i := utils.ReadInput()
	e.Vote(i, 1)
}
