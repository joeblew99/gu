package app

import (
	"github.com/go-humble/detect"
	"github.com/influx6/gu/design"
	"github.com/influx6/gu/redom"
	"honnef.co/go/js/dom"
)

// New returns a new design.Resources activating the DOMRenderer if its gets called
// on the browser or else on the server.
func New() *design.Resources {
	if detect.IsBrowser() {
		return design.New(&redom.DOMRenderer{
			Document: dom.GetWindow().Document(),
		})
	}

	return design.New()
}
