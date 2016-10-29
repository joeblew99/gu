package trees

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/faux/loop"
	"github.com/influx6/faux/mque"
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

// EventDebroadcast defines a struct which gets a unsubscription for the events.
type EventDebroadcast struct {
	TreeID  string
	EventID string
	Link    func(*js.Object)
	Handle  mque.End
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
	eid     string
	Target  string
	Type    string
	EventID string
	Handle  loop.Looper
	Tree    *Markup
}

// NewEvent returns a event object that allows registering events to eventlisteners
func NewEvent(evtype string, evtarget string) *Event {
	return &Event{
		Target: evtarget,
		Type:   evtarget,
	}
}

// ID returns the uique event id string for the event.
func (e *Event) ID() string {
	return fmt.Sprintf("%s#%s", e.Target, e.Type)
}

// Clone  returns a new Event object from this.
func (e *Event) Clone() *Event {
	return &Event{
		Target: e.Target,
		Type:   e.Type,
	}
}

// Apply adds the event into the elements events lists
func (e *Event) Apply(ex *Markup) {
	if !ex.allowEvents {
		return
	}

	e.Tree = ex

	if e.Target == "" {
		e.Target = ex.EventID()
	}

	e.EventID = fmt.Sprintf("%s-%s", e.Type, ex.UID())
	ex.AddEvent(*e)
}

//==============================================================================
