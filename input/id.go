package input

import ()

type Id struct {
	id int
}

func (i *Id) Validate(role Role) bool {
	return role.ID() != i.id
}

func NewId(id int) *Id {
	return &Id{id}
}
