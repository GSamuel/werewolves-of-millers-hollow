package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type Slut struct {
	*BaseRole
	target int
}

func (s *Slut) Name() string {
	return SLUT
}

func (s *Slut) OnNightStarted(g *game.Game, e *events.NightStartedEvent) {

	if !s.Alive() {
		return
	}

	p := g.Player(s.ID())

	done := false
	for !done {
		fmt.Printf("Slut %v:", s.ID())
		i := p.ReadInt()
		if i >= 0 && i < g.Count() {
			if g.Player(i).Alive() && i != s.target && i != s.ID() {
				s.target = i
				done = true
			}
		}
	}
}

func (s *Slut) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {
	if e.ID() == s.ID() && e.WerewolfAttack() {
		e.PreventDeath()
	}

	if e.ID() == s.target && e.WerewolfAttack() {
		s.Die(g, false)
	}

}
