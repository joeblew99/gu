package trees

import (
	"fmt"

	"github.com/gu-io/gu/notifications/mque"
)

// EventBroadcast defines a struct which gets published for the events.
type EventBroadcast struct {
	EventID string
	Event   EventObject
}

// EventObject defines a interface for the basic methods which events needs
// expose.
type EventObject interface {
	RemoveEvent()
	PreventDefault()
	StopPropagation()
	Underlying() interface{}
	StopImmediatePropagation()
}

// Event provide a meta registry for helps in registering events for dom markups
// which is translated to the nodes themselves
type Event struct {
	Type      string
	EventID   string
	secTarget string
	Tree      *Markup
	Handle    mque.End
}

// NewEvent returns a event object that allows registering events to eventlisteners
func NewEvent(evtype string, evtarget string) *Event {
	return &Event{
		Type:      evtype,
		secTarget: evtarget,
	}
}

// Target returns the target of the giving event.
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
