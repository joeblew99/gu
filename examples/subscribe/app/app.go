package app

import (
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/elems"
	. "github.com/influx6/gu/trees/events"
	. "github.com/influx6/gu/trees/property"
)

// Subscription defines a component which collects the subscribers email received through
// a subscription form to be submitted to a API.
type Subscription struct {
	Email string `json:"email"`
}

// SubscriptionNotification defines a type which exposes the notification and status
// of a subscription submission.
type SubscriptionNotification struct {
	Status bool `json:"status"`
}

// Subscriber defines the Subscriber component which renders a subscriber
// submission form which collects the data received and submits it to the API.
type Subscriber struct {
	SubmitBtnColor string
}

// Render returns the markup for the subscription component.
func (s *Subscriber) Render() *Markup {
	return Section(
		CSS(SubscribeCSS, s),
		ClassAttr("subscription"),
		Form(
			ClassAttr("form", "form-control"),
			Section(
				ClassAttr("email"),
				Input(
					ClassAttr("email"),
					PlaceholderAttr("example@mail.com"),
					TypeAttr("email"),
				),
			),
			Section(
				ClassAttr("buttons"),
				Button(
					Text("Subscribe"),
					ClassAttr("button", "named"),
					ClickEvent(func(event EventObject) {

					}, ""),
				),
			),
		),
	)
}
