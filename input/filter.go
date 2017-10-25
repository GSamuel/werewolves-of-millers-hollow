package input

import ()

type Filter interface {
	Validate(Role) bool
}

type Role interface {
	Name() string
	ID() int
	Alive() bool
}
