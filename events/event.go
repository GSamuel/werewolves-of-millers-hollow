package events

import (
	"github.com/GSamuel/werewolvesmillershollow/voting"
)

type NightStartedEvent struct {
}

type GameStartedEvent struct {
}

type WerewolfVoteEvent struct {
	*voting.BallotBox
}

type DailyVoteEvent struct {
	*voting.BallotBox
}
