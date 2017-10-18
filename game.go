package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Game struct {
	players []*Player
}

func (g *Game) run() {
	for !g.isOver() {
		g.printPlayers()
		fmt.Println("The night starts, Everyone goes to sleep.")
		fmt.Println("The werewolves wake up and choose a victim.")
		i, err := g.readInput()

		if err != nil {
			panic(err)
		}

		if g.players[i].alive && g.players[i].alliance != WEREWOLF {
			fmt.Println("Player ", i, " has died this night. It was a ", g.players[i].alliance)
			g.players[i].alive = false
		} else {
			fmt.Println("There are no victims this night.")
		}

		g.printPlayers()

		if g.isOver() {
			break
		}
		fmt.Println("Vote for a player to be executed.")

		i, err = g.readInput()

		if err != nil {
			panic(err)
		}

		if g.players[i].alive {
			g.players[i].alive = false
			fmt.Println("Player ", i, " has died by public execution. It was a ", g.players[i].alliance)
		}
	}

	g.printPlayers()
	fmt.Println("The game is over")

	//night
	//wake werewolves
	//werewolves vote on victim
	//day
	//announce werewolve victim
	//check win conditions
	//vote for player to be eliminated
	//announce death of player
	//check win conditions
	//Repeat
}

func (g *Game) isOver() bool {

	stillHumans := false
	stillWerewolves := false

	for i := 0; i < len(g.players); i++ {
		if g.players[i].alive {
			if g.players[i].alliance == WEREWOLF {
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
		if g.players[i].alive {
			fmt.Println("Player ", i, " alliance:", g.players[i].alliance)
		}
	}
}

func (g *Game) readInput() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	return strconv.Atoi(text)
}
