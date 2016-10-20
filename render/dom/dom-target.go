// Package dom contains the DOM RenderTarget for rendering gu markup into a
// browser DOM.
package dom

import (
	hdom "honnef.co/go/js/dom"

	"github.com/influx6/gu"
	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/js"
)

// DOMTarget defines a DOM RenderingTarget to handle rendering and update cycles
// for rendering gu markups into the provided DOM.
type DOMTarget struct {
	Target     hdom.Element
	AutoUpdate bool
}

// RenderView manages the initialization and management of the rendering and
// update of the passed in view.
func (dt *DOMTarget) RenderView(view gu.RenderView) {
	if ev, ok := view.(gu.EventableRenderView); ok {
		ev.LoadEvents(dt.Target.Underlying())
	}

	if dt.AutoUpdate {
		dispatch.Subscribe(func(update gu.ViewUpdate) {
			if update.ID != view.UUID() {
				return
			}

			js.Patch(js.CreateFragment(view.Render().HTML()), dt.Target.Underlying(), false)
		})
	}

	js.Patch(js.CreateFragment(view.Render().HTML()), dt.Target.Underlying(), true)
}

// TreeTarget defines a DOM renderer which  handles rendering and update cycles different from
// the DOMTarget, it ensures to move all contents of <head> and <body> into their proper positions
// for rendering.
type TreeTarget struct {
	MountTarget string
	AutoUpdate  bool
	doc         hdom.Document
}

// RenderView manages the initialization and management of the rendering and
// update of the passed in view.
func (tree *TreeTarget) RenderView(view gu.RenderView) {
	if tree.doc == nil {
		window := hdom.GetWindow()
		tree.doc = window.Document()
	}

	for _, target := range tree.doc.QuerySelectorAll(tree.MountTarget) {

		if ev, ok := view.(gu.EventableRenderView); ok {
			ev.LoadEvents(target.Underlying())
		}

		if tree.AutoUpdate {
			dispatch.Subscribe(func(update gu.ViewUpdate) {
				if update.ID != view.UUID() {
					return
				}

				js.Patch(js.CreateFragment(view.Render().HTML()), target.Underlying(), false)
			})
		}

		js.Patch(js.CreateFragment(view.Render().HTML()), target.Underlying(), true)
	}
}