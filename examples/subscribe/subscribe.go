package main

import (
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/examples/subscribe/app"
	. "github.com/influx6/gu/redom"
	. "github.com/influx6/gu/trees/elems"
	. "github.com/influx6/gu/trees/property"
	"honnef.co/go/js/dom"
)

var _ = Resource(func() {

	DoTitle("App Subscribe")

	DoMarkup(Div(
		ClassAttr("main-header"),
		Header1(
			Text("Become A Subscriber"),
		),
	), "")

	var sub Subscriber
	DoView(&sub, "")

})

func main() {
	New(&DOMRenderer{Document: dom.GetWindow().Document()}).Init(true)
}
