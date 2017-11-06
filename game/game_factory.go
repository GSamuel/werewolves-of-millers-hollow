package game

import (
	"github.com/GSamuel/werewolvesmillershollow/deck"
	"github.com/GSamuel/werewolvesmillershollow/io"
)

type RoleFactory interface {
	New(role string, id int) (Role, error)
}

func New(deck *deck.Deck, roleFactory RoleFactory) *Game {
	players := make([]*Player, 0, deck.Count())

	console := &io.Console{}

	for i := 0; i < deck.Count(); i++ {
		role, err := roleFactory.New(deck.Get(i), i)

		if err != nil {
			panic(err)
		}

		players = append(players, NewPlayer(role, console, console))
	}

	group := &PlayerGroup{players}
	eventSystem := NewEventSystem()

	game := NewGame(group, eventSystem)

	return game
}
