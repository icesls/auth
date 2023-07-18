package bll

import (
	"auth/event"
)

type eventHandle struct {
}

func (h *eventHandle) init() func() {
	// resource event handler
	event.Bus.Register(event.ExampleEvent, h.ExampleEventHandler)

	return func() {}
}

func (h *eventHandle) ExampleEventHandler(e event.Event, data interface{}) {
	// handle event
}
