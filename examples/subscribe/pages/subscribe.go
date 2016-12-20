package pages

import (
	. "github.com/gu-io/gu/design"
	. "github.com/gu-io/gu/examples/subscribe/app"
	. "github.com/gu-io/gu/trees/elems"
	. "github.com/gu-io/gu/trees/property"
)

var _ = Resource(func() {

	UseRoute("#")

	DoMarkup(Div(
		CSS(RootCSS, nil),
		ClassAttr("root"),
		Header1(
			Text("Become A Subscriber"),
		),
	), "", false, false)

	DoView(&Subscriber{
		SubmitBtnColor: "",
	}, "", false, false)

})
