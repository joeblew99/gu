package guviews

import (
	"fmt"
	"strings"
	"testing"

	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
)

var success = "\u2713"
var failed = "\u2717"

var treeRenderlen = 272

type videoList []map[string]string

func (v videoList) Render(m ...string) gutrees.Markup {
	dom := elems.Div()
	for _, data := range v {
		gutrees.Augment(dom, elems.Video(
			attrs.Src(data["src"]),
			elems.Text(data["name"]),
		))
	}
	return dom
}

func TestView(t *testing.T) {
	videos := NewView(videoList([]map[string]string{
		map[string]string{
			"src":  "https://youtube.com/xF5R32YF4",
			"name": "Joyride Lewis!",
		},
		map[string]string{
			"src":  "https://youtube.com/dox32YF4",
			"name": "Wonderlust Bombs!",
		},
	}))

	bo := videos.RenderHTML()

	if len(bo) != treeRenderlen {
		t.Fatalf("\t%s\t Rendered result with invalid length, expected %d but got %d -> \n %s", failed, treeRenderlen, len(bo), bo)
	}

	t.Logf("\t%s\t Rendered result accurated with length %d", success, treeRenderlen)
}

type item string

func (i item) Render(m ...string) gutrees.Markup {
	return elems.Span(elems.Text(fmt.Sprintf("+ %s", i)))
}

func TestSequenceView(t *testing.T) {
	items := SequenceView(SequenceMeta{Tag: "div"}, item("Book"), item("Funch"), item("Fudder"))

	out := string(items.RenderHTML())

	if !strings.Contains(out, "+ Book") {
		t.Errorf("\t%s\tShould contain %q inside rendered output", failed, "+ Book")
	}

	if !strings.Contains(out, "+ Book") {
		t.Errorf("\t%s\tShould contain %q inside rendered output", failed, "+ Funch")
	}

	if !strings.Contains(out, "+ Book") {
		t.Errorf("\t%s\tShould contain %q inside rendered output", failed, "+ Fudder")
	}

	t.Logf("\t%s\tShould contain %q inside rendered output", success, []string{"+ Book", "+ Funch", "+ Fudder"})
}
