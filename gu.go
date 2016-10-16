// Package gu provides a UI framework for Go.
package gu

import (
	"fmt"
	"html/template"
	"sync"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/guevents"
	"github.com/influx6/gu/gutrees"
)

var countKeeper = struct {
	baseKey   string
	ml        sync.Mutex
	baseCount int
}{
	baseKey: "3bce4931-6c75-41ab-afe0-2ec108a30860",
}

func newKey() string {
	countKeeper.ml.Lock()
	countKeeper.baseCount++
	countKeeper.ml.Unlock()

	return fmt.Sprintf("%d-%s", countKeeper.baseCount, countKeeper.baseKey)
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
	gudispatch.Resolvable

	UUID() string
}

// Eventable exposes the events manager provided by the structure that implements
// the interface.
type Eventable interface {
	Events() guevents.EventManagers
}

// EventableRenderView defines a composite of a RenderView which provides access to its Events manager.
type EventableRenderView interface {
	LoadEvents(*js.Object)
}

// RenderingGroup defines an interface that exposes it's set of children RenderViews.
type RenderingGroup interface {
	Views() []RenderView
}

// Renderer defines an interface which takes responsiblity in translating
// the provided markup into the appropriate media.
type Renderer interface {
	RenderView(RenderView)
}

// RenderingTargetGroup defines an interface which takes the target to be handle
// for rendering.
type RenderingTarget interface {
	UseRendering(Renderer)
}

// RenderGroup provides a central backbone through which RenderViews are rendered as one.
// RenderGroup implements the RenderView, RenderingGroup, Eventable, EventableRenderView and
// Rendering interfaces.
type RenderGroup struct {
	views      []RenderView
	baseTag    string
	uuid       string
	baseEvents guevents.EventManagers
}

// New returns a new RenderGroup which allows grouping RenderView into a set of
// one rendering.
func New() *RenderGroup {
	return &RenderGroup{
		baseTag:    "div",
		uuid:       newKey(),
		baseEvents: guevents.NewEventManager(),
	}
}

// Views returns the giving underline children views, it implements the RenderingGroup interface.
func (rg *RenderGroup) Views() []RenderView {
	return rg.views
}


// AddRenderView adds the giving RenderView set into the RenderGroups views
// list.
func (rg *RenderGroup) AddRenderView(rs ...RenderView) {
	rg.views = append(rg.views, rs...)

	for _, view := range rs {
		if ev, ok := view.(Eventable); ok {
			rg.baseEvents.AttachManager(ev.Events())
		}
	}
}

// View adds the giving renderables into the provided a view managed by
// the RenderGroup.
func (rg *RenderGroup) View(r ...Renderable) RenderView{
	esm := guevents.NewEventManager()
	rg.baseEvents.AttachManager(esm)

	csv := customView("section", esm, r...)
	rg.views = append(rg.views, csv)
	return csv
}

// CustomView does a similar operation as the .View method but allows the user
// to specify the tag used to wrap more than one Renderable.
func (rg *RenderGroup) CustomView(tag string, r ...Renderable) RenderView {
	esm := guevents.NewEventManager()
	rg.baseEvents.AttachManager(esm)

	csv := customView(tag, esm, r...)
	rg.views = append(rg.views, csv)
	return csv
}

// UseRendering requests the render group initialize all internal views
// into the provided render target.
func (rg *RenderGroup) UseRendering(render Renderer) {
	render.RenderView(rg)
}

// LoadEvents loads the giving RenderGroup with the provided object.
func (rg *RenderGroup) LoadEvents(target *js.Object) {
	rg.baseEvents.OffloadDOM()
	rg.baseEvents.LoadDOM(target)
}

// Resolve resolves the internal RenderViews attached to this RenderGroup with
// the provided path.
func (rg *RenderGroup) Resolve(path gudispatch.Path) {
	for _, vmr := range rg.views {
		if rs, ok := vmr.(gudispatch.Resolvable); ok {
			rs.Resolve(path)
		}
	}
}

// Events returns the provided Event manages the event manager connected to the
// group.
func (rg *RenderGroup) Events() guevents.EventManagers {
	return rg.Events()
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
