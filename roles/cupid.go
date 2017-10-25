package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type Cupid struct {
	*BaseRole
	lover1 int
	lover2 int
}

func (c *Cupid) Name() string {
	return CUPID
}

func (c *Cupid) OnGameStarted(e *events.GameStartedEvent) {
	done := false

	for !done {
		fmt.Printf("Cupid %v, lover 1:", c.ID())
		i := utils.ReadInput()
		if i >= 0 && i < c.eventSystem.Count() {
			if c.eventSystem.Role(i).Alive() {
				c.lover1 = i
				done = true
			}
		}
	}

	done = false

	for !done {
		fmt.Printf("Cupid %v, lover 2:", c.ID())
		i := utils.ReadInput()
		if i >= 0 && i < c.eventSystem.Count() {
			if c.eventSystem.Role(i).Alive() && i != c.lover1 {
				c.lover2 = i
				done = true
			}
		}
	}

	fmt.Printf("Player %v and %v fell in love\n", c.lover1, c.lover2)
}

func (c *Cupid) OnPlayerRevealed(e *events.PlayerRevealedEvent) {
	if e.ID() == c.lover1 && c.eventSystem.Role(c.lover2).Alive() {
		c.eventSystem.Role(c.lover2).Die(false)
	}

	if e.ID() == c.lover2 && c.eventSystem.Role(c.lover1).Alive() {
		c.eventSystem.Role(c.lover1).Die(false)
	}
}
