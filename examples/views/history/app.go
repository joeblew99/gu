package main

import (
	"bytes"
	"math"
	"time"

	"github.com/influx6/gu/gucss"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/guviews"
	"honnef.co/go/js/dom"
)

// StatItem defines a component which defines a stat item on a dashboard.
type StatItem struct {
	guviews.Reactive
	Name  string
	Value int
}

// Set sets the value of the StatItem to the provided value.
func (s *StatItem) Set(d int) {
	s.Value = d
	s.Publish()
}

// NewStatItem returns a new instance of a StatItem.
func NewStatItem(name string, val int) *StatItem {
	return &StatItem{
		Reactive: guviews.NewReactive(),
		Name:     name,
		Value:    val,
	}
}

// Render returns the markup for a StatItem.
func (s *StatItem) Render() gutrees.Markup {
	return elems.Section(
		attrs.Class("stat"),
		elems.Label(
			attrs.Class("statDetail header"),
			elems.Header3(elems.Text(s.Name)),
		),
		elems.Label(
			attrs.Class("statDetail value"),
			elems.Text(s.Value),
		),
	)
}

// Dashboard defines a dashboard component which displays different stats.
type Dashboard struct {
	guviews.Reactive
	stats []*StatItem
}

// NewDashboard returns a new instance of a Dashboard.
func NewDashboard() *Dashboard {
	return &Dashboard{Reactive: guviews.NewReactive()}
}

// Add adds a new stat into the list of stats displayed by the dashboard
func (d *Dashboard) Add(stat *StatItem) {
	d.stats = append(d.stats, stat)
	stat.Subscribe(d.Publish)
}

// Render returns the markup for the Dashboard component.
func (d *Dashboard) Render() gutrees.Markup {
	root := elems.Div(attrs.Class("dashboard"))

	for _, stat := range d.stats {
		stat.Render().Apply(root)
	}

	return root
}

func wave(factor float64, val int) float64 {
	vm := math.Max(float64(val), 2.5)
	return vm + (math.Sin(math.Tan(factor)*math.Pow(math.Sin(factor), 10)) * vm)
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
	doc := dom.GetWindow().Document()

	css := gucss.NewRoot()

	css.Child("*", gucss.Properties{
		"padding": "0px",
		"margin":  "0px",
	})

	gucss.Extend(css, gucss.Properties{
		"width": "100%",
	}, "html", "html body", ".dashboard")

	css.Child(".dashboard", gucss.Properties{
		"height": "400px",
	}).Child(".stat", gucss.Properties{
		"display":   "inline-block",
		"width":     "30%",
		"color":     "#000",
		"border":    "1px dotted rgba(20,20,20,0.5)",
		"padding":   "30px",
		"font-size": "3em",
	})

	makeCSS(doc, css)

	dashboard := NewDashboard()

	visitors := NewStatItem("Visitor's Count", 200)
	clickRate := NewStatItem("Click Rates", 100)
	bounceRate := NewStatItem("Bounce Rate", 50)

	dashboard.Add(visitors)
	dashboard.Add(clickRate)
	dashboard.Add(bounceRate)

	dashboardView := guviews.New(dashboard)
	dashboardView.Mount(doc.QuerySelector("body").Underlying())

	factor := 0.05

	go func() {
		for {
			<-time.After(300 * time.Millisecond)
			visitors.Set(int(wave(factor, visitors.Value)))
			clickRate.Set(int(wave(factor, clickRate.Value)))
			bounceRate.Set(int(wave(factor, bounceRate.Value)))
			factor += 0.1
		}
	}()

}
