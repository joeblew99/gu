package resources_test

import (
	"testing"

	"github.com/influx6/gu"
	"github.com/influx6/gu/design"
	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/tests"
	"github.com/influx6/gu/trees"
	"github.com/influx6/gu/trees/elems"
)

func TestResource(t *testing.T) {
	trees.SetMode(trees.Testing)
	defer trees.SetMode(trees.Normal)

	_ = design.Resource(func() {

		design.DoOrder(design.Any)
		design.UseRoute("/home")

		design.DoCSS("../some.css", false)
		design.DoScript("../some.js", "text/javascript", false)

		design.DoMarkup(elems.Header1(elems.Text("Speed Dashboard")), "", false)

		design.DoMarkup(func() *trees.Markup {
			return elems.Div(
				elems.Section(
					elems.Label(elems.Text("Total Current Speed")),
				),
			)
		}, "", false)

	})

	master := design.New()
	master.Init()

	root := master.MustCurrentResource()

	if root.Order != design.Any {
		tests.Failed(t, "Should have 'Any' RenderingOrder")
	}
	tests.Passed(t, "Should have 'Any' RenderingOrder")

	if len(root.Links) != 2 {
		tests.Failed(t, "Should have added two links into the resource")
	}
	tests.Passed(t, "Should have added two links into the resource")

	if len(root.Renderables) != 2 {
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

	if name := root.Renderables[0].View.Render().Name(); name != "h1" {
		tests.Failed(t, "Should have added a h1(header) tag static view as the second item: %q", name)
	}
	tests.Passed(t, "Should have added a style tag static view as the second item")

	if name := root.Renderables[1].View.Render().Name(); name != "div" {
		tests.Failed(t, "Should have added a div tag static view as the second item: %q", name)
	}
	tests.Passed(t, "Should have added a style tag static view as the second item")
}

var flatRender = "<html data-gen=\"gu\">\n<head data-gen=\"gu\">\n</head>\n<body data-gen=\"gu\">\n</body>\n</html>\n"
var fullRender = "<html data-gen=\"gu\">\n<head data-gen=\"gu\">\n<link data-gen=\"gu\"  href=\"../some.css\"  rel=\"stylesheet\"  type=\"text/css\">\n</link>\n<script data-gen=\"gu\"  src=\"../some.js\"  type=\"text/javascript\">\n</script>\n</head>\n<body data-gen=\"gu\">\n<h1 data-gen=\"gu\"  resource-id=\"2-3bce4931-6c75-41ab-afe0-2ec108a30860\">\nSpeed Dashboard</h1>\n<div data-gen=\"gu\"  resource-id=\"2-3bce4931-6c75-41ab-afe0-2ec108a30860\">\n<section data-gen=\"gu\">\n<label data-gen=\"gu\">\nCurrent Speed</label>\n</section>\n</div>\n<div data-gen=\"gu\">\n</div>\n</body>\n</html>\n"
var routedRender = "<html data-gen=\"gu\">\n<head data-gen=\"gu\">\n<link data-gen=\"gu\"  href=\"../some.css\"  rel=\"stylesheet\"  type=\"text/css\">\n</link>\n<script data-gen=\"gu\"  src=\"../some.js\"  type=\"text/javascript\">\n</script>\n</head>\n<body data-gen=\"gu\">\n<h1 data-gen=\"gu\"  resource-id=\"2-3bce4931-6c75-41ab-afe0-2ec108a30860\">\nSpeed Dashboard</h1>\n<div data-gen=\"gu\"  resource-id=\"2-3bce4931-6c75-41ab-afe0-2ec108a30860\">\n<section data-gen=\"gu\">\n<label data-gen=\"gu\">\nCurrent Speed</label>\n</section>\n</div>\n<div data-gen=\"gu\">\n<section data-gen=\"gu\">\n<label data-gen=\"gu\">\nTotal Speed</label>\n</section>\n</div>\n</body>\n</html>\n"

func TestResourceRendering(t *testing.T) {
	trees.SetMode(trees.Testing)
	defer trees.SetMode(trees.Normal)

	_ = design.Resource(func() {
		design.DoOrder(design.Any)
		design.UseRoute("/home/*")

		design.DoCSS("../some.css", false)
		design.DoScript("../some.js", "text/javascript", false)

		design.DoMarkup(elems.Header1(elems.Text("Speed Dashboard")), "", false)

		design.DoMarkup(func() *trees.Markup {
			return elems.Div(
				elems.Section(
					elems.Label(elems.Text("Current Speed")),
				),
			)
		}, "", false)

		design.DoView(gu.Static(elems.Div(
			elems.Section(
				design.DoRoute("/models/*", "/:speed"),
				elems.Label(elems.Text("Total Speed")),
			),
		)), "", false)

	})

	master := design.New().Init()

	testRender := master.Render(dispatch.UseLocation("/home/"))
	if testRender.HTML() != fullRender {
		t.Logf("\t\tExpected: %q", fullRender)
		t.Logf("\t\tReceived: %q", testRender.HTML())
		tests.Failed(t, "Should have rendered the expected full markup")
	}
	tests.Passed(t, "Should have rendered the expected full markup")

	testRender2 := master.Render(dispatch.UseLocation("/home/models/340mph"))
	if testRender2.HTML() != routedRender {
		t.Logf("\t\tExpected: %q", routedRender)
		t.Logf("\t\tReceived: %q", testRender2.HTML())
		tests.Failed(t, "Should have rendered the routed full markup")
	}
	tests.Passed(t, "Should have rendered the routed full markup")

	testFlatRender := master.Render(dispatch.UseLocation("/root"))
	if testFlatRender.HTML() != flatRender {
		t.Logf("\t\tExpected: %q", flatRender)
		t.Logf("\t\tReceived: %q", testFlatRender.HTML())
		tests.Failed(t, "Should have rendered the expected markup")
	}
	tests.Passed(t, "Should have rendered the expected markup")

}
