// Package gu provides a UI framework for Go.
package gu

import (
	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/gu/gujs"
	"github.com/influx6/gu/guviews"
)

// RenderAs assigns the appropriate container to be used for rendering.
func RenderAs(v guviews.Views, o *js.Object) {
	if !detect.IsBrowser() {
		return
	}

	v.Mount(o)
}

// RenderSelector renders the giving view by using query selector to get the
// dom element to be called. If the selector is not found then it would fail to
// bind the view.
func RenderSelector(v guviews.Views, selector string) {
	if !detect.IsBrowser() {
		return
	}

	elem := gujs.QuerySelector(js.Global.Get("document"), selector)
	if elem == nil {
		return
	}

	v.Mount(elem)
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

// AddStylesheet adds an external stylesheet to the document into the document
// DOM using the provide URL.
func AddStylesheet(url string) {
	if !detect.IsBrowser() {
		return
	}

	link := js.Global.Get("document").Call("createElement", "link")
	link.Set("rel", "stylesheet")
	link.Set("href", url)
	js.Global.Get("document").Get("head").Call("appendChild", link)
}
