// The generation of this package was inspired by Neelance work on DOM (https://github.com/neelance/dom)

//go:generate go run generate.go

// Documentation source: "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.

//Package events defines the event binding system that combines different libraries to create a interesting event system.
package events

import (
	"github.com/influx6/gu/trees"
)

// AbortEvent defines the expected type to be received when using the Abort event.
type AbortEvent struct {
	Underline *js.Object
}

// Abort Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/abort
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Abort(callback func(AbortEvent, trees.Markup), selectorOverride string) trees.Event {
	return trees.NewEvent("abort", selectorOverride, fx)
}

// BeforeUnloadEvent defines the expected type to be received when using the BeforeUnload event.
type BeforeUnloadEvent struct {
	Underline *js.Object
}

// BeforeUnload Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/beforeunload
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func BeforeUnload(callback func(BeforeUnloadEvent, trees.Markup), selectorOverride string) trees.Event {
	return trees.NewEvent("beforeunload", selectorOverride, fx)
}

// CachedEvent defines the expected type to be received when using the Cached event.
type CachedEvent struct {
	Underline *js.Object
}

// Cached Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/cached
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Cached(callback func(CachedEvent, trees.Markup), selectorOverride string) trees.Event {
	return trees.NewEvent("cached", selectorOverride, fx)
}

// ErrorEvent defines the expected type to be received when using the Error event.
type ErrorEvent struct {
	Underline *js.Object
}

// Error Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/error
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Error(callback func(ErrorEvent, trees.Markup), selectorOverride string) trees.Event {
	return trees.NewEvent("error", selectorOverride, fx)
}

// LoadEvent defines the expected type to be received when using the Load event.
type LoadEvent struct {
	Underline *js.Object
}

// Load Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/load
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Load(callback func(LoadEvent, trees.Markup), selectorOverride string) trees.Event {
	return trees.NewEvent("load", selectorOverride, fx)
}

// UnloadEvent defines the expected type to be received when using the Unload event.
type UnloadEvent struct {
	Underline *js.Object
}

// Unload Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/unload
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Unload(callback func(UnloadEvent, trees.Markup), selectorOverride string) trees.Event {
	return trees.NewEvent("unload", selectorOverride, fx)
}
