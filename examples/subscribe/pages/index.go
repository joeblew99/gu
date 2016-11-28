package pages

import (
  . "github.com/influx6/gu/design"
  . "github.com/influx6/gu/examples/subscribe/app"
)

var _ = Resource(func(){

  DoTitle("App Subscription Submission")
  DoLink("https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto","stylesheet",false)

  DoStyle(IndexCSS, nil,false)

})
