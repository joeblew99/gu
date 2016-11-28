package gu

import (
	"html/template"

	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/trees"
)

//==============================================================================

// StaticView defines a MarkupRenderer implementing structure which returns its Content has
// its markup.
type StaticView struct {
	Content *trees.Markup
	Morph   bool
}

// Static defines a toplevel function which returns a new instance of a StaticView using the
// provided markup as its content.
func Static(tree *trees.Markup) *StaticView {
	return &StaticView{
		Content: tree,
	}
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

// CustomView generates a RenderView for the provided Renderable.
func CustomView(tag string, r ...Renderable) RenderView {
	var vw view
	vw.tag = tag
	vw.renders = r
	vw.uuid = NewKey()
	vw.Reactive = NewReactive()

	for _, vr := range r {
		if rws, ok := vr.(ReactiveSubscription); ok {
			rws.React(vw.Reactive.Publish)
		}
	}

	return &vw
}

//==============================================================================

// view defines a base level implementation for a set of Renderables.
type view struct {
	Reactive
	hide    bool
	tag     string
	uuid    string
	renderedBefore    bool
	live    *trees.Markup
	renders []Renderable
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
