package main

import (
	. "github.com/gu-io/gu/design"
	_ "github.com/gu-io/gu/examples/hello/apps"
	. "github.com/gu-io/gu/redom"
	"honnef.co/go/js/dom"
)

func main() {

	New(&DOMRenderer{
		Document: dom.GetWindow().Document(),
	}).Init(true)

}
