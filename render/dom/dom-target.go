// Package dom contains the DOM RenderTarget for rendering gu markup into a
// browser DOM.
package dom

import (
	hdom "honnef.co/go/js/dom"

	"github.com/influx6/gu"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/gujs"
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
		if rg, ok := view.(gu.RenderingGroup); ok {
			for _, view := range rg.Views() {
				gudispatch.Subscribe(func(update gu.ViewUpdate) {
					if update.ID != view.UUID() {
						return
					}

					gujs.Patch(gujs.CreateFragment(view.Render().HTML()), dt.Target.Underlying(), false)
				})
			}
		}
	}

	gujs.Patch(gujs.CreateFragment(view.Render().HTML()), dt.Target.Underlying(), true)
}

// TreeTarget defines a DOM renderer which  handles rendering and update cycles different from 
// the DOMTarget, it ensures to move all contents of <head> and <body> into their proper positions 
// for rendering.
type TreeTarget struct {
	MountTarget string
	AutoUpdate bool
	doc hdom.Document
}

// RenderView manages the initialization and management of the rendering and
// update of the passed in view.
func (tree *TreeTarget) RenderView(view gu.RenderView) {
	if ev, ok := view.(gu.EventableRenderView); ok {
		ev.LoadEvents(tree.Target.Underlying())
	}

	if tree.doc == nil {
		window := hdom.GetWindow()
		tree.doc = window.Document()
	}

	target = tree.doc.QuerySelectorAll(tree.MountTarget)

	if tree.AutoUpdate {
		if rg, ok := view.(gu.RenderingGroup); ok {
			for _, view := range rg.Views() {
				gudispatch.Subscribe(func(update gu.ViewUpdate) {
					if update.ID != view.UUID() {
						return
					}

					gujs.Patch(gujs.CreateFragment(view.Render().HTML()), target.Underlying(), false)
				})
			}
		}
	}

	gujs.Patch(gujs.CreateFragment(view.Render().HTML()), target.Underlying(), true)
}
