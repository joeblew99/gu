package pages

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/examples/subscribe/app"
)

var _ = gu.Resource(func() {
	gu.GlobalRoute("#subscriptions/submit")
	gu.View(app.NewNotifier(), "", false, false)
})
