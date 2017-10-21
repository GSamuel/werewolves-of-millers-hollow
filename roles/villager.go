package roles

type Villager struct {
	*BaseRole
}

func (v *Villager) Name() string {
	return VILLAGER
}
