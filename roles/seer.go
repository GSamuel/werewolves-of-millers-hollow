package roles

import (
	"fmt"
	"github.com/GSamuel/werewolves-of-millers-hollow/events"
	"github.com/GSamuel/werewolves-of-millers-hollow/game"
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

	fmt.Printf("Seer %v:", s.ID())
	i := p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, p.ID())
	})

	fmt.Printf("Player %v is a %v\n", i, g.Player(i).Name())
}
