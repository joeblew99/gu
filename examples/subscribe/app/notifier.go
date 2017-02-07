package app

import (
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/notifications"
	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
	"github.com/gu-io/gu/trees/property"
)

// SubmissionNotifier defines the handler which displays the notification on the success or
// failure of a subscription.
type SubmissionNotifier struct {
	gu.Reactive
	c SubscriptionSubmitEvent
}

// NewNotifier returns a new instance of a page email subscription notifier.
func NewNotifier() *SubmissionNotifier {
	sn := SubmissionNotifier{
		Reactive: gu.NewReactive(),
	}

	sn.start()
	return &sn
}

// starts begin listening for a SubmitNotification which gets displayed.
func (s *SubmissionNotifier) start() {
	notifications.Subscribe(func(sme SubscriptionSubmitEvent) {
		s.c = sme
		s.Publish()
	})
}

// Render returns the markup for the page which displays the end result of a
// subscription submission.
func (s *SubmissionNotifier) Render() *trees.Markup {
	return elems.Div(
		elems.CSS(NotificationCSS, s),
		elems.Header1(property.ClassAttr("submission", "title"), elems.Text("App Subscription")),
		elems.Section(
			property.ClassAttr("submission", "content"),
			trees.MarkupWhen((s.c.Status && s.c.Email != ""),
				elems.Header2(property.ClassAttr("header", "roboto", "passed"), elems.Text("Done!")),
				elems.Header2(property.ClassAttr("header", "roboto", "failed"), elems.Text("Failed!")),
			),
			trees.MarkupWhen(s.c.Email != "",
				elems.Paragraph(
					property.ClassAttr("desc"),
					elems.Text("Welcome "),
					elems.Span(property.ClassAttr("email"), elems.Text("%q", s.c.Email)),
					elems.Text(" to the App."),
				),
				elems.Paragraph(
					property.ClassAttr("desc"),
					elems.Text("We are sorry but subscription failed."),
				),
			),
		),
	)
}
