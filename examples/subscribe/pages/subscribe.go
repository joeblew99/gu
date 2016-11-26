package pages

import (
  . "github.com/influx6/gu/design"
  . "github.com/influx6/gu/examples/subscribe/app"
  . "github.com/influx6/gu/trees/elems"
  . "github.com/influx6/gu/trees/property"
)

var _ = Resource(func() {

  UseRoute("#")
  DoTitle("App Subscribe")

  DoMarkup(Div(
    CSS(RootCSS, nil),
    ClassAttr("root"),
    Header1(
      Text("Become A Subscriber"),
    ),
  ), "",false, false)

  DoView(&Subscriber{
    SubmitBtnColor: "",
  }, "", false, false)

})

