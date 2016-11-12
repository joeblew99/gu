package main

import (
	. "github.com/influx6/gu/design"
	_ "github.com/influx6/gu/examples/subscribe/app"
	. "github.com/influx6/gu/redom"
	"honnef.co/go/js/dom"
)

var _ = Resource(func() {

	DoTitle("App Subscribe")

	var sub Subscriber
	DoView(sub, "")

})

func main() {
	New(&DOMRenderer{Document: dom.GetWindow().Document()}).Init(true)
}
