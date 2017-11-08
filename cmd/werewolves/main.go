package main

import (
	"github.com/GSamuel/werewolves-of-millers-hollow/deck"
	"github.com/GSamuel/werewolves-of-millers-hollow/game"
	"github.com/GSamuel/werewolves-of-millers-hollow/roles"
)

func main() {

	were := roles.WEREWOLF
	vil := roles.VILLAGER
	hun := roles.HUNTER
	heal := roles.HEALER
	//elder := roles.VILLAGE_ELDER
	witch := roles.WITCH
	cupid := roles.CUPID
	seer := roles.SEER
	slut := roles.SLUT

	deck := deck.New()
	deck.Add(vil)
	deck.Add(hun)
	deck.Add(vil)
	deck.Add(were)
	deck.Add(heal)
	deck.Add(witch)
	deck.Add(cupid)
	deck.Add(seer)
	deck.Add(slut)
	deck.Shuffle()

	g := game.New(deck, roles.NewFactory())
	g.Run()
}
