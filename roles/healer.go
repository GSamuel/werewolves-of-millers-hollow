package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type Healer struct {
	*BaseRole
	target int
}

func (h *Healer) Name() string {
	return HEALER
}

func (h *Healer) OnNightStarted(g *game.Game, e *events.NightStartedEvent) {
	if !h.Alive() {
		return
	}

	p := g.Player(h.ID())

	done := false

	for !done {
		fmt.Printf("Healer %v:", h.ID())
		i := p.ReadInt()
		if i >= 0 && i < g.Count() {
			if g.Player(i).Alive() && i != h.target {
				h.target = i
				done = true
			}
		}
	}

}

func (h *Healer) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {
	if e.ID() == h.target && e.WerewolfAttack() {
		e.PreventDeath()
	}
}
