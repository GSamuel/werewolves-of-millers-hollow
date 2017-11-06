package game

import ()

type Group interface {
	Count() int
	Role(int) Role
}
