package roles

import (
	"fmt"
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
)

func New(role string, id int) (Role, error) {

	base := &BaseRole{NewState(id)}

	switch role {
	case WEREWOLF:
		return &Werewolf{base}, nil
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
	}
	return nil, fmt.Errorf("Role %v does not exist in role factory", role)
}
