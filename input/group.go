package input

import ()

type Group interface {
	Count() int
	Role(int) Role
}
