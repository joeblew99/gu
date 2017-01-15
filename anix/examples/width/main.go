package main

import (
	"github.com/gu-io/gu/anix"
	"honnef.co/go/js/dom"
)

func main() {
	win := dom.GetWindow()
	doc := win.Document()

	tick := anix.Animate(Width(400))
	tick(doc.QuerySelector("div"))

	tick.Start()
}
