// The generation of this package was inspired by Neelance work on DOM (https://github.com/neelance/dom)

//go:generate go run generate.go

// Documentation source: "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.

//Package events defines the event binding system that combines different libraries to create a interesting event system.
package events

import (
	"github.com/gu-io/gu/notifications"
	"github.com/gu-io/gu/trees"
)

// EventHandler defines a function type for event callbacks.
type EventHandler func(trees.EventObject, *trees.Markup)

// WrapHandler wraps the function returning a EventHandler to call the provided
// function to be called when the event occurs without need for the arguments.
func WrapHandler(callback func()) EventHandler {
	return func(ev trees.EventObject, root *trees.Markup) {
		callback()
	}
}

// WrapEventOnlyHandler wraps the function returning a EventHandler to call the provided
// function to be called when the event occurs without need for the arguments.
func WrapEventOnlyHandler(callback func(trees.EventObject)) EventHandler {
	return func(ev trees.EventObject, root *trees.Markup) {
		callback(ev)
	}
}

// AbortEvent Documentation is as below: "A transaction has been aborted."
// https://developer.mozilla.org/docs/Web/Reference/Events/abort_indexedDB
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AbortEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("abort", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AfterPrintEvent Documentation is as below: "The associated document has started printing or the print preview has been closed."
// https://developer.mozilla.org/docs/Web/Events/afterprint
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AfterPrintEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("afterprint", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AfterScriptExecuteEvent Documentation is as below: "A script has been executed."
// https://developer.mozilla.org/docs/Web/Events/afterscriptexecute
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AfterScriptExecuteEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("afterscriptexecute", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AlertActiveEvent Documentation is as below: "A notification element is shown."
// https://developer.mozilla.org/docs/Web/Reference/Events/AlertActive
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AlertActiveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("AlertActive", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AlertCloseEvent Documentation is as below: "A notification element is closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/AlertClose
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AlertCloseEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("AlertClose", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AlertingEvent Documentation is as below: "The correspondent is being alerted (his/her phone is ringing)."
// https://developer.mozilla.org/docs/Web/Events/alerting
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AlertingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("alerting", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AnimationEndEvent Documentation is as below: "A CSS animation has completed."
// https://developer.mozilla.org/docs/Web/Events/animationend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AnimationEndEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("animationend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AnimationIterationEvent Documentation is as below: "A CSS animation is repeated."
// https://developer.mozilla.org/docs/Web/Events/animationiteration
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AnimationIterationEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("animationiteration", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AnimationStartEvent Documentation is as below: "A CSS animation has started."
// https://developer.mozilla.org/docs/Web/Events/animationstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AnimationStartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("animationstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AudioProcessEvent Documentation is as below: "The input buffer of a ScriptProcessorNode is ready to be processed."
// https://developer.mozilla.org/docs/Web/Events/audioprocess
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AudioProcessEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("audioprocess", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AudioendEvent Documentation is as below: "The user agent has finished capturing audio for speech recognition."
// https://developer.mozilla.org/docs/Web/Events/audioend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AudioendEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("audioend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// AudiostartEvent Documentation is as below: "The user agent has started to capture audio for speech recognition."
// https://developer.mozilla.org/docs/Web/Events/audiostart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func AudiostartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("audiostart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BeforeInstallPromptEvent Documentation is as below: "A user is prompted to save a web site to a home screen on mobile."
// https://developer.mozilla.org/docs/Web/Events/beforeinstallprompt
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BeforeInstallPromptEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("beforeinstallprompt", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BeforePrintEvent Documentation is as below: "The associated document is about to be printed or previewed for printing."
// https://developer.mozilla.org/docs/Web/Events/beforeprint
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BeforePrintEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("beforeprint", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BeforeScriptExecuteEvent Documentation is as below: "A script is about to be executed."
// https://developer.mozilla.org/docs/Web/Events/beforescriptexecute
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BeforeScriptExecuteEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("beforescriptexecute", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BeforeUnloadEvent Documentation is as below: "The window, the document and its resources are about to be unloaded."
// https://developer.mozilla.org/docs/Web/Events/beforeunload
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BeforeUnloadEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("beforeunload", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BeginEventEvent Documentation is as below: "A SMIL animation element begins."
// https://developer.mozilla.org/docs/Web/Events/beginEvent
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BeginEventEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("beginEvent", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BlockedEvent Documentation is as below: "An open connection to a database is blocking a versionchange transaction on the same database."
// https://developer.mozilla.org/docs/Web/Reference/Events/blocked_indexedDB
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BlockedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("blocked", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BlurEvent Documentation is as below: "An element has lost focus (does not bubble)."
// https://developer.mozilla.org/docs/Web/Events/blur
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BlurEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("blur", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BoundaryEvent Documentation is as below: "The spoken utterance reaches a word or sentence boundary"
// https://developer.mozilla.org/docs/Web/Events/boundary
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BoundaryEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("boundary", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BroadcastEvent Documentation is as below: "An observer noticed a change to the attributes of a watched broadcaster."
// https://developer.mozilla.org/docs/Web/Events/broadcast
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BroadcastEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("broadcast", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// BusyEvent Documentation is as below: "The line of the correspondent is busy."
// https://developer.mozilla.org/docs/Web/Events/busy
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func BusyEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("busy", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CachedEvent Documentation is as below: "The resources listed in the manifest have been downloaded, and the application is now cached."
// https://developer.mozilla.org/docs/Web/Events/cached
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CachedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("cached", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CallschangedEvent Documentation is as below: "A call has been added or removed from the list of current calls."
// https://developer.mozilla.org/docs/Web/Events/callschanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CallschangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("callschanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CanPlayEvent Documentation is as below: "The user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content."
// https://developer.mozilla.org/docs/Web/Events/canplay
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CanPlayEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("canplay", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CanPlayThroughEvent Documentation is as below: "The user agent can play the media up to its end without having to stop for further buffering of content."
// https://developer.mozilla.org/docs/Web/Events/canplaythrough
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CanPlayThroughEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("canplaythrough", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CardstatechangeEvent Documentation is as below: "The MozMobileConnection.cardState property changes value."
// https://developer.mozilla.org/docs/Web/Events/cardstatechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CardstatechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("cardstatechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CfstatechangeEvent Documentation is as below: "The call forwarding state changes."
// https://developer.mozilla.org/docs/Web/Events/cfstatechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CfstatechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("cfstatechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ChangeEvent Documentation is as below: "This event is triggered each time a file is created, modified or deleted on a given storage area."
// https://developer.mozilla.org/docs/Web/Events/change
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("change", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ChargingChangeEvent Documentation is as below: "The battery begins or stops charging."
// https://developer.mozilla.org/docs/Web/Events/chargingchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ChargingChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("chargingchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ChargingTimeChangeEvent Documentation is as below: "The chargingTime attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/chargingtimechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ChargingTimeChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("chargingtimechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CheckboxStateChangeEvent Documentation is as below: "The state of a checkbox has been changed either by a user action or by a script (useful for accessibility)."
// https://developer.mozilla.org/docs/Web/Events/CheckboxStateChange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CheckboxStateChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("CheckboxStateChange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CheckingEvent Documentation is as below: "The user agent is checking for an update, or attempting to download the cache manifest for the first time."
// https://developer.mozilla.org/docs/Web/Events/checking
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CheckingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("checking", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ClickEvent Documentation is as below: "A pointing device button has been pressed and released on an element."
// https://developer.mozilla.org/docs/Web/Events/click
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ClickEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("click", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CloseEvent Documentation is as below: "The close button of the window has been clicked."
// https://developer.mozilla.org/docs/Web/Reference/Events/close_event
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CloseEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("close", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CommandEvent Documentation is as below: "An element has been activated."
// https://developer.mozilla.org/docs/Web/Events/command
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CommandEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("command", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CommandupdateEvent Documentation is as below: "A command update occurred on a commandset element."
// https://developer.mozilla.org/docs/Web/Events/commandupdate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CommandupdateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("commandupdate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CompleteEvent Documentation is as below: "The rendering of an OfflineAudioContext is terminated."
// https://developer.mozilla.org/docs/Web/Events/complete
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CompleteEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("complete", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CompositionEndEvent Documentation is as below: "The composition of a passage of text has been completed or canceled."
// https://developer.mozilla.org/docs/Web/Events/compositionend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CompositionEndEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("compositionend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CompositionStartEvent Documentation is as below: "The composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition)."
// https://developer.mozilla.org/docs/Web/Events/compositionstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CompositionStartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("compositionstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CompositionUpdateEvent Documentation is as below: "A character is added to a passage of text being composed."
// https://developer.mozilla.org/docs/Web/Events/compositionupdate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CompositionUpdateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("compositionupdate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ConnectingEvent Documentation is as below: "A call is about to connect."
// https://developer.mozilla.org/docs/Web/Events/connecting
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ConnectingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("connecting", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ConnectionInfoUpdateEvent Documentation is as below: "The informations about the signal strength and the link speed have been updated."
// https://developer.mozilla.org/docs/Web/Events/connectionInfoUpdate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ConnectionInfoUpdateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("connectionInfoUpdate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ContextMenuEvent Documentation is as below: "The right button of the mouse is clicked (before the context menu is displayed)."
// https://developer.mozilla.org/docs/Web/Events/contextmenu
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ContextMenuEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("contextmenu", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CopyEvent Documentation is as below: "The text selection has been added to the clipboard."
// https://developer.mozilla.org/docs/Web/Events/copy
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CopyEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("copy", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CSSRuleViewCSSLinkClickedEvent Documentation is as below: "A link to a CSS file has been clicked in the \"Rules\" view of the style inspector."
// https://developer.mozilla.org/docs/Web/Reference/Events/CssRuleViewCSSLinkClicked
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CSSRuleViewCSSLinkClickedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("CssRuleViewCSSLinkClicked", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CSSRuleViewChangedEvent Documentation is as below: "The \"Rules\" view of the style inspector has been changed."
// https://developer.mozilla.org/docs/Web/Reference/Events/CssRuleViewChanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CSSRuleViewChangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("CssRuleViewChanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CSSRuleViewRefreshedEvent Documentation is as below: "The \"Rules\" view of the style inspector has been updated."
// https://developer.mozilla.org/docs/Web/Reference/Events/CssRuleViewRefreshed
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CSSRuleViewRefreshedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("CssRuleViewRefreshed", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// CutEvent Documentation is as below: "The text selection has been removed from the document and added to the clipboard."
// https://developer.mozilla.org/docs/Web/Events/cut
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func CutEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("cut", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMAutoCompleteEvent Documentation is as below: "The content of an element has been auto-completed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMAutoComplete
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMAutoCompleteEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMAutoComplete", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMContentLoadedEvent Documentation is as below: "The document has finished loading (but not its dependent resources)."
// https://developer.mozilla.org/docs/Web/Events/DOMContentLoaded
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMContentLoadedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMContentLoaded", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMFrameContentLoadedEvent Documentation is as below: "The frame has finished loading (but not its dependent resources)."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMFrameContentLoaded
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMFrameContentLoadedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMFrameContentLoaded", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMLinkAddedEvent Documentation is as below: "A link has been added a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMLinkAdded
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMLinkAddedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMLinkAdded", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMLinkRemovedEvent Documentation is as below: "A link has been removed inside from a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMLinkRemoved
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMLinkRemovedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMLinkRemoved", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMMenuItemActiveEvent Documentation is as below: "A menu or menuitem has been hovered or highlighted."
// https://developer.mozilla.org/docs/Web/Events/DOMMenuItemActive
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMMenuItemActiveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMMenuItemActive", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMMenuItemInactiveEvent Documentation is as below: "A menu or menuitem is no longer hovered or highlighted."
// https://developer.mozilla.org/docs/Web/Events/DOMMenuItemInactive
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMMenuItemInactiveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMMenuItemInactive", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMMetaAddedEvent Documentation is as below: "A meta element has been added to a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMMetaAdded
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMMetaAddedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMMetaAdded", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMMetaRemovedEvent Documentation is as below: "A meta element has been removed from a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMMetaRemoved
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMMetaRemovedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMMetaRemoved", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMModalDialogClosedEvent Documentation is as below: "A modal dialog has been closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMModalDialogClosed
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMModalDialogClosedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMModalDialogClosed", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMPopupBlockedEvent Documentation is as below: "A popup has been blocked"
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMPopupBlocked
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMPopupBlockedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMPopupBlocked", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMTitleChangedEvent Documentation is as below: "The title of a window has changed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMTitleChanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMTitleChangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMTitleChanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMWillOpenModalDialogEvent Documentation is as below: "A modal dialog is about to open."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMWillOpenModalDialog
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMWillOpenModalDialogEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMWillOpenModalDialog", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMWindowCloseEvent Documentation is as below: "A window is about to be closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMWindowClose
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMWindowCloseEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMWindowClose", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DOMWindowCreatedEvent Documentation is as below: "A window has been created."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMWindowCreated
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DOMWindowCreatedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("DOMWindowCreated", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DatachangeEvent Documentation is as below: "The MozMobileConnection.data object changes values."
// https://developer.mozilla.org/docs/Web/Events/datachange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DatachangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("datachange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DataerrorEvent Documentation is as below: "The MozMobileConnection.data object receive an error from the RIL."
// https://developer.mozilla.org/docs/Web/Events/dataerror
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DataerrorEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dataerror", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DblClickEvent Documentation is as below: "A pointing device button is clicked twice on an element."
// https://developer.mozilla.org/docs/Web/Events/dblclick
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DblClickEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dblclick", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DeliveredEvent Documentation is as below: "An SMS has been successfully delivered."
// https://developer.mozilla.org/docs/Web/Events/delivered
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DeliveredEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("delivered", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DeviceLightEvent Documentation is as below: "Fresh data is available from a light sensor."
// https://developer.mozilla.org/docs/Web/Events/devicelight
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DeviceLightEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("devicelight", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DeviceMotionEvent Documentation is as below: "Fresh data is available from a motion sensor."
// https://developer.mozilla.org/docs/Web/Events/devicemotion
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DeviceMotionEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("devicemotion", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DeviceOrientationEvent Documentation is as below: "Fresh data is available from an orientation sensor."
// https://developer.mozilla.org/docs/Web/Events/deviceorientation
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DeviceOrientationEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("deviceorientation", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DeviceProximityEvent Documentation is as below: "Fresh data is available from a proximity sensor (indicates an approximated distance between the device and a nearby object)."
// https://developer.mozilla.org/docs/Web/Events/deviceproximity
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DeviceProximityEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("deviceproximity", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DevicechangeEvent Documentation is as below: "A media device such as a camera, microphone, or speaker is connected or removed from the system."
// https://developer.mozilla.org/docs/Web/Events/devicechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DevicechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("devicechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DialingEvent Documentation is as below: "The number of a correspondent has been dialed."
// https://developer.mozilla.org/docs/Web/Events/dialing
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DialingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dialing", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DisabledEvent Documentation is as below: "Wifi has been disabled on the device."
// https://developer.mozilla.org/docs/Web/Events/disabled
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DisabledEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("disabled", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DischargingTimeChangeEvent Documentation is as below: "The dischargingTime attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/dischargingtimechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DischargingTimeChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dischargingtimechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DisconnectedEvent Documentation is as below: "A call has been disconnected."
// https://developer.mozilla.org/docs/Web/Events/disconnected
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DisconnectedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("disconnected", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DisconnectingEvent Documentation is as below: "A call is about to disconnect."
// https://developer.mozilla.org/docs/Web/Events/disconnecting
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DisconnectingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("disconnecting", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DownloadingEvent Documentation is as below: "The user agent has found an update and is fetching it, or is downloading the resources listed by the cache manifest for the first time."
// https://developer.mozilla.org/docs/Web/Events/downloading
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DownloadingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("downloading", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DragEvent Documentation is as below: "An element or text selection is being dragged (every 350ms)."
// https://developer.mozilla.org/docs/Web/Events/drag
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DragEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("drag", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DragEndEvent Documentation is as below: "A drag operation is being ended (by releasing a mouse button or hitting the escape key)."
// https://developer.mozilla.org/docs/Web/Events/dragend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DragEndEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dragend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DragEnterEvent Documentation is as below: "A dragged element or text selection enters a valid drop target."
// https://developer.mozilla.org/docs/Web/Events/dragenter
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DragEnterEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dragenter", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DragLeaveEvent Documentation is as below: "A dragged element or text selection leaves a valid drop target."
// https://developer.mozilla.org/docs/Web/Events/dragleave
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DragLeaveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dragleave", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DragOverEvent Documentation is as below: "An element or text selection is being dragged over a valid drop target (every 350ms)."
// https://developer.mozilla.org/docs/Web/Events/dragover
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DragOverEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dragover", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DragStartEvent Documentation is as below: "The user starts dragging an element or text selection."
// https://developer.mozilla.org/docs/Web/Events/dragstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DragStartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("dragstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DropEvent Documentation is as below: "An element is dropped on a valid drop target."
// https://developer.mozilla.org/docs/Web/Events/drop
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DropEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("drop", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// DurationChangeEvent Documentation is as below: "The duration attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/durationchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func DurationChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("durationchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// EmptiedEvent Documentation is as below: "The media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load()\u00a0method is called to reload it."
// https://developer.mozilla.org/docs/Web/Events/emptied
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func EmptiedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("emptied", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// EnabledEvent Documentation is as below: "Wifi has been enabled on the device."
// https://developer.mozilla.org/docs/Web/Events/enabled
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func EnabledEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("enabled", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// EndEvent Documentation is as below: "The utterance has finished being spoken."
// https://developer.mozilla.org/docs/Web/Events/end_(SpeechSynthesis)
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func EndEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("end", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// EndEventEvent Documentation is as below: "A SMIL animation element ends."
// https://developer.mozilla.org/docs/Web/Events/endEvent
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func EndEventEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("endEvent", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// EndedEvent Documentation is as below: "Playback has stopped because the end of the media was reached."
// https://developer.mozilla.org/docs/Web/Events/ended_(Web_Audio)
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func EndedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("ended", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// FocusEvent Documentation is as below: "An element has received focus (does not bubble)."
// https://developer.mozilla.org/docs/Web/Events/focus
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func FocusEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("focus", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// FocusInEvent Documentation is as below: "An element is about to receive focus (bubbles)."
// https://developer.mozilla.org/docs/Web/Events/focusin
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func FocusInEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("focusin", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// FocusOutEvent Documentation is as below: "An element is about to lose focus (bubbles)."
// https://developer.mozilla.org/docs/Web/Events/focusout
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func FocusOutEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("focusout", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// FullScreenChangeEvent Documentation is as below: "An element was turned to fullscreen mode or back to normal mode."
// https://developer.mozilla.org/docs/Web/Events/fullscreenchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func FullScreenChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("fullscreenchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// FullScreenErrorEvent Documentation is as below: "It was impossible to switch to fullscreen mode for technical reasons or because the permission was denied."
// https://developer.mozilla.org/docs/Web/Events/fullscreenerror
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func FullScreenErrorEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("fullscreenerror", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// FullscreenEvent Documentation is as below: "Browser fullscreen mode has been entered or left."
// https://developer.mozilla.org/docs/Web/Reference/Events/fullscreen
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func FullscreenEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("fullscreen", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// GamepadConnectedEvent Documentation is as below: "A gamepad has been connected."
// https://developer.mozilla.org/docs/Web/Events/gamepadconnected
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func GamepadConnectedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("gamepadconnected", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// GamepadDisconnectedEvent Documentation is as below: "A gamepad has been disconnected."
// https://developer.mozilla.org/docs/Web/Events/gamepaddisconnected
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func GamepadDisconnectedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("gamepaddisconnected", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// GotpointercaptureEvent Documentation is as below: "Element receives pointer capture."
// https://developer.mozilla.org/docs/Web/Events/gotpointercapture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func GotpointercaptureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("gotpointercapture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// HashChangeEvent Documentation is as below: "The fragment identifier of the URL has changed (the part of the URL after the #)."
// https://developer.mozilla.org/docs/Web/Events/hashchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func HashChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("hashchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// HeldEvent Documentation is as below: "A call has been held."
// https://developer.mozilla.org/docs/Web/Events/held
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func HeldEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("held", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// HoldingEvent Documentation is as below: "A call is about to be held."
// https://developer.mozilla.org/docs/Web/Events/holding
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func HoldingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("holding", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// IcccardlockerrorEvent Documentation is as below: "the MozMobileConnection.unlockCardLock() or MozMobileConnection.setCardLock() methods fails."
// https://developer.mozilla.org/docs/Web/Events/icccardlockerror
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func IcccardlockerrorEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("icccardlockerror", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// IccinfochangeEvent Documentation is as below: "The MozMobileConnection.iccInfo object changes."
// https://developer.mozilla.org/docs/Web/Events/iccinfochange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func IccinfochangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("iccinfochange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// IncomingEvent Documentation is as below: "A call is being received."
// https://developer.mozilla.org/docs/Web/Events/incoming
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func IncomingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("incoming", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// InputEvent Documentation is as below: "The value of an element changes or the content of an element with the attribute contenteditable is modified."
// https://developer.mozilla.org/docs/Web/Events/input
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func InputEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("input", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// InvalidEvent Documentation is as below: "A submittable element has been checked and doesn't satisfy its constraints."
// https://developer.mozilla.org/docs/Web/Events/invalid
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func InvalidEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("invalid", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// KeyDownEvent Documentation is as below: "A key is pressed down."
// https://developer.mozilla.org/docs/Web/Events/keydown
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func KeyDownEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("keydown", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// KeyPressEvent Documentation is as below: "A key is pressed down and that key normally produces a character value (use input instead)."
// https://developer.mozilla.org/docs/Web/Events/keypress
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func KeyPressEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("keypress", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// KeyUpEvent Documentation is as below: "A key is released."
// https://developer.mozilla.org/docs/Web/Events/keyup
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func KeyUpEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("keyup", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LanguageChangeEvent Documentation is as below: "The user's preferred languages have changed."
// https://developer.mozilla.org/docs/Web/Events/languagechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LanguageChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("languagechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LevelChangeEvent Documentation is as below: "The level attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/levelchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LevelChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("levelchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LoadEvent Documentation is as below: "Progression has been successful."
// https://developer.mozilla.org/docs/Web/Reference/Events/load_(ProgressEvent)
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LoadEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("load", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LoadEndEvent Documentation is as below: "Progress has stopped (after \"error\", \"abort\" or \"load\" have been dispatched)."
// https://developer.mozilla.org/docs/Web/Events/loadend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LoadEndEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("loadend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LoadStartEvent Documentation is as below: "Progress has begun."
// https://developer.mozilla.org/docs/Web/Events/loadstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LoadStartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("loadstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LoadedDataEvent Documentation is as below: "The first frame of the media has finished loading."
// https://developer.mozilla.org/docs/Web/Events/loadeddata
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LoadedDataEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("loadeddata", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LoadedMetadataEvent Documentation is as below: "The metadata has been loaded."
// https://developer.mozilla.org/docs/Web/Events/loadedmetadata
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LoadedMetadataEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("loadedmetadata", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LocalizedEvent Documentation is as below: "The page has been localized using data-l10n-* attributes."
// https://developer.mozilla.org/docs/Web/Events/localized
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LocalizedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("localized", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// LostpointercaptureEvent Documentation is as below: "Element lost pointer capture."
// https://developer.mozilla.org/docs/Web/Events/lostpointercapture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func LostpointercaptureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("lostpointercapture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MarkEvent Documentation is as below: "The spoken utterance reaches a named SSML \"mark\" tag."
// https://developer.mozilla.org/docs/Web/Events/mark
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MarkEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mark", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MessageEvent Documentation is as below: "A message is received from a service worker, or a message is received in a service worker from another context."
// https://developer.mozilla.org/docs/Web/Events/message_(ServiceWorker)
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MessageEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("message", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MouseDownEvent Documentation is as below: "A pointing device button (usually a mouse) is pressed on an element."
// https://developer.mozilla.org/docs/Web/Events/mousedown
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MouseDownEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mousedown", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MouseEnterEvent Documentation is as below: "A pointing device is moved onto the element that has the listener attached."
// https://developer.mozilla.org/docs/Web/Events/mouseenter
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MouseEnterEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mouseenter", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MouseLeaveEvent Documentation is as below: "A pointing device is moved off the element that has the listener attached."
// https://developer.mozilla.org/docs/Web/Events/mouseleave
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MouseLeaveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mouseleave", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MouseMoveEvent Documentation is as below: "A pointing device is moved over an element."
// https://developer.mozilla.org/docs/Web/Events/mousemove
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MouseMoveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mousemove", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MouseOutEvent Documentation is as below: "A pointing device is moved off the element that has the listener attached or off one of its children."
// https://developer.mozilla.org/docs/Web/Events/mouseout
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MouseOutEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mouseout", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MouseOverEvent Documentation is as below: "A pointing device is moved onto the element that has the listener attached or onto one of its children."
// https://developer.mozilla.org/docs/Web/Events/mouseover
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MouseOverEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mouseover", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MouseUpEvent Documentation is as below: "A pointing device button is released over an element."
// https://developer.mozilla.org/docs/Web/Events/mouseup
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MouseUpEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mouseup", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozAfterPaintEvent Documentation is as below: "Content has been repainted."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozAfterPaint
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozAfterPaintEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozAfterPaint", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozAudioAvailableEvent Documentation is as below: "The audio buffer is full and the corresponding raw samples are available."
// https://developer.mozilla.org/docs/Web/Events/MozAudioAvailable
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozAudioAvailableEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozAudioAvailable", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozBeforeResizeEvent Documentation is as below: "A window is about to be resized."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozBeforeResize
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozBeforeResizeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozBeforeResize", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozEdgeUIGestureEvent Documentation is as below: "A touch point is swiped across the touch surface to invoke the edge UI (Win8 only)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozEdgeUIGesture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozEdgeUIGestureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozEdgeUIGesture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozEnteredDomFullscreenEvent Documentation is as below: "DOM fullscreen mode has been entered."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozEnteredDomFullscreen
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozEnteredDomFullscreenEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozEnteredDomFullscreen", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozGamepadButtonDownEvent Documentation is as below: "A gamepad button is pressed down."
// https://developer.mozilla.org/docs/Web/Events/MozGamepadButtonDown
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozGamepadButtonDownEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozGamepadButtonDown", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozGamepadButtonUpEvent Documentation is as below: "A gamepad button is released."
// https://developer.mozilla.org/docs/Web/Events/MozGamepadButtonUp
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozGamepadButtonUpEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozGamepadButtonUp", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozMagnifyGestureEvent Documentation is as below: "Two touch points moved away from each other (after a sequence of MozMagnifyGestureUpdate)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozMagnifyGesture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozMagnifyGestureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozMagnifyGesture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozMagnifyGestureStartEvent Documentation is as below: "Two touch points start to move away from each other."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozMagnifyGestureStart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozMagnifyGestureStartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozMagnifyGestureStart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozMagnifyGestureUpdateEvent Documentation is as below: "Two touch points move away from each other (after a MozMagnifyGestureStart)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozMagnifyGestureUpdate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozMagnifyGestureUpdateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozMagnifyGestureUpdate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozPressTapGestureEvent Documentation is as below: "A \"press-tap\" gesture happened on the touch surface (first finger down, second finger down, second finger up, first finger up)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozPressTapGesture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozPressTapGestureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozPressTapGesture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozRotateGestureEvent Documentation is as below: "Two touch points rotate around a point (after a sequence of MozRotateGestureUpdate)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozRotateGesture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozRotateGestureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozRotateGesture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozRotateGestureStartEvent Documentation is as below: "Two touch points start to rotate around a point."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozRotateGestureStart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozRotateGestureStartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozRotateGestureStart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozRotateGestureUpdateEvent Documentation is as below: "Two touch points rotate around a point (after a MozRotateGestureStart)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozRotateGestureUpdate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozRotateGestureUpdateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozRotateGestureUpdate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozScrolledAreaChangedEvent Documentation is as below: "The document view has been scrolled or resized."
// https://developer.mozilla.org/docs/Web/Events/MozScrolledAreaChanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozScrolledAreaChangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozScrolledAreaChanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozSwipeGestureEvent Documentation is as below: "A touch point is swiped across the touch surface"
// https://developer.mozilla.org/docs/Web/Reference/Events/MozSwipeGesture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozSwipeGestureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozSwipeGesture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozTapGestureEvent Documentation is as below: "Two touch points are tapped on the touch surface."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozTapGesture
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozTapGestureEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("MozTapGesture", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowseractivitydoneEvent Documentation is as below: "Sent when some activity has been completed (complete description TBD.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowseractivitydone
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowseractivitydoneEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowseractivitydone", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserasyncscrollEvent Documentation is as below: "Sent when the scroll position within a browser <iframe> changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserasyncscroll
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserasyncscrollEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserasyncscroll", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowseraudioplaybackchangeEvent Documentation is as below: "Sent when audio starts or stops playing within the browser <iframe> content."
// https://developer.mozilla.org/docs/Web/Events/mozbrowseraudioplaybackchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowseraudioplaybackchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowseraudioplaybackchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsercaretstatechangedEvent Documentation is as below: "Sent when the text selected inside the browser <iframe> content changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsercaretstatechanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsercaretstatechangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsercaretstatechanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsercloseEvent Documentation is as below: "Sent when window.close() is called within a browser <iframe>."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserclose
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsercloseEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserclose", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsercontextmenuEvent Documentation is as below: "Sent when a browser <iframe> try to open a context menu."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsercontextmenu
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsercontextmenuEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsercontextmenu", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserdocumentfirstpaintEvent Documentation is as below: "Sent when a new paint occurs on any document in the browser <iframe>."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserdocumentfirstpaint
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserdocumentfirstpaintEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserdocumentfirstpaint", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsererrorEvent Documentation is as below: "Sent when an error occured while trying to load a content within a browser iframe"
// https://developer.mozilla.org/docs/Web/Events/mozbrowsererror
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsererrorEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsererror", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserfindchangeEvent Documentation is as below: "Sent when a search operation is performed on the browser <iframe> content (see HTMLIFrameElement search methods.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowserfindchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserfindchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserfindchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserfirstpaintEvent Documentation is as below: "Sent when the <iframe> paints content for the first time (this doesn't include the initial paint from about:blank.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowserfirstpaint
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserfirstpaintEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserfirstpaint", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsericonchangeEvent Documentation is as below: "Sent when the favicon of a browser iframe changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsericonchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsericonchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsericonchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserloadendEvent Documentation is as below: "Sent when the browser iframe has finished loading all its assets."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserloadend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserloadendEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserloadend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserloadstartEvent Documentation is as below: "Sent when the browser iframe starts to load a new page."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserloadstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserloadstartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserloadstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserlocationchangeEvent Documentation is as below: "Sent when an browser iframe's location changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserlocationchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserlocationchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserlocationchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsermanifestchangeEvent Documentation is as below: "Sent when a the path to the app manifest changes, in the case of a browser <iframe> with an open web app embedded in it."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsermanifestchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsermanifestchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsermanifestchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsermetachangeEvent Documentation is as below: "Sent when a <meta> elelment is added to, removed from or changed in the browser <iframe>'s content."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsermetachange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsermetachangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsermetachange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowseropensearchEvent Documentation is as below: "Sent when a link to a search engine is found."
// https://developer.mozilla.org/docs/Web/Events/mozbrowseropensearch
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowseropensearchEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowseropensearch", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowseropentabEvent Documentation is as below: "Sent when a new tab is opened within a browser <iframe> as a result of the user issuing a command to open a link target in a new tab (for example ctrl/cmd + click.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowseropentab
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowseropentabEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowseropentab", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowseropenwindowEvent Documentation is as below: "Sent when window.open() is called within a browser iframe."
// https://developer.mozilla.org/docs/Web/Events/mozbrowseropenwindow
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowseropenwindowEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowseropenwindow", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserresizeEvent Documentation is as below: "Sent when the browser <iframe>'s window size has changed."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserresize
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserresizeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserresize", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserscrollEvent Documentation is as below: "Sent when the browser <iframe> content scrolls."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserscroll
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserscrollEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserscroll", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserscrollareachangedEvent Documentation is as below: "Sent when the available scrolling area\u00a0 in the browser <iframe> changes. This can occur on resize and when the page size changes (while loading for example.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowserscrollareachanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserscrollareachangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserscrollareachanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserscrollviewchangeEvent Documentation is as below: "Sent when asynchronous scrolling (i.e. APCZ) starts or stops."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserscrollviewchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserscrollviewchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserscrollviewchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsersecuritychangeEvent Documentation is as below: "Sent when the SSL state changes within a browser iframe."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsersecuritychange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsersecuritychangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsersecuritychange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserselectionstatechangedEvent Documentation is as below: "Sent when the text selected inside the browser <iframe> content changes. Note that this is deprecated, and newer implementations use mozbrowsercaretstatechanged instead."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserselectionstatechanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserselectionstatechangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserselectionstatechanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsershowmodalpromptEvent Documentation is as below: "Sent when alert(), confirm() or prompt() are called within a browser iframe"
// https://developer.mozilla.org/docs/Web/Events/mozbrowsershowmodalprompt
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsershowmodalpromptEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsershowmodalprompt", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowsertitlechangeEvent Documentation is as below: "Sent when the document.title changes within a browser iframe."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsertitlechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowsertitlechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowsertitlechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowserusernameandpasswordrequiredEvent Documentation is as below: "Sent when an HTTP authentification is requested."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserusernameandpasswordrequired
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowserusernameandpasswordrequiredEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowserusernameandpasswordrequired", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MozbrowservisibilitychangeEvent Documentation is as below: "Sent when the visibility state of the current browser iframe <iframe> changes, for example due to a call to setVisible()."
// https://developer.mozilla.org/docs/Web/Events/mozbrowservisibilitychange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MozbrowservisibilitychangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("mozbrowservisibilitychange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// MoztimechangeEvent Documentation is as below: "The time of the device has been changed."
// https://developer.mozilla.org/docs/Web/Events/moztimechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func MoztimechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("moztimechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// NoUpdateEvent Documentation is as below: "The manifest hadn't changed."
// https://developer.mozilla.org/docs/Web/Events/noupdate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func NoUpdateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("noupdate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// NomatchEvent Documentation is as below: "The speech recognition service returns a final result with no significant recognition."
// https://developer.mozilla.org/docs/Web/Events/nomatch
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func NomatchEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("nomatch", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// NotificationclickEvent Documentation is as below: "A system notification spawned by ServiceWorkerRegistration.showNotification() has been clicked."
// https://developer.mozilla.org/docs/Web/Events/notificationclick
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func NotificationclickEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("notificationclick", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ObsoleteEvent Documentation is as below: "The manifest was found to have become a 404 or 410 page, so the application cache is being deleted."
// https://developer.mozilla.org/docs/Web/Events/obsolete
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ObsoleteEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("obsolete", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// OfflineEvent Documentation is as below: "The browser has lost access to the network."
// https://developer.mozilla.org/docs/Web/Events/offline
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func OfflineEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("offline", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// OnconnectedEvent Documentation is as below: "A call has been connected."
// https://developer.mozilla.org/docs/DOM/onconnected
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func OnconnectedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("onconnected", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// OnlineEvent Documentation is as below: "The browser has gained access to the network (but particular websites might be unreachable)."
// https://developer.mozilla.org/docs/Web/Events/online
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func OnlineEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("online", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// OpenEvent Documentation is as below: "An event source connection has been established."
// https://developer.mozilla.org/docs/Web/Reference/Events/open_serversentevents
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func OpenEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("open", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// OrientationChangeEvent Documentation is as below: "The orientation of the device (portrait/landscape) has changed"
// https://developer.mozilla.org/docs/Web/Events/orientationchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func OrientationChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("orientationchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// OverflowEvent Documentation is as below: "An element has been overflowed by its content or has been rendered for the first time in this state (only works for elements styled with overflow != visible)."
// https://developer.mozilla.org/docs/Web/Events/overflow
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func OverflowEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("overflow", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PageHideEvent Documentation is as below: "A session history entry is being traversed from."
// https://developer.mozilla.org/docs/Web/Events/pagehide
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PageHideEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pagehide", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PageShowEvent Documentation is as below: "A session history entry is being traversed to."
// https://developer.mozilla.org/docs/Web/Events/pageshow
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PageShowEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pageshow", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PasteEvent Documentation is as below: "Data has been transferred from the system clipboard to the document."
// https://developer.mozilla.org/docs/Web/Events/paste
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PasteEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("paste", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PauseEvent Documentation is as below: "The utterance is paused part way through."
// https://developer.mozilla.org/docs/Web/Events/pause_(SpeechSynthesis)
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PauseEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pause", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PlayEvent Documentation is as below: "Playback has begun."
// https://developer.mozilla.org/docs/Web/Events/play
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PlayEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("play", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PlayingEvent Documentation is as below: "Playback is ready to start after having been paused or delayed due to lack of data."
// https://developer.mozilla.org/docs/Web/Events/playing
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PlayingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("playing", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointerLockChangeEvent Documentation is as below: "The pointer was locked or released."
// https://developer.mozilla.org/docs/Web/Events/pointerlockchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointerLockChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerlockchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointerLockErrorEvent Documentation is as below: "It was impossible to lock the pointer for technical reasons or because the permission was denied."
// https://developer.mozilla.org/docs/Web/Events/pointerlockerror
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointerLockErrorEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerlockerror", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointercancelEvent Documentation is as below: "The pointer is unlikely to produce any more events."
// https://developer.mozilla.org/docs/Web/Events/pointercancel
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointercancelEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointercancel", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointerdownEvent Documentation is as below: "The pointer enters the active buttons state."
// https://developer.mozilla.org/docs/Web/Events/pointerdown
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointerdownEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerdown", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointerenterEvent Documentation is as below: "Pointing device is moved inside the hit-testing boundary."
// https://developer.mozilla.org/docs/Web/Events/pointerenter
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointerenterEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerenter", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointerleaveEvent Documentation is as below: "Pointing device is moved out of the hit-testing boundary."
// https://developer.mozilla.org/docs/Web/Events/pointerleave
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointerleaveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerleave", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointermoveEvent Documentation is as below: "The pointer changed coordinates."
// https://developer.mozilla.org/docs/Web/Events/pointermove
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointermoveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointermove", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointeroutEvent Documentation is as below: "The pointing device moved out of hit-testing boundary or leaves detectable hover range."
// https://developer.mozilla.org/docs/Web/Events/pointerout
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointeroutEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerout", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointeroverEvent Documentation is as below: "The pointing device is moved into the hit-testing boundary."
// https://developer.mozilla.org/docs/Web/Events/pointerover
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointeroverEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerover", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PointerupEvent Documentation is as below: "The pointer leaves the active buttons state."
// https://developer.mozilla.org/docs/Web/Events/pointerup
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PointerupEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pointerup", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PopStateEvent Documentation is as below: "A session history entry is being navigated to (in certain cases)."
// https://developer.mozilla.org/docs/Web/Events/popstate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PopStateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("popstate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PopuphiddenEvent Documentation is as below: "A menupopup, panel or tooltip has been hidden."
// https://developer.mozilla.org/docs/Web/Events/popuphidden
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PopuphiddenEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("popuphidden", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PopuphidingEvent Documentation is as below: "A menupopup, panel or tooltip is about to be hidden."
// https://developer.mozilla.org/docs/Web/Events/popuphiding
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PopuphidingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("popuphiding", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PopupshowingEvent Documentation is as below: "A menupopup, panel or tooltip is about to become visible."
// https://developer.mozilla.org/docs/Web/Events/popupshowing
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PopupshowingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("popupshowing", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PopupshownEvent Documentation is as below: "A menupopup, panel or tooltip has become visible."
// https://developer.mozilla.org/docs/Web/Events/popupshown
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PopupshownEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("popupshown", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ProgressEvent Documentation is as below: "The user agent is downloading resources listed by the manifest."
// https://developer.mozilla.org/docs/Web/Reference/Events/progress_(appcache_event)
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ProgressEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("progress", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PushEvent Documentation is as below: "A Service Worker has received a push message."
// https://developer.mozilla.org/docs/Web/Events/push
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PushEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("push", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// PushsubscriptionchangeEvent Documentation is as below: "A PushSubscription has expired."
// https://developer.mozilla.org/docs/Web/Events/pushsubscriptionchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func PushsubscriptionchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("pushsubscriptionchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// RadioStateChangeEvent Documentation is as below: "The state of a radio has been changed either by a user action or by a script (useful for accessibility)."
// https://developer.mozilla.org/docs/Web/Events/RadioStateChange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func RadioStateChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("RadioStateChange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// RateChangeEvent Documentation is as below: "The playback rate has changed."
// https://developer.mozilla.org/docs/Web/Events/ratechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func RateChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("ratechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ReadystateChangeEvent Documentation is as below: "The readyState attribute of a document has changed."
// https://developer.mozilla.org/docs/Web/Events/readystatechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ReadystateChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("readystatechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ReceivedEvent Documentation is as below: "An SMS has been received."
// https://developer.mozilla.org/docs/Web/Events/received
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ReceivedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("received", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// RepeatEventEvent Documentation is as below: "A SMIL animation element is repeated."
// https://developer.mozilla.org/docs/Web/Events/repeatEvent
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func RepeatEventEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("repeatEvent", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ResetEvent Documentation is as below: "A form is reset."
// https://developer.mozilla.org/docs/Web/Events/reset
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ResetEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("reset", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ResizeEvent Documentation is as below: "The document view has been resized."
// https://developer.mozilla.org/docs/Web/Events/resize
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ResizeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("resize", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ResourcetimingbufferfullEvent Documentation is as below: "The browser's resource timing buffer is full."
// https://developer.mozilla.org/docs/Web/Events/resourcetimingbufferfull
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ResourcetimingbufferfullEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("resourcetimingbufferfull", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ResultEvent Documentation is as below: "The speech recognition service returns a result — a word or phrase has been positively recognized and this has been communicated back to the app."
// https://developer.mozilla.org/docs/Web/Events/result
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ResultEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("result", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ResumeEvent Documentation is as below: "A paused utterance is resumed."
// https://developer.mozilla.org/docs/Web/Events/resume
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ResumeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("resume", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ResumingEvent Documentation is as below: "A call is about to resume."
// https://developer.mozilla.org/docs/Web/Events/resuming
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ResumingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("resuming", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SSTabClosingEvent Documentation is as below: "The session store will stop tracking this tab."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSTabClosing
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SSTabClosingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SSTabClosing", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SSTabRestoredEvent Documentation is as below: "A tab has been restored."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSTabRestored
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SSTabRestoredEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SSTabRestored", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SSTabRestoringEvent Documentation is as below: "A tab is about to be restored."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSTabRestoring
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SSTabRestoringEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SSTabRestoring", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SSWindowClosingEvent Documentation is as below: "The session store will stop tracking this window."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSWindowClosing
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SSWindowClosingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SSWindowClosing", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SSWindowStateBusyEvent Documentation is as below: "A window state has switched to \"busy\"."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSWindowStateBusy
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SSWindowStateBusyEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SSWindowStateBusy", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SSWindowStateReadyEvent Documentation is as below: "A window state has switched to \"ready\"."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSWindowStateReady
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SSWindowStateReadyEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SSWindowStateReady", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SVGAbortEvent Documentation is as below: "Page loading has been stopped before the SVG was loaded."
// https://developer.mozilla.org/docs/Web/Events/SVGAbort
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SVGAbortEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SVGAbort", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SVGErrorEvent Documentation is as below: "An error has occurred before the SVG was loaded."
// https://developer.mozilla.org/docs/Web/Events/SVGError
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SVGErrorEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SVGError", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SVGLoadEvent Documentation is as below: "An SVG document has been loaded and parsed."
// https://developer.mozilla.org/docs/Web/Events/SVGLoad
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SVGLoadEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SVGLoad", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SVGResizeEvent Documentation is as below: "An SVG document is being resized."
// https://developer.mozilla.org/docs/Web/Events/SVGResize
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SVGResizeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SVGResize", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SVGScrollEvent Documentation is as below: "An SVG document is being scrolled."
// https://developer.mozilla.org/docs/Web/Events/SVGScroll
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SVGScrollEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SVGScroll", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SVGUnloadEvent Documentation is as below: "An SVG document has been removed from a window or frame."
// https://developer.mozilla.org/docs/Web/Events/SVGUnload
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SVGUnloadEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SVGUnload", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SVGZoomEvent Documentation is as below: "An SVG document is being zoomed."
// https://developer.mozilla.org/docs/Web/Events/SVGZoom
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SVGZoomEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("SVGZoom", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ScrollEvent Documentation is as below: "The document view or an element has been scrolled."
// https://developer.mozilla.org/docs/Web/Events/scroll
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ScrollEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("scroll", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SeekedEvent Documentation is as below: "A seek operation completed."
// https://developer.mozilla.org/docs/Web/Events/seeked
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SeekedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("seeked", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SeekingEvent Documentation is as below: "A seek operation began."
// https://developer.mozilla.org/docs/Web/Events/seeking
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SeekingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("seeking", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SelectEvent Documentation is as below: "Some text is being selected."
// https://developer.mozilla.org/docs/Web/Events/select
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SelectEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("select", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SelectionchangeEvent Documentation is as below: "The selection in the document has been changed."
// https://developer.mozilla.org/docs/Web/Events/selectionchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SelectionchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("selectionchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SelectstartEvent Documentation is as below: "A selection just started."
// https://developer.mozilla.org/docs/Web/Events/selectstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SelectstartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("selectstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SentEvent Documentation is as below: "An SMS has been sent."
// https://developer.mozilla.org/docs/Web/Events/sent
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SentEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("sent", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ShowEvent Documentation is as below: "A contextmenu event was fired on/bubbled to an element that has a contextmenu attribute"
// https://developer.mozilla.org/docs/Web/Events/show
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ShowEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("show", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SizemodechangeEvent Documentation is as below: "Window has entered/left fullscreen mode, or has been minimized/unminimized."
// https://developer.mozilla.org/docs/Web/Reference/Events/sizemodechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SizemodechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("sizemodechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SmartCardInsertEvent Documentation is as below: "A smartcard has been inserted."
// https://developer.mozilla.org/docs/Web/Events/smartcard-insert
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SmartCardInsertEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("smartcard-insert", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SmartCardRemoveEvent Documentation is as below: "A smartcard has been removed."
// https://developer.mozilla.org/docs/Web/Events/smartcard-remove
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SmartCardRemoveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("smartcard-remove", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SoundendEvent Documentation is as below: "Any sound — recognisable speech or not — has stopped being detected."
// https://developer.mozilla.org/docs/Web/Events/soundend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SoundendEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("soundend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SoundstartEvent Documentation is as below: "Any sound — recognisable speech or not — has been detected."
// https://developer.mozilla.org/docs/Web/Events/soundstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SoundstartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("soundstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SpeechendEvent Documentation is as below: "Speech recognised by the speech recognition service has stopped being detected."
// https://developer.mozilla.org/docs/Web/Events/speechend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SpeechendEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("speechend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SpeechstartEvent Documentation is as below: "Sound that is recognised by the speech recognition service as speech has been detected."
// https://developer.mozilla.org/docs/Web/Events/speechstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SpeechstartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("speechstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// StalledEvent Documentation is as below: "The user agent is trying to fetch media data, but data is unexpectedly not forthcoming."
// https://developer.mozilla.org/docs/Web/Events/stalled
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func StalledEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("stalled", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// StartEvent Documentation is as below: "The utterance has begun to be spoken."
// https://developer.mozilla.org/docs/Web/Events/start_(SpeechSynthesis)
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func StartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("start", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// StatechangeEvent Documentation is as below: "The state of a call has changed."
// https://developer.mozilla.org/docs/Web/Events/statechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func StatechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("statechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// StatuschangeEvent Documentation is as below: "The status of the Wifi connection changed."
// https://developer.mozilla.org/docs/Web/Events/statuschange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func StatuschangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("statuschange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// StkcommandEvent Documentation is as below: "The STK Proactive Command is issued from ICC."
// https://developer.mozilla.org/docs/Web/Events/stkcommand
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func StkcommandEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("stkcommand", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// StksessionendEvent Documentation is as below: "The STK Session is terminated by ICC."
// https://developer.mozilla.org/docs/Web/Events/stksessionend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func StksessionendEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("stksessionend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// StorageEvent Documentation is as below: "A storage area (localStorage or sessionStorage) has changed."
// https://developer.mozilla.org/docs/Web/Events/storage
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func StorageEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("storage", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SubmitEvent Documentation is as below: "A form is submitted."
// https://developer.mozilla.org/docs/Web/Events/submit
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SubmitEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("submit", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SuccessEvent Documentation is as below: "A request successfully completed."
// https://developer.mozilla.org/docs/Web/Reference/Events/success_indexedDB
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SuccessEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("success", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// SuspendEvent Documentation is as below: "Media data loading has been suspended."
// https://developer.mozilla.org/docs/Web/Events/suspend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func SuspendEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("suspend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TabCloseEvent Documentation is as below: "A tab has been closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabClose
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TabCloseEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("TabClose", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TabHideEvent Documentation is as below: "A tab has been hidden."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabHide
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TabHideEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("TabHide", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TabOpenEvent Documentation is as below: "A tab has been opened."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabOpen
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TabOpenEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("TabOpen", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TabPinnedEvent Documentation is as below: "A tab has been pinned."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabPinned
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TabPinnedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("TabPinned", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TabSelectEvent Documentation is as below: "A tab has been selected."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabSelect
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TabSelectEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("TabSelect", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TabShowEvent Documentation is as below: "A tab has been shown."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabShow
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TabShowEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("TabShow", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TabUnpinnedEvent Documentation is as below: "A tab has been unpinned."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabUnpinned
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TabUnpinnedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("TabUnpinned", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TimeUpdateEvent Documentation is as below: "The time indicated by the currentTime attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/timeupdate
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TimeUpdateEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("timeupdate", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TimeoutEvent Documentation is as below: "(no documentation)"
// https://developer.mozilla.org/docs/Web/Events/timeout
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TimeoutEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("timeout", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TouchCancelEvent Documentation is as below: "A touch point has been disrupted in an implementation-specific manners (too many touch points for example)."
// https://developer.mozilla.org/docs/Web/Events/touchcancel
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TouchCancelEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("touchcancel", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TouchEndEvent Documentation is as below: "A touch point is removed from the touch surface."
// https://developer.mozilla.org/docs/Web/Events/touchend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TouchEndEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("touchend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TouchEnterEvent Documentation is as below: "(no documentation)"
// https://developer.mozilla.org/docs/Web/Events/touchenter
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TouchEnterEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("touchenter", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TouchLeaveEvent Documentation is as below: "(no documentation)"
// https://developer.mozilla.org/docs/Web/Events/touchleave
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TouchLeaveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("touchleave", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TouchMoveEvent Documentation is as below: "A touch point is moved along the touch surface."
// https://developer.mozilla.org/docs/Web/Events/touchmove
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TouchMoveEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("touchmove", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TouchStartEvent Documentation is as below: "A touch point is placed on the touch surface."
// https://developer.mozilla.org/docs/Web/Events/touchstart
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TouchStartEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("touchstart", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// TransitionEndEvent Documentation is as below: "A CSS transition has completed."
// https://developer.mozilla.org/docs/Web/Events/transitionend
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func TransitionEndEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("transitionend", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// UnderflowEvent Documentation is as below: "An element is no longer overflowed by its content (only works for elements styled with overflow != visible)."
// https://developer.mozilla.org/docs/Web/Events/underflow
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func UnderflowEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("underflow", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// UnloadEvent Documentation is as below: "The document or a dependent resource is being unloaded."
// https://developer.mozilla.org/docs/Web/Events/unload
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func UnloadEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("unload", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// UpdateReadyEvent Documentation is as below: "The resources listed in the manifest have been newly redownloaded, and the script can use swapCache() to switch to the new cache."
// https://developer.mozilla.org/docs/Web/Events/updateready
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func UpdateReadyEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("updateready", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// UpgradeNeededEvent Documentation is as below: "An attempt was made to open a database with a version number higher than its current version. A versionchange transaction has been created."
// https://developer.mozilla.org/docs/Web/Reference/Events/upgradeneeded_indexedDB
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func UpgradeNeededEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("upgradeneeded", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// UserProximityEvent Documentation is as below: "Fresh data is available from a proximity sensor (indicates whether the nearby object is near the device or not)."
// https://developer.mozilla.org/docs/Web/Events/userproximity
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func UserProximityEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("userproximity", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// UssdreceivedEvent Documentation is as below: "A new USSD message is received"
// https://developer.mozilla.org/docs/Web/Events/ussdreceived
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func UssdreceivedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("ussdreceived", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// ValueChangeEvent Documentation is as below: "The value of an element has changed (a progress bar for example, useful for accessibility)."
// https://developer.mozilla.org/docs/Web/Events/ValueChange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func ValueChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("ValueChange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VersionChangeEvent Documentation is as below: "A versionchange transaction completed."
// https://developer.mozilla.org/docs/Web/Reference/Events/versionchange_indexedDB
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VersionChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("versionchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VisibilityChangeEvent Documentation is as below: "The content of a tab has become visible or has been hidden."
// https://developer.mozilla.org/docs/Web/Events/visibilitychange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VisibilityChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("visibilitychange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VoicechangeEvent Documentation is as below: "The MozMobileConnection.voice object changes values."
// https://developer.mozilla.org/docs/Web/Events/voicechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VoicechangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("voicechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VoiceschangedEvent Documentation is as below: "The list of SpeechSynthesisVoice objects that would be returned by the SpeechSynthesis.getVoices() method has changed (when the voiceschanged event fires.)"
// https://developer.mozilla.org/docs/Web/Events/voiceschanged
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VoiceschangedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("voiceschanged", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VolumeChangeEvent Documentation is as below: "The volume has changed."
// https://developer.mozilla.org/docs/Web/Events/volumechange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VolumeChangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("volumechange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VrdisplayconnectedEvent Documentation is as below: "A compatible VR device has been connected to the computer."
// https://developer.mozilla.org/docs/Web/Events/vrdisplayconnected
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VrdisplayconnectedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("vrdisplayconnected", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VrdisplaydisconnectedEvent Documentation is as below: "A compatible VR device has been disconnected from the computer."
// https://developer.mozilla.org/docs/Web/Events/vrdisplaydisconnected
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VrdisplaydisconnectedEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("vrdisplaydisconnected", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// VrdisplaypresentchangeEvent Documentation is as below: "The presenting state of a VR device has changed — i.e. from presenting to not presenting, or vice versa."
// https://developer.mozilla.org/docs/Web/Events/vrdisplaypresentchange
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func VrdisplaypresentchangeEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("vrdisplaypresentchange", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// WaitingEvent Documentation is as below: "Playback has stopped because of a temporary lack of data."
// https://developer.mozilla.org/docs/Web/Events/waiting
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func WaitingEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("waiting", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}

// WheelEvent Documentation is as below: "A wheel button of a pointing device is rotated in any direction."
// https://developer.mozilla.org/docs/Web/Events/wheel
// This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector
// mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an
// appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if
// the selector value is not empty then that becomes the default selector used match the event with.
func WheelEvent(callback interface{}, sel string) *trees.Event {
	var handler EventHandler

	switch cb := callback.(type) {
	case func():
		handler = WrapHandler(cb)
	case func(trees.EventObject):
		handler = WrapEventOnlyHandler(cb)
	case func(trees.EventObject, *trees.Markup):
		handler = cb
	default:
		panic("Unacceptable type for event callback")
	}

	ev := trees.NewEvent("wheel", sel)
	ev.Handle = notifications.Subscribe(func(evm trees.EventBroadcast) {
		if ev.EventID != evm.EventID {
			return
		}

		handler(evm.Event, ev.Tree)
	})

	return ev
}
