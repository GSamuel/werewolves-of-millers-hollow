package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type Witch struct {
	*BaseRole
	potionOfLife bool
	poison       bool
}

func (w *Witch) Name() string {
	return WITCH
}

func (w *Witch) OnPlayerDead(e *events.PlayerDeadEvent) {

	if !e.WerewolfAttack() {
		return
	}

	if w.potionOfLife {
		fmt.Printf("Witch %v: safe %v with life potion? ", w.ID(), e.ID())
		revive := utils.ReadYesNo()
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
	useIt := utils.ReadYesNo()
	if !useIt {
		return
	}

	done := false
	for !done {
		fmt.Printf("Witch %v:", w.ID())
		i := utils.ReadInput()
		if i >= 0 && i < w.eventSystem.Count() {
			if w.eventSystem.Role(i).Alive() {
				w.eventSystem.Role(i).Die(false)
				done = true
				w.poison = false
			}
		}
	}

}
