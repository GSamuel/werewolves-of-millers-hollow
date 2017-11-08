package game

import (
	"github.com/GSamuel/werewolves-of-millers-hollow/io"
)

type Player struct {
	Role
	io.PlayerWriter
	io.PlayerReader
}

func NewPlayer(role Role, w io.PlayerWriter, r io.PlayerReader) *Player {
	return &Player{role, w, r}
}
