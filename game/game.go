package game

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/roles"
)

type Game struct {
	*PlayerGroup
	*roles.EventSystem
}

func (g *Game) Run() {
	g.startGame()

	for !g.isOver() {

		g.printPlayers()

		g.startNight()

		votes := g.startWerewolfVote()
		if len(votes) == 1 {
			id := votes[0]
			g.Player(id).Die(true)
		}

		g.printPlayers()

		if g.isOver() {
			break
		}

		votes = g.starDailyVote()

		if len(votes) == 1 {
			id := votes[0]
			g.Player(id).Die(false)
		}
	}

	g.printPlayers()
	fmt.Println("The game is over")
}

func (g *Game) startGame() {
	fmt.Println("The Game has started.")
	event := events.NewGameStartedEvent()
	g.GameStartedEvent(event)
}

func (g *Game) startNight() {
	fmt.Println("The night starts, Everyone goes to sleep.")
	event := events.NewNightStartedEvent()
	g.NightStartedEvent(event)
}

func (g *Game) startWerewolfVote() []int {
	fmt.Println("The werewolves wake up and choose a victim.")
	event := events.NewWerewolfVoteEvent()
	g.WerewolfVoteEvent(event)

	return event.Result()
}

func (g *Game) starDailyVote() []int {
	fmt.Println("Vote for a player to be executed.")
	event := events.NewDailyVoteEvent()
	g.DailyVoteEvent(event)
	return event.Result()
}

func (g *Game) isOver() bool {

	stillHumans := false
	stillWerewolves := false

	for i := 0; i < g.Count(); i++ {
		if g.Player(i).Alive() {
			if g.Player(i).Name() == roles.WEREWOLF {
				stillWerewolves = true
			} else {
				stillHumans = true
			}
		}
	}

	return !(stillHumans && stillWerewolves)
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

func NewGame(group *PlayerGroup, eventSystem *roles.EventSystem) *Game {
	return &Game{group, eventSystem}
}
