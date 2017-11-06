package game

import ()

type Player struct {
	Role
}

func NewPlayer(role Role) *Player {
	return &Player{role}
}
