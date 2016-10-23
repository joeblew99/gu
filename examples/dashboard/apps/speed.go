package apps

import (
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/trees/elems"
)

// _ defines the resource page associated with the speed resource.
var _ = Resource(func() {

	Scripts("https://code.jquery.com/jquery-3.1.1.js", "text/javascript", false)
	CSS("https://raw.githubusercontent.com/csswizardry/csswizardry-grids/master/csswizardry-grids.scss", false)

	Markup(Header1(Text("Speed Dashboard")), "", false)

	Markup(func() Viewable {
		return Div(
			Section(
				Label(Text("Current Speed")),
			),
		)
	}, "", false)

})
