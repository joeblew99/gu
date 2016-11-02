package apps

import (
	"github.com/gopherjs/gopherjs/js"
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/attrs"
	. "github.com/influx6/gu/trees/elems"
	events "github.com/influx6/gu/trees/events"
)

var _ = Resource(func() {

	DoTitle("Hello App")
	DoMarkup(func() *Markup {
		return Div(
			CSS(`
				${
					width:100%;
					height: 100%;
				}

				$ h1 {
					font-size: {{ .Size }};
					text-align: center;
					margin: 20% auto;
				}

				$ h1::after{
					content:"!";
					display: inline-block;
				}

				$ h1:hover::after{
					content:"*";
					color: red;
					display: inline-block;
				}

			`, struct{ Size string }{Size: "130px"}),
			ID("hello"),
			Header1(
				Text("Hello"),
				events.Click(func(ev EventObject, tree *Markup) {
					js.Global.Call("alert", "I just got clicked, Yaay!!!")
				}, ""),
			),
		)
	}, "", false)

})
