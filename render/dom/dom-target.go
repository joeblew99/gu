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

// HandleView manages the initialization and management of the rendering and
// update of the passed in view.
func (dt DOMTarget) HandleView(view gu.RenderView) {
  if evr, ok := view.(gu.EventableRenderView){
    view.Events().OffloadDOM()
    view.Events().LoadDOM(dt.Target)
  }

	if dt.AutoUpdate {
		gudispatch.Subscribe(func(update gu.ViewUpdate) {
			if update.ID != view.UUID() {
				return
			}

			gujs.Patch(gujs.CreateFragment(view.Render().HTML()), dt.Target, !firstRender)
			firstRender = true
		})
	}

	gujs.Patch(gujs.CreateFragment(view.Render().HTML()), dt.Target, true)
}
