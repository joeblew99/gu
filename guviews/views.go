package guviews

import (
	"html/template"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/guevents"
	"github.com/influx6/gu/gujs"
	"github.com/influx6/gu/gustates"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/pborman/uuid"
)

//==============================================================================

// Renderable provides a interface for a renderable type.
type Renderable interface {
	Render(...string) gutrees.Markup
}

//==============================================================================

// Behaviour provides a state changers for haiku.
type Behaviour interface {
	Hide()
	Show()
}

//==============================================================================

// Views define a Haiku Component
type Views interface {
	Behaviour
	gustates.States
	MarkupRenderer

	UUID() string
	UID() string

	BindView(Views)
	Mount(*js.Object)
	Events() guevents.EventManagers
}

//==============================================================================

// ViewStates defines the two possible behavioral state of a view's markup
type ViewStates interface {
	Render(gutrees.Markup)
}

// HideView provides a ViewStates for Views inactive state
type HideView struct{}

// Render marks the given markup as display:none
func (v HideView) Render(m gutrees.Markup) {
	//if we are allowed to query for styles then check and change display
	if mm, ok := m.(gutrees.MarkupProps); ok {
		if ds, err := gutrees.GetStyle(mm, "display"); err == nil {
			if !strings.Contains(ds.Value, "none") {
				ds.Value = "none"
			}
		}
	}
}

// ShowView provides a ViewStates for Views active state
type ShowView struct{}

// Render marks the given markup with a display: block
func (v ShowView) Render(m gutrees.Markup) {
	//if we are allowed to query for styles then check and change display
	if mm, ok := m.(gutrees.MarkupProps); ok {
		if ds, err := gutrees.GetStyle(mm, "display"); err == nil {
			if strings.Contains(ds.Value, "none") {
				ds.Value = "block"
			}
		}
	}
}

//==============================================================================

// ViewUpdate defines a view update notification which contains the name of the
// view to be notified for an update.
type ViewUpdate struct {
	ID string
}

//==============================================================================

// hider defines a global hide renderer.
var hider HideView

// shower defines a global display renderer.
var shower ShowView

// view defines a basic struture for building UI view.
type view struct {
	gustates.States
	loaded      int64
	uid         string
	uuid        string
	dom         *js.Object
	renders     []Renderable
	liveMarkup  gutrees.Markup
	encoder     gutrees.MarkupWriter
	events      guevents.EventManagers
	rl          sync.RWMutex
	activeState ViewStates
}

// View returns a View instance. The view is giving a customID string, which
// gets associated with this view, and provides a convenient means of dispatching
// events directly to it, if this is a empty string, a random one will be
// generated for it.
func View(customID string, view ...Renderable) Views {
	return CustomView(customID, gutrees.SimpleMarkupWriter, view...)
}

// CustomView returns a gu.Views implementing struct that provides the ability to
// render and update UI efficiently. This function allows greater control of
// the customId for which the views and it's dom will be identified with and
// the writer used to decode our dom structures into valid html.
func CustomView(cid string, writer gutrees.MarkupWriter, vw ...Renderable) Views {
	if cid == "" {
		cid = gutrees.RandString(8)
	}

	vm := &view{
		States:  gustates.NewState(),
		events:  guevents.NewEventManager(),
		encoder: writer,
		renders: vw,
		uid:     cid,
		uuid:    uuid.New(),
	}

	// Subscribe for view update requests from the central dispatcher.
	gudispatch.Subscribe(func(v *ViewUpdate) {
		if v.ID != vm.UUID() && v.ID != vm.UID() {
			return
		}

		//if we are not domless then patch
		if vm.dom == nil {
			return
		}

		replaceOnly := atomic.LoadInt64(&vm.loaded) == 0
		html := vm.RenderHTML()
		gujs.Patch(gujs.CreateFragment(string(html)), vm.dom, replaceOnly)
	})

	// Subscribe for URI updates.
	gudispatch.Subscribe(func(p *gudispatch.PathDirective) {
		vm.Engine().All(p.Sequence)
	})

	// Connect the corresponding state methods to the state manager.
	vm.UseActivator(vm.Show)
	vm.UseDeactivator(vm.Hide)

	return vm
}

// UID returns the custom id associated with this view.
func (v *view) UID() string {
	return v.uuid
}

// UUID returns the identification number associated with this view.
func (v *view) UUID() string {
	return v.uuid
}

// BindView binds the given views together,were the view provided as argument will notify this view of change and to act according
func (v *view) BindView(vs Views) {
	gudispatch.Subscribe(func(vm *ViewUpdate) {
		if vm.ID != vs.UUID() && vm.ID != vs.UID() {
			return
		}

		// Notify this view of change.
		gudispatch.Dispatch(&ViewUpdate{ID: v.UUID()})

	})
}

// Mount is to be called in the browser to loadup this view with a dom
func (v *view) Mount(dom *js.Object) {
	v.dom = dom
	v.events.OffloadDOM()
	v.events.LoadDOM(dom)
	atomic.StoreInt64(&v.loaded, 1)

	// Notify for update to dom.
	gudispatch.Dispatch(&ViewUpdate{
		ID: v.UUID(),
	})
}

// Show activates the view to generate a visible markup
func (v *view) Show() {
	v.rl.Lock()
	defer v.rl.Unlock()
	v.activeState = shower
}

// Hide deactivates the view
func (v *view) Hide() {
	v.rl.Lock()
	defer v.rl.Unlock()
	v.activeState = hider
}

// Events returns the views events manager
func (v *view) Events() guevents.EventManagers {
	return v.events
}

//==============================================================================

// MarkupRenderer provides a interface for a types capable of rendering dom markup.
type MarkupRenderer interface {
	Renderable
	RenderHTML(...string) template.HTML
}

// Render renders the generated markup for this view, if the renderers are more
// than one then all are rendered into a div(as we need this to maintain sanity
// during reconciliation and updates) of rendered dom.
func (v *view) Render(m ...string) gutrees.Markup {
	if len(m) <= 0 {
		m = []string{"."}
	}

	v.Engine().All(m[0])

	if len(v.renders) == 0 {
		return elems.Div()
	}

	var dom gutrees.Markup

	// If we are more than 1 then run through and apply all to a div.
	if len(v.renders) > 1 {
		dom = elems.Div()

		for _, rv := range v.renders {
			rv.Render(m...).Apply(dom)
		}

	} else {
		dom = v.renders[0].Render(m...)
	}

	v.rl.RLock()
	v.activeState.Render(dom)
	v.rl.RUnlock()

	// // swap the uid for the new dom
	// // to ensure we keep the sync between backend and frontend in sync.
	if backdoor, ok := dom.(gutrees.SwappableIdentity); ok {
		backdoor.SwapUID(v.uid)
	}

	if v.liveMarkup != nil {
		dom.Reconcile(v.liveMarkup)
	}

	if eventdoor, ok := dom.(gutrees.Eventers); ok {
		eventdoor.UseEventManager(v.events)
	}

	v.events.LoadUpEvents()
	v.liveMarkup = dom

	return dom
}

// RenderHTML renders out the views markup as a string wrapped with template.HTML
func (v *view) RenderHTML(m ...string) template.HTML {
	ma, _ := v.encoder.Write(v.Render(m...))
	return template.HTML(ma)
}
