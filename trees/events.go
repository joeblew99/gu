package trees

import "github.com/influx6/faux/loop"

// Event provide a meta registry for helps in registering events for dom markups
// which is translated to the nodes themselves
type Event struct {
	Target string
	Type   string
	Handle loop.Looper
}

// NewEvent returns a event object that allows registering events to eventlisteners
func NewEvent(evtype, evtarget string) *Event {
	return &Event{
		Target: evtarget,
		Type:   evtarget,
	}
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

	if e.Target == "" {
		e.Target = ex.EventID()
	}

	ex.AddEvent(*e)
}
