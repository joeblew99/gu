package main

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/examples/greeter/components"
	"github.com/gu-io/gu/trees/elems"
)

func main() {
	app := gu.App(AppSettings)

	index := app.View(gu.ViewAttr{
		Name:  "Greeter.Index",
		Route: "/",
		Base: elems.Parse(`
      <div class="greeter-view view wrapper">
        <h1 class="view-header">Greeter App<h1>

        <div class="greeter-app" id="greeter-app-component">
        </div>
      </div>
    `),
	})

	index.Component(gu.ComponentAttr{
		Route:  "*",
		Target: "#greeter-app-component",
		Base:   components.NewGreeter(),
	})
}
