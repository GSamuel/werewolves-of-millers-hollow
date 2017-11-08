package roles

import (
	"fmt"
	"github.com/GSamuel/werewolves-of-millers-hollow/events"
	"github.com/GSamuel/werewolves-of-millers-hollow/game"
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

	fmt.Printf("Hunter %v:", h.ID())
	i := p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, p.ID())
	})
	g.Player(i).Die(g, false)

}
