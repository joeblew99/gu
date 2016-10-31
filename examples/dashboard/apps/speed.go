package apps

import (
	"github.com/gopherjs/gopherjs/js"
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/attrs"
	. "github.com/influx6/gu/trees/elems"
	"github.com/influx6/gu/trees/events"
)

type feel struct {
	Color      string
	Background string
}

var _ = Resource(func() {

	DoTitle("Dashboard App")

	DoScript("https://code.jquery.com/jquery-3.1.1.js", "text/javascript", false)
	DoScript("https://code.jquery.com/jquery-ui.js", "text/javascript", true)
	DoCSS("https://raw.githubusercontent.com/csswizardry/csswizardry-grids/master/csswizardry-grids.scss", false)

	DoMarkup(Header1(Text("Dashboard App")), "", false)

	DoMarkup(func() *Markup {
		return Div(
			ID("hello"),
			CSS(`
				${
					width: 500px;
					height: 500px;
				}

				$:hover{
					background: {{ .Background }};
				}

				section.super{
					width: 50%;
					height: 50%;
					margin: 0 auto;
					background: {{ .Color }};
				}

				@media (max-width: 400px){

					$:hover {
						background: red;
					}

				}
			`, feel{
				Color:      "#eee",
				Background: "pink",
			}),
			Section(
				events.Click(func(ev EventObject, root *Markup) {
					js.Global.Call("alert", "Am clicked!")
				}, ""),
				Class("super"),
				Label(Text("Hello")),
			),
		)
	}, "", false)

})
