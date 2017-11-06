package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type Hunter struct {
	*BaseRole
}

func (h *Hunter) Name() string {
	return HUNTER
}

func (h *Hunter) OnPlayerRevealed(g *game.Game, e *events.PlayerRevealedEvent) {
	if e.ID() != h.ID() {
		return
	}

	p := g.Player(h.ID())

	done := false
	for !done {
		fmt.Printf("Hunter %v:", h.ID())
		i := p.ReadInt()
		if i >= 0 && i < g.Count() {
			if g.Player(i).Alive() {
				g.Player(i).Die(g, false)
				done = true
			}
		}
	}

}
