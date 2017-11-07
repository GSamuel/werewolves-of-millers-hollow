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

	fmt.Printf("Slut %v:", s.ID())
	i := p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, p.ID()) && NotEqual(g, i, s.target)
	})
	s.target = i
}

func (s *Slut) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {
	if e.ID() == s.ID() && e.WerewolfAttack() {
		e.PreventDeath()
	}

	if e.ID() == s.target && e.WerewolfAttack() {
		s.Die(g, false)
	}

}
