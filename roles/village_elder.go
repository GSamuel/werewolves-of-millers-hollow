package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
)

type VillageElder struct {
	*BaseRole
	extraLife bool
}

func (v *VillageElder) Name() string {
	return VILLAGE_ELDER
}

func (v *VillageElder) OnPlayerDead(e *events.PlayerDeadEvent) {
	if v.extraLife && e.ID() == v.ID() && e.WerewolfAttack() {
		v.extraLife = false
		e.PreventDeath()
	}
}
