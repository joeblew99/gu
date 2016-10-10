// Package gu provides a UI framework for Go.
package gu

import (
	"fmt"
	"html/template"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/gutrees"
)

var (
	// baseCount defines the current total count delivered up for use and once
	// this is increasing then will keep increasing.
	baseCount int

	// baseKey provides a compliant version 4 uuid which we can use to create
	// incremental uuid counts for renderings.
	baseKey = "3bce4931-6c75-41ab-afe0-2ec108a30860"
)

func newKey() string {
	baseCount++
	return fmt.Sprintf("%d-%s", baseCount, baseKey)
}

//==============================================================================

// Renderable provides a interface for a renderable type.
type Renderable interface {
	Render() gutrees.Markup
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

// Viewable defines a generic interface for a generic return type. It exists to
// give symantic representation in the areas it is used to express the expected
// returned to be one of a Viewable souce in the context of the gu library.
type Viewable interface{}

// RenderView defines an interface through which you gain access into a rendering
// branch of the current rendered markup view.
type RenderView interface {
	MarkupRenderer
	UUID() string
	Events() guevents.EventManagers()
}


// RenderingTarget defines an interface which takes responsiblity in translating
// the provided markup into the appropriate media.
type RenderingTarget interface {
	HandleView(RenderView)
}

// RenderGroup provides a central backbone through which RenderViews are rendered
// as one.
type RenderGroup struct {
	views   []RenderView
	baseTag string
	uuid    string
}

// New returns a new RenderGroup which allows grouping RenderView into a set of
// one rendering.
func New() *RenderGroup {
	return &RenderGroup{
		baseTag: "div",
		uuid:    newKey(),
	}
}

// View adds the giving renderables into the provided a view managed by 
// the RenderGroup.
func (rg *RenderGroup) View(r ...Renderable) {
	rg.views = append(rg.views, customView("div",guevents.NewEventManager(), r...))
}

// CustomView does a similar operation as the .View method but allows the user 
// to specify the tag used to wrap more than one Renderable.
func (rg *RenderGroup) CustomView(tag string, r ...Renderable) {
	rg.views = append(rg.views, customView(tag,guevents.NewEventManager(), r...))
}

// UseRenderingTarget requests the render group initialize all internal views 
// into the provided render target.
func (rg *RenderGroup) UseRenderingTarget(target RenderingTarget) {
	for _, view := range rg.views {
		target.HandleView(view)
	}
}

// UUID returns the RenderGroup UUID for identification.
func (rg *RenderGroup) UUID() string {
	return rg.uuid
}

// SetBaseTag sets the name of the base tag which is generated to wrap the
// contents of the children RenderView.
func (rg *RenderGroup) SetBaseTag(tag string) {
	rg.baseTag = tag
}

// RenderHTML returns the markup converted into a compliant html markup.
func (rg *RenderGroup) RenderHTML() template.HTML {
	return rg.Render().EHTML()
}

// Render returns the groups markup for the giving render group.
func (rg *RenderGroup) Render() gutrees.Markup {
	root := gutrees.NewElement(rg.baseTag, false)

	for _, view := range rg.views {
		view.Render().Apply(root)
	}

	return root
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
func AttachURL(pattern string, activeFn, inactiveFn func(gudispatch.Path)) {
	gudispatch.ResolveAttachURL(pattern, activeFn, inactiveFn)
}

// AttachHash attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the URL hash.
// failFn must either be a FailNormal, FailPath or nil.
func AttachHash(pattern string, activeFn, inactiveFn func(gudispatch.Path)) {
	gudispatch.ResolveAttachHash(pattern, activeFn, inactiveFn)
}

//==============================================================================
