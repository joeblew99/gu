package main

import (
	"fmt"
	"time"

	"github.com/influx6/gu"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/guviews"
)

// SpeedCounter defines a Speedometer component.
type SpeedCounter struct {
	Speed float64
}

// Increment adds the giving value to the speed.
func (sc *SpeedCounter) Increment(m float64) {
	sc.Speed += m
}

// Render renders the markup for the speed counter.
func (sc *SpeedCounter) Render() gutrees.Markup {
	root := elems.Div(
		elems.Label(attrs.Class("speed-title"), elems.Text("Speed ")),
		elems.Label(attrs.Class("speed-value"), elems.Text(fmt.Sprintf("%.4f mps", sc.Speed))),
	)
	return root
}

// Dashbaord defines a Dashboard component.
type Dashbaord struct {
	Speed *SpeedCounter
}

// Render renders the markup for the dashboard.
func (sc *Dashbaord) Render() gutrees.Markup {
	root := elems.Div(
		elems.Section(
			elems.Header2(elems.Text("Speedometer")),
			sc.Speed.Render(),
		),
	)

	return root
}

func main() {

	dashy := &Dashbaord{
		Speed: &SpeedCounter{Speed: 40.072},
	}

	dashyView := guviews.New(dashy)

	gu.RenderAsBody(dashyView)

	for {
		time.Sleep(4 * time.Second)

		// Increase the speed count.
		dashy.Speed.Increment(40.0)

		// Notify the view of new update.
		gudispatch.Dispatch(&guviews.ViewUpdate{ID: dashyView.UUID()})
	}

}
