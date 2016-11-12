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

// Subscriber defines the
type Subscriber struct{}

// Render returns the markup for the subscription component.
func (s Subscriber) Render() *Markup {
	return Section(
		CSS(`
      $ {
        width: 100%;
        height: 100%;
      }

      $.subscription form{
        width: 100%;
        height: 100%;
      }

      $.subscription .form .form-control{

      }

      $.subscription .form .form-control .submit-wrapper {
        width: 100%;
        height: 100%;
      }

      $.subscription .form .form-control .submit-wrapper .submit-email {
        width: 70%;
        height: 100%;
      }

      $.subscription .form .form-control .submit-wrapper .submit-button {
        width: 30%;
        height: 100%;
        float: left;
      }

      $.subscription .form .form-control .submit-wrapper .submit-arrow {
        display: none;
        width: 300px;
        height: 100%;
        float: left;
      }
    `, nil),
		ClassAttr("subscription"),
		Form(
			ClassAttr("form", "form-control"),
			Section(
				ClassAttr("submit-wrapper"),
				Input(
					ClassAttr("submit-email"),
					PlaceholderAttr("Email"),
					TypeAttr("email"),
				),
				Div(
					ClassAttr("submit-wrapper-button"),
					Input(
						ClassAttr("submit-arrow"),
						TypeAttr("submit"),
						ValueAttr(">"),
						ClickEvent(func(event EventObject) {

						}, ""),
					),
					Input(
						ClassAttr("submit-button"),
						TypeAttr("submit"),
						ValueAttr("Subscribe"),
						ClickEvent(func(event EventObject) {

						}, ""),
					),
				),
			),
		),
	)
}
