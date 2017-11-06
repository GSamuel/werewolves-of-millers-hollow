package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type Seer struct {
	*BaseRole
}

func (s *Seer) Name() string {
	return SEER
}

func (s *Seer) OnNightStarted(g *game.Game, e *events.NightStartedEvent) {

	if !s.Alive() {
		return
	}

	p := g.Player(s.ID())

	done := false
	for !done {
		fmt.Printf("Seer %v:", s.ID())
		i := p.ReadInt()
		if i >= 0 && i < g.Count() {
			if g.Player(i).Alive() {
				fmt.Printf("Player %v is a %v\n", i, g.Player(i).Name())
				done = true
			}
		}
	}
}
