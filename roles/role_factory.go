package roles

import (
	"fmt"
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
	}
	return nil, fmt.Errorf("Role %v does not exist in role factory", role)
}
