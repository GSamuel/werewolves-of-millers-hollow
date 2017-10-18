package main

import ()

func main() {

	deck := NewDeck()
	deck.add(&Villager{})
	deck.add(&Villager{})
	deck.add(&Villager{})
	deck.add(&Villager{})
	deck.add(&Villager{})
	deck.add(&Villager{})
	deck.add(&Werewolf{})
	deck.add(&Werewolf{})
	deck.shuffle()

	players := make([]*Player, 0, deck.count())

	for i := 0; i < deck.count(); i++ {
		players = append(players, &Player{deck.get(i), true})
	}

	game := &Game{players}
	game.run()
}
