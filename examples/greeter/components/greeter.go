package components

import (
	"fmt"

	"github.com/gu-io/gu"
	"github.com/gu-io/gu/eventx"
	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
	"github.com/gu-io/gu/trees/events"
	"github.com/gu-io/gu/trees/property"
)

// Greeter defines a component which greets the name from a input event.
type Greeter struct {
	gu.Reactive
	Name string
}

// NewGreeter returns a new instance of a Greeter.
func NewGreeter() *Greeter {
	return &Greeter{
		Reactive: gu.NewReactive(),
	}
}

// Render returns the markup for the greeter.
func (g *Greeter) Render() *trees.Markup {
	return elems.CustomElement("greeter",
		property.ClassAttr("greeter"),
		elems.CSS(`
      &{
        width: 100%;
        height: auto;
				margin: 0px auto;
				font-size: 100%;
      }


      & div.intro{
        width: 90%;
        padding: 10px;
				font-size: 3.0em;
      }

      & div.intro span.person{
				text-align: center;
			}

      & div.receiver{
        width: 80%;
        height: 40px;
        padding: 10px;
				font-size: 1.0em;
      }

      & div.receiver input{
        display: block;
        padding: 10px;
        border: none;
        outline: none;
				background: none;
        border-bottom:5px solid rgba(255,255,255,0.3);
      }

    `, nil),
		elems.Div(
			property.ClassAttr("intro"),
			elems.Text("Welcome "),
			elems.Span(
				property.ClassAttr("person"),
				trees.MarkupWhen(g.Name == "",
					elems.SpaceCharacter(3),
					elems.Text("%q, Great Explorer!", g.Name),
				),
			),
		),
		elems.Div(
			property.ClassAttr("receiver"),
			elems.Input(
				property.TypeAttr("text"),
				property.ValueAttr(g.Name),
				property.PlaceholderAttr("Enter your Name"),
				events.ChangeEvent(func(event trees.EventObject, _ *trees.Markup) {
					if change, ok := event.Underlying().(*eventx.ChangeEvent); ok {
						fmt.Printf("Changed Occured: %#v\n", change)
						g.Name = change.Value
						g.Publish()
					}
				}, ""),
			),
		),
	)
}
