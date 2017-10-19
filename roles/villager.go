package roles

//Basic Villager role
type Villager struct {
	*BaseRole
}

func (v *Villager) Name() string {
	return VILLAGER
}
