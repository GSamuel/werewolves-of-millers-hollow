package events

import ()

type NightStartedEvent struct {
}

func NewNightStartedEvent() NightStartedEvent {
	return NightStartedEvent{}
}
