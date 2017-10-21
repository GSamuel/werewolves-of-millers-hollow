package roles

import ()

type Role interface {
	Name() string
	ID() int
	SetID(int)
	Alive() bool
	SetAlive(bool)
	SetEventSystem(*EventSystem)
	Die(bool)
	EventListener
}
