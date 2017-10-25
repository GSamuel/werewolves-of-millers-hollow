package input

import ()

type Alive struct {
}

func (a *Alive) Validate(role Role) bool {
	return role.Alive()
}

func NewAlive() *Alive {
	return &Alive{}
}
