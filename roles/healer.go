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

	fmt.Printf("Healer %v:", h.ID())
	h.target = p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, h.target)
	})

}

func (h *Healer) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {
	if e.ID() == h.target && e.WerewolfAttack() {
		e.PreventDeath()
	}
}
