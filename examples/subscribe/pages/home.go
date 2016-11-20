package pages

import (
  . "github.com/influx6/gu/design"
  . "github.com/influx6/gu/examples/subscribe/app"
  . "github.com/influx6/gu/trees/elems"
  . "github.com/influx6/gu/trees/property"
)

var _ = Resource(func() {

  DoTitle("App Subscribe")

  DoMarkup(Div(
    CSS(RootCSS, nil),
    ClassAttr("root"),
    Header1(
      Text("Become A Subscriber"),
    ),
  ), "")

  DoView(&Subscriber{
    SubmitBtnColor: "",
  }, "")

})