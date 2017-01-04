package main

import (
	"github.com/gu-io/gu"
	_ "github.com/gu-io/gu/examples/hello/apps"
	"github.com/gu-io/gu/redom"
	"honnef.co/go/js/dom"
)

func main() {

	gu.New("hello.v1", &redom.DOMRenderer{
		Document: dom.GetWindow().Document(),
	}).Init(true)

}
