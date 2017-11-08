package roles

import (
	"fmt"
	"github.com/GSamuel/werewolves-of-millers-hollow/events"
	"github.com/GSamuel/werewolves-of-millers-hollow/game"
)

type BaseRole struct {
	*state
}

func (b *BaseRole) Name() string {
	return UNDEFINED
}

func (b *BaseRole) Die(g *game.Game, werewolves bool) {
	deathEvent := events.NewPlayerDeadEvent(b.ID())
	deathEvent.SetWerewolfAttack(werewolves)
	g.PlayerDeadEvent(g, deathEvent)

	if deathEvent.DeathPrevented() {
		return
	}

	b.SetAlive(false)
	revealEvent := events.NewPlayerRevealedEvent(b.ID())
	g.PlayerRevealedEvent(g, revealEvent)

	fmt.Printf("Player %v died.\n", b.ID())
}

func (b *BaseRole) OnNightStarted(g *game.Game, e *events.NightStartedEvent) {
}

func (b *BaseRole) OnGameStarted(g *game.Game, e *events.GameStartedEvent) {
}

func (b *BaseRole) OnWerewolfVote(g *game.Game, e *events.WerewolfVoteEvent) {
}

func (b *BaseRole) OnPlayerDead(g *game.Game, e *events.PlayerDeadEvent) {
}

func (b *BaseRole) OnPlayerRevealed(g *game.Game, e *events.PlayerRevealedEvent) {
}

func (b *BaseRole) OnDailyVote(g *game.Game, e *events.DailyVoteEvent) {
	if !b.Alive() {
		return
	}

	p := g.Player(b.ID())

	p.Write(fmt.Sprintf("Player %v:", b.ID()))

	vote := p.ReadInt(func(i int) bool {
		return Alive(g, i) && NotEqual(g, i, p.ID())
	})

	e.Vote(vote, 1)
}
