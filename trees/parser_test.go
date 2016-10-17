package trees_test

import (
	"strings"
	"testing"

	"github.com/influx6/gu/trees"
)

var success = "\u2713"
var failed = "\u2717"

func TestParser(t *testing.T) {
	trees.SetMode(trees.Testing)
	defer trees.SetMode(trees.Normal)

	result, err := trees.ParseTree(`
		<!doctype html>
		<html>
			<head></head>
			<body>
				<div class="racket" id="racket-wrapper">
		      <a href="#" rel="bounce postive">Bounce +</a>
		    </div>

		    <!--thertorial words-->

				<div class="racket" id="racket-wrapper-2">
		      <a href="#" rel="bounce negative">Bounce -</a>
		    </div>
			</body>
		</html>
  `)

	if err != nil {
		t.Fatalf("\t%s\t Parser should have produced markup for html: %q", failed, err)
	}

	var html []string

	for _, res := range result {
		html = append(html, res.HTML())
	}

	t.Logf("\t%s\t Parser should have produced markup for html: %s", success, strings.Join(html, ""))
}
