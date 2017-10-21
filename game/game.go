package game

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/roles"
)

type Game struct {
	players []*Player
}

func (g *Game) Run() {
	g.startGame()

	for !g.isOver() {

		g.printPlayers()

		g.startNight()

		votes := g.startWerewolfVote()

		if len(votes) == 1 && g.players[votes[0]].Alive() && g.players[votes[0]].Name() != roles.WEREWOLF {
			id := votes[0]
			g.startPlayerDeath(id)
			if !g.players[id].Alive() {
				g.startPlayerRevealed(id)
			}
		} else {
			fmt.Println("There are no victims this night.")
		}

		g.printPlayers()

		if g.isOver() {
			break
		}

		votes = g.starDailyVote()

		if len(votes) == 1 && g.players[votes[0]].Alive() {
			id := votes[0]
			g.startPlayerDeath(id)
			if !g.players[id].Alive() {
				g.startPlayerRevealed(id)
			}
		}
	}

	g.printPlayers()
	fmt.Println("The game is over")

	//night
	//wake werewolves
	//werewolves vote on victim
	//day
	//announce werewolf victim
	//check win conditions
	//vote for player to be eliminated
	//announce death of player
	//check win conditions
	//Repeat
}

func (g *Game) startPlayerDeath(id int) {
	g.players[id].SetAlive(false)
	event := events.NewPlayerDeadEvent(id)
	for i := 0; i < len(g.players); i++ {
		g.players[i].OnPlayerDead(event)
	}
	//
}

func (g *Game) startPlayerRevealed(id int) {
	event := events.NewPlayerRevealedEvent(id)
	for i := 0; i < len(g.players); i++ {
		g.players[i].OnPlayerRevealed(event)
	}
	fmt.Println("Player ", id, " has died. It was a ", g.players[id].Name())
}

func (g *Game) startGame() {
	fmt.Println("The Game has started.")
	event := events.NewGameStartedEvent()
	for i := 0; i < len(g.players); i++ {
		g.players[i].OnGameStarted(event)
	}
}

func (g *Game) startNight() {
	fmt.Println("The night starts, Everyone goes to sleep.")
	event := events.NewNightStartedEvent()
	for i := 0; i < len(g.players); i++ {
		g.players[i].OnNightStarted(event)
	}
}

func (g *Game) startWerewolfVote() []int {

	fmt.Println("The werewolves wake up and choose a victim.")
	event := events.NewWerewolfVoteEvent()
	for i := 0; i < len(g.players); i++ {
		g.players[i].OnWerewolfVote(event)
	}

	return event.Result()
}

func (g *Game) starDailyVote() []int {
	fmt.Println("Vote for a player to be executed.")
	event := events.NewDailyVoteEvent()

	for i := 0; i < len(g.players); i++ {
		g.players[i].OnDailyVote(event)
	}

	return event.Result()
}

func (g *Game) isOver() bool {

	stillHumans := false
	stillWerewolves := false

	for i := 0; i < len(g.players); i++ {
		if g.players[i].Alive() {
			if g.players[i].Name() == roles.WEREWOLF {
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
	for i := 0; i < len(g.players); i++ {
		if g.players[i].Alive() {
			fmt.Println("Player ", i, " alliance:", g.players[i].Name())
		}
	}
}

func NewGame(players []*Player) *Game {
	return &Game{players}
}
