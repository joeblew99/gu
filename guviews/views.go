package guviews

import (
	"html/template"
	"strings"
	"sync/atomic"

	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/faux/pub"
	"github.com/influx6/gu/guevents"
	"github.com/influx6/gu/gujs"
	"github.com/influx6/gu/gustates"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
)

//==============================================================================

// MarkupRenderer provides a interface for a types capable of rendering dom markup.
type MarkupRenderer interface {
	Render(...string) gutrees.Markup
	RenderHTML(...string) template.HTML
}

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
	pub.Publisher
	gustates.States
	Behaviour
	MarkupRenderer

	Events() guevents.EventManagers
	Mount(*js.Object)
	BindView(Views)
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
func (v *HideView) Render(m gutrees.Markup) {
	//if we are allowed to query for styles then check and change display
	if ds, err := gutrees.GetStyle(m, "display"); err == nil {
		if !strings.Contains(ds.Value, "none") {
			ds.Value = "none"
		}
	}
}

// ShowView provides a ViewStates for Views active state
type ShowView struct{}

// Render marks the given markup with a display: block
func (v *ShowView) Render(m gutrees.Markup) {
	//if we are allowed to query for styles then check and change display
	if ds, err := gutrees.GetStyle(m, "display"); err == nil {
		if strings.Contains(ds.Value, "none") {
			ds.Value = "block"
		}
	}
}

//==============================================================================

// View represent a basic Haiku view
type View struct {
	gustates.States
	pub.Publisher
	HideState   ViewStates
	ShowState   ViewStates
	activeState ViewStates
	encoder     gutrees.MarkupWriter
	events      guevents.EventManagers
	dom         *js.Object
	rview       Renderable
	liveMarkup  gutrees.Markup //liveMarkup represent the current rendered markup
	backdoor    gutrees.MutableBackdoor
	loaded      int64
	uid         string
}

// NewView returns a View instance. The view is giving a empty uid string,
// which it will generate itself. Use NewScopedView for more control especially
// when rendering from server.
func NewView(view Renderable) *View {
	return MakeView("", gutrees.SimpleMarkupWriter, view)
}

// NewScopedView returns a View instance with a custom UID.
// These allows managing effective transisitions from rendering on the backends
// to rendering on the frontend. The uid provides a unchanging identification
// of the server-rendered dom node the view is concerned about and it loads off
// that dom node on the client side to ensure correct transition.
func NewScopedView(uid string, view Renderable) *View {
	return MakeView(uid, gutrees.SimpleMarkupWriter, view)
}

// SequenceView returns a new  View instance rendered through a sequence renderer.
func SequenceView(meta SequenceMeta, rs ...Renderable) *View {
	return MakeView(meta.UID, gutrees.SimpleMarkupWriter, Sequence(meta, rs...))
}

// MakeView returns a Components style
func MakeView(uid string, writer gutrees.MarkupWriter, vw Renderable) (vm *View) {
	if uid == "" {
		uid = gutrees.RandString(8)
	}

	vm = &View{
		Publisher: pub.Always(vm),
		States:    gustates.NewState(),
		HideState: &HideView{},
		ShowState: &ShowView{},
		events:    guevents.NewEventManager(),
		encoder:   writer,
		rview:     vw,
		uid:       uid,
	}

	// If its a ReactiveRenderable type then bind the view
	if rxv, ok := vw.(ReactiveRenderable); ok {
		rxv.Bind(vm, true)
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

	vm.States.UseActivator(func() {
		vm.Show()
	})

	vm.States.UseDeactivator(func() {
		vm.Hide()
	})

	return
}

// BindView binds the given views together,were the view provided as argument will notify this view of change and to act according
func (v *View) BindView(vs Views) {
	vs.Bind(v, true)
}

// Mount is to be called in the browser to loadup this view with a dom
func (v *View) Mount(dom *js.Object) {
	v.dom = dom
	v.events.OffloadDOM()
	v.events.LoadDOM(dom)
	v.Send(true)
	atomic.StoreInt64(&v.loaded, 1)
}

// Show activates the view to generate a visible markup
func (v *View) Show() {
	if v.ShowState == nil {
		v.ShowState = &ShowView{}
	}
	v.activeState = v.ShowState
}

// Hide deactivates the view
func (v *View) Hide() {
	if v.HideState == nil {
		v.HideState = &HideView{}
	}
	v.activeState = v.HideState
}

// Events returns the views events manager
func (v *View) Events() guevents.EventManagers {
	return v.events
}

// Render renders the generated markup for this view
func (v *View) Render(m ...string) gutrees.Markup {
	if len(m) <= 0 {
		m = []string{"."}
	}

	v.Engine().All(m[0])

	if v.rview == nil {
		return elems.Div()
	}

	dom := v.rview.Render(m...)

	if dom == nil {
		return elems.Div()
	}

	// // swap the uid for the new dom
	// // to ensure we keep the sync between backend and frontend in sync.
	v.backdoor.M = dom
	v.backdoor.SwapUID(v.uid)
	v.backdoor.M = nil

	if v.liveMarkup != nil {
		dom.Reconcile(v.liveMarkup)
	}

	dom.UseEventManager(v.events)
	v.events.LoadUpEvents()
	v.liveMarkup = dom

	return dom
}

// RenderHTML renders out the views markup as a string wrapped with template.HTML
func (v *View) RenderHTML(m ...string) template.HTML {
	ma, _ := v.encoder.Write(v.Render(m...))
	return template.HTML(ma)
}

//==============================================================================

// SequenceMeta  provides a configuration object for SequenceRenderer.
type SequenceMeta struct {
	Tag   string   // Name of the root tag.
	UID   string   // Custom UID of the root dom tag.
	ID    string   // Id of the root tag.
	Class []string // Class list of the root tag.
}

// SequenceRenderer provides a rendering lists of Renderables to be rendered in
// their added sequence/order. SequenceRenderer embedds pub.Publisher and will
// automatically bind with any ReactiveRenderable being added if its `Bind` is
// set to true.
type SequenceRenderer struct {
	pub.Publisher
	*SequenceMeta
	stack []Renderable
}

// Sequence returns a new sequence renderer instance.
func Sequence(meta SequenceMeta, r ...Renderable) *SequenceRenderer {
	if meta.Tag == "" {
		meta.Tag = "div"
	}

	s := SequenceRenderer{
		Publisher:    pub.Identity(),
		SequenceMeta: &meta,
	}

	s.Add(r...)

	return &s
}

// Add adds new renders into the publisher lists.
func (s *SequenceRenderer) Add(r ...Renderable) {
	for _, rm := range r {
		if rx, ok := rm.(ReactiveRenderable); ok {
			rx.Bind(s, true)
		}
		s.stack = append(s.stack, rm)
	}
}

// Render renders the giving giving lists of views.
func (s *SequenceRenderer) Render(m ...string) gutrees.Markup {
	root := gutrees.NewElement(s.Tag, false)

	attrs.Class(strings.Join(s.Class, " ")).Apply(root)
	attrs.ID(s.ID).Apply(root)

	for _, st := range s.stack {
		st.Render(m...).Apply(root)
	}

	return root
}

//==============================================================================
