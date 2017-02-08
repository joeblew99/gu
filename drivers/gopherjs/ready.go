package gopherjs

import (
	"github.com/gopherjs/gopherjs/js"
)

// DOMReady fires the provided function when the dom is ready.
func DOMReady(fn func()) {
	if callReady(fn) {
		return
	}

	doc := js.Global.Get("document")
	if doc.Get("addEventListener") != js.Undefined {

		// first choice is DOMContentLoaded event
		doc.Call("addEventListener", "DOMContentLoaded", func() { callReady(fn) }, false)
		// backup is window load event
		js.Global.Call("addEventListener", "load", func() { callReady(fn) }, false)

	} else {

		// Must be IE
		doc.Call("attachEvent", "onreadystatechange", callReady)

		js.Global.Call("attachEvent", "onload", func() {
			callReady(fn)
		})
	}
}

// callReady returns true/false if the document had reached a ready state.
func callReady(fn func()) bool {
	doc := js.Global.Get("document")
	if doc.Get("readyState").String() == "complete" {
		go fn()
		return true
	}

	return false
}
