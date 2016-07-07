package main

import (
	"bytes"
	"fmt"

	"github.com/influx6/gu"
	"github.com/influx6/gu/gucss"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/gutrees/styles"
	"github.com/influx6/gu/guviews"
	"honnef.co/go/js/dom"
)

type backgrounds []*background

func (b backgrounds) Resolve(path gudispatch.Path) {
	for _, bg := range b {
		bg.Resolve(path)
	}
}

func (b backgrounds) Render() gutrees.Markup {
	root := elems.Div(attrs.Class("backgrounds"))

	for _, bg := range b {
		bg.Render().Apply(root)
	}

	return root
}

type background struct {
	gudispatch.Resolver
	color string
	show  bool
}

func newBackground(pattern string, color string) *background {
	bg := &background{
		color:    color,
		Resolver: gudispatch.NewResolver(pattern),
	}

	bg.Subscribe(func(ps gudispatch.Path) {
		fmt.Printf("Path: %+s -> %s\n", ps, color)
		if !bg.show {
			bg.show = false
			return
		}
		bg.show = true
	})

	return bg
}

// Render returns the markup for the background box.
func (b *background) Render() gutrees.Markup {
	return elems.Div(
		gutrees.When(b.show, styles.Display("block"), styles.Display("none")),
		attrs.Class("box"),
		styles.Background(b.color),
		styles.Width("100%"),
		styles.Height("100%"),
	)
}

// css takes a document and a gucss.Render, creating the need dom element
// with the content of the rendered css stylesheet.
func makeCSS(doc dom.Document, css gucss.Render) {
	styl := doc.CreateElement("style")
	doc.QuerySelector("head").AppendChild(styl)

	var data bytes.Buffer
	css.Render(&data)

	styl.SetTextContent(string(data.Bytes()))
}

func main() {
	css := gucss.NewRoot()

	css.Child("*", gucss.Properties{
		"padding": "0px",
		"margin":  "0px",
	})

	gucss.Extend(css, gucss.Properties{
		"width":  "100%",
		"Height": "100%",
	}, "html", "html body", ".backgrounds")

	doc := dom.GetWindow().Document()
	makeCSS(doc, css)

	colors := []string{"red", "black", "cyan"}

	var bgs backgrounds

	for i := 0; i < 3; i++ {
		bgs = append(bgs, newBackground(colors[i], colors[i]))
	}

	bgView := guviews.New(bgs)
	gu.RenderAsBody(bgView)

	// Register the view to watch for Hash changes and also
	// when it fails to match, it should hide.
	guviews.AttachHash("/colors/*", bgView, guviews.WrapNormal(bgView.Hide))
}
