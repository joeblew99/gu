package main

import (
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/guviews"
)

// StatItem defines a component which defines a stat item on a dashboard.
type StatItem struct {
	guviews.Reactive
	Name  string
	Value interface{}
}

// NewStatItem returns a new instance of a StatItem.
func NewStatItem(name string, val interface{}) *StatItem {
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
	stats []StatItem
}

// NewDashboard returns a new instance of a Dashboard.
func NewDashboard() *Dashboard {
	return &Dashboard{Reactive: guviews.NewReactive()}
}

// Add adds a new stat into the list of stats displayed by the dashboard
func (d Dashboard) Add(stat StatItem) {
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

func main() {
	dashboard := NewDashboard()

	visitors := NewStatItem("Visitor's Count", 200)
	clickRate := NewStatItem("Click Rates", 100)
	bounceRate := NewStatItem("Bounce Rate", 100)

	dashboard.Add(visitors)
	dashboard.Add(clickRate)
	dashboard.Add(bounceRate)
}
