package main

import (
	"github.com/GSamuel/werewolvesmillershollow/events"
	"github.com/GSamuel/werewolvesmillershollow/voting"
)

const (
	VILLAGER = "Villager"
	WEREWOLF = "Werewolf"
)

type Role interface {
	Name() string
	OnGameStarted(events.GameStartedEvent)
	OnNightStarted(events.NightStartedEvent)
	OnWerewolfVote(events.WerewolfVoteEvent)
	OnDailyVote(events.DailyVoteEvent)
}

type BaseRole struct {
}

func (b *BaseRole) Name() string {
	return "base"
}

func (b *BaseRole) OnNightStarted(e events.NightStartedEvent) {
}

func (b *BaseRole) OnGameStarted(e events.GameStartedEvent) {
}

func (b *BaseRole) OnWerewolfVote(e events.WerewolfVoteEvent) {
}

func (b *BaseRole) OnDailyVote(e events.DailyVoteEvent) {
	i := ReadInput()
	e.Vote(voting.NewVote(i, 1))
}

type Werewolf struct {
	BaseRole
}

func (w *Werewolf) Name() string {
	return WEREWOLF
}

func (w *Werewolf) OnWerewolfVote(e events.WerewolfVoteEvent) {
	i := ReadInput()
	e.Vote(voting.NewVote(i, 1))
}

type Villager struct {
	BaseRole
}

func (v *Villager) Name() string {
	return VILLAGER
}
