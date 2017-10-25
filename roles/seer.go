package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type Seer struct {
	*BaseRole
}

func (s *Seer) Name() string {
	return SEER
}

func (s *Seer) OnNightStarted(e *events.NightStartedEvent) {

	if !s.Alive() {
		return
	}

	done := false
	for !done {
		fmt.Printf("Seer %v:", s.ID())
		i := utils.ReadInput()
		if i >= 0 && i < s.eventSystem.Count() {
			if s.eventSystem.Role(i).Alive() {
				fmt.Printf("Player %v is a %v\n", i, s.eventSystem.Role(i).Name())
				done = true
			}
		}
	}
}
