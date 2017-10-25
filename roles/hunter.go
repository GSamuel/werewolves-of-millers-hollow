package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/input"
)

type Hunter struct {
	*BaseRole
}

func (h *Hunter) Name() string {
	return HUNTER
}

func (h *Hunter) OnPlayerRevealed(e *events.PlayerRevealedEvent) {
	if e.ID() != h.ID() {
		return
	}

	done := false
	for !done {
		fmt.Printf("Hunter %v:", h.ID())
		i, _ := input.ReadInput()
		if i >= 0 && i < h.eventSystem.Count() {
			if h.eventSystem.Role(i).Alive() {
				h.eventSystem.Role(i).Die(false)
				done = true
			}
		}
	}

}
