package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/utils"
)

type Hunter struct {
	*BaseRole
}

func (h *Hunter) Name() string {
	return HUNTER
}

func (h *Hunter) OnPlayerRevealed(e events.PlayerRevealedEvent) {
	if e.ID() != h.ID() {
		return
	}

	done := false
	for !done {
		fmt.Printf("Hunter %v:", h.ID())
		i := utils.ReadInput()
		if i >= 0 && i < h.eventSystem.Count() {
			if h.eventSystem.Role(i).Alive() {
				h.eventSystem.Role(i).Die()
				done = true
			}
		}
	}

}
