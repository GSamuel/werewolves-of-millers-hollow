package roles

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/game"
)

const (
	UNDEFINED     = "Undefined"
	VILLAGER      = "Villager"
	WEREWOLF      = "Werewolf"
	HUNTER        = "Hunter"
	HEALER        = "Healer"
	VILLAGE_ELDER = "Village Elder"
	WITCH         = "Witch"
	CUPID         = "Cupid"
	SEER          = "Seer"
	SLUT          = "Slut"
)

type RoleFactory struct {
}

func (r RoleFactory) New(role string, id int) (game.Role, error) {

	base := &BaseRole{NewState(id)}
	var result game.Role

	switch role {
	case WEREWOLF:
		result = &Werewolf{base}
		result.SetAlliance(game.ALLIANCE_EVIL)
		return result, nil
	case VILLAGER:
		return &Villager{base}, nil
	case HUNTER:
		return &Hunter{base}, nil
	case HEALER:
		return &Healer{base, -1}, nil
	case VILLAGE_ELDER:
		return &VillageElder{base, true}, nil
	case WITCH:
		return &Witch{base, true, true}, nil
	case CUPID:
		return &Cupid{base, -1, -1}, nil
	case SEER:
		return &Seer{base}, nil
	case SLUT:
		return &Slut{base, -1}, nil
	}
	return nil, fmt.Errorf("Role %v does not exist in role factory", role)
}

func NewFactory() RoleFactory {
	return RoleFactory{}
}
