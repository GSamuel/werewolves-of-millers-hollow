package main

import (
	"github.com/GSamuel/werewolvesmillershollow/game"
	"github.com/GSamuel/werewolvesmillershollow/roles"
)

func main() {

	were := roles.WEREWOLF
	vil := roles.VILLAGER
	hun := roles.HUNTER

	deck := game.NewDeck()
	deck.Add(vil)
	deck.Add(vil)
	deck.Add(vil)
	deck.Add(vil)
	deck.Add(hun)
	deck.Add(vil)
	deck.Add(were)
	deck.Add(were)
	deck.Shuffle()

	players := make([]*game.Player, 0, deck.Count())

	for i := 0; i < deck.Count(); i++ {
		role, err := roles.New(deck.Get(i))

		if err != nil {
			panic(err)
		}

		players = append(players, game.NewPlayer(role))
	}

	game := game.NewGame(players)
	game.Run()
}
