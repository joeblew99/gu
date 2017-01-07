package pages

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/examples/subscribe/app"
)

var _ = gu.Resource(func() {

	gu.UseRoute("#subscriptions/submit")

	notifier := app.NewNotifier()
	gu.DoView(notifier, "", false, false)

})
