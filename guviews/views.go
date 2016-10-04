package guviews

import (
	"html/template"
	"sync/atomic"

	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/guevents"
	"github.com/influx6/gu/gujs"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/elems"
)

//==============================================================================

// Behaviour provides a state changers for haiku.
type Behaviour interface {
	Hide()
	Show()
}

// Views define the rendering core for the gu library.
type Views interface {
	Behaviour
	MarkupRenderer

	UUID() string
	UID() string

	Bind(Views)
	Sync(Views)
	BindSync(Views)
	Path(gudispatch.Path)

	Mount(*js.Object)
	Events() guevents.EventManagers
}

// New returns a instance of a Views with the customID set to a random string.
func New(r ...Renderable) Views {
	return CustomView("", r...)
}

// NewWithID returns a View instance. The view is giving a customID string, which
// gets associated with this view, and provides a convenient means of dispatching
// events directly to it, if this is a empty string, a random one will be
// generated for it.
func NewWithID(customID string, r ...Renderable) Views {
	return CustomView(customID, r...)
}

//==============================================================================

// hider defines a global hide renderer.
var hider HideView

// shower defines a global display renderer.
var shower ShowView

// ViewStates defines the two possible behavioral state of a view's markup
type ViewStates interface {
	Render(gutrees.Markup)
}

// HideView provides a ViewStates for Views inactive state
type HideView struct{}

// Render marks the given markup as display:none
func (v HideView) Render(m gutrees.Markup) {
	gutrees.ReplaceORAddStyle(m, "display", "none")
}

// ShowView provides a ViewStates for Views active state
type ShowView struct{}

// Render marks the given markup with a display: block
func (v ShowView) Render(m gutrees.Markup) {
	gutrees.ReplaceORAddStyle(m, "display", "block")
}

//==============================================================================

// view defines a basic struture for building UI view.
type view struct {
	ready        int64
	switchActive int64
	uid          string
	uuid         string
	dom          *js.Object
	renders      []Renderable
	liveMarkup   gutrees.Markup
	events       guevents.EventManagers
	activeState  ViewStates
}

// CustomView returns a gu.Views implementing struct that provides the ability to
// render and update UI efficiently. This function allows greater control of
// the customId for which the views.
func CustomView(cid string, vw ...Renderable) Views {
	if cid == "" {
		cid = gutrees.RandString(8)
	}

	uuid := gutrees.RandString(20)

	vm := &view{
		renders:     vw,
		uid:         cid,
		activeState: shower,
		events:      guevents.NewEventManager(),
		uuid:        uuid,
	}

	for _, vws := range vw {

		// Check if we its a vieww then bind.
		if vms, ok := vws.(Views); ok {
			vm.Bind(vms)
			continue
		}

		// Check if we can react and then subscribe.
		if rws, ok := vws.(ReactiveSubscription); ok {
			rws.React(func() {
				gudispatch.Dispatch(ViewUpdate{ID: uuid})
			})
		}
	}

	// Subscribe for view update requests from the central dispatcher.
	gudispatch.Subscribe(func(v ViewUpdate) {
		if v.ID != vm.UUID() && v.ID != vm.UID() {
			return
		}

		// If we are not domless then patch.
		if vm.dom == nil {
			return
		}

		replaceOnly := atomic.LoadInt64(&vm.ready) == 0

		html := vm.RenderHTML()
		// fmt.Printf("NewHTML %s\n", html)
		gujs.Patch(gujs.CreateFragment(string(html)), vm.dom, replaceOnly)

		// If we have just updated then
		// Set the ready signal as on.
		atomic.StoreInt64(&vm.ready, 1)
	})

	return vm
}

// Path sends a Path report to all renderables that are Resolvable.
func (v *view) Path(path gudispatch.Path) {
	for _, vm := range v.renders {
		if rs, ok := vm.(gudispatch.Resolvable); ok {
			rs.Resolve(path)
		}
	}
}

// UID returns the custom id associated with this view.
func (v *view) UID() string {
	return v.uuid
}

// UUID returns the identification number associated with this view.
func (v *view) UUID() string {
	return v.uuid
}

// ViewUpdate defines a view update notification which contains the name of the
// view to be notified for an update.
type ViewUpdate struct {
	ID string
}

