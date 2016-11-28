package app

import (
	"github.com/influx6/gu/dispatch"
	. "github.com/influx6/gu/trees"
	. "github.com/influx6/gu/trees/elems"
	. "github.com/influx6/gu/trees/events"
	. "github.com/influx6/gu/trees/property"
	"honnef.co/go/js/dom"
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
					PlaceholderAttr("example@mail.com"),
					TypeAttr("email"),
				),
			),
			Section(
				ClassAttr("buttons"),
				Button(
					Text("Subscribe"),
					ClassAttr("button", "named"),
					ClickEvent(func(event EventObject, tree *Markup) {
						event.PreventDefault()
						event.StopPropagation()

						doc := dom.GetWindow().Document()
						input, ok := doc.QuerySelector(".subscription .form .email input").(*dom.HTMLInputElement)

						if !ok {
							dispatch.Dispatch(SubscriptionSubmitEvent{
								Status: false,
							})

							return
						}

						dispatch.NavigateHash("/", "#subscriptions/submit", "#")

						dispatch.Dispatch(SubscriptionSubmitEvent{
							Email:  input.Value,
							Status: true,
						})

					}, ""),
				),
			),
		),
	)
}
