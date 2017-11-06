package game

import (
	"github.com/GSamuel/werewolvesmillershollow/deck"
)

type RoleFactory interface {
	New(role string, id int) (Role, error)
}

func New(deck *deck.Deck, roleFactory RoleFactory) *Game {
	players := make([]*Player, 0, deck.Count())

	for i := 0; i < deck.Count(); i++ {
		role, err := roleFactory.New(deck.Get(i), i)

		if err != nil {
			panic(err)
		}

		players = append(players, NewPlayer(role))
	}

	group := &PlayerGroup{players}
	eventSystem := NewEventSystem(group)

	for i := 0; i < group.Count(); i++ {
		group.Player(i).SetEventSystem(eventSystem)
	}

	return NewGame(group, eventSystem)
}
