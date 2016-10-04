package gutrees_test

import (
	"testing"

	"github.com/influx6/gu/gutrees"
)

// TestFinders validates the utility functions which help in retrieving underline
// markup.
func TestFinders(t *testing.T) {
	body := generateMarkup()

	if _, err := gutrees.GetStyle(body, "width"); err != nil {
		t.Fatalf("\t%s\t  Should have been able to find style in markup", failed)
	}
	t.Logf("\t%s\t  Should have been able to find style in markup", success)

	if _, err := gutrees.GetAttr(body, "id"); err != nil {
		t.Fatalf("\t%s\t  Should have been able to find attr in markup", failed)
	}
	t.Logf("\t%s\t  Should have been able to find attr in markup", success)

	if len(gutrees.GetStyles(body, "width", "auto")) < 1 {
		t.Fatalf("\t%s\t  Should have been able to find style in markup with value of auto and name of width", failed)
	}
	t.Logf("\t%s\t  Should have been able to find style in markup with value of auto and name of width", success)

	if len(gutrees.ElementsUsingStyle(body, "width", "")) > 2 {
		t.Fatalf("\t%s\t  Should have been able to found more than two elements using width style", failed)
	}
	t.Logf("\t%s\t  Should have been able to found more than two elements using width style", success)

	if len(gutrees.ElementsUsingStyle(body, "id", "")) > 2 {
		t.Fatalf("\t%s\t  Should have been able to found more than two elements using id attr", failed)
	}
	t.Logf("\t%s\t  Should have been able to found more than two elements using id attr", success)

	if len(gutrees.ElementsUsingStyle(body, "width", "200px")) > 2 {
		t.Fatalf("\t%s\t  Should have been able to found more than two elements using width style with value of 200px", failed)
	}
	t.Logf("\t%s\t  Should have been able to found more than two elements using width style with value of 200px", success)

	if len(gutrees.ElementsWithTag(body, "div")) < 2 {
		t.Fatalf("\t%s\t  Should have been able to found two div element", failed)
	}
	t.Logf("\t%s\t  Should have been able to found two div element", success)

	if len(gutrees.ElementsWithTag(body, "label")) != 1 {
		t.Fatalf("\t%s\t  Should have been able to find a label element", failed)
	}
	t.Logf("\t%s\t  Should have been able to find a label element", success)
}

func generateMarkup() gutrees.Markup {
	body := gutrees.NewElement("body", false)
	gutrees.NewStyle("width", "auto").Apply(body)
	gutrees.NewAttr("id", "main-wrapper").Apply(body)

	root := gutrees.NewElement("div", false)
	gutrees.NewAttr("id", "root-div").Apply(root)
	gutrees.NewAttr("class", "roots").Apply(root)
	gutrees.NewStyle("width", "100px").Apply(root)
	gutrees.NewElement("section", false).Apply(root)
	root.Apply(body)

	root2 := gutrees.NewElement("div", false)
	gutrees.NewAttr("id", "root-div-2").Apply(root2)
	gutrees.NewAttr("class", "roots").Apply(root2)
	root2.Apply(body)

	label := gutrees.NewElement("label", false)
	gutrees.NewStyle("width", "200px").Apply(label)
	label.Apply(root2)

	return body
}
