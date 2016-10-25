// The generation of this package was inspired by Neelance work on DOM (https://github.com/neelance/dom)

//go:generate go run generate.go

// Documentation source: "Event reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/Events, licensed under CC-BY-SA 2.5.

//Package events defines the event binding system that combines different libraries to create a interesting event system.
package events

import (
	"github.com/influx6/gu/trees"
)

// Abort Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/abort
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Abort(fx trees.EventHandler,selectorOverride string) trees.Event {
	return trees.NewEvent("abort",selectorOverride,fx)
}

// BeforeUnload Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/beforeunload
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func BeforeUnload(fx trees.EventHandler,selectorOverride string) trees.Event {
	return trees.NewEvent("beforeunload",selectorOverride,fx)
}

// Cached Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/cached
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Cached(fx trees.EventHandler,selectorOverride string) trees.Event {
	return trees.NewEvent("cached",selectorOverride,fx)
}

// Error Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/error
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Error(fx trees.EventHandler,selectorOverride string) trees.Event {
	return trees.NewEvent("error",selectorOverride,fx)
}

// Load Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/load
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Load(fx trees.EventHandler,selectorOverride string) trees.Event {
	return trees.NewEvent("load",selectorOverride,fx)
}

// Unload Documentation is as below:
// (no documentation)
// https://developer.mozilla.org/docs/Web/Events/unload
/* This event provides options() to be called when the events is triggered and an optional selector which will override the internal selector mechanism of the domtrees.Element i.e if the selectorOverride argument is an empty string then domtrees.Element will create an appropriate selector matching its type and uid value in this format  (ElementType[uid='UID_VALUE']) but if the selector value is not empty then that becomes the default selector used
match the event with. */
func Unload(fx trees.EventHandler,selectorOverride string) trees.Event {
	return trees.NewEvent("unload",selectorOverride,fx)
}
