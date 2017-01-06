package app

import (
	"github.com/gu-io/gu/dispatch"
	. "github.com/gu-io/gu/trees"
	. "github.com/gu-io/gu/trees/elems"
	. "github.com/gu-io/gu/trees/events"
	. "github.com/gu-io/gu/trees/property"
	"honnef.co/go/js/dom"
)

// SubscriptionSubmitEvent defines a event sent out to define the status of a subscription
// event.
//
// shell:component
//
// Resource {
//  Name: roboto.font
//	Path: https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto
//	Hook: link-embed
// }
//
type SubscriptionSubmitEvent struct {
	Email  string `json:"email"`
	Status bool   `json:"status"`
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
