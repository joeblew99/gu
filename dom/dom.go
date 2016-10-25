// Package dom contains the DOM RenderTarget for rendering gu markup into a
// browser DOM.
package dom

import (
	hdom "honnef.co/go/js/dom"

	"github.com/influx6/gu"
	"github.com/influx6/gu/design"
	"github.com/influx6/gu/js"
)

// DOMRenderer defines an implementation for gu.design.ResourceRenderere
// and handles rendering of a giving group of resources into the live DOM body root.
type DOMRenderer struct {
	Document hdom.Document
}

// Render renders the giving set of resources into the provided body and header
// of the DOM.
func (dm *DOMRenderer) Render(rs ...design.ResourceDefinition) {
	head := dm.Document.QuerySelector("head")
	body := dm.Document.QuerySelector("body")

	// clear all children of head and body if the belong to us.
	for _, item := range head.QuerySelectorAll("[data-gen*='gu']") {
		item.ParentNode().RemoveChild(item)
	}

	for _, item := range body.QuerySelectorAll("[data-gen*='gu']") {
		item.ParentNode().RemoveChild(item)
	}

	for _, res := range rs {
		res.Events.LoadDOM(body.Underlying())

		// Render the normal links first.
		for _, item := range res.Links {
			js.Patch(js.CreateFragment(item.Render().HTML()), head.Underlying(), false)
		}

		// Render all basic views which are not to be deffered.
		for _, render := range res.Renderables {
			dm.RenderUpdate(render.View, render.Targets)
		}

		// Render all defered views.
		for _, render := range res.DRenderables {
			dm.RenderUpdate(render.View, render.Targets)
		}

		// Render the defered links.
		for _, item := range res.DeferLinks {
			js.Patch(js.CreateFragment(item.Render().HTML()), body.Underlying(), false)
		}
	}
}

// RenderUpdate handles rendering calls for individual renderers which may have
// determined targets within the body.
func (dm *DOMRenderer) RenderUpdate(rv gu.Renderable, targets string) {
	body := dm.Document.QuerySelector("body")

	if targets == "" {
		if ev, ok := rv.(gu.EventableRenderView); ok {
			ev.LoadEvents(body.Underlying())
		}

		js.Patch(js.CreateFragment(rv.Render().HTML()), body.Underlying(), false)
		return
	}

	kernels := body.QuerySelectorAll(targets)

	if len(kernels) == 0 {
		if ev, ok := rv.(gu.EventableRenderView); ok {
			ev.LoadEvents(kernels[0].Underlying())
		}
	} else {
		if ev, ok := rv.(gu.EventableRenderView); ok {
			ev.LoadEvents(body.Underlying())
		}
	}

	markup := rv.Render()
	for _, targetDOM := range kernels {
		js.Patch(js.CreateFragment(markup.HTML()), targetDOM.Underlying(), false)
	}
}
