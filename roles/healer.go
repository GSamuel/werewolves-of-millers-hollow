package roles

import (
	"fmt"
	"github.com/GSamuel/werewolves-of-millers-hollow/events"
	"github.com/GSamuel/werewolves-of-millers-hollow/game"
)

type Healer struct {
	*BaseRole
	target int
	round  int
}

func (h *Healer) Name() string {
	return HEALER
}

func (h *Healer) OnNightStarted(g *game.Game, e *events.NightStartedEvent) {
	if !h.Alive() {
		return
	}

	p := g.Player(h.ID())

	fmt.Printf("Healer %v:", h.ID())
	h.target = p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, h.target)
	})
	h.round = g.Round()
}

func (h *Healer) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {
	if e.ID() == h.target && e.WerewolfAttack() && g.Round() == h.round {
		e.PreventDeath()
	}
}
