Components
===========
Creating components is the core reason Gu exists as a package. It's major aim is
to provide a base library that allows rendering these components easily and efficiently.

Gu takes a different approach to components and how they should work. Has Gu does not
try to be a React version in Go, it takes advantage of the simple concepts that makes
the Go language very powerful.

- Composition over Inheritance
  Where components compose each other to create a larger component than using a form
  of inheritance or inter-logic where components are separately rendered and communicate with
  each other. Although such is possible, but when capable components are advised to let
  a root or parent component compose and initialize others and also handles the rendering
  of these components to it's view.

- Interfaces compliance for upgrades
  By adopting the strategy and powerful concept of interfaces components can provide
  the capability to expose themselves to higher functionality or gain access to objects
  such has the internal caching and resource request `Fetch` objects.
  More so this system allows components to declare themselves reactive and notify
  themselves and their views of change to be updated by the driver.


By sticking to such basic ideas and principles, it allows constructing components
with the standard construct provided by the Go language to the maximum capability
allowed.

## Basics
Creating a component is comparatively easy, has you are only required to meet a
single interface by which the rendering markup for the component is retrieved.

```go

import (
  "github.com/gu-io/gu/trees"
  "github.com/gu-io/gu/trees/elems"
  "github.com/gu-io/gu/trees/property"
)

// Greeter takes a giving name and generates a greeting.
type Greeter struct{
  Name string
}

// change updates the greeters name field.
func (g *Greeting) change(name string){
  g.Name = name
}

// Render returns the Gu's tree structures which declares the markup for
// the greeter.
func (g *Greeting) Render() *trees.Markup {
  return elems.Div(
    property.ClassAttr("greeter"),
    elems.Div(
      property.ClassAttr("greeting"),
      elems.Text("Welcome to the %s!",g.Name),
    ),
    elems.Div(
      property.ClassAttr("box", "input"),
      elems.Input(
        property.PlaceholderAttr("Enter your Name"),
        property.TypeAttr("text"),
      ),
    ),
  )
}
```
