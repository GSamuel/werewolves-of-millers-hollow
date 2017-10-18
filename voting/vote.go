package voting

import ()

type Vote struct {
	id     int
	weight float32
}

func NewVote(id int, weight float32) Vote {
	return Vote{id, weight}
}
