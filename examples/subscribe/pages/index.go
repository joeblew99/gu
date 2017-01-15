package pages

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/examples/subscribe/app"
)

var _ = gu.Resource(func() {
	gu.Title("App Subscription Submission")
	gu.Style(app.IndexCSS, nil, false)
})
