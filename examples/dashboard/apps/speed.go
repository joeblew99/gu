package apps

import (
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/elems"
)

// _ defines the resource page associated with the speed resource.
var _ = Resource(func() {

	DoTitle("Dashboard App")

	DoScript("https://code.jquery.com/jquery-3.1.1.js", "text/javascript", false)
	DoScript("https://code.jquery.com/jquery-ui.js", "text/javascript", true)
	DoCSS("https://raw.githubusercontent.com/csswizardry/csswizardry-grids/master/csswizardry-grids.scss", false)

	DoMarkup(Header1(Text("Dashboard App")), "", false)

	DoMarkup(func() *Markup {
		return Div(
			Style(
				Text("* { font-size: 1em; }"),
			),
			Section(
				Label(Text("Hello")),
			),
		)
	}, "", false)

})
