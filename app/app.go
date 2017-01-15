package app

import (
	"github.com/go-humble/detect"
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/redom"
	"honnef.co/go/js/dom"
)

// New returns a new design.Resources activating the DOMRenderer if its gets called
// on the browser or else on the server.
func New(appName string, op *gu.Options) *gu.Resources {
	if detect.IsBrowser() {
		return gu.New(appName, op, &redom.DOMRenderer{
			Document: dom.GetWindow().Document(),
		})
	}

	return gu.New(appName, op)
}
