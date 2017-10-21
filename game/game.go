package game

import (
	"fmt"
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/roles"
)

type Game struct {
	*PlayerGroup
	*events.EventSystem
}

func (g *Game) Run() {
	g.startGame()

	for !g.isOver() {

		g.printPlayers()

		g.startNight()

		votes := g.startWerewolfVote()
		if len(votes) == 1 && g.Player(votes[0]).Alive() && g.Player(votes[0]).Name() != roles.WEREWOLF {
			id := votes[0]
			g.startPlayerDeath(id)
			if !g.Player(id).Alive() {
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

		if len(votes) == 1 && g.Player(votes[0]).Alive() {
			id := votes[0]
			g.startPlayerDeath(id)
			if !g.Player(id).Alive() {
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
	g.Player(id).SetAlive(false)
	event := events.NewPlayerDeadEvent(id)
	g.PlayerDeadEvent(event)
}

func (g *Game) startPlayerRevealed(id int) {
	event := events.NewPlayerRevealedEvent(id)
	g.PlayerRevealedEvent(event)
	fmt.Println("Player ", id, " has died. It was a ", g.players[id].Name())
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
}

func NewGame(players []*Player) *Game {
	group := &PlayerGroup{players}
	eventSystem := events.NewEventSystem()
	for i := 0; i < group.Count(); i++ {
		eventSystem.Add(group.Player(i))
	}
	return &Game{group, eventSystem}
}
