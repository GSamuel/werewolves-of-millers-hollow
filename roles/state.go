package roles

import ()

type State struct {
	id          int
	alive       bool
	eventSystem *EventSystem
}

func (s *State) ID() int {
	return s.id
}

func (s *State) SetID(id int) {
	s.id = id
}

func (s *State) Alive() bool {
	return s.alive
}

func (s *State) SetAlive(alive bool) {
	s.alive = alive
}

func (s *State) SetEventSystem(eventSystem *EventSystem) {
	s.eventSystem = eventSystem
}

func NewState(id int) *State {
	return &State{id, true, nil}
}
