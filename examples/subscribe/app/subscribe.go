package app

import (
	"github.com/gu-io/gu/dispatch"
	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
	"github.com/gu-io/gu/trees/events"
	"github.com/gu-io/gu/trees/property"
	"honnef.co/go/js/dom"
)

// SubscriptionSubmitEvent defines a event sent out to define the status of a subscription
// event.
type SubscriptionSubmitEvent struct {
	Email  string `json:"email"`
	Status bool   `json:"status"`
}

// Subscriber defines the Subscriber component which renders a subscriber
// submission form which collects the data received and submits it to the API.
//
// shell:component
//
// Resource {
//  Name: roboto.font.css
//	Path: https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto
//	Hook: css-embed
//	Localize: true
// }
//
type Subscriber struct {
	SubmitBtnColor string
}

// Render returns the markup for the subscription component.
func (s *Subscriber) Render() *trees.Markup {
	return elems.Section(
		elems.CSS(SubscribeCSS, s),
		property.ClassAttr("subscription"),
		elems.Form(
			property.ClassAttr("form", "form-control"),
			elems.Section(
				property.ClassAttr("email"),
				elems.Input(
					property.PlaceholderAttr("example@mail.com"),
					property.TypeAttr("email"),
				),
			),
			elems.Section(
				property.ClassAttr("buttons"),
				elems.Button(
					elems.Text("Subscribe"),
					property.ClassAttr("button", "named"),
					events.ClickEvent(func(event trees.EventObject, tree *trees.Markup) {
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
