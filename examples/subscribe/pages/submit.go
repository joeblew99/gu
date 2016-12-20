package pages

import (
	. "github.com/gu-io/gu/design"
	. "github.com/gu-io/gu/examples/subscribe/app"
)

var _ = Resource(func() {

	UseRoute("#subscriptions/submit")

	notifier := NewNotifier()
	DoView(notifier, "", false, false)

})
