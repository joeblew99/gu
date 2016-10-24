package main

import (
	. "github.com/influx6/gu/design"
	_ "github.com/influx6/gu/examples/dashboard/apps"
	. "github.com/influx6/gu/render/dom"
	"honnef.co/go/js/dom"
)

// DashboardApp defines the central application resource renderer.
var DashboardApp = NewResources()

func main() {
	DashboardApp.Init()

	New(DashboardApp, &DOMResourceRender{
		Document: dom.GetWindow().Document(),
	}).Begin()

}
