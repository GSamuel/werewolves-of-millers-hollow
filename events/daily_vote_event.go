package events

import (
	"github.com/GSamuel/werewolvesmillershollow/voting"
)

type DailyVoteEvent struct {
	*voting.BallotBox
}

func NewDailyVoteEvent() *DailyVoteEvent {
	return &DailyVoteEvent{voting.NewBallotBox()}
}
