package events

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// JSEventMux represents a js.listener function which is returned when attached
// using AddEventListeners and is used for removals with RemoveEventListeners
type JSEventMux func(*js.Object)

// Event defines the base interface for browser events and defines the basic interface methods they must provide
type Event interface {
	Bubbles() bool
	Cancelable() bool
	CurrentTarget() *js.Object
	DefaultPrevented() bool
	EventPhase() int
	Target() *js.Object
	Timestamp() int
	Type() string
	Core() *js.Object
	StopPropagation()
	StopImmediatePropagation()
	PreventDefault()
}

// EventObject implements the Event interface and is embedded by
// concrete event types.
type EventObject struct {
	Bubbles       bool
	EventPhase    int
	Type          string
	Target        interface{}
	CurrentTarget interface{}
	core          *js.Object
}

// Timestamp returns the event timestamp value as a int (in seconds)
func (ev *EventObject) Timestamp() int {
	var ms int

	if ev.core != nil {
		ms := ev.core.Get("timeStamp").Int()
	} else {
		ms = int(time.Now().Unix())
	}

	s := ms / 1000
	return s
}

// PreventDefault prevents the default value of the event
func (ev *EventObject) PreventDefault() {
	if ev.core != nil {
		ev.core.Call("preventDefault")
	}
}

// StopImmediatePropagation stops the propagation of the event forcefully
func (ev *EventObject) StopImmediatePropagation() {
	if ev.core != nil {
		ev.core.Call("stopImmediatePropagation")
	}
}

// StopPropagation stops the propagation of the event
func (ev *EventObject) StopPropagation() {
	if ev.core != nil {
		ev.core.Call("stopPropagation")
	}
}
