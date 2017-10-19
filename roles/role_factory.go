package roles

import (
	"fmt"
)

func New(role string) (Role, error) {

	switch role {
	case WEREWOLF:
		return &Werewolf{}, nil
	case VILLAGER:
		return &Villager{}, nil
	case HUNTER:
		return &Hunter{}, nil
	}
	return nil, fmt.Errorf("Role %v does not exist in role factory", role)
}
