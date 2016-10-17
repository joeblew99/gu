package gu

import (
	"html/template"

	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/events"
	"github.com/influx6/gu/trees"
)

//==============================================================================

// StaticView defines a MarkupRenderer implementing structure which returns its Content has
// its markup.
type StaticView struct {
	Content trees.Markup
}

// Render returns the markup for the static view.
func (s *StaticView) Render() trees.Markup {
	return s.Content
}

// RenderHTML returns the html template version of the StaticView content.
func (s *StaticView) RenderHTML() template.HTML {
	return s.Content.EHTML()
}

//==============================================================================

// ViewUpdate defines a view update notification which contains the name of the
// view to be notified for an update.
type ViewUpdate struct {
	ID string
}

// CustomView generates a RenderView for the provided Renderable.
func CustomView(tag string, ev events.EventManagers, r ...Renderable) RenderView {
	var vw view
	vw.tag = tag
	vw.renders = r
	vw.uuid = newKey()
	vw.events = ev

	for _, vr := range r {
		if rws, ok := vr.(ReactiveSubscription); ok {
			rws.React(func() {
				dispatch.Dispatch(ViewUpdate{ID: vw.uuid})
			})
		}
	}

	return &vw
}

//==============================================================================

// view defines a base level implementation for a set of Renderables.
type view struct {
	tag     string
	uuid    string
	hide    bool
	live    trees.Markup
	renders []Renderable
	events  events.EventManagers
}

// Events returns the events.EventManager attached with this view.
func (v *view) Events() events.EventManagers {
	return v.events
}

// Resolves exposes the internal renderables and passes the supplied path
// to allow any desired behaviour to be initiated.
func (v *view) Resolve(path dispatch.Path) {
	v.live.Router().Resolve(path)

	for _, vmr := range v.renders {
		if rs, ok := vmr.(dispatch.Resolvable); ok {
			rs.Resolve(path)
		}
	}
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
func (v *view) Render() trees.Markup {
	if len(v.renders) == 0 {
		return trees.NewElement(v.tag, false)
	}

	var root trees.Markup
	if len(v.renders) > 1 {
		root = trees.NewElement(v.tag, false)

		for _, view := range v.renders {
			view.Render().Apply(root)
		}

	} else {
		root = v.renders[0].Render()
	}

	if v.live != nil {
		root.Reconcile(v.live)
	}

	if swapper, ok := root.(trees.SwappableIdentity); ok {
		swapper.SwapUID(v.uuid)
	}

	if eventers, ok := root.(trees.Eventers); ok {
		eventers.UseEventManager(v.events)
	}

	root = root.ApplyMorphers()

	v.events.LoadUpEvents()
	v.live = root

	return root
}
