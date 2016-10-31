package main

import (
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/dom"
	_ "github.com/influx6/gu/examples/hello/apps"
	"honnef.co/go/js/dom"
)

func main() {

	New(&DOMRenderer{
		Document: dom.GetWindow().Document(),
	}).Init()

}
