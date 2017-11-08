package events

import (
	"github.com/GSamuel/werewolves-of-millers-hollow/voting"
)

type WerewolfVoteEvent struct {
	*voting.BallotBox
}

func NewWerewolfVoteEvent() *WerewolfVoteEvent {
	return &WerewolfVoteEvent{voting.NewBallotBox()}
}
