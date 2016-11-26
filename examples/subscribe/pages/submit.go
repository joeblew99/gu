package pages

import (
  . "github.com/influx6/gu/design"
  . "github.com/influx6/gu/examples/subscribe/app"
)

var _ = Resource(func(){

  UseRoute("#subscriptions/submit")
  DoTitle("App Subscription Submission")

  notifier := NewNotifier()
  DoView(notifier, "", false, false)

})