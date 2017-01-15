package gu

import (
	"html/template"

	"github.com/gu-io/gu/dispatch"
	"github.com/gu-io/gu/shell"
	"github.com/gu-io/gu/trees"
	"github.com/influx6/faux/reflection"
)

//==============================================================================

// StaticView defines a MarkupRenderer implementing structure which returns its Content has
// its markup.
type StaticView struct {
	uid      string
	Content  *trees.Markup
	Mounted  Subscriptions
	Rendered Subscriptions
	Morph    bool
}

// Static defines a toplevel function which returns a new instance of a StaticView using the
// provided markup as its content.
func Static(tree *trees.Markup) *StaticView {
	return &StaticView{
		Content: tree,
		uid:     NewKey(),
	}
}

// Identity defines an interface which expoese the identity of a giving object.
type Identity interface {
	UUID() string
}

// UUID returns the RenderGroup UUID for identification.
func (s *StaticView) UUID() string {
	return s.uid
}

// Dependencies returns the list of all internal dependencies of the given view.
// It returns the names of the structs and their internals composed values/fields
// to help conditional resource loading.
func (s *StaticView) Dependencies() []RenderableData {
	return nil
}

// Render returns the markup for the static view.
func (s *StaticView) Render() *trees.Markup {
	if s.Morph {
		return s.Content.ApplyMorphers()
	}

	return s.Content
}

// RenderHTML returns the html template version of the StaticView content.
func (s *StaticView) RenderHTML() template.HTML {
	return s.Content.EHTML()
}

//==============================================================================

// ViewSubscriptions defines an interface for structures which expose a subscription
// hooks to be used to register hooks for callers.
type ViewSubscriptions interface {
	OnSubscriptions(mounts, renders, unmount Subscriptions)
}

// RenderView defines an interface through which you gain access into a rendering
// branch of the current rendered markup view.
type RenderView interface {
	MarkupRenderer
	dispatch.Resolvable

	UUID() string
	RenderedBefore() bool
	Dependencies() []RenderableData
}

// Renderer defines an interface which takes responsiblity in translating
// the provided markup into the appropriate media.
type Renderer interface {
	RenderView(RenderView)
}

// CustomView generates a RenderView for the provided Renderable.
func CustomView(tag string, r ...Renderable) RenderView {
	var vw view
	vw.tag = tag
	vw.renders = r
	vw.uuid = NewKey()
	vw.Reactive = NewReactive()
	vw.mounted = NewSubscriptions()
	vw.unmounted = NewSubscriptions()
	vw.rendered = NewSubscriptions()

	for _, vr := range r {
		if renderField, _, err := reflection.StructAndEmbeddedTypeNames(vr); err == nil {
			if !vw.hasRenderable(renderField.TypeName) {
				vw.dependencies = append(vw.dependencies, RenderableData{
					Name: renderField.TypeName,
					Pkg:  renderField.Pkg,
				})
			}
		}

		if vs, ok := vr.(ViewSubscriptions); ok {
			vs.OnSubscriptions(vw.mounted, vw.rendered, vw.unmounted)
		}

		if rws, ok := vr.(Reactor); ok {
			rws.React(vw.Reactive.Publish)
		}
	}

	return &vw
}

//==============================================================================

// view defines a base level implementation for a set of Renderables.
type view struct {
	Reactive
	mounted        Subscriptions
	unmounted      Subscriptions
	rendered       Subscriptions
	hide           bool
	tag            string
	uuid           string
	renderedBefore bool
	live           *trees.Markup
	renders        []Renderable
	dependencies   []RenderableData
}

// RenderableData defines a struct which contains the name of a giving renderable
// and it's package.
type RenderableData struct {
	Name string
	Pkg  string
}

// hasRenderable returns true/false if a giving dependencies has been identified
// and is in the views dependencies list.
func (v *view) hasRenderable(name string) bool {
	for _, rd := range v.dependencies {
		if rd.Name == name {
			return true
		}
	}

	return false
}

// Dependencies returns the list of all internal dependencies of the given view.
// It returns the names of the structs and their internals composed values/fields
// to help conditional resource loading.
func (v *view) Dependencies() []RenderableData {
	return v.dependencies
}

// Resolves exposes the internal renderables and passes the supplied path
// to allow any desired behaviour to be initiated.
func (v *view) Resolve(path dispatch.Path) {
	for _, vmr := range v.renders {
		if rs, ok := vmr.(dispatch.Resolvable); ok {
			rs.Resolve(path)
		}
	}
}

// UseFetch exposes a UseFetch which passes the shell.Fetch and shell.Cache
// instance into a view.
func (v *view) UseFetch(f shell.Fetch, c shell.Cache) {
	for _, vmr := range v.renders {
		if rs, ok := vmr.(Fetchable); ok {
			rs.UseFetch(f, c)
		}
	}
}

// ViewHooks defines an interface which exposes a view internal hooks.
type ViewHooks interface {
	Hooks() (mounted Subscriptions, rendered Subscriptions, unmount Subscriptions)
}

// Hooks return the Subscriptions hooks used by the view.
func (v *view) Hooks() (Subscriptions, Subscriptions, Subscriptions) {
	return v.mounted, v.rendered, v.unmounted
}

// RenderedBefore returns true/false if the view has been rendered before.
func (v *view) RenderedBefore() bool {
	return v.renderedBefore
}

// UUID returns the RenderGroup UUID for identification.
func (v *view) UUID() string {
	return v.uuid
}

// RenderHTML returns the markup converted into a compliant html markup.
func (v *view) RenderHTML() template.HTML {
	return v.Render().EHTML()
}

// Render returns the groups markup for the giving render group.
func (v *view) Render() *trees.Markup {
	defer v.rendered.Publish()

	v.renderedBefore = true

	if len(v.renders) == 0 {
		return trees.NewMarkup(v.tag, false)
	}

	var root *trees.Markup

	if len(v.renders) > 1 {
		root = trees.NewMarkup(v.tag, false)

		for _, view := range v.renders {
			view.Render().Apply(root)
		}

	} else {
		root = v.renders[0].Render()
	}

	if v.live != nil {
		live := v.live
		live.EachEvent(func(e *trees.Event, _ *trees.Markup) {
			if e.Handle != nil {
				e.Handle.End()
			}
		})

		v.live = nil
		root.Reconcile(live)

		// Clear out internal references with the current live markup.
		// TODO: Check if this will cause unknown side effects.
		if live != root {
			live.Empty()
		}
	}

	root.SwapUID(v.uuid)
	root = root.ApplyMorphers()

	v.live = root

	return root
}
