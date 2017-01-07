package pages

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/examples/subscribe/app"
)

var _ = gu.Resource(func() {

	gu.DoTitle("App Subscription Submission")
	gu.DoLink("https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto", "stylesheet", false)

	gu.DoStyle(app.IndexCSS, nil, false)

})
