package input

import ()

type Input struct {
	group   Group
	filters []Filter
}

func (i *Input) Read() int {
	done := false
	var n int
	for !done {
		n, err := ReadInput()
		if err == nil {
			continue
		}

		if n >= 0 && n < i.group.Count() {
			check := true
			for j := 0; j < len(i.filters); j++ {
				if !i.filters[j].Validate(i.group.Role(j)) {
					check = false
				}
			}
			if check {
				done = true
			}
		}
	}
	return n
}

func (i *Input) AddFilter(filter Filter) {
	i.filters = append(i.filters, filter)
}

func New(group Group) *Input {
	return &Input{group, make([]Filter, 0)}
}
