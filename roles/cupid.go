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
	p := g.Player(c.ID())

	fmt.Printf("Cupid %v, lover 1:", c.ID())
	c.lover1 = p.ReadInt(func(i int) bool {
		return Alive(g, i)
	})

	fmt.Printf("Cupid %v, lover 2:", c.ID())
	c.lover2 = p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, c.lover1)
	})

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
