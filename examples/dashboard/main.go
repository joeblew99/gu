package main

import (
	. "github.com/influx6/gu/design"
	_ "github.com/influx6/gu/examples/dashboard/apps"
)

// DashboardApp defines the central application resource renderer.
var DashboardApp = New()

func main() {
	DashboardApp.Init()
}
