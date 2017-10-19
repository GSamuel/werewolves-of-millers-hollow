package game

import (
	"math/rand"
)

type Deck struct {
	roles    []string
	shuffled []int
}

func (d *Deck) Add(role string) {
	d.roles = append(d.roles, role)
}

func (d *Deck) Shuffle() {
	d.shuffled = rand.Perm(len(d.roles))
}

func (d *Deck) Count() int {
	return len(d.roles)
}

func (d *Deck) Get(i int) string {
	return d.roles[d.shuffled[i]]
}

func NewDeck() *Deck {
	return &Deck{make([]string, 0), make([]int, 0)}
}
