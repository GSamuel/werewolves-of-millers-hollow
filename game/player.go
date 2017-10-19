package game

import (
	"github.com/GSamuel/werewolvesmillershollow/roles"
)

type Player struct {
	roles.Role
}

func NewPlayer(role roles.Role) *Player {
	return &Player{role}
}
