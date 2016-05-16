package guviews_test

import (
	"testing"

	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/guviews"
)

var success = "\u2713"
var failed = "\u2717"

var treeRenderlen = 275

type videoList []map[string]string

func (v videoList) Render() gutrees.Markup {
	dom := elems.Div()
	for _, data := range v {
		elems.Video(
			attrs.Src(data["src"]),
			elems.Text(data["name"]),
		).Apply(dom)
	}
	return dom
}

func TestView(t *testing.T) {
	gutrees.SetMode(gutrees.Testing)
	defer gutrees.SetMode(gutrees.Normal)

	var expectedDOM = `<div style=" display:block;"><video src="https://youtube.com/xF5R32YF4" style="">Joyride Lewis!</video><video src="https://youtube.com/dox32YF4" style="">Wonderlust Bombs!</video></div>`

	videos := guviews.NewWithID("video-vabbs", videoList([]map[string]string{
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

	if string(bo) != expectedDOM {
		t.Logf("Given:  %s\n", bo)
		t.Logf("Wanted: %s\n", expectedDOM)
		t.Fatalf("\t%s\t Rendered result does match expected", failed)
	}

	t.Logf("\t%s\t Rendered result does match expected", success)
}
