package main

import (
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/guviews"
)

// menu defines a map of name and url links.
type menu map[string]string

func (v menu) Render() gutrees.Markup {
	dom := elems.Div(attrs.ID("bat"))

	for name, src := range v {
		gutrees.Augment(dom, elems.Image(
			attrs.Src(src),
			attrs.Rel(name),
		))
	}

	return dom
}

func main() {

	guviews.Register("video-components/vabbs", func(links map[string]string) guviews.Renderable {
		return menu(links)
	})

	// When deployed either on the frontend or backend, the view is able
	// to reconciled the already rendered content from the backend and
	// take over.
	guviews.MustCreate(guviews.ViewConfig{
		Name: "video-components/vabbs",
		ID:   "02-vabbs",
		Elem: "body",
		Param: map[string]string{
			"Joyride": "https://pbs.twimg.com/media/CbtrJu9UAAAXhSs.jpg",
			"Bombs":   "https://pbs.twimg.com/media/Cbs8Ia0VAAAYyfR.jpg",
		},
	})
}
