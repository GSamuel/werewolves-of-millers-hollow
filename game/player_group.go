package game

import ()

type PlayerGroup struct {
	players []*Player
}

func (p *PlayerGroup) Count() int {
	return len(p.players)
}

func (p *PlayerGroup) Player(id int) *Player {
	return p.players[id]
}
