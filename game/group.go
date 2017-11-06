package game

import ()

type Group interface {
	Count() int
	Player(int) *Player
}
