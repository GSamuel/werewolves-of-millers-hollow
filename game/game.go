package game

import (
	"fmt"
	"github.com/GSamuel/werewolves-of-millers-hollow/events"
)

type Game struct {
	*PlayerGroup
	*EventSystem
	round int
}

func (g *Game) Run() {

	g.printPlayers()
	g.startGame()

	for !g.isOver() {

		g.startNight()

		votes := g.startWerewolfVote()
		if len(votes) == 1 {
			id := votes[0]
			g.Player(id).Die(g, true)
		}

		g.printPlayers()

		if g.isOver() {
			break
		}

		votes = g.starDailyVote()

		if len(votes) == 1 {
			id := votes[0]
			g.Player(id).Die(g, false)
		}

		g.printPlayers()
		g.round = g.round + 1
	}

	fmt.Println("The game is over")
}

func (g *Game) startGame() {
	fmt.Println("The Game has started.")
	event := events.NewGameStartedEvent()
	g.GameStartedEvent(g, event)
}

func (g *Game) startNight() {
	fmt.Println("The night starts, Everyone goes to sleep.")
	event := events.NewNightStartedEvent()
	g.NightStartedEvent(g, event)
}

func (g *Game) startWerewolfVote() []int {
	fmt.Println("The werewolves wake up and choose a victim.")
	event := events.NewWerewolfVoteEvent()
	g.WerewolfVoteEvent(g, event)

	return event.Result()
}

func (g *Game) starDailyVote() []int {
	fmt.Println("Vote for a player to be executed.")
	event := events.NewDailyVoteEvent()
	g.DailyVoteEvent(g, event)
	return event.Result()
}

func (g *Game) isOver() bool {

	stillHumans := false
	stillWerewolves := false

	for i := 0; i < g.Count(); i++ {
		if g.Player(i).Alive() {
			if g.Player(i).Alliance() == ALLIANCE_EVIL { //TODO: Wincondition check should be the responsible for the roles.
				stillWerewolves = true
			} else {
				stillHumans = true
			}
		}
	}

	return !(stillHumans && stillWerewolves)
}

func (g *Game) Round() int {
	return g.round
}

func (g *Game) printPlayers() {
	fmt.Println()
	for i := 0; i < g.Count(); i++ {
		if g.Player(i).Alive() {
			fmt.Println("Player ", i, " alliance:", g.Player(i).Name())
		}
	}
	fmt.Println()
}

func NewGame(group *PlayerGroup, eventSystem *EventSystem) *Game {
	return &Game{group, eventSystem, 0}
}
