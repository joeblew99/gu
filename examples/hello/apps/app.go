package apps

import (
	. "github.com/influx6/gu/design"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/attrs"
	. "github.com/influx6/gu/trees/elems"
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
			`, struct{ Size string }{Size: "130px"}),
			ID("hello"),
			Header1(Text("Hello")),
		)
	}, "", false)

})
