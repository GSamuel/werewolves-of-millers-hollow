package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type Healer struct {
	*BaseRole
	target int
}

func (h *Healer) Name() string {
	return HEALER
}

func (h *Healer) OnNightStarted(e *events.NightStartedEvent) {
	if !h.Alive() {
		return
	}

	done := false

	for !done {
		fmt.Printf("Healer %v:", h.ID())
		i := utils.ReadInput()
		if i >= 0 && i < h.eventSystem.Count() {
			if h.eventSystem.Role(i).Alive() && i != h.target {
				h.target = i
				done = true
			}
		}
	}

}

func (h *Healer) OnPlayerDead(e *events.PlayerDeadEvent) {
	if e.ID() == h.target && e.WerewolfAttack() {
		e.PreventDeath()
	}
}
