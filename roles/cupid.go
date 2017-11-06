package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type Cupid struct {
	*BaseRole
	lover1 int
	lover2 int
}

func (c *Cupid) Name() string {
	return CUPID
}

func (c *Cupid) OnGameStarted(g *game.Game, e *events.GameStartedEvent) {
	done := false

	p := g.Player(c.ID())

	for !done {
		fmt.Printf("Cupid %v, lover 1:", c.ID())
		i := p.ReadInt()
		if i >= 0 && i < g.Count() {
			if g.Player(i).Alive() {
				c.lover1 = i
				done = true
			}
		}
	}

	done = false

	for !done {
		fmt.Printf("Cupid %v, lover 2:", c.ID())
		i := p.ReadInt()
		if i >= 0 && i < g.Count() {
			if g.Player(i).Alive() && i != c.lover1 {
				c.lover2 = i
				done = true
			}
		}
	}

	fmt.Printf("Player %v and %v fell in love\n", c.lover1, c.lover2)
}

func (c *Cupid) OnPlayerRevealed(g *game.Game, e *events.PlayerRevealedEvent) {
	if e.ID() == c.lover1 && g.Player(c.lover2).Alive() {
		g.Player(c.lover2).Die(g, false)
	}

	if e.ID() == c.lover2 && g.Player(c.lover1).Alive() {
		g.Player(c.lover1).Die(g, false)
	}
}
