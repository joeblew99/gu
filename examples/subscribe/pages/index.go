package pages

import (
	. "github.com/gu-io/gu/design"
	. "github.com/gu-io/gu/examples/subscribe/app"
)

var _ = Resource(func() {

	DoTitle("App Subscription Submission")
	DoLink("https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto", "stylesheet", false)

	DoStyle(IndexCSS, nil, false)

})
