package trees

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/dispatch/mque"
)

// EventObject defines a interface for the basic methods which events needs
// expose.
type EventObject interface {
	PreventDefault()
	StopPropagation()
	StopImmediatePropagation()
	Underlying() *js.Object
}

// EventBroadcast defines a struct which gets published for the events.
type EventBroadcast struct {
	EventID string
	Event   EventObject
}

//==============================================================================

// WrapperEvent implements the EventObject interface.
type WrapperEvent struct {
	dummy *js.Object
	isDum bool
}

// NewEvent creates a new Event object useful to hold in place of a wrapper
// object.
func NewWrapperEvent(dm *js.Object) *WrapperEvent {
	return &WrapperEvent{
		dummy: dm,
	}
}

// NewDummyEvent returns a WrapperEvent instance wrapping a dummy object.
func NewDummy() *WrapperEvent {
	return &WrapperEvent{
		dummy: js.Global.Get("Object").New(),
		isDum: true,
	}
}

// Underlying returns the internal wrapper object exposes by this event.
func (d *WrapperEvent) Underlying() *js.Object {
	return d.dummy
}

// PreventDefault implements the PreventDefault of the event object interface.
func (d *WrapperEvent) PreventDefault() {
	if !d.isDum {
		d.dummy.Call("preventDefault")
	}
}

// StopPropagation implements the StopPropagation of the event object interface.
func (d *WrapperEvent) StopPropagation() {
	if !d.isDum {
		d.dummy.Call("stopPropagation")
	}
}

// StopImmediatePropagation implements the StopPropagation of the event object interface.
func (d *WrapperEvent) StopImmediatePropagation() {
	if !d.isDum {
		d.dummy.Call("stopImmediatePropagation")
	}
}

//==============================================================================

// Event provide a meta registry for helps in registering events for dom markups
// which is translated to the nodes themselves
type Event struct {
	Type      string
	EventID   string
	secTarget string
	Tree      *Markup
	Handle    mque.End
	Link      func(*js.Object)
}

// NewEvent returns a event object that allows registering events to eventlisteners
func NewEvent(evtype string, evtarget string) *Event {
	return &Event{
		Type:      evtype,
		secTarget: evtarget,
	}
}

func (e *Event) Target() string {
	if e.Tree != nil {
		return e.Tree.EventID()
	}

	return e.secTarget
}

// ID returns the uique event id string for the event.
func (e *Event) ID() string {
	return fmt.Sprintf("%s#%s", e.Target(), e.Type)
}

// Clone  returns a new Event object from this.
func (e *Event) Clone() *Event {
	return &Event{
		Type:      e.Type,
		secTarget: e.secTarget,
	}
}

// Apply adds the event into the elements events lists
func (e *Event) Apply(ex *Markup) {
	if !ex.allowEvents {
		return
	}

	e.Tree = ex

	e.EventID = fmt.Sprintf("%s-%s", e.Type, ex.UID())
	ex.AddEvent(*e)
}

//==============================================================================
