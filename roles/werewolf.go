package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

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

	done := false

	for !done {
		fmt.Printf("Werewolf %v:", w.ID())
		i := utils.ReadInput()
		if i >= 0 && i < w.eventSystem.Count() {
			if w.eventSystem.Role(i).Alive() && w.eventSystem.Role(i).Name() != WEREWOLF {
				e.Vote(i, 1)
				done = true
			}
		}
	}

}
