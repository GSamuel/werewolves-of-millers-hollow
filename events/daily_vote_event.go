package events

import (
	"github.com/GSamuel/werewolves-of-millers-hollow/voting"
)

type DailyVoteEvent struct {
	*voting.BallotBox
}

func NewDailyVoteEvent() *DailyVoteEvent {
	return &DailyVoteEvent{voting.NewBallotBox()}
}
