package apps

import (
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/elems"
)

type feel struct{ Color string }

var _ = Resource(func() {

	DoTitle("Dashboard App")

	DoScript("https://code.jquery.com/jquery-3.1.1.js", "text/javascript", false)
	DoScript("https://code.jquery.com/jquery-ui.js", "text/javascript", true)
	DoCSS("https://raw.githubusercontent.com/csswizardry/csswizardry-grids/master/csswizardry-grids.scss", false)

	DoMarkup(Header1(Text("Dashboard App")), "", false)

	DoMarkup(func() *Markup {
		return Div(
			CSS(`
				section{
					width: 50%;
					height: 50%;
					margin: 0 auto;
					background: {{ .Color }};
				}
			`, feel{Color: "#eee"}, "#hello-wrapper"),
			Section(
				Label(Text("Hello")),
			),
		)
	}, "", false)

})
