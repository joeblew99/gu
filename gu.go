// Package gu provides a UI framework for Go.
package gu

import (
	"fmt"
	"html/template"
	"sync"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/trees"
)

// countKeeper handles management of the keys being generating. Guards the incrementation
// using a mutex.
var countKeeper = struct {
	ml        sync.Mutex
	baseKey   string
	baseCount int
}{
	baseKey: "3bce4931-6c75-41ab-afe0-2ec108a30860",
}

// NewKey returns a new string key which is path of the incremental key which once initializes
// constantly increases.
func NewKey() string {
	countKeeper.ml.Lock()
	countKeeper.baseCount++
	countKeeper.ml.Unlock()

	return fmt.Sprintf("%d-%s", countKeeper.baseCount, countKeeper.baseKey)
}

//==============================================================================

// Renderable provides a interface for a renderable type.
type Renderable interface {
	Render() *trees.Markup
}

// Renderables defines a lists of Renderable structures.
type Renderables []Renderable

// MarkupRenderer provides a interface for a types capable of rendering dom markup.
type MarkupRenderer interface {
	Renderable
	RenderHTML() template.HTML
}

// ReactiveSubscription defines an interface for functions subscribing for
// notifications to react.
type ReactiveSubscription interface {
	React(func())
}

// ReactiveRenderable defines an interface of a Renderable capable of
// notifying subscribe functions of changes to allow them to React.
type ReactiveRenderable interface {
	Renderable
	ReactiveSubscription
}

// Reactive extends the ReactiveRenderable by exposing a Publish method
// which allows calling the update notifications list of a ReactiveRenderable.
type Reactive interface {
	ReactiveSubscription
	Publish()
}

// NewReactive returns an instance of a Reactive struct.
func NewReactive() Reactive {
	var rc reactive
	return &rc
}

// reactive defines a baseline structure that can be composed into
// any struct to provide a reactive view.
type reactive struct {
	subs []func()
}

// React adds a function into the subscription list for this reactor.
func (r *reactive) React(sub func()) {
	r.subs = append(r.subs, sub)
}

// Publish runs a through the subscription list and calls the registerd functions.
func (r *reactive) Publish() {
	for _, sub := range r.subs {
		sub()
	}
}

//==============================================================================

// RenderView defines an interface through which you gain access into a rendering
// branch of the current rendered markup view.
type RenderView interface {
	MarkupRenderer
	dispatch.Resolvable

	UUID() string
}

// Renderer defines an interface which takes responsiblity in translating
// the provided markup into the appropriate media.
type Renderer interface {
	RenderView(RenderView)
}

//==============================================================================

// AddStylesheet adds an external stylesheet to the document into the document
// DOM using the provide URL.
func AddStylesheet(url string) {
	if !detect.IsBrowser() {
		return
	}

	link := js.Global.Get("document").Call("createElement", "link")
	link.Set("rel", "stylesheet")
	link.Set("href", url)
	js.Global.Get("document").Get("head").Call("appendChild", link)
}

// AttachURL attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the full URL(Path+Hash).
// failFn must either be a FailNormal, FailPath or nil.
func AttachURL(pattern string, activeFn, inactiveFn func(dispatch.Path)) {
	dispatch.ResolveAttachURL(pattern, activeFn, inactiveFn)
}

// AttachHash attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the URL hash.
// failFn must either be a FailNormal, FailPath or nil.
func AttachHash(pattern string, activeFn, inactiveFn func(dispatch.Path)) {
	dispatch.ResolveAttachHash(pattern, activeFn, inactiveFn)
}

//==============================================================================
