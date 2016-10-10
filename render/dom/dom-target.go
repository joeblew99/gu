// Package dom contains the DOM RenderTarget for rendering gu markup into a 
// browser DOM.
package dom

import (
  hdom "honnef.co/go/js/dom"
)

// DOMTarget defines a DOM RenderingTarget to handle rendering and update cycles 
// for rendering gu markups into the provided DOM.
type DOMTarget struct{
  Target hdom.Element
  AutoUpdate bool
}

// HandleView manages the initialization and management of the rendering and 
// update of the passed in view.
func (dt DOMTarget) HandleView(view gu.RenderView){
  view.Events().OffloadDOM()
  view.Events().LoadDOM(dt.Target)

  if dt.AutoUpdate {
    gudispatch.Subscribe(function(update gu.ViewUpdate){
      if update.ID != view.UUID(){
        return
      }

      gujs.Patch(gujs.CreateFragment(view.Render().HTML()), dt.Target, !firstRender)
      firstRender = true
    })

    gudispatch.Publish(gu.ViewUpdate{
      ID: view.UUID(),
    })

    return
  }

  gujs.Patch(gujs.CreateFragment(view.Render().HTML()), dt.Target, true)
}