// The generation of this package was inspired by Neelance work on DOM (https://github.com/neelance/dom)

//go:generate go run generate.go

// Documentation source: "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.

//Package events defines the event binding system that combines different libraries to create a interesting event system.
package events

import (
	"github.com/influx6/gu/trees"
	"github.com/influx6/gu/dispatch"
)

// EventHandler defines a function type for event callbacks.
type EventHandler func(trees.EventObject, *trees.Markup)
 

// Abort Documentation is as below: "A transaction has been aborted."
// https://developer.mozilla.org/docs/Web/Reference/Events/abort_indexedDB
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Abort(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("abort",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// AfterPrint Documentation is as below: "The associated document has started printing or the print preview has been closed."
// https://developer.mozilla.org/docs/Web/Events/afterprint
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func AfterPrint(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("afterprint",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Afterscriptexecute Documentation is as below: "A script has been executed."
// https://developer.mozilla.org/docs/Web/Events/afterscriptexecute
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Afterscriptexecute(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("afterscriptexecute",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// AlertActive Documentation is as below: "A notification element is shown."
// https://developer.mozilla.org/docs/Web/Reference/Events/AlertActive
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func AlertActive(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("AlertActive",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// AlertClose Documentation is as below: "A notification element is closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/AlertClose
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func AlertClose(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("AlertClose",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Alerting Documentation is as below: "The correspondent is being alerted (his/her phone is ringing)."
// https://developer.mozilla.org/docs/Web/Events/alerting
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Alerting(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("alerting",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// AnimationEnd Documentation is as below: "A CSS animation has completed."
// https://developer.mozilla.org/docs/Web/Events/animationend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func AnimationEnd(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("animationend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// AnimationIteration Documentation is as below: "A CSS animation is repeated."
// https://developer.mozilla.org/docs/Web/Events/animationiteration
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func AnimationIteration(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("animationiteration",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// AnimationStart Documentation is as below: "A CSS animation has started."
// https://developer.mozilla.org/docs/Web/Events/animationstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func AnimationStart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("animationstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// AudioProcess Documentation is as below: "The input buffer of a ScriptProcessorNode is ready to be processed."
// https://developer.mozilla.org/docs/Web/Events/audioprocess
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func AudioProcess(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("audioprocess",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Audioend Documentation is as below: "The user agent has finished capturing audio for speech recognition."
// https://developer.mozilla.org/docs/Web/Events/audioend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Audioend(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("audioend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Audiostart Documentation is as below: "The user agent has started to capture audio for speech recognition."
// https://developer.mozilla.org/docs/Web/Events/audiostart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Audiostart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("audiostart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// BeforeInstallPrompt Documentation is as below: "A user is prompted to save a web site to a home screen on mobile."
// https://developer.mozilla.org/docs/Web/Events/beforeinstallprompt
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func BeforeInstallPrompt(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("beforeinstallprompt",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// BeforePrint Documentation is as below: "The associated document is about to be printed or previewed for printing."
// https://developer.mozilla.org/docs/Web/Events/beforeprint
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func BeforePrint(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("beforeprint",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// BeforeScriptExecute Documentation is as below: "A script is about to be executed."
// https://developer.mozilla.org/docs/Web/Events/beforescriptexecute
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func BeforeScriptExecute(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("beforescriptexecute",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// BeforeUnload Documentation is as below: "The window, the document and its resources are about to be unloaded."
// https://developer.mozilla.org/docs/Web/Events/beforeunload
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func BeforeUnload(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("beforeunload",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// BeginEvent Documentation is as below: "A SMIL animation element begins."
// https://developer.mozilla.org/docs/Web/Events/beginEvent
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func BeginEvent(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("beginEvent",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Blocked Documentation is as below: "An open connection to a database is blocking a versionchange transaction on the same database."
// https://developer.mozilla.org/docs/Web/Reference/Events/blocked_indexedDB
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Blocked(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("blocked",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Blur Documentation is as below: "An element has lost focus (does not bubble)."
// https://developer.mozilla.org/docs/Web/Events/blur
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Blur(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("blur",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Boundary Documentation is as below: "The spoken utterance reaches a word or sentence boundary"
// https://developer.mozilla.org/docs/Web/Events/boundary
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Boundary(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("boundary",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Broadcast Documentation is as below: "An observer noticed a change to the attributes of a watched broadcaster."
// https://developer.mozilla.org/docs/Web/Events/broadcast
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Broadcast(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("broadcast",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Busy Documentation is as below: "The line of the correspondent is busy."
// https://developer.mozilla.org/docs/Web/Events/busy
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Busy(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("busy",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Cached Documentation is as below: "The resources listed in the manifest have been downloaded, and the application is now cached."
// https://developer.mozilla.org/docs/Web/Events/cached
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Cached(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("cached",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Callschanged Documentation is as below: "A call has been added or removed from the list of current calls."
// https://developer.mozilla.org/docs/Web/Events/callschanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Callschanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("callschanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CanPlay Documentation is as below: "The user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content."
// https://developer.mozilla.org/docs/Web/Events/canplay
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CanPlay(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("canplay",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CanPlayThrough Documentation is as below: "The user agent can play the media up to its end without having to stop for further buffering of content."
// https://developer.mozilla.org/docs/Web/Events/canplaythrough
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CanPlayThrough(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("canplaythrough",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Cardstatechange Documentation is as below: "The MozMobileConnection.cardState property changes value."
// https://developer.mozilla.org/docs/Web/Events/cardstatechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Cardstatechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("cardstatechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Cfstatechange Documentation is as below: "The call forwarding state changes."
// https://developer.mozilla.org/docs/Web/Events/cfstatechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Cfstatechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("cfstatechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Change Documentation is as below: "This event is triggered each time a file is created, modified or deleted on a given storage area."
// https://developer.mozilla.org/docs/Web/Events/change
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Change(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("change",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// ChargingChange Documentation is as below: "The battery begins or stops charging."
// https://developer.mozilla.org/docs/Web/Events/chargingchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func ChargingChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("chargingchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// ChargingTimeChange Documentation is as below: "The chargingTime attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/chargingtimechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func ChargingTimeChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("chargingtimechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CheckboxStateChange Documentation is as below: "The state of a checkbox has been changed either by a user action or by a script (useful for accessibility)."
// https://developer.mozilla.org/docs/Web/Events/CheckboxStateChange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CheckboxStateChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("CheckboxStateChange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Checking Documentation is as below: "The user agent is checking for an update, or attempting to download the cache manifest for the first time."
// https://developer.mozilla.org/docs/Web/Events/checking
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Checking(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("checking",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Click Documentation is as below: "A pointing device button has been pressed and released on an element."
// https://developer.mozilla.org/docs/Web/Events/click
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Click(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("click",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Close Documentation is as below: "The close button of the window has been clicked."
// https://developer.mozilla.org/docs/Web/Reference/Events/close_event
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Close(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("close",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Command Documentation is as below: "An element has been activated."
// https://developer.mozilla.org/docs/Web/Events/command
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Command(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("command",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Commandupdate Documentation is as below: "A command update occurred on a commandset element."
// https://developer.mozilla.org/docs/Web/Events/commandupdate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Commandupdate(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("commandupdate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Complete Documentation is as below: "The rendering of an OfflineAudioContext is terminated."
// https://developer.mozilla.org/docs/Web/Events/complete
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Complete(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("complete",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CompositionEnd Documentation is as below: "The composition of a passage of text has been completed or canceled."
// https://developer.mozilla.org/docs/Web/Events/compositionend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CompositionEnd(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("compositionend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CompositionStart Documentation is as below: "The composition of a passage of text is prepared (similar to keydown for a keyboard input, but works with other inputs such as speech recognition)."
// https://developer.mozilla.org/docs/Web/Events/compositionstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CompositionStart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("compositionstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CompositionUpdate Documentation is as below: "A character is added to a passage of text being composed."
// https://developer.mozilla.org/docs/Web/Events/compositionupdate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CompositionUpdate(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("compositionupdate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Connecting Documentation is as below: "A call is about to connect."
// https://developer.mozilla.org/docs/Web/Events/connecting
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Connecting(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("connecting",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// ConnectionInfoUpdate Documentation is as below: "The informations about the signal strength and the link speed have been updated."
// https://developer.mozilla.org/docs/Web/Events/connectionInfoUpdate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func ConnectionInfoUpdate(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("connectionInfoUpdate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// ContextMenu Documentation is as below: "The right button of the mouse is clicked (before the context menu is displayed)."
// https://developer.mozilla.org/docs/Web/Events/contextmenu
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func ContextMenu(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("contextmenu",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Copy Documentation is as below: "The text selection has been added to the clipboard."
// https://developer.mozilla.org/docs/Web/Events/copy
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Copy(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("copy",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CssRuleViewCSSLinkClicked Documentation is as below: "A link to a CSS file has been clicked in the \"Rules\" view of the style inspector."
// https://developer.mozilla.org/docs/Web/Reference/Events/CssRuleViewCSSLinkClicked
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CssRuleViewCSSLinkClicked(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("CssRuleViewCSSLinkClicked",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CssRuleViewChanged Documentation is as below: "The \"Rules\" view of the style inspector has been changed."
// https://developer.mozilla.org/docs/Web/Reference/Events/CssRuleViewChanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CssRuleViewChanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("CssRuleViewChanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// CssRuleViewRefreshed Documentation is as below: "The \"Rules\" view of the style inspector has been updated."
// https://developer.mozilla.org/docs/Web/Reference/Events/CssRuleViewRefreshed
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func CssRuleViewRefreshed(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("CssRuleViewRefreshed",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Cut Documentation is as below: "The text selection has been removed from the document and added to the clipboard."
// https://developer.mozilla.org/docs/Web/Events/cut
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Cut(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("cut",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMAutoComplete Documentation is as below: "The content of an element has been auto-completed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMAutoComplete
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMAutoComplete(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMAutoComplete",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMContentLoaded Documentation is as below: "The document has finished loading (but not its dependent resources)."
// https://developer.mozilla.org/docs/Web/Events/DOMContentLoaded
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMContentLoaded(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMContentLoaded",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMFrameContentLoaded Documentation is as below: "The frame has finished loading (but not its dependent resources)."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMFrameContentLoaded
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMFrameContentLoaded(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMFrameContentLoaded",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMLinkAdded Documentation is as below: "A link has been added a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMLinkAdded
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMLinkAdded(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMLinkAdded",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMLinkRemoved Documentation is as below: "A link has been removed inside from a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMLinkRemoved
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMLinkRemoved(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMLinkRemoved",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMMenuItemActive Documentation is as below: "A menu or menuitem has been hovered or highlighted."
// https://developer.mozilla.org/docs/Web/Events/DOMMenuItemActive
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMMenuItemActive(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMMenuItemActive",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMMenuItemInactive Documentation is as below: "A menu or menuitem is no longer hovered or highlighted."
// https://developer.mozilla.org/docs/Web/Events/DOMMenuItemInactive
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMMenuItemInactive(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMMenuItemInactive",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMMetaAdded Documentation is as below: "A meta element has been added to a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMMetaAdded
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMMetaAdded(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMMetaAdded",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMMetaRemoved Documentation is as below: "A meta element has been removed from a document."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMMetaRemoved
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMMetaRemoved(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMMetaRemoved",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMModalDialogClosed Documentation is as below: "A modal dialog has been closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMModalDialogClosed
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMModalDialogClosed(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMModalDialogClosed",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMPopupBlocked Documentation is as below: "A popup has been blocked"
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMPopupBlocked
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMPopupBlocked(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMPopupBlocked",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMTitleChanged Documentation is as below: "The title of a window has changed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMTitleChanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMTitleChanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMTitleChanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMWillOpenModalDialog Documentation is as below: "A modal dialog is about to open."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMWillOpenModalDialog
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMWillOpenModalDialog(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMWillOpenModalDialog",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMWindowClose Documentation is as below: "A window is about to be closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMWindowClose
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMWindowClose(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMWindowClose",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DOMWindowCreated Documentation is as below: "A window has been created."
// https://developer.mozilla.org/docs/Web/Reference/Events/DOMWindowCreated
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DOMWindowCreated(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("DOMWindowCreated",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Datachange Documentation is as below: "The MozMobileConnection.data object changes values."
// https://developer.mozilla.org/docs/Web/Events/datachange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Datachange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("datachange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Dataerror Documentation is as below: "The MozMobileConnection.data object receive an error from the RIL."
// https://developer.mozilla.org/docs/Web/Events/dataerror
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Dataerror(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dataerror",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DblClick Documentation is as below: "A pointing device button is clicked twice on an element."
// https://developer.mozilla.org/docs/Web/Events/dblclick
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DblClick(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dblclick",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Delivered Documentation is as below: "An SMS has been successfully delivered."
// https://developer.mozilla.org/docs/Web/Events/delivered
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Delivered(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("delivered",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DeviceLight Documentation is as below: "Fresh data is available from a light sensor."
// https://developer.mozilla.org/docs/Web/Events/devicelight
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DeviceLight(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("devicelight",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DeviceMotion Documentation is as below: "Fresh data is available from a motion sensor."
// https://developer.mozilla.org/docs/Web/Events/devicemotion
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DeviceMotion(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("devicemotion",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DeviceOrientation Documentation is as below: "Fresh data is available from an orientation sensor."
// https://developer.mozilla.org/docs/Web/Events/deviceorientation
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DeviceOrientation(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("deviceorientation",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DeviceProximity Documentation is as below: "Fresh data is available from a proximity sensor (indicates an approximated distance between the device and a nearby object)."
// https://developer.mozilla.org/docs/Web/Events/deviceproximity
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DeviceProximity(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("deviceproximity",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Devicechange Documentation is as below: "A media device such as a camera, microphone, or speaker is connected or removed from the system."
// https://developer.mozilla.org/docs/Web/Events/devicechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Devicechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("devicechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Dialing Documentation is as below: "The number of a correspondent has been dialed."
// https://developer.mozilla.org/docs/Web/Events/dialing
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Dialing(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dialing",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Disabled Documentation is as below: "Wifi has been disabled on the device."
// https://developer.mozilla.org/docs/Web/Events/disabled
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Disabled(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("disabled",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DischargingTimeChange Documentation is as below: "The dischargingTime attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/dischargingtimechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DischargingTimeChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dischargingtimechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Disconnected Documentation is as below: "A call has been disconnected."
// https://developer.mozilla.org/docs/Web/Events/disconnected
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Disconnected(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("disconnected",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Disconnecting Documentation is as below: "A call is about to disconnect."
// https://developer.mozilla.org/docs/Web/Events/disconnecting
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Disconnecting(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("disconnecting",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Downloading Documentation is as below: "The user agent has found an update and is fetching it, or is downloading the resources listed by the cache manifest for the first time."
// https://developer.mozilla.org/docs/Web/Events/downloading
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Downloading(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("downloading",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Drag Documentation is as below: "An element or text selection is being dragged (every 350ms)."
// https://developer.mozilla.org/docs/Web/Events/drag
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Drag(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("drag",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DragEnd Documentation is as below: "A drag operation is being ended (by releasing a mouse button or hitting the escape key)."
// https://developer.mozilla.org/docs/Web/Events/dragend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DragEnd(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dragend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DragEnter Documentation is as below: "A dragged element or text selection enters a valid drop target."
// https://developer.mozilla.org/docs/Web/Events/dragenter
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DragEnter(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dragenter",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DragLeave Documentation is as below: "A dragged element or text selection leaves a valid drop target."
// https://developer.mozilla.org/docs/Web/Events/dragleave
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DragLeave(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dragleave",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DragOver Documentation is as below: "An element or text selection is being dragged over a valid drop target (every 350ms)."
// https://developer.mozilla.org/docs/Web/Events/dragover
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DragOver(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dragover",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DragStart Documentation is as below: "The user starts dragging an element or text selection."
// https://developer.mozilla.org/docs/Web/Events/dragstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DragStart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("dragstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Drop Documentation is as below: "An element is dropped on a valid drop target."
// https://developer.mozilla.org/docs/Web/Events/drop
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Drop(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("drop",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// DurationChange Documentation is as below: "The duration attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/durationchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func DurationChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("durationchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Emptied Documentation is as below: "The media has become empty; for example, this event is sent if the media has already been loaded (or partially loaded), and the load()\u00a0method is called to reload it."
// https://developer.mozilla.org/docs/Web/Events/emptied
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Emptied(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("emptied",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Enabled Documentation is as below: "Wifi has been enabled on the device."
// https://developer.mozilla.org/docs/Web/Events/enabled
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Enabled(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("enabled",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// End Documentation is as below: "The utterance has finished being spoken."
// https://developer.mozilla.org/docs/Web/Events/end_(SpeechSynthesis)
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func End(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("end",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// EndEvent Documentation is as below: "A SMIL animation element ends."
// https://developer.mozilla.org/docs/Web/Events/endEvent
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func EndEvent(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("endEvent",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Ended Documentation is as below: "Playback has stopped because the end of the media was reached."
// https://developer.mozilla.org/docs/Web/Events/ended_(Web_Audio)
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Ended(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("ended",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Focus Documentation is as below: "An element has received focus (does not bubble)."
// https://developer.mozilla.org/docs/Web/Events/focus
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Focus(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("focus",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// FocusIn Documentation is as below: "An element is about to receive focus (bubbles)."
// https://developer.mozilla.org/docs/Web/Events/focusin
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func FocusIn(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("focusin",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// FocusOut Documentation is as below: "An element is about to lose focus (bubbles)."
// https://developer.mozilla.org/docs/Web/Events/focusout
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func FocusOut(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("focusout",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// FullScreenChange Documentation is as below: "An element was turned to fullscreen mode or back to normal mode."
// https://developer.mozilla.org/docs/Web/Events/fullscreenchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func FullScreenChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("fullscreenchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// FullScreenError Documentation is as below: "It was impossible to switch to fullscreen mode for technical reasons or because the permission was denied."
// https://developer.mozilla.org/docs/Web/Events/fullscreenerror
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func FullScreenError(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("fullscreenerror",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Fullscreen Documentation is as below: "Browser fullscreen mode has been entered or left."
// https://developer.mozilla.org/docs/Web/Reference/Events/fullscreen
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Fullscreen(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("fullscreen",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// GamepadConnected Documentation is as below: "A gamepad has been connected."
// https://developer.mozilla.org/docs/Web/Events/gamepadconnected
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func GamepadConnected(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("gamepadconnected",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// GamepadDisconnected Documentation is as below: "A gamepad has been disconnected."
// https://developer.mozilla.org/docs/Web/Events/gamepaddisconnected
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func GamepadDisconnected(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("gamepaddisconnected",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Gotpointercapture Documentation is as below: "Element receives pointer capture."
// https://developer.mozilla.org/docs/Web/Events/gotpointercapture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Gotpointercapture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("gotpointercapture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// HashChange Documentation is as below: "The fragment identifier of the URL has changed (the part of the URL after the #)."
// https://developer.mozilla.org/docs/Web/Events/hashchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func HashChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("hashchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Held Documentation is as below: "A call has been held."
// https://developer.mozilla.org/docs/Web/Events/held
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Held(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("held",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Holding Documentation is as below: "A call is about to be held."
// https://developer.mozilla.org/docs/Web/Events/holding
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Holding(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("holding",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Icccardlockerror Documentation is as below: "the MozMobileConnection.unlockCardLock() or MozMobileConnection.setCardLock() methods fails."
// https://developer.mozilla.org/docs/Web/Events/icccardlockerror
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Icccardlockerror(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("icccardlockerror",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Iccinfochange Documentation is as below: "The MozMobileConnection.iccInfo object changes."
// https://developer.mozilla.org/docs/Web/Events/iccinfochange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Iccinfochange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("iccinfochange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Incoming Documentation is as below: "A call is being received."
// https://developer.mozilla.org/docs/Web/Events/incoming
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Incoming(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("incoming",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Input Documentation is as below: "The value of an element changes or the content of an element with the attribute contenteditable is modified."
// https://developer.mozilla.org/docs/Web/Events/input
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Input(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("input",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Invalid Documentation is as below: "A submittable element has been checked and doesn't satisfy its constraints."
// https://developer.mozilla.org/docs/Web/Events/invalid
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Invalid(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("invalid",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// KeyDown Documentation is as below: "A key is pressed down."
// https://developer.mozilla.org/docs/Web/Events/keydown
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func KeyDown(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("keydown",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// KeyPress Documentation is as below: "A key is pressed down and that key normally produces a character value (use input instead)."
// https://developer.mozilla.org/docs/Web/Events/keypress
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func KeyPress(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("keypress",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// KeyUp Documentation is as below: "A key is released."
// https://developer.mozilla.org/docs/Web/Events/keyup
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func KeyUp(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("keyup",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// LanguageChange Documentation is as below: "The user's preferred languages have changed."
// https://developer.mozilla.org/docs/Web/Events/languagechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func LanguageChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("languagechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// LevelChange Documentation is as below: "The level attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/levelchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func LevelChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("levelchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Load Documentation is as below: "Progression has been successful."
// https://developer.mozilla.org/docs/Web/Reference/Events/load_(ProgressEvent)
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Load(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("load",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// LoadEnd Documentation is as below: "Progress has stopped (after \"error\", \"abort\" or \"load\" have been dispatched)."
// https://developer.mozilla.org/docs/Web/Events/loadend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func LoadEnd(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("loadend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// LoadStart Documentation is as below: "Progress has begun."
// https://developer.mozilla.org/docs/Web/Events/loadstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func LoadStart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("loadstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// LoadedData Documentation is as below: "The first frame of the media has finished loading."
// https://developer.mozilla.org/docs/Web/Events/loadeddata
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func LoadedData(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("loadeddata",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// LoadedMetadata Documentation is as below: "The metadata has been loaded."
// https://developer.mozilla.org/docs/Web/Events/loadedmetadata
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func LoadedMetadata(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("loadedmetadata",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Localized Documentation is as below: "The page has been localized using data-l10n-* attributes."
// https://developer.mozilla.org/docs/Web/Events/localized
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Localized(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("localized",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Lostpointercapture Documentation is as below: "Element lost pointer capture."
// https://developer.mozilla.org/docs/Web/Events/lostpointercapture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Lostpointercapture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("lostpointercapture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mark Documentation is as below: "The spoken utterance reaches a named SSML \"mark\" tag."
// https://developer.mozilla.org/docs/Web/Events/mark
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mark(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mark",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Message Documentation is as below: "A message is received from a service worker, or a message is received in a service worker from another context."
// https://developer.mozilla.org/docs/Web/Events/message_(ServiceWorker)
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Message(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("message",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MouseDown Documentation is as below: "A pointing device button (usually a mouse) is pressed on an element."
// https://developer.mozilla.org/docs/Web/Events/mousedown
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MouseDown(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mousedown",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MouseEnter Documentation is as below: "A pointing device is moved onto the element that has the listener attached."
// https://developer.mozilla.org/docs/Web/Events/mouseenter
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MouseEnter(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mouseenter",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MouseLeave Documentation is as below: "A pointing device is moved off the element that has the listener attached."
// https://developer.mozilla.org/docs/Web/Events/mouseleave
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MouseLeave(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mouseleave",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MouseMove Documentation is as below: "A pointing device is moved over an element."
// https://developer.mozilla.org/docs/Web/Events/mousemove
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MouseMove(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mousemove",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MouseOut Documentation is as below: "A pointing device is moved off the element that has the listener attached or off one of its children."
// https://developer.mozilla.org/docs/Web/Events/mouseout
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MouseOut(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mouseout",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MouseOver Documentation is as below: "A pointing device is moved onto the element that has the listener attached or onto one of its children."
// https://developer.mozilla.org/docs/Web/Events/mouseover
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MouseOver(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mouseover",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MouseUp Documentation is as below: "A pointing device button is released over an element."
// https://developer.mozilla.org/docs/Web/Events/mouseup
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MouseUp(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mouseup",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozAfterPaint Documentation is as below: "Content has been repainted."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozAfterPaint
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozAfterPaint(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozAfterPaint",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozAudioAvailable Documentation is as below: "The audio buffer is full and the corresponding raw samples are available."
// https://developer.mozilla.org/docs/Web/Events/MozAudioAvailable
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozAudioAvailable(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozAudioAvailable",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozBeforeResize Documentation is as below: "A window is about to be resized."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozBeforeResize
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozBeforeResize(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozBeforeResize",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozEdgeUIGesture Documentation is as below: "A touch point is swiped across the touch surface to invoke the edge UI (Win8 only)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozEdgeUIGesture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozEdgeUIGesture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozEdgeUIGesture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozEnteredDomFullscreen Documentation is as below: "DOM fullscreen mode has been entered."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozEnteredDomFullscreen
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozEnteredDomFullscreen(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozEnteredDomFullscreen",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozGamepadButtonDown Documentation is as below: "A gamepad button is pressed down."
// https://developer.mozilla.org/docs/Web/Events/MozGamepadButtonDown
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozGamepadButtonDown(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozGamepadButtonDown",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozGamepadButtonUp Documentation is as below: "A gamepad button is released."
// https://developer.mozilla.org/docs/Web/Events/MozGamepadButtonUp
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozGamepadButtonUp(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozGamepadButtonUp",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozMagnifyGesture Documentation is as below: "Two touch points moved away from each other (after a sequence of MozMagnifyGestureUpdate)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozMagnifyGesture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozMagnifyGesture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozMagnifyGesture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozMagnifyGestureStart Documentation is as below: "Two touch points start to move away from each other."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozMagnifyGestureStart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozMagnifyGestureStart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozMagnifyGestureStart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozMagnifyGestureUpdate Documentation is as below: "Two touch points move away from each other (after a MozMagnifyGestureStart)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozMagnifyGestureUpdate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozMagnifyGestureUpdate(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozMagnifyGestureUpdate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozPressTapGesture Documentation is as below: "A \"press-tap\" gesture happened on the touch surface (first finger down, second finger down, second finger up, first finger up)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozPressTapGesture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozPressTapGesture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozPressTapGesture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozRotateGesture Documentation is as below: "Two touch points rotate around a point (after a sequence of MozRotateGestureUpdate)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozRotateGesture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozRotateGesture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozRotateGesture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozRotateGestureStart Documentation is as below: "Two touch points start to rotate around a point."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozRotateGestureStart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozRotateGestureStart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozRotateGestureStart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozRotateGestureUpdate Documentation is as below: "Two touch points rotate around a point (after a MozRotateGestureStart)."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozRotateGestureUpdate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozRotateGestureUpdate(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozRotateGestureUpdate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozScrolledAreaChanged Documentation is as below: "The document view has been scrolled or resized."
// https://developer.mozilla.org/docs/Web/Events/MozScrolledAreaChanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozScrolledAreaChanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozScrolledAreaChanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozSwipeGesture Documentation is as below: "A touch point is swiped across the touch surface"
// https://developer.mozilla.org/docs/Web/Reference/Events/MozSwipeGesture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozSwipeGesture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozSwipeGesture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// MozTapGesture Documentation is as below: "Two touch points are tapped on the touch surface."
// https://developer.mozilla.org/docs/Web/Reference/Events/MozTapGesture
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func MozTapGesture(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("MozTapGesture",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowseractivitydone Documentation is as below: "Sent when some activity has been completed (complete description TBD.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowseractivitydone
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowseractivitydone(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowseractivitydone",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserasyncscroll Documentation is as below: "Sent when the scroll position within a browser <iframe> changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserasyncscroll
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserasyncscroll(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserasyncscroll",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowseraudioplaybackchange Documentation is as below: "Sent when audio starts or stops playing within the browser <iframe> content."
// https://developer.mozilla.org/docs/Web/Events/mozbrowseraudioplaybackchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowseraudioplaybackchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowseraudioplaybackchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsercaretstatechanged Documentation is as below: "Sent when the text selected inside the browser <iframe> content changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsercaretstatechanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsercaretstatechanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsercaretstatechanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserclose Documentation is as below: "Sent when window.close() is called within a browser <iframe>."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserclose
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserclose(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserclose",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsercontextmenu Documentation is as below: "Sent when a browser <iframe> try to open a context menu."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsercontextmenu
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsercontextmenu(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsercontextmenu",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserdocumentfirstpaint Documentation is as below: "Sent when a new paint occurs on any document in the browser <iframe>."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserdocumentfirstpaint
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserdocumentfirstpaint(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserdocumentfirstpaint",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsererror Documentation is as below: "Sent when an error occured while trying to load a content within a browser iframe"
// https://developer.mozilla.org/docs/Web/Events/mozbrowsererror
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsererror(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsererror",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserfindchange Documentation is as below: "Sent when a search operation is performed on the browser <iframe> content (see HTMLIFrameElement search methods.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowserfindchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserfindchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserfindchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserfirstpaint Documentation is as below: "Sent when the <iframe> paints content for the first time (this doesn't include the initial paint from about:blank.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowserfirstpaint
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserfirstpaint(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserfirstpaint",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsericonchange Documentation is as below: "Sent when the favicon of a browser iframe changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsericonchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsericonchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsericonchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserloadend Documentation is as below: "Sent when the browser iframe has finished loading all its assets."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserloadend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserloadend(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserloadend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserloadstart Documentation is as below: "Sent when the browser iframe starts to load a new page."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserloadstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserloadstart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserloadstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserlocationchange Documentation is as below: "Sent when an browser iframe's location changes."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserlocationchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserlocationchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserlocationchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsermanifestchange Documentation is as below: "Sent when a the path to the app manifest changes, in the case of a browser <iframe> with an open web app embedded in it."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsermanifestchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsermanifestchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsermanifestchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsermetachange Documentation is as below: "Sent when a <meta> elelment is added to, removed from or changed in the browser <iframe>'s content."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsermetachange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsermetachange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsermetachange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowseropensearch Documentation is as below: "Sent when a link to a search engine is found."
// https://developer.mozilla.org/docs/Web/Events/mozbrowseropensearch
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowseropensearch(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowseropensearch",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowseropentab Documentation is as below: "Sent when a new tab is opened within a browser <iframe> as a result of the user issuing a command to open a link target in a new tab (for example ctrl/cmd + click.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowseropentab
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowseropentab(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowseropentab",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowseropenwindow Documentation is as below: "Sent when window.open() is called within a browser iframe."
// https://developer.mozilla.org/docs/Web/Events/mozbrowseropenwindow
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowseropenwindow(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowseropenwindow",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserresize Documentation is as below: "Sent when the browser <iframe>'s window size has changed."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserresize
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserresize(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserresize",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserscroll Documentation is as below: "Sent when the browser <iframe> content scrolls."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserscroll
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserscroll(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserscroll",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserscrollareachanged Documentation is as below: "Sent when the available scrolling area\u00a0 in the browser <iframe> changes. This can occur on resize and when the page size changes (while loading for example.)"
// https://developer.mozilla.org/docs/Web/Events/mozbrowserscrollareachanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserscrollareachanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserscrollareachanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserscrollviewchange Documentation is as below: "Sent when asynchronous scrolling (i.e. APCZ) starts or stops."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserscrollviewchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserscrollviewchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserscrollviewchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsersecuritychange Documentation is as below: "Sent when the SSL state changes within a browser iframe."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsersecuritychange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsersecuritychange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsersecuritychange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserselectionstatechanged Documentation is as below: "Sent when the text selected inside the browser <iframe> content changes. Note that this is deprecated, and newer implementations use mozbrowsercaretstatechanged instead."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserselectionstatechanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserselectionstatechanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserselectionstatechanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsershowmodalprompt Documentation is as below: "Sent when alert(), confirm() or prompt() are called within a browser iframe"
// https://developer.mozilla.org/docs/Web/Events/mozbrowsershowmodalprompt
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsershowmodalprompt(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsershowmodalprompt",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowsertitlechange Documentation is as below: "Sent when the document.title changes within a browser iframe."
// https://developer.mozilla.org/docs/Web/Events/mozbrowsertitlechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowsertitlechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowsertitlechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowserusernameandpasswordrequired Documentation is as below: "Sent when an HTTP authentification is requested."
// https://developer.mozilla.org/docs/Web/Events/mozbrowserusernameandpasswordrequired
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowserusernameandpasswordrequired(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowserusernameandpasswordrequired",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Mozbrowservisibilitychange Documentation is as below: "Sent when the visibility state of the current browser iframe <iframe> changes, for example due to a call to setVisible()."
// https://developer.mozilla.org/docs/Web/Events/mozbrowservisibilitychange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Mozbrowservisibilitychange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("mozbrowservisibilitychange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Moztimechange Documentation is as below: "The time of the device has been changed."
// https://developer.mozilla.org/docs/Web/Events/moztimechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Moztimechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("moztimechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// NoUpdate Documentation is as below: "The manifest hadn't changed."
// https://developer.mozilla.org/docs/Web/Events/noupdate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func NoUpdate(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("noupdate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Nomatch Documentation is as below: "The speech recognition service returns a final result with no significant recognition."
// https://developer.mozilla.org/docs/Web/Events/nomatch
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Nomatch(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("nomatch",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Notificationclick Documentation is as below: "A system notification spawned by ServiceWorkerRegistration.showNotification() has been clicked."
// https://developer.mozilla.org/docs/Web/Events/notificationclick
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Notificationclick(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("notificationclick",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Obsolete Documentation is as below: "The manifest was found to have become a 404 or 410 page, so the application cache is being deleted."
// https://developer.mozilla.org/docs/Web/Events/obsolete
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Obsolete(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("obsolete",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Offline Documentation is as below: "The browser has lost access to the network."
// https://developer.mozilla.org/docs/Web/Events/offline
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Offline(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("offline",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Onconnected Documentation is as below: "A call has been connected."
// https://developer.mozilla.org/docs/DOM/onconnected
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Onconnected(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("onconnected",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Online Documentation is as below: "The browser has gained access to the network (but particular websites might be unreachable)."
// https://developer.mozilla.org/docs/Web/Events/online
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Online(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("online",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Open Documentation is as below: "An event source connection has been established."
// https://developer.mozilla.org/docs/Web/Reference/Events/open_serversentevents
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Open(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("open",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// OrientationChange Documentation is as below: "The orientation of the device (portrait/landscape) has changed"
// https://developer.mozilla.org/docs/Web/Events/orientationchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func OrientationChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("orientationchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Overflow Documentation is as below: "An element has been overflowed by its content or has been rendered for the first time in this state (only works for elements styled with overflow != visible)."
// https://developer.mozilla.org/docs/Web/Events/overflow
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Overflow(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("overflow",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// PageHide Documentation is as below: "A session history entry is being traversed from."
// https://developer.mozilla.org/docs/Web/Events/pagehide
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func PageHide(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pagehide",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// PageShow Documentation is as below: "A session history entry is being traversed to."
// https://developer.mozilla.org/docs/Web/Events/pageshow
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func PageShow(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pageshow",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Paste Documentation is as below: "Data has been transferred from the system clipboard to the document."
// https://developer.mozilla.org/docs/Web/Events/paste
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Paste(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("paste",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pause Documentation is as below: "The utterance is paused part way through."
// https://developer.mozilla.org/docs/Web/Events/pause_(SpeechSynthesis)
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pause(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pause",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Play Documentation is as below: "Playback has begun."
// https://developer.mozilla.org/docs/Web/Events/play
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Play(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("play",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Playing Documentation is as below: "Playback is ready to start after having been paused or delayed due to lack of data."
// https://developer.mozilla.org/docs/Web/Events/playing
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Playing(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("playing",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// PointerLockChange Documentation is as below: "The pointer was locked or released."
// https://developer.mozilla.org/docs/Web/Events/pointerlockchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func PointerLockChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerlockchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// PointerLockError Documentation is as below: "It was impossible to lock the pointer for technical reasons or because the permission was denied."
// https://developer.mozilla.org/docs/Web/Events/pointerlockerror
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func PointerLockError(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerlockerror",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointercancel Documentation is as below: "The pointer is unlikely to produce any more events."
// https://developer.mozilla.org/docs/Web/Events/pointercancel
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointercancel(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointercancel",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointerdown Documentation is as below: "The pointer enters the active buttons state."
// https://developer.mozilla.org/docs/Web/Events/pointerdown
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointerdown(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerdown",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointerenter Documentation is as below: "Pointing device is moved inside the hit-testing boundary."
// https://developer.mozilla.org/docs/Web/Events/pointerenter
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointerenter(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerenter",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointerleave Documentation is as below: "Pointing device is moved out of the hit-testing boundary."
// https://developer.mozilla.org/docs/Web/Events/pointerleave
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointerleave(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerleave",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointermove Documentation is as below: "The pointer changed coordinates."
// https://developer.mozilla.org/docs/Web/Events/pointermove
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointermove(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointermove",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointerout Documentation is as below: "The pointing device moved out of hit-testing boundary or leaves detectable hover range."
// https://developer.mozilla.org/docs/Web/Events/pointerout
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointerout(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerout",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointerover Documentation is as below: "The pointing device is moved into the hit-testing boundary."
// https://developer.mozilla.org/docs/Web/Events/pointerover
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointerover(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerover",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pointerup Documentation is as below: "The pointer leaves the active buttons state."
// https://developer.mozilla.org/docs/Web/Events/pointerup
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pointerup(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pointerup",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// PopState Documentation is as below: "A session history entry is being navigated to (in certain cases)."
// https://developer.mozilla.org/docs/Web/Events/popstate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func PopState(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("popstate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Popuphidden Documentation is as below: "A menupopup, panel or tooltip has been hidden."
// https://developer.mozilla.org/docs/Web/Events/popuphidden
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Popuphidden(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("popuphidden",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Popuphiding Documentation is as below: "A menupopup, panel or tooltip is about to be hidden."
// https://developer.mozilla.org/docs/Web/Events/popuphiding
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Popuphiding(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("popuphiding",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Popupshowing Documentation is as below: "A menupopup, panel or tooltip is about to become visible."
// https://developer.mozilla.org/docs/Web/Events/popupshowing
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Popupshowing(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("popupshowing",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Popupshown Documentation is as below: "A menupopup, panel or tooltip has become visible."
// https://developer.mozilla.org/docs/Web/Events/popupshown
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Popupshown(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("popupshown",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Progress Documentation is as below: "The user agent is downloading resources listed by the manifest."
// https://developer.mozilla.org/docs/Web/Reference/Events/progress_(appcache_event)
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Progress(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("progress",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Push Documentation is as below: "A Service Worker has received a push message."
// https://developer.mozilla.org/docs/Web/Events/push
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Push(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("push",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Pushsubscriptionchange Documentation is as below: "A PushSubscription has expired."
// https://developer.mozilla.org/docs/Web/Events/pushsubscriptionchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Pushsubscriptionchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("pushsubscriptionchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// RadioStateChange Documentation is as below: "The state of a radio has been changed either by a user action or by a script (useful for accessibility)."
// https://developer.mozilla.org/docs/Web/Events/RadioStateChange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func RadioStateChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("RadioStateChange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// RateChange Documentation is as below: "The playback rate has changed."
// https://developer.mozilla.org/docs/Web/Events/ratechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func RateChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("ratechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// ReadystateChange Documentation is as below: "The readyState attribute of a document has changed."
// https://developer.mozilla.org/docs/Web/Events/readystatechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func ReadystateChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("readystatechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Received Documentation is as below: "An SMS has been received."
// https://developer.mozilla.org/docs/Web/Events/received
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Received(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("received",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// RepeatEvent Documentation is as below: "A SMIL animation element is repeated."
// https://developer.mozilla.org/docs/Web/Events/repeatEvent
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func RepeatEvent(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("repeatEvent",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Reset Documentation is as below: "A form is reset."
// https://developer.mozilla.org/docs/Web/Events/reset
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Reset(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("reset",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Resize Documentation is as below: "The document view has been resized."
// https://developer.mozilla.org/docs/Web/Events/resize
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Resize(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("resize",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Resourcetimingbufferfull Documentation is as below: "The browser's resource timing buffer is full."
// https://developer.mozilla.org/docs/Web/Events/resourcetimingbufferfull
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Resourcetimingbufferfull(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("resourcetimingbufferfull",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Result Documentation is as below: "The speech recognition service returns a result — a word or phrase has been positively recognized and this has been communicated back to the app."
// https://developer.mozilla.org/docs/Web/Events/result
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Result(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("result",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Resume Documentation is as below: "A paused utterance is resumed."
// https://developer.mozilla.org/docs/Web/Events/resume
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Resume(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("resume",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Resuming Documentation is as below: "A call is about to resume."
// https://developer.mozilla.org/docs/Web/Events/resuming
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Resuming(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("resuming",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SSTabClosing Documentation is as below: "The session store will stop tracking this tab."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSTabClosing
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SSTabClosing(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SSTabClosing",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SSTabRestored Documentation is as below: "A tab has been restored."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSTabRestored
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SSTabRestored(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SSTabRestored",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SSTabRestoring Documentation is as below: "A tab is about to be restored."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSTabRestoring
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SSTabRestoring(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SSTabRestoring",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SSWindowClosing Documentation is as below: "The session store will stop tracking this window."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSWindowClosing
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SSWindowClosing(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SSWindowClosing",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SSWindowStateBusy Documentation is as below: "A window state has switched to \"busy\"."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSWindowStateBusy
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SSWindowStateBusy(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SSWindowStateBusy",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SSWindowStateReady Documentation is as below: "A window state has switched to \"ready\"."
// https://developer.mozilla.org/docs/Web/Reference/Events/SSWindowStateReady
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SSWindowStateReady(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SSWindowStateReady",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SVGAbort Documentation is as below: "Page loading has been stopped before the SVG was loaded."
// https://developer.mozilla.org/docs/Web/Events/SVGAbort
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SVGAbort(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SVGAbort",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SVGError Documentation is as below: "An error has occurred before the SVG was loaded."
// https://developer.mozilla.org/docs/Web/Events/SVGError
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SVGError(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SVGError",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SVGLoad Documentation is as below: "An SVG document has been loaded and parsed."
// https://developer.mozilla.org/docs/Web/Events/SVGLoad
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SVGLoad(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SVGLoad",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SVGResize Documentation is as below: "An SVG document is being resized."
// https://developer.mozilla.org/docs/Web/Events/SVGResize
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SVGResize(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SVGResize",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SVGScroll Documentation is as below: "An SVG document is being scrolled."
// https://developer.mozilla.org/docs/Web/Events/SVGScroll
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SVGScroll(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SVGScroll",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SVGUnload Documentation is as below: "An SVG document has been removed from a window or frame."
// https://developer.mozilla.org/docs/Web/Events/SVGUnload
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SVGUnload(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SVGUnload",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SVGZoom Documentation is as below: "An SVG document is being zoomed."
// https://developer.mozilla.org/docs/Web/Events/SVGZoom
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SVGZoom(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("SVGZoom",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Scroll Documentation is as below: "The document view or an element has been scrolled."
// https://developer.mozilla.org/docs/Web/Events/scroll
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Scroll(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("scroll",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Seeked Documentation is as below: "A seek operation completed."
// https://developer.mozilla.org/docs/Web/Events/seeked
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Seeked(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("seeked",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Seeking Documentation is as below: "A seek operation began."
// https://developer.mozilla.org/docs/Web/Events/seeking
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Seeking(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("seeking",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Select Documentation is as below: "Some text is being selected."
// https://developer.mozilla.org/docs/Web/Events/select
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Select(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("select",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Selectionchange Documentation is as below: "The selection in the document has been changed."
// https://developer.mozilla.org/docs/Web/Events/selectionchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Selectionchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("selectionchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Selectstart Documentation is as below: "A selection just started."
// https://developer.mozilla.org/docs/Web/Events/selectstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Selectstart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("selectstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Sent Documentation is as below: "An SMS has been sent."
// https://developer.mozilla.org/docs/Web/Events/sent
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Sent(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("sent",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Show Documentation is as below: "A contextmenu event was fired on/bubbled to an element that has a contextmenu attribute"
// https://developer.mozilla.org/docs/Web/Events/show
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Show(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("show",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Sizemodechange Documentation is as below: "Window has entered/left fullscreen mode, or has been minimized/unminimized."
// https://developer.mozilla.org/docs/Web/Reference/Events/sizemodechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Sizemodechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("sizemodechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SmartCardInsert Documentation is as below: "A smartcard has been inserted."
// https://developer.mozilla.org/docs/Web/Events/smartcard-insert
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SmartCardInsert(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("smartcard-insert",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// SmartCardRemove Documentation is as below: "A smartcard has been removed."
// https://developer.mozilla.org/docs/Web/Events/smartcard-remove
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func SmartCardRemove(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("smartcard-remove",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Soundend Documentation is as below: "Any sound — recognisable speech or not — has stopped being detected."
// https://developer.mozilla.org/docs/Web/Events/soundend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Soundend(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("soundend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Soundstart Documentation is as below: "Any sound — recognisable speech or not — has been detected."
// https://developer.mozilla.org/docs/Web/Events/soundstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Soundstart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("soundstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Speechend Documentation is as below: "Speech recognised by the speech recognition service has stopped being detected."
// https://developer.mozilla.org/docs/Web/Events/speechend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Speechend(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("speechend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Speechstart Documentation is as below: "Sound that is recognised by the speech recognition service as speech has been detected."
// https://developer.mozilla.org/docs/Web/Events/speechstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Speechstart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("speechstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Stalled Documentation is as below: "The user agent is trying to fetch media data, but data is unexpectedly not forthcoming."
// https://developer.mozilla.org/docs/Web/Events/stalled
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Stalled(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("stalled",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Start Documentation is as below: "The utterance has begun to be spoken."
// https://developer.mozilla.org/docs/Web/Events/start_(SpeechSynthesis)
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Start(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("start",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Statechange Documentation is as below: "The state of a call has changed."
// https://developer.mozilla.org/docs/Web/Events/statechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Statechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("statechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Statuschange Documentation is as below: "The status of the Wifi connection changed."
// https://developer.mozilla.org/docs/Web/Events/statuschange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Statuschange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("statuschange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Stkcommand Documentation is as below: "The STK Proactive Command is issued from ICC."
// https://developer.mozilla.org/docs/Web/Events/stkcommand
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Stkcommand(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("stkcommand",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Stksessionend Documentation is as below: "The STK Session is terminated by ICC."
// https://developer.mozilla.org/docs/Web/Events/stksessionend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Stksessionend(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("stksessionend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Storage Documentation is as below: "A storage area (localStorage or sessionStorage) has changed."
// https://developer.mozilla.org/docs/Web/Events/storage
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Storage(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("storage",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Submit Documentation is as below: "A form is submitted."
// https://developer.mozilla.org/docs/Web/Events/submit
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Submit(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("submit",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Success Documentation is as below: "A request successfully completed."
// https://developer.mozilla.org/docs/Web/Reference/Events/success_indexedDB
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Success(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("success",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Suspend Documentation is as below: "Media data loading has been suspended."
// https://developer.mozilla.org/docs/Web/Events/suspend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Suspend(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("suspend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TabClose Documentation is as below: "A tab has been closed."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabClose
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TabClose(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("TabClose",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TabHide Documentation is as below: "A tab has been hidden."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabHide
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TabHide(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("TabHide",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TabOpen Documentation is as below: "A tab has been opened."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabOpen
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TabOpen(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("TabOpen",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TabPinned Documentation is as below: "A tab has been pinned."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabPinned
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TabPinned(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("TabPinned",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TabSelect Documentation is as below: "A tab has been selected."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabSelect
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TabSelect(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("TabSelect",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TabShow Documentation is as below: "A tab has been shown."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabShow
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TabShow(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("TabShow",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TabUnpinned Documentation is as below: "A tab has been unpinned."
// https://developer.mozilla.org/docs/Web/Reference/Events/TabUnpinned
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TabUnpinned(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("TabUnpinned",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TimeUpdate Documentation is as below: "The time indicated by the currentTime attribute has been updated."
// https://developer.mozilla.org/docs/Web/Events/timeupdate
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TimeUpdate(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("timeupdate",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Timeout Documentation is as below: "(no documentation)"
// https://developer.mozilla.org/docs/Web/Events/timeout
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Timeout(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("timeout",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TouchCancel Documentation is as below: "A touch point has been disrupted in an implementation-specific manners (too many touch points for example)."
// https://developer.mozilla.org/docs/Web/Events/touchcancel
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TouchCancel(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("touchcancel",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TouchEnd Documentation is as below: "A touch point is removed from the touch surface."
// https://developer.mozilla.org/docs/Web/Events/touchend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TouchEnd(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("touchend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TouchEnter Documentation is as below: "(no documentation)"
// https://developer.mozilla.org/docs/Web/Events/touchenter
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TouchEnter(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("touchenter",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TouchLeave Documentation is as below: "(no documentation)"
// https://developer.mozilla.org/docs/Web/Events/touchleave
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TouchLeave(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("touchleave",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TouchMove Documentation is as below: "A touch point is moved along the touch surface."
// https://developer.mozilla.org/docs/Web/Events/touchmove
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TouchMove(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("touchmove",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TouchStart Documentation is as below: "A touch point is placed on the touch surface."
// https://developer.mozilla.org/docs/Web/Events/touchstart
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TouchStart(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("touchstart",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// TransitionEnd Documentation is as below: "A CSS transition has completed."
// https://developer.mozilla.org/docs/Web/Events/transitionend
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func TransitionEnd(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("transitionend",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Underflow Documentation is as below: "An element is no longer overflowed by its content (only works for elements styled with overflow != visible)."
// https://developer.mozilla.org/docs/Web/Events/underflow
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Underflow(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("underflow",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Unload Documentation is as below: "The document or a dependent resource is being unloaded."
// https://developer.mozilla.org/docs/Web/Events/unload
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Unload(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("unload",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// UpdateReady Documentation is as below: "The resources listed in the manifest have been newly redownloaded, and the script can use swapCache() to switch to the new cache."
// https://developer.mozilla.org/docs/Web/Events/updateready
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func UpdateReady(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("updateready",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// UpgradeNeeded Documentation is as below: "An attempt was made to open a database with a version number higher than its current version. A versionchange transaction has been created."
// https://developer.mozilla.org/docs/Web/Reference/Events/upgradeneeded_indexedDB
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func UpgradeNeeded(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("upgradeneeded",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// UserProximity Documentation is as below: "Fresh data is available from a proximity sensor (indicates whether the nearby object is near the device or not)."
// https://developer.mozilla.org/docs/Web/Events/userproximity
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func UserProximity(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("userproximity",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Ussdreceived Documentation is as below: "A new USSD message is received"
// https://developer.mozilla.org/docs/Web/Events/ussdreceived
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Ussdreceived(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("ussdreceived",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// ValueChange Documentation is as below: "The value of an element has changed (a progress bar for example, useful for accessibility)."
// https://developer.mozilla.org/docs/Web/Events/ValueChange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func ValueChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("ValueChange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// VersionChange Documentation is as below: "A versionchange transaction completed."
// https://developer.mozilla.org/docs/Web/Reference/Events/versionchange_indexedDB
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func VersionChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("versionchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// VisibilityChange Documentation is as below: "The content of a tab has become visible or has been hidden."
// https://developer.mozilla.org/docs/Web/Events/visibilitychange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func VisibilityChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("visibilitychange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Voicechange Documentation is as below: "The MozMobileConnection.voice object changes values."
// https://developer.mozilla.org/docs/Web/Events/voicechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Voicechange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("voicechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Voiceschanged Documentation is as below: "The list of SpeechSynthesisVoice objects that would be returned by the SpeechSynthesis.getVoices() method has changed (when the voiceschanged event fires.)"
// https://developer.mozilla.org/docs/Web/Events/voiceschanged
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Voiceschanged(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("voiceschanged",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// VolumeChange Documentation is as below: "The volume has changed."
// https://developer.mozilla.org/docs/Web/Events/volumechange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func VolumeChange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("volumechange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Vrdisplayconnected Documentation is as below: "A compatible VR device has been connected to the computer."
// https://developer.mozilla.org/docs/Web/Events/vrdisplayconnected
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Vrdisplayconnected(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("vrdisplayconnected",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Vrdisplaydisconnected Documentation is as below: "A compatible VR device has been disconnected from the computer."
// https://developer.mozilla.org/docs/Web/Events/vrdisplaydisconnected
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Vrdisplaydisconnected(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("vrdisplaydisconnected",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Vrdisplaypresentchange Documentation is as below: "The presenting state of a VR device has changed — i.e. from presenting to not presenting, or vice versa."
// https://developer.mozilla.org/docs/Web/Events/vrdisplaypresentchange
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Vrdisplaypresentchange(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("vrdisplaypresentchange",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Waiting Documentation is as below: "Playback has stopped because of a temporary lack of data."
// https://developer.mozilla.org/docs/Web/Events/waiting
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Waiting(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("waiting",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
 

// Wheel Documentation is as below: "A wheel button of a pointing device is rotated in any direction."
// https://developer.mozilla.org/docs/Web/Events/wheel
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Wheel(callback EventHandler, sel string) *trees.Event {
	ev := trees.NewEvent("wheel",sel)
	ev.Handle = dispatch.Subscribe(func(evm trees.EventBroadcast){
		if ev.EventID != evm.EventID{
			return
		}

		callback(evm.Event, ev.Tree)
	})

	return ev
}
