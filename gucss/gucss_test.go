package gucss_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/influx6/gu/gucss"
)

var (
	success = "\u2713"
	failed  = "\u2717"
)

// TestBasicCSS validates the behaviour of the gucss generation parser.
func TestBasicCSS(t *testing.T) {

	var expected = []byte(`* {
	margin: 0px;
	padding: 0px;
}

html {
	width: 100%;
}

html body {
	font-size: 1em;
	width: 100%;
}

div {
}

div ul {
	list-style-type: none;
}

div ul, div ol {
	list-style-type: none;
}

div ul ~ li {
	padding: 10px;
}

div ul ~ li + span {
	padding: 3px;
}

div ul ~ li > a {
	color: #ccc;
}

div ul:hover  {
	padding: 30px;
}

`)

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
			root.Render(&b)
			fmt.Printf("%s", b.Bytes())

			if !bytes.Equal(expected, b.Bytes()) {
				t.Logf("Expected: %q", expected)
				t.Logf("Recieved: %q", b.Bytes())
				t.Fatalf("\t%s\t Expected CSS Output to be equal with expected", failed)
			}
			t.Logf("\t%s\t Expected CSS Output to be equal with expected", success)

		}
	}
}
