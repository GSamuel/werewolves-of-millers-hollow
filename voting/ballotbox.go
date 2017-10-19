package voting

import ()

type BallotBox struct {
	votes []Vote
}

func (b *BallotBox) Vote(id int, weight float32) {
	b.votes = append(b.votes, NewVote(id, weight))
}

func (b *BallotBox) Count() int {
	return len(b.votes)
}

//Result gives the id that has the most votes
//Not complete yet. undefined behaviour for tied votes.
func (b *BallotBox) Result() int {

	m := make(map[int]float32)

	for i := 0; i < len(b.votes); i++ {
		id := b.votes[i].id
		val := m[id]
		m[id] = b.votes[i].weight + val
	}

	first := true
	id := -1
	value := float32(-1)

	for key, val := range m {
		if first || val >= value {
			id = key
			value = val
		}
	}

	return id
}

func NewBallotBox() *BallotBox {
	return &BallotBox{}
}
