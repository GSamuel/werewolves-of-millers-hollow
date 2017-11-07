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

	fmt.Printf("Werewolf %v:", w.ID())

	i := p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, p.ID()) && g.Player(i).Alliance() != game.ALLIANCE_EVIL
	})

	e.Vote(i, 1)

}
