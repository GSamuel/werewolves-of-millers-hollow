package game

import ()

type PlayerGroup struct {
	players []*Player
}

func (p *PlayerGroup) Count() int {
	return len(p.players)
}

func (p *PlayerGroup) Role(id int) Role {
	return p.players[id]
}

func (p *PlayerGroup) Player(id int) *Player {
	return p.players[id]
}
