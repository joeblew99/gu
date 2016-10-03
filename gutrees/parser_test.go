package gutrees_test

import (
	"strings"
	"testing"

	"github.com/influx6/gu/gutrees"
)

var success = "\u2713"
var failed = "\u2717"

func TestParser(t *testing.T) {
	gutrees.SetMode(gutrees.Testing)
	defer gutrees.SetMode(gutrees.Normal)

	result, err := gutrees.ParseTree(`<div class="racket" id="racket-wrapper">
            <a href="#" rel="bounce postive">Bounce +</a>
          </div>`)

	if err != nil {
		t.Fatalf("\t%s\t Parser should have produced markup for html: %q", failed, err)
	}

	var html []string

	for _, res := range result {
		html = append(html, res.HTML())
	}

	t.Logf("\t%s\t Parser should have produced markup for html: %s", success, strings.Join(html, ""))
}
