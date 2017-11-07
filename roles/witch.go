package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type Witch struct {
	*BaseRole
	potionOfLife bool
	poison       bool
}

func (w *Witch) Name() string {
	return WITCH
}

func (w *Witch) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {

	if !w.Alive() {
		return
	}

	if !e.WerewolfAttack() {
		return
	}

	p := g.Player(w.ID())

	if w.potionOfLife {
		fmt.Printf("Witch %v: safe %v with life potion? ", w.ID(), e.ID())
		revive := p.ReadBool()
		if revive {
			e.PreventDeath()
			w.potionOfLife = false
			fmt.Println("death prevented by witch")
		}
	}

	if !w.poison {
		return
	}

	fmt.Printf("Witch %v: want to use poison?", w.ID())
	useIt := p.ReadBool()
	if !useIt {
		return
	}

	fmt.Printf("Witch %v:", w.ID())
	i := p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, p.ID())
	})

	g.Player(i).Die(g, false)
	w.poison = false

}
