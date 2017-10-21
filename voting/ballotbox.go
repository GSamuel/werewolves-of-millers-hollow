package voting

import (
	"sort"
)

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

//Result gives the id that has the most votes
//Not complete yet. undefined behaviour for tied votes.
func (b *BallotBox) Result() []int {

	result := make([]int, 0)

	if len(b.votes) == 0 {
		return result
	}

	value := -1

	for _, val := range b.votes {
		if val > value {
			value = val
		}
	}

	for id, val := range b.votes {
		if val == value {
			result = append(result, id)
		}
	}

	sort.Ints(result)

	return result
}

func NewBallotBox() *BallotBox {
	return &BallotBox{make(map[int]int), 0}
}
