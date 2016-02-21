package guviews

import (
	"html/template"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/faux/pub"
	"github.com/influx6/gu/guevents"
	"github.com/influx6/gu/gujs"
	"github.com/influx6/gu/gustates"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/elems"
)

//==============================================================================

// Renderable provides a interface for a renderable type.
type Renderable interface {
	Render(...string) gutrees.Markup
}

//==============================================================================

// ReactiveRenderable provides a interface for a reactive renderable type.
type ReactiveRenderable interface {
	pub.Publisher
	Renderable
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
	pub.Publisher
	gustates.States
	MarkupRenderer

	BindView(Views)
	Mount(*js.Object)
	Events() guevents.EventManagers
}

//==============================================================================

// ViewStates defines the two possible behavioral state of a view's markup
type ViewStates interface {
	Render(gutrees.Markup)
}

//==============================================================================

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

// hider defines a global hide renderer.
var hider HideView

// shower defines a global display renderer.
var shower ShowView

// view defines a basic struture for building UI view.
type view struct {
	gustates.States
	pub.Publisher
	loaded      int64
	uid         string
	dom         *js.Object
	renders     []Renderable
	liveMarkup  gutrees.Markup
	encoder     gutrees.MarkupWriter
	events      guevents.EventManagers
	rl          sync.RWMutex
	activeState ViewStates
}

// View returns a View instance. The view is giving a empty uid string,
// which it will generate itself. Use NewScopedView for more control especially
// when rendering from server.
func View(view ...Renderable) Views {
	return CustomView("", gutrees.SimpleMarkupWriter, view...)
}

// CustomView returns a gu.Views implementing struct that provides the ability to
// render and update UI efficiently. This function allows greater control of
// the uid for which the dom will be identified with and the writer used to
// decode our dom structures into valid html.
func CustomView(uid string, writer gutrees.MarkupWriter, vw ...Renderable) Views {
	if uid == "" {
		uid = gutrees.RandString(8)
	}

	vm := &view{
		States:  gustates.NewState(),
		events:  guevents.NewEventManager(),
		encoder: writer,
		renders: vw,
		uid:     uid,
	}

	vm.Publisher = pub.Always(vm)

	for _, rv := range vw {
		// If its a ReactiveRenderable type then bind the view
		if rxv, ok := rv.(ReactiveRenderable); ok {
			rxv.Bind(vm, true)
		}
	}

	//set up the reaction chain, if we have node attach then render to it
	vm.React(func(r pub.Publisher, _ error, _ interface{}) {
		//if we are not domless then patch
		if vm.dom != nil {
			replaceOnly := atomic.LoadInt64(&vm.loaded) == 0
			html := vm.RenderHTML()
			gujs.Patch(gujs.CreateFragment(string(html)), vm.dom, replaceOnly)
		}
	}, true)

	vm.UseActivator(vm.Show)
	vm.UseDeactivator(vm.Hide)

	return vm
}

// BindView binds the given views together,were the view provided as argument will notify this view of change and to act according
func (v *view) BindView(vs Views) {
	vs.Bind(v, true)
}

// Mount is to be called in the browser to loadup this view with a dom
func (v *view) Mount(dom *js.Object) {
	v.dom = dom
	v.events.OffloadDOM()
	v.events.LoadDOM(dom)
	atomic.StoreInt64(&v.loaded, 1)
	v.Send(true)
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
