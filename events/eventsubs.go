package events

import (
	"fmt"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	gujs "github.com/influx6/gu/js"
)

// EventHandler provides the function type for event callbacks when subscribing
// to events in Haiku.
type EventHandler func(Event)

// Events defines a package level variable to access the event registeration.
var Events eventCtrl

type eventCtrl struct{}

// DOM sets up the event subs for listening
func (eventCtrl) DOM(dom *js.Object) {
	e.Offload()
	e.dom = dom
	e.jslink = func(o *js.Object) { e.TriggerMatch(&EventObject{o}) }
	e.dom.Call("addEventListener", e.Type(), e.jslink, true)
}

// Offload removes all event bindings from current dom element
func (eventCtrl) Offload(dom *js.Object) {
	if dom == nil || dom == js.Undefined {
		return
	}

	e.dom.Call("removeEventListener", e.Type(), e.jslink, true)
	e.jslink = nil
}

// Trigger provides bypass for triggering this event sub by passing down an event
// directly without matching target or selector
func (e *EventSub) Trigger(h Event) {
	e.Run(h)
}

// TriggerMatch check if the current event from a specific parent matches the
// eventarget by using the eventsub selector,if the target is within the results for
// that selector then it triggers the event subscribers
func (e *EventSub) TriggerMatch(h Event) {
	// if e.dom != nil
	if strings.ToLower(h.Type()) != strings.ToLower(e.Type()) {
		return
	}

	//get the current event target
	target := h.Target()

	//get the targets parent
	// parent := target.Get("parentElement")
	parent := e.dom

	var match bool

	children := parent.Call("querySelectorAll", e.Target())

	if children == nil || children == js.Undefined {
		return
	}

	// log.Printf("children -> %s  -> %t %t", children, children == nil, children == js.Undefined)

	//get all possible matches of this query
	// posis := parent.QuerySelectorAll(e.EventSelector())
	posis := gujs.DOMObjectToList(children)

	// log.Printf("Checking: %s for %s -> %+s", target, e.ID(), posis)

	//is our target part of those that match the selector
	for _, item := range posis {
		// log.Printf("taget %+s and item %+s -> %t", target, item, target == item)
		if item != target {
			continue
		}
		match = true
		break
	}

	//if we match then run the listeners registered
	if match {
		//if we dont want immediatepropagation kill it else check propagation also
		if e.StopImmediatePropagation() {
			h.StopImmediatePropagation()
		} else {
			//if we dont want propagation kill it
			if e.StopPropagation() {
				h.StopPropagation()
			}
		}

		//we want to PreventDefault then stop default action
		if e.PreventDefault() {
			h.PreventDefault()
		}

		e.Run(h)
	}
}

// ID returns the event id that EventManager use for this event
func (e *EventSub) ID() string {
	return GetEventID(e)
}

// GetEventID returns the id for a ElemEvent object
func GetEventID(m EventSubs) string {
	sel := strings.TrimSpace(m.Target())
	return BuildEventID(sel, m.Type())
}

// BuildEventID returns the string represent of the values using the select#event format
func BuildEventID(etype, eselect string) string {
	return fmt.Sprintf("%s#%s", eselect, etype)
}
