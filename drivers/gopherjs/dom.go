// Package gopherjs contains the DOM RenderTarget for rendering gu markup into a
// browser DOM.
package gopherjs

import (
	"fmt"

	"honnef.co/go/js/dom"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/notifications"
	"github.com/gu-io/gu/trees"
)

// DOM defines an package level variable for rendering.
var DOM domRenderer

type domRenderer struct{}

// Render drives the function to render the resource.
func (dm DOMRenderer) Render(rs ...*gu.ResourceDefinition) {
	Render(dom.GetWindow().Document(), rs...)
}

// Update drives the function to update the resource render.
func (dm DOMRenderer) Update(r gu.Renderable, target string, update bool) {
	RenderUpdate(dom.GetWindow().Document(), r, target, update)
}

// Render renders the giving set of resources into the provided body and header
// of the DOM.
func Render(doc dom.Document, rs ...*gu.ResourceDefinition) {
	head := doc.QuerySelector("head")
	body := doc.QuerySelector("body")

	// clear all children of head and body if the belong to us.
	for _, item := range head.QuerySelectorAll("[data-gen*='gu']") {
		if !item.HasAttribute("gu-resource-root") {
			item.ParentNode().RemoveChild(item)
		}
	}

	for _, item := range body.QuerySelectorAll("[data-gen*='gu']") {
		if !item.HasAttribute("gu-resource-root") {
			item.ParentNode().RemoveChild(item)
		}
	}

	// Load global Resources first.
	if len(rs) != 0 {
		globalHead, globalBody := rs[0].Root.GenerateResources()

		for _, item := range globalHead {
			Patch(CreateFragment(item.HTML()), head.Underlying(), false)
		}

		for _, item := range globalBody {
			Patch(CreateFragment(item.HTML()), body.Underlying(), false)
		}
	}

	// Render the important resources and normal links first.
	for _, res := range rs {

		// Retrieve base resources which must be rendered for each component and
		// add them according to the head and body.
		toHead, toBody := res.Resources()

		for _, item := range toHead {
			Patch(CreateFragment(item.HTML()), head.Underlying(), false)
		}

		for _, item := range toBody {
			Patch(CreateFragment(item.HTML()), body.Underlying(), false)
		}

		for _, item := range res.Links {
			markup := item.Render()

			markup.EachEvent(func(ev *trees.Event, root *trees.Markup) {
				dm.BindEvent(ev, head.Underlying())
			})

			Patch(CreateFragment(markup.HTML()), head.Underlying(), false)
		}
	}

	// Render all basic views which are not to be deffered.
	for _, res := range rs {
		for _, render := range res.Renderables {
			RenderUpdate(doc, render.View, render.Targets, false)
		}
	}

	// Render all defered views.
	for _, res := range rs {
		for _, render := range res.DRenderables {
			RenderUpdate(doc, render.View, render.Targets, false)
		}
	}

	// Render the defered links.
	for _, res := range rs {
		for _, item := range res.DeferLinks {
			markup := item.Render()

			markup.EachEvent(func(ev *trees.Event, root *trees.Markup) {
				BindEvent(ev, body.Underlying())
			})

			Patch(CreateFragment(markup.HTML()), body.Underlying(), false)
		}
	}
}

// RenderUpdate handles rendering calls for individual renderers which may have
// determined targets within the body.
func RenderUpdate(document dom.Documnt, rv gu.Renderable, targets string, update bool) {
	body := Document.QuerySelector("body")

	if targets == "" {
		markup := rv.Render()

		if !update {
			markup.EachEvent(func(ev *trees.Event, root *trees.Markup) {
				BindEvent(ev, body.Underlying())
			})
		}

		if kvr, ok := rv.(gu.RenderView); ok {
			Patch(CreateFragment(markup.HTML()), body.Underlying(), !kvr.RenderedBefore())

			if cvs, ok := rv.(gu.ViewHooks); ok {
				mounted, _, _ := cvs.Hooks()
				if !mounted.Used() {
					mounted.Publish()
				}
			}

			return
		}

		Patch(CreateFragment(markup.HTML()), body.Underlying(), false)

		if cvs, ok := rv.(gu.ViewHooks); ok {
			mounted, _, _ := cvs.Hooks()
			if !mounted.Used() {
				mounted.Publish()
			}
		}

		return
	}

	if kernels := body.QuerySelectorAll(targets); kernels == nil {
		fmt.Printf("Unable to mount to target %q\n", targets)
		return
	}

	for _, targetDOM := range kernels {
		markup := rv.Render()

		if !update {
			markup.EachEvent(func(ev *trees.Event, root *trees.Markup) {
				BindEvent(ev, targetDOM.Underlying())
			})
		}

		if kvr, ok := rv.(gu.RenderView); ok {
			Patch(CreateFragment(markup.HTML()), targetDOM.Underlying(), !kvr.RenderedBefore())
			continue
		}

		Patch(CreateFragment(markup.HTML()), targetDOM.Underlying(), false)
	}

	if cvs, ok := rv.(gu.ViewHooks); ok {
		mounted, _, _ := cvs.Hooks()

		if !mounted.Used() {
			mounted.Publish()
		}
	}
}

// BindEvent connects the event with the provided event object and root.
func BindEvent(source *trees.Event, root *js.Object) {
	link = func(ev *js.Object) { TriggerBindEvent(ev, root, source) }

	root.Call("addEventListener", source.Type, link, true)

	source.Handle.AddEnd(func() {
		root.Call("removeEventListener", source.Type, link, true)
	})
}

// TriggerBindEvent connects the giving event with the provided dom target.
func TriggerBindEvent(event *js.Object, root *js.Object, source *trees.Event) {
	target := event.Get("target")

	children := root.Call("querySelectorAll", source.Target())
	if children == nil || children == js.Undefined {
		return
	}

	kids := DOMObjectToList(children)
	var match bool

	for _, item := range kids {
		if item != target {
			continue
		}

		match = true
		break
	}

	// if we match then run the listeners registered.
	if match {
		notifications.Dispatch(trees.EventBroadcast{
			EventID: source.EventID,
			Event:   NewWrapperEvent(event),
		})
	}
}
