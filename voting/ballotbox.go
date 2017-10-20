package voting

import ()

type BallotBox struct {
	votes map[int]int
	count int
}

func (b *BallotBox) Vote(id, weight int) {
	v := b.votes[id]
	b.votes[id] = weight + v
	b.count = b.count + 1
}

func (b *BallotBox) Count() int {
	return b.count
}

func (b BallotBox) Marjority() bool {
	return false
}

//Result gives the id that has the most votes
//Not complete yet. undefined behaviour for tied votes.
func (b *BallotBox) Result() int {

	first := true
	id := -1
	value := -1

	for key, val := range b.votes {
		if first || val >= value {
			id = key
			value = val
			first = false
		}
	}

	return id
}

func NewBallotBox() *BallotBox {
	return &BallotBox{make(map[int]int), 0}
}
