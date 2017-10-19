package roles

type Hunter struct {
	*BaseRole
}

func (h *Hunter) Name() string {
	return HUNTER
}
