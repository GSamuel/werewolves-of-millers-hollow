package events

import (
	"github.com/GSamuel/werewolvesmillershollow/voting"
)

type WerewolfVoteEvent struct {
	*voting.BallotBox
}

func NewWerewolfVoteEvent() WerewolfVoteEvent {
	return WerewolfVoteEvent{voting.NewBallotBox()}
}
