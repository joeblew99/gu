package resources_test

import (
	"testing"

	"github.com/influx6/gu/design"
	"github.com/influx6/gu/tests"
	"github.com/influx6/gu/trees"
	"github.com/influx6/gu/trees/elems"
)

var testRes = design.New()

func TestResource(t *testing.T) {
	_ = design.Resource(func() {

		design.Order(design.Any)
		design.UseRoute("/home")

		design.CSS("../some.css", false)
		design.Scripts("../some.js", "text/javascript", false)

		design.Markup(elems.Header1(elems.Text("Speed Dashboard")), "", false)

		design.Markup(func() trees.Markup {
			return elems.Div(
				elems.Section(
					elems.Label(elems.Text("Current Speed")),
				),
			)
		}, "", false)

	})

	master := design.GetCurrentResources()
	root := master.MustCurrentResource()

	root.Init()

	if root.Order != design.Any {
		tests.Failed(t, "Should have 'Any' RenderingOrder")
	}
	tests.Passed(t, "Should have 'Any' RenderingOrder")

	if len(root.Links) != 2 {
		tests.Failed(t, "Should have added two links into the resource")
	}
	tests.Passed(t, "Should have added two links into the resource")

	if len(root.Markups) != 2 {
		tests.Failed(t, "Should have added two markups into the resource")
	}
	tests.Passed(t, "Should have added two markups into the resource")

	if _, _, state := root.Resolver.Test("/home"); !state {
		tests.Failed(t, "Should have matched given route: %q", "/home")
	}
	tests.Passed(t, "Should have matched given route: %q", "/home")

	if name := root.Links[0].Content.Name(); name != "link" {
		tests.Failed(t, "Should have added a style tag static view as the second item: %q", name)
	}
	tests.Passed(t, "Should have added a style tag static view as the first item")

	if name := root.Markups[0].View.Content.Name(); name != "h1" {
		tests.Failed(t, "Should have added a h1(header) tag static view as the second item: %q", name)
	}
	tests.Passed(t, "Should have added a style tag static view as the second item")

	if name := root.Markups[1].View.Content.Name(); name != "div" {
		tests.Failed(t, "Should have added a div tag static view as the second item: %q", name)
	}
	tests.Passed(t, "Should have added a style tag static view as the second item")

}
