// Package gopherjs contains the DOM RenderTarget for rendering gu markup into a
// browser DOM.
package gopherjs

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/notifications"
	"github.com/gu-io/gu/trees"
)

// BindEvent connects the event with the provided event object and root.
func BindEvent(source *trees.Event, root *js.Object) {
	link := func(o *js.Object) { TriggerBindEvent(o, root, source) }

	root.Call("addEventListener", source.Type, link, source.UseCapture)

	source.Handle.AddEnd(func() {
		root.Call("removeEventListener", source.Type, link, source.UseCapture)
	})
}

// TriggerBindEvent connects the giving event with the provided dom target.
func TriggerBindEvent(event *js.Object, root *js.Object, source *trees.Event) {
	target := event.Get("target")

	children := root.Call("querySelectorAll", source.Target())
	if children == nil || children == js.Undefined {
		return
	}

	kids := DOMObjectToList(children)
	var match bool

	for _, item := range kids {
		if item != target {
			continue
		}

		match = true
		break
	}

	// if we match then run the listeners registered.
	if match {
		if source.PreventDefault {
			event.Call("preventDefault", nil)
		}

		if source.StopImmediatePropagation {
			event.Call("stopImmediatePropagation", nil)
		}

		if source.StopPropagation {
			event.Call("stopPropagation", nil)
		}

		notifications.Dispatch(trees.EventBroadcast{
			EventName: source.EventName(),
			EventID:   source.EventID,
			Event:     GetEvent(event, source.Handle),
		})
	}
}
