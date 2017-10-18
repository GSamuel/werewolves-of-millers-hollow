package main

import (
	"math/rand"
)

const (
	HUMAN    = "Human"
	WEREWOLF = "Werewolf"
)

type Deck struct {
	cards    []string
	shuffled []int
}

func (d *Deck) add(card string) {
	d.cards = append(d.cards, card)
}

func (d *Deck) shuffle() {
	d.shuffled = rand.Perm(len(d.cards))
}

func (d *Deck) count() int {
	return len(d.cards)
}

func (d *Deck) get(i int) string {
	return d.cards[d.shuffled[i]]
}

func NewDeck() *Deck {
	return &Deck{make([]string, 0), make([]int, 0)}
}
