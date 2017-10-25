package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type Slut struct {
	*BaseRole
	target int
}

func (s *Slut) Name() string {
	return SLUT
}

func (s *Slut) OnNightStarted(e *events.NightStartedEvent) {
	done := false
	for !done {
		fmt.Printf("Slut %v:", s.ID())
		i := utils.ReadInput()
		if i >= 0 && i < s.eventSystem.Count() {
			if s.eventSystem.Role(i).Alive() && i != s.target && i != s.ID() {
				s.target = i
				done = true
			}
		}
	}
}

func (s *Slut) OnPlayerDead(e *events.PlayerDeadEvent) {
	if e.ID() == s.ID() && e.WerewolfAttack() {
		e.PreventDeath()
	}

	if e.ID() == s.target && e.WerewolfAttack() {
		s.Die(false)
	}

}
