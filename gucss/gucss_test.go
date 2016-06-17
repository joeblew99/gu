package gucss_test

import (
	"bytes"
	"testing"

	"github.com/influx6/gu/gucss"
)

var (
	success = "\u2713"
	failed  = "\u2717"
)

// TestBasicCSS validates the behaviour of the gucss generation parser.
func TestBasicCSS(t *testing.T) {

	var expected = []byte("@media (width: 30em) and (height: 40em) screen {\n* {\n\tmargin: 0px;\n\tpadding: 0px;\n}\n\nhtml {\n\twidth: 100%;\n}\n\nhtml body {\n\twidth: 100%;\n\tfont-size: 1em;\n}\n\ndiv {\n}\n\ndiv ul {\n\tlist-style-type: none;\n}\n\ndiv ul, div ol {\n\tlist-style-type: none;\n}\n\ndiv ul ~ li {\n\tpadding: 10px;\n}\n\ndiv ul ~ li + span {\n\tpadding: 3px;\n}\n\ndiv ul ~ li > a {\n\tcolor: #ccc;\n}\n\ndiv ul:hover  {\n\tpadding: 30px;\n}\n\n}\n\n")

	t.Logf("Given the need to generate css styles")
	{
		t.Logf("\tWhen giving a gucss.RootGroup")
		{
			root := gucss.NewRoot()

			root.Within("*", gucss.Properties{
				"margin":  "0px",
				"padding": "0px",
			})

			root.Within("html", gucss.Properties{
				"width": "100%",
			}).Within("body", gucss.Properties{
				"width":     "100%",
				"font-size": "1em",
			})

			root.Child("div", gucss.Properties{}).Within("ul", gucss.Properties{
				"list-style-type": "none",
			}).Extend("ol", gucss.Properties{
				"list-style-type": "none",
			}).PreSibling("li", gucss.Properties{
				"padding": "10px",
			}).PostSibling("span", gucss.Properties{
				"padding": "3px",
			}).NthParent(1).Child("a", gucss.Properties{
				"color": "#ccc",
			}).NthParent(2).NS(":hover", gucss.Properties{
				"padding": "30px",
			})

			var b bytes.Buffer
			gucss.NewMedia("screen", "(width: 30em) and (height: 40em)").Render(root, &b)

			if b.Len() != len(expected) {
				t.Logf("Expected: %q", expected)
				t.Logf("Recieved: %q", b.Bytes())
				t.Fatalf("\t%s\t Expected CSS Output to be equal with expected", failed)
			}
			t.Logf("\t%s\t Expected CSS Output to be equal with expected", success)

		}
	}
}
