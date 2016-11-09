package main

import (
	. "github.com/influx6/gu/design"
	_ "github.com/influx6/gu/examples/hello/apps"
	. "github.com/influx6/gu/redom"
	"honnef.co/go/js/dom"
)

func main() {

	New(&DOMRenderer{
		Document: dom.GetWindow().Document(),
	}).Init()

}
