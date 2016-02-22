// Package gudom contains utility functions that provide easy interaction with
// the browser dom.
package gudom

import (
	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/gu/guviews"
)

// RenderAs assigns the appropriate container to be used for rendering.
func RenderAs(v guviews.Views, o *js.Object) {
	if !detect.IsBrowser() {
		return
	}

	v.Mount(o)
}

// RenderAsBody connects a specific as part of the DOM body.
func RenderAsBody(v guviews.Views) {
	if !detect.IsBrowser() {
		return
	}

	body := js.Global.Get("document").Get("body")
	body.Set("innerHTML", "")
	v.Mount(body)
}

// AddStylesheet adds an external stylesheet to the document.
func AddStylesheet(url string) {
	if !detect.IsBrowser() {
		return
	}

	link := js.Global.Get("document").Call("createElement", "link")
	link.Set("rel", "stylesheet")
	link.Set("href", url)
	js.Global.Get("document").Get("head").Call("appendChild", link)
}
