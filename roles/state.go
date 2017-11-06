package roles

import (
	"github.com/GSamuel/werewolvesmillershollow/game"
)

type state struct {
	id          int
	alive       bool
	eventSystem *game.EventSystem
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

func (s *state) SetEventSystem(eventSystem *game.EventSystem) {
	s.eventSystem = eventSystem
}

func NewState(id int) *state {
	return &state{id, true, nil}
}
