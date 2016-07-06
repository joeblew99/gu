package main

import (
	"github.com/influx6/gu"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/gutrees/styles"
	"github.com/influx6/gu/guviews"
)

// Background defines a background component which displays different colors.
type Background struct {
	guviews.Resolvers
	color string
}

// Render returns the markup for the background box.
func (b Background) Render() gutrees.Markup {
	return elems.Div(
		attrs.Class("box"),
		styles.Background(color),
		styles.Width("100%"),
		styles.Height("100%"),
	)
}

func main() {
	// doc := dom.GetWindow().Document()

	var bgs []Background

	colors := []string{"red", "black", "cyan"}

	for i := 0; i < 3; i++ {
		bgs = append(bgs, Background{
			Resolvers: guviews.NewResolver("red"),
			color:     colors[i],
		})
	}

	bgView := guviews.New(bgs...)

	// Register the view to watch for Hash changes and also
	// when it fails to match, it should hide.
	guviews.AttachHash("/colors/:color", bgView, bgView.Hide)

	gu.RenderAsBody(bgView)
}
