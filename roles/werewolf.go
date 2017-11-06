package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type Werewolf struct {
	*BaseRole
}

func (w *Werewolf) Name() string {
	return WEREWOLF
}

func (w *Werewolf) OnWerewolfVote(g *game.Game, e *events.WerewolfVoteEvent) {
	if !w.Alive() {
		return
	}

	p := g.Player(w.ID())

	done := false

	for !done {
		fmt.Printf("Werewolf %v:", w.ID())
		i := p.ReadInt()
		if i >= 0 && i < g.Count() {
			if g.Player(i).Alive() && g.Player(i).Name() != WEREWOLF {
				e.Vote(i, 1)
				done = true
			}
		}
	}

}
