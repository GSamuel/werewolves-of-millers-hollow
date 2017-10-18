package main

import ()

func main() {

	deck := NewDeck()
	deck.add(HUMAN)
	deck.add(HUMAN)
	deck.add(HUMAN)
	deck.add(HUMAN)
	deck.add(HUMAN)
	deck.add(HUMAN)
	deck.add(WEREWOLF)
	deck.add(WEREWOLF)
	deck.shuffle()

	players := make([]*Player, 0, deck.count())

	for i := 0; i < deck.count(); i++ {
		players = append(players, &Player{true, deck.get(i)})
	}

	game := &Game{players}
	game.run()
}
