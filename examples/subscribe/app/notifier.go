package app

import (
	"fmt"
	"github.com/influx6/gu"
	"github.com/influx6/gu/dispatch"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/elems"
	. "github.com/influx6/gu/trees/property"
)

// SubscriptionSubmitEvent defines a event sent out to define the status of a subscription
// event.
type SubscriptionSubmitEvent struct {
	Email  string `json:"email"`
	Status bool   `json:"status"`
}

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
	dispatch.Subscribe(func(sme SubscriptionSubmitEvent) {
		fmt.Printf("Received submit %#v\n", sme)
		s.c = sme
		s.Publish()
	})
}

// Render returns the markup for the page which displays the end result of a
// subscription submission.
func (s *SubmissionNotifier) Render() *Markup {
	return Div(
		CSS(NotificationCSS, s),
		Header1(ClassAttr("submission", "title"), Text("App Subscription")),
		Section(
			ClassAttr("submission", "content"),
			MarkupWhen((s.c.Status && s.c.Email != ""), Header2(ClassAttr("header", "roboto", "passed"), Text("Done!")), Header2(ClassAttr("header", "roboto", "failed"), Text("Failed!"))),
			MarkupWhen(s.c.Email != "", Paragraph(
				ClassAttr("desc"),
				Text("Welcome "),
				Span(ClassAttr("email"), Text("%q", s.c.Email)),
				Text(" to the App."),
			),
				Paragraph(
					ClassAttr("desc"),
					Text("We are sorry but subscription failed."),
				)),
		),
	)
}
