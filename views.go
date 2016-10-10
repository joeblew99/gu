package gu

import (
	"html/template"

	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/guevents"
	"github.com/influx6/gu/gutrees"
)


//==============================================================================

// ViewState defines a notification struct of the state of the view wether it
// is active or not.
type ViewState struct {
	ID string
	On bool
}

// ViewUpdate defines a view update notification which contains the name of the
// view to be notified for an update.
type ViewUpdate struct {
	ID string
}

// CustomView generates a RenderView for the provided Renderable.
func customView(tag string, events guevents.EventManagers, r ...Renderable) RenderView {
	var vw view
	vw.uuid = newKey()
	vw.events = events
	vw.tag = tag
	vw.renders = r

  for _, vr := range r {
    if rws, ok := vr.(ReactiveSubscription); ok {
      rws.React(func() {
        gudispatch.Dispatch(ViewUpdate{ID: vw.uuid})
      })
    }
  }


	gudispatch.Subscribe(func(path gudispatch.Path) {
		for _, vmr := range r {
			if rs, ok := vmr.(gudispatch.Resolvable); ok {
				rs.Resolve(path)
			}
		}
	})

	gudispatch.Subscribe(func(state ViewState) {
		if state.ID != vw.uuid {
			return
		}

		vw.hide = !state.On
	})

	return &vw
}

//==============================================================================

// view defines a base level implementation for a set of Renderables.
type view struct {
	tag     string
	uuid    string
	hide    bool
	live    gutrees.Markup
	renders []Renderable
	events  guevents.EventManagers
}

// Events returns the guevents.EventManager attached with this view.
func (v *view) Events() guevents.EventManagers {
  return v.events
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
func (v *view) Render() gutrees.Markup {
	if len(v.renders) == 0 {
		return gutrees.NewElement("div", false)
	}

	var root gutrees.Markup

	if len(v.renders) > 1 {
		root = gutrees.NewElement(v.tag, false)

		for _, view := range v.renders {
			view.Render().Apply(root)
		}

	} else {
		root = v.renders[0].Render()
	}

	if v.live != nil {
		root.Reconcile(v.live)
	}

	if swapper, ok := root.(gutrees.SwappableIdentity); ok {
		swapper.SwapUID(v.uuid)
	}

	if eventers, ok := root.(gutrees.Eventers); ok {
		eventers.UseEventManager(v.events)
	}

	root.Finalize(nil)

	v.events.LoadUpEvents()
	v.live = root

	if v.hide {
		gutrees.Hide.Mode(root)
	}

	return root
}
