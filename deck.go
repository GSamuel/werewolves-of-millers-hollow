package main

import (
	"math/rand"
)

type Deck struct {
	roles    []Role
	shuffled []int
}

func (d *Deck) add(role Role) {
	d.roles = append(d.roles, role)
}

func (d *Deck) shuffle() {
	d.shuffled = rand.Perm(len(d.roles))
}

func (d *Deck) count() int {
	return len(d.roles)
}

func (d *Deck) get(i int) Role {
	return d.roles[d.shuffled[i]]
}

func NewDeck() *Deck {
	return &Deck{make([]Role, 0), make([]int, 0)}
}
