package css_test

import (
	"testing"

	"github.com/influx6/gu/css"
	"github.com/influx6/gu/tests"
)

func TestBasicCSS(t *testing.T) {
	expected := "#galatica:hover {\n  color: red;\n}\n#galatica::before {\n  content: \"bugger\";\n}\n#galatica div a {\n  color: black;\n  font-family: Helvetica;\n}\n@media (max-width: 400px) {\n  #galatica:hover {\n    color: blue;\n    font-family: Helvetica;\n  }\n}"

	csr := css.New(`

    $:hover {
      color: red;
    }

    $::before {
      content: "bugger";
    }

    $ div a {
      color: black;
      font-family: {{ .Font }}
    }

    @media (max-width: 400px){

      $:hover {
        color: blue;
        font-family: {{ .Font }}
      }

    }
`)

	sheet, err := csr.Stylesheet(struct {
		Font string
	}{Font: "Helvetica"}, "#galatica")

	if err != nil {
		tests.Failed(t, "Should have successfully processed stylesheet for rule")
	}
	tests.Passed(t, "Should have successfully processed stylesheet for rule")

	if val := sheet.String(); val != expected {
		t.Logf("\t\tRecieved: %q\n", val)
		t.Logf("\t\tExpected: %q\n", expected)
		tests.Failed(t, "Should have rendered expected stylesheet")
	}
	tests.Passed(t, "Should have rendered expected stylesheet")
}

func TestLinkedCSS(t *testing.T) {
	expected := "#galatica block {\n  Helvetica\n      color: Pink;\n}\n#galatica::before {\n  content: \"bugger\";\n}\n#galatica div a {\n  color: black;\n  font-family: Helvetica;\n}\n@media (max-width: 400px) {\n  #galatica:hover {\n    color: blue;\n    font-family: Helvetica;\n  }\n}"

	csr := css.New(`
    block {
      font-family: {{ .Font }}
      color: {{ .Color }}
    }
  `)

	csx := css.New(`

    ::before {
      content: "bugger";
    }

    div a {
      color: black;
      font-family: {{ .Font }}
    }

    @media (max-width: 400px){

      :hover {
        color: blue;
        font-family: {{ .Font }}
      }

    }
`, csr)

	sheet, err := csx.Stylesheet(struct {
		Font  string
		Color string
	}{
		Font:  "Helvetica",
		Color: "Pink",
	}, "#galatica")

	if err != nil {
		tests.Failed(t, "Should have successfully processed stylesheet for rule")
	}
	tests.Passed(t, "Should have successfully processed stylesheet for rule")

	if val := sheet.String(); val != expected {
		t.Logf("\t\tRecieved: %q\n", val)
		t.Logf("\t\tExpected: %q\n", expected)
		tests.Failed(t, "Should have rendered expected stylesheet")
	}
	tests.Passed(t, "Should have rendered expected stylesheet")
}
