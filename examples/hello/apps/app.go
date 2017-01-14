package apps

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
	"github.com/gu-io/gu/trees/events"
	"github.com/gu-io/gu/trees/property"
)

var _ = gu.Resource(func() {

	gu.Title("Hello App")

	gu.View(func() *trees.Markup {
		return elems.Div(
			elems.CSS(`
				${
					width:100%;
					height: 100%;
				}

				$ h1 {
					font-size: {{ .Size }};
					text-align: center;
					margin: 20% 0  0 5%;
				}

				$ span {
					text-align: center;
					font-weight: bold;
					margin: 0 0 0 50%;
					color: rgba(0,0,0,0.5);
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
			property.IDAttr("hello"),
			elems.Header1(
				elems.Text("Hello"),
				events.ClickEvent(func(ev trees.EventObject, tree *trees.Markup) {
					js.Global.Call("alert", "I just got clicked, Yaay!!!")
				}, ""),
			),
			elems.Span(elems.Text("Click me")),
		)
	}, "", false, false)

})