// Bind connects this giving view to both the update notification of the
// provided view. Whenever the provided view updates, so will this as well.
func (v *view) Bind(vs Views) {
	gudispatch.Subscribe(func(vm ViewUpdate) {
		if vm.ID != vs.UUID() && vm.ID != vs.UID() {
			return
		}

		// Notify this view of change.
		gudispatch.Dispatch(ViewUpdate{ID: v.UUID()})
	})
}

// ViewState defines a notification struct of the state of the view wether it
// is active or not.
type ViewState struct {
	ID string
	On bool
}

// Sync connects this giving view to the display state of the provided view.
// Whenever the provided view shows/hides, so will this view.
func (v *view) Sync(vs Views) {
	gudispatch.Subscribe(func(vm ViewState) {
		if vm.ID != vs.UUID() && vm.ID != vs.UID() {
			return
		}

		if !vm.On {
			vs.Hide()
			return
		}

		vs.Show()
	})
}

// BindSync connects this giving view to both the display state and update
// notification of the provided view.
// Whenever the provided view shows/hides, so will this view and when the
// provided view updates, so will this as well.
func (v *view) BindSync(vs Views) {
	v.Bind(vs)
	gudispatch.Subscribe(func(vm ViewState) {
		if vm.ID != vs.UUID() && vm.ID != vs.UID() {
			return
		}

		if !vm.On {
			vs.Hide()
			return
		}

		vs.Show()
	})
}

// Mount is to be called in the browser to loadup this view with a dom
func (v *view) Mount(dom *js.Object) {
	v.dom = dom
	v.events.OffloadDOM()
	v.events.LoadDOM(dom)

	// Set the ready state as zero.
	atomic.StoreInt64(&v.ready, 0)

	// Notify for update to dom.
	gudispatch.Dispatch(ViewUpdate{
		ID: v.UUID(),
	})
}

// Show activates the view to generate a visible markup.
func (v *view) Show() {
	atomic.StoreInt64(&v.switchActive, 1)
	{
		v.activeState = shower
	}
	atomic.StoreInt64(&v.switchActive, 0)

	gudispatch.Dispatch(ViewUpdate{ID: v.UUID()})

	gudispatch.Dispatch(ViewState{
		ID: v.UUID(),
		On: true,
	})
}

// Hide deactivates the view by adding a 'display:none'.
func (v *view) Hide() {
	atomic.StoreInt64(&v.switchActive, 1)
	{
		v.activeState = hider
	}
	atomic.StoreInt64(&v.switchActive, 0)

	gudispatch.Dispatch(ViewUpdate{ID: v.UUID()})

	gudispatch.Dispatch(ViewState{
		ID: v.UUID(),
		On: false,
	})
}

// Events returns the views events manager.
func (v *view) Events() guevents.EventManagers {
	return v.events
}

//==============================================================================

// Render renders the generated markup for this view, if the renderers are more
// than one then all are rendered into a div(as we need this to maintain sanity
// during reconciliation and updates) of rendered dom.
func (v *view) Render() gutrees.Markup {
	if len(v.renders) == 0 {
		return elems.Div()
	}

	var dom gutrees.Markup

	// If we have more than 1 then run through and apply all to a div.
	if len(v.renders) > 1 {
		dom = elems.Div()

		for _, rv := range v.renders {
			rv.Render().Apply(dom)
		}

	} else {
		dom = v.renders[0].Render()
	}

	atomic.StoreInt64(&v.switchActive, 1)
	{
		v.activeState.Render(dom)
	}
	atomic.StoreInt64(&v.switchActive, 0)

	if v.liveMarkup != nil {
		dom.Reconcile(v.liveMarkup)
	}

	// swap the uid for the new dom
	// to ensure we keep the sync between backend and frontend in sync.
	if backdoor, ok := dom.(gutrees.SwappableIdentity); ok {
		backdoor.SwapUID(v.uid)
	}

	if eventdoor, ok := dom.(gutrees.Eventers); ok {
		eventdoor.UseEventManager(v.events)
	}

	v.events.LoadUpEvents()
	v.liveMarkup = dom

	v.liveMarkup.Finalize(nil)

	return dom
}

// RenderHTML renders out the views markup as a string wrapped with template.HTML
func (v *view) RenderHTML() template.HTML {
	return v.Render().EHTML()
}
