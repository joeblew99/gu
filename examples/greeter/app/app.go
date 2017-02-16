package app

import (
	"github.com/gu-io/gu/eventx"
	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
	"github.com/gu-io/gu/trees/events"
	"github.com/gu-io/gu/trees/property"
)

// Greeter defines a component which greets the name from a input event.
//
// shell:component
// Resource{
//   Name: spacewall.png
//   URL:
//   Localize:true
// }
type Greeter struct {
	Name string
}

// Render returns the markup for the greeter.
func (g *Greeter) Render() *trees.Markup {
	return elems.CustomElement("greeter",
		property.ClassAttr("greeter"),
		elems.CSS(`
      &{
        width: 100%;
        height: auto;
      }


      & div.intro{
        width: 90%;
        padding: 10px;
      }

      & div.receiver{
        width: 50%;
        height: 40px;
        padding: 10px;
      }

      & div.receiver input{
        display: block;
        width: 100%;
        height: 100%;
        border: none;
        outline: none;
        border-bottom:5px solid #fff;
      }

    `, nil),
		elems.Div(
			property.ClassAttr("intro"),
			elems.Text("Welcome to the Club"),
			elems.Span(
				property.ClassAttr("person"),
				trees.MarkupWhen(g.Name == "",
					elems.SpaceCharacter(3),
					elems.Text("%s", g.Name),
				),
			),
		),
		elems.Div(
			property.ClassAttr("receiver"),
			elems.Input(
				property.TypeAttr("text"),
				property.PlaceholderAttr("Enter your Name"),
				events.ChangeEvent(func(event trees.EventObject, _ *trees.Markup) {
					if change, ok := event.Underlying().(*eventx.ChangeEvent); ok {
						g.Name = change.Value
					}
				}, ""),
			),
		),
	)
}
