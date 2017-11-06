package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type state struct {
	id       int
	alive    bool
	alliance game.Alliance
}

func (s *state) ID() int {
	return s.id
}

func (s *state) SetID(id int) {
	s.id = id
}

func (s *state) Alive() bool {
	return s.alive
}

func (s *state) SetAlive(alive bool) {
	s.alive = alive
}

func (s *state) Alliance() game.Alliance {
	return s.alliance
}

func (s *state) SetAlliance(alliance game.Alliance) {
	s.alliance = alliance
}

func NewState(id int) *state {
	return &state{id, true, game.ALLIANCE_GOOD}
}
