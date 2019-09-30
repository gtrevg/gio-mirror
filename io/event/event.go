// SPDX-License-Identifier: Unlicense OR MIT

/*
Package event contains the types for event handling.

The Queue interface is the protocol for receiving external events.

For example:

	var queue event.Queue = ...

	for _, e := range queue.Events(h) {
		switch e.(type) {
			...
		}
	}

In general, handlers must be declared before events become
available. Other packages such as pointer and key provide
the means for declaring handlers for specific event types.

The following example declares a handler ready for key input:

	import gioui.org/io/key

	ops := new(ui.Ops)
	var h *Handler = ...
	key.InputOp{Key: h}.Add(ops)

*/
package event

// Queue maps an event handler key to the events
// available to the handler.
type Queue interface {
	// Events returns the available events for a
	// Key.
	Events(k Key) []Event
}

// Key is the stable identifier for an event handler.
// For a handler h, the key is typically &h.
type Key interface{}

// Event is the marker interface for events.
type Event interface {
	ImplementsEvent()
}
