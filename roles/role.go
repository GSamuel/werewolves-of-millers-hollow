package roles

import ()

const (
	UNDEFINED = "Undefined"
	VILLAGER  = "Villager"
	WEREWOLF  = "Werewolf"
	HUNTER    = "Hunter"
	HEALER    = "Healer"
)

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
