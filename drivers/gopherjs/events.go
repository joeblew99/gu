package gopherjs

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/faux/mque"
)

// WrapperEvent implements the EventObject interface.
type WrapperEvent struct {
	*js.Object
	isDum  bool
	handle mque.End
}

// NewWrapperEvent creates a new Event object useful to hold in place of a wrapper
// object.
func NewWrapperEvent(event *js.Object, ender mque.End) *WrapperEvent {
	return &WrapperEvent{
		Object: event,
		handle: ender,
	}
}

// NewDummy returns a WrapperEvent instance wrapping a dummy object.
func NewDummy(event string) *WrapperEvent {
	return &WrapperEvent{
		Object: js.Global.Get("Object").New(),
		isDum:  true,
	}
}

// Underlying returns the internal wrapper object exposes by this event.
func (d *WrapperEvent) Underlying() interface{} {
	return d.Object
}

// PreventDefault implements the PreventDefault of the event object interface.
func (d *WrapperEvent) PreventDefault() {
	if !d.isDum {
		d.Object.Call("preventDefault")
	}
}

// RemoveEvent removes the event from it's root parent.
func (d *WrapperEvent) RemoveEvent() {
	if d.handle != nil {
		d.handle.End()
	}
}

// StopPropagation implements the StopPropagation of the event object interface.
func (d *WrapperEvent) StopPropagation() {
	if !d.isDum {
		d.Object.Call("stopPropagation")
	}
}

// StopImmediatePropagation implements the StopPropagation of the event object interface.
func (d *WrapperEvent) StopImmediatePropagation() {
	if !d.isDum {
		d.Object.Call("stopImmediatePropagation")
	}
}
