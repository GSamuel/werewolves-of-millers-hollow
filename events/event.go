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

type PlayerDeadEvent struct {
	id int
}

type PlayerRevealedEvent struct {
	id int
}

func NewNightStartedEvent() NightStartedEvent {
	return NightStartedEvent{}
}

func NewGameStartedEvent() GameStartedEvent {
	return GameStartedEvent{}
}

func NewWerewolfVoteEvent() WerewolfVoteEvent {
	return WerewolfVoteEvent{voting.NewBallotBox()}
}

func NewDailyVoteEvent() DailyVoteEvent {
	return DailyVoteEvent{voting.NewBallotBox()}
}

func NewPlayerDeadEvent(id int) PlayerDeadEvent {
	return PlayerDeadEvent{id}
}

func NewPlayerRevealedEvent(id int) PlayerRevealedEvent {
	return PlayerRevealedEvent{id}
}
