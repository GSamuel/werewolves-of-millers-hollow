package game

import (
	"github.com/GSamuel/werewolvesmillershollow/roles"
)

func New(deck *Deck) *Game {
	players := make([]*Player, 0, deck.Count())

	for i := 0; i < deck.Count(); i++ {
		role, err := roles.New(deck.Get(i), i)

		if err != nil {
			panic(err)
		}

		players = append(players, NewPlayer(role))
	}

	return NewGame(players)
}
