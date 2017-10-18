package voting

import ()

type BallotBox struct {
	votes []Vote
}

func (b *BallotBox) Vote(vote Vote) {
	b.votes = append(b.votes, vote)
}

func (b *BallotBox) Count() int {
	return len(b.votes)
}

func (b *BallotBox) Result() int {
	return b.votes[0].id
}

func NewBallotBox() *BallotBox {
	return &BallotBox{}
}
