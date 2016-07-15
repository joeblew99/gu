# Views And Routes
Although the internal routing system for `gu` is built to be an optional addition just as the use of the Reactive interface. It does provide an nice and elegant solution for routing.

Taking inspiration from the idea of delayed routing and insulated routes that are not flat but nest-able regardless of their parent (i.e where a route can be attached to another high level route without loosing functionality), still the approach still leaves the developer to either use or adopt another suitable means.

*Previous versions of the routing system actually were opinionated and had routes flip the view state of views, but the new system allows more control on what should happen when a route actually changes*

## Resolvers
Taking inspiration from Polymer and other approaches, `Resolvers` exists to provide a custom registering system of nest-able routes, where a Component/Renderable can embedded a resolver within itself, then have its view or parent view register to a root route. This way anytime, the root route matches, the resolver gets called to match against the remaining of the path after the view matches its criteria.

This approach will allow any component to be moved or initialized into another without needing to change internal routing details.

*This routing system is as much an experimental one and will evolve as time goes.*

Resolvers provide a very simple approach by simple matching the route giving
within their level then passing on the remainder of the route down the stack if
there exists anymore path left from the route.

For example: If we had a route of `/colors/*`, where the root resolver was
giving the `/colors` path and we had multiple components with routes like `/red`,
`/green`, `/black`, etc. If the resolver matches against a provided route like
`/colors/red`, the root resolvers matches `/colors`, leaving `/red` as its remaining
portion as indicted by the `/*`, which alerts the internal pattern matching system
to allow extra paths within the received path.
Most routers will stop here, but `Resolvers` passes the remaining piece (i.e `/red`)
down to child `Resolvers`, which then match and react(that is call their pass or fail)
handlers. This in effect allows us to provide independent components with behaviors
free from their parents and can be attached to any root component with its own independent
routing rules.

You can see sample code in [History Example](../examples/history)

## Example
We want to create components which display or hide themselves based on the colors
specified within the browser route as explained above.

```go
package main

import (
	"bytes"

	"github.com/influx6/gu"
	"github.com/influx6/gu/gucss"
	"github.com/influx6/gu/gudispatch"
	"github.com/influx6/gu/gutrees"
	"github.com/influx6/gu/gutrees/attrs"
	"github.com/influx6/gu/gutrees/elems"
	"github.com/influx6/gu/gutrees/styles"
	"github.com/influx6/gu/guviews"
	"honnef.co/go/js/dom"
)

type backgrounds []*background

func (b backgrounds) Resolve(path gudispatch.Path) {
	for _, bg := range b {
		bg.Resolve(path)
	}
}

func (b backgrounds) React(fx func()) {
	for _, bg := range b {
		bg.React(fx)
	}
}

func (b backgrounds) Render() gutrees.Markup {
	root := elems.Div(attrs.Class("backgrounds"))

	for _, bg := range b {
		bg.Render().Apply(root)
	}

	return root
}

type background struct {
	gudispatch.Resolver
	guviews.Reactive
	color string
	show  bool
}

func newBackground(pattern string, color string) *background {
	bg := &background{
		color:    color,
		Resolver: gudispatch.NewResolver(pattern),
		Reactive: guviews.NewReactive(),
	}

	bg.ResolvedFailed(gudispatch.WrapHandler(bg.updateFail))
	bg.ResolvedPassed(gudispatch.WrapHandler(bg.updatePass))

	return bg
}

// Render returns the markup for the background box.
func (b *background) Render() gutrees.Markup {
	return elems.Div(
		gutrees.When(b.show, styles.Display("block"), styles.Display("none")),
		attrs.Class("box"),
		styles.Background(b.color),
		styles.Margin("10% auto"),
		styles.Width("50%"),
		styles.Height("50%"),
	)
}

func (b *background) updateFail() {
	b.show = false
	b.Publish()
}

func (b *background) updatePass() {
	b.show = true
	b.Publish()
}

// css takes a document and a gucss.Render, creating the need dom element
// with the content of the rendered css stylesheet.
func makeCSS(doc dom.Document, css gucss.Render) {
	styl := doc.CreateElement("style")
	doc.QuerySelector("head").AppendChild(styl)

	var data bytes.Buffer
	css.Render(&data)

	styl.SetTextContent(string(data.Bytes()))
}

func main() {
	css := gucss.NewRoot()

	css.Child("*", gucss.Properties{
		"padding": "0px",
		"margin":  "0px",
	})

	gucss.Extend(css, gucss.Properties{
		"width":  "100%",
		"Height": "100%",
	}, "html", "html body", ".backgrounds")

	doc := dom.GetWindow().Document()
	makeCSS(doc, css)

	colors := []string{"red", "black", "cyan", "yellow", "blue"}

	var bgs backgrounds

	for _, color := range colors {
		bgs = append(bgs, newBackground("/"+color, color))
	}

	bgView := guviews.New(bgs)
	gu.RenderAsBody(bgView)

	// Register the view to watch for Hash changes and also
	// when it fails to match, it should hide.
	guviews.AttachHash("/colors/*", bgView, nil)
	guviews.AttachHash("/bgcolors/*", bgView, nil)
}
```

Jumping into the above code, the details within the `main` function to seem familiar
except for a few parts.


- The Backgrounds
The background code is simple. It creates a component which has within it both  
`guviews.Reactive` and `gudispatch.Resolver`. As already explained, the first
allows a form of active reactivity, but the one we are concerned about is the  
`gudispatch.Resolver`.

```go

type background struct {
	gudispatch.Resolver
	guviews.Reactive
	color string
	show  bool
}

func newBackground(pattern string, color string) *background {
	bg := &background{
		color:    color,
		Resolver: gudispatch.NewResolver(pattern),
		Reactive: guviews.NewReactive(),
	}

	bg.ResolvedFailed(gudispatch.WrapHandler(bg.updateFail))
	bg.ResolvedPassed(gudispatch.WrapHandler(bg.updatePass))

	return bg
}

func (b *background) updateFail() {
	b.show = false
	b.Publish()
}

func (b *background) updatePass() {
	b.show = true
	b.Publish()
}
```
Using the `gudispatch.NewResolver(pattern)`, which returns a new instance of
a object matching the `gudispatch.Resolver` interface, by using the giving pattern
as it's section ruling (i.e the part of the route it matches).

*The `gudispatch.WrapHandler` returns a function with a type `func(gudispatch.Path)`,
hence wrapping a function with no argument to allow binding to functions that need
such a type*

```go
	bg.ResolvedFailed(gudispatch.WrapHandler(bg.updateFail))
	bg.ResolvedPassed(gudispatch.WrapHandler(bg.updatePass))
```

Within the code we also register two functions, which flip switches for the background
when to set its internal style to allow it either be displayed or hidden. When the
resolver passes its match test, the `passed` handlers are called and if failed then
the failed handlers are called instead.

Next is the rendering for the component, which returns a markup and using a helper
package level functioon `gutrees.When`, depending on the state of a boolean value
returns the first property or the second.

*The usage of the `gutrees.When` is a matter of choice, but this author prefers
it, as it makes the rendering code alot cleaner.*

```go
// Render returns the markup for the background box.
func (b *background) Render() gutrees.Markup {
	return elems.Div(
		gutrees.When(b.show, styles.Display("block"), styles.Display("none")),
		attrs.Class("box"),
		styles.Background(b.color),
		styles.Margin("10% auto"),
		styles.Width("50%"),
		styles.Height("50%"),
	)
}

```

- The background List
*The `backgrounds` slice type defines a list of `background` pointers.*
This indeed may seem odd, but it was created for convenience sake and to allow more
control on the output of the view rendering structure. In a sense, this shows the
one weakness of `gu` approach, where to provide more control, you will need to wrap
individual `Renderables` to suit your needs.

The `backgrounds` is nothing special, because it only exposes standard functions
to allow the view to believe its a reactable and resolvable component. But in truth,
simply passes control to it's internals.

```go
type backgrounds []*background

func (b backgrounds) Resolve(path gudispatch.Path) {
	for _, bg := range b {
		bg.Resolve(path)
	}
}

func (b backgrounds) React(fx func()) {
	for _, bg := range b {
		bg.React(fx)
	}
}

func (b backgrounds) Render() gutrees.Markup {
	root := elems.Div(attrs.Class("backgrounds"))

	for _, bg := range b {
		bg.Render().Apply(root)
	}

	return root
}

```

- The Route
This final part within the `main` function, defines the binding for which URL the
view must respond to. Luckily with the `Resolvers` we are not bound to just one and
can bind to multiple, since it's all about letting the parent define its own piece.

```go
	guviews.AttachHash("/colors/*", bgView, nil)
	guviews.AttachHash("/bgcolors/*", bgView, nil)
```

*One thing left out here is the piece set to `nil`, the last argument requires
a function that accepts a `gudispatch` and this is called when the current path
does not match the attached route, allowing the view to hide or perform some other
action once, it's route is out of scope.*

This piece sets up the view for our background to two different routes based on
the browser hash changes (i.e when the hash part of the browser route changes),
these is used to match if the view should be notified. Both binds to a root
piece of `/colors` and `/bgcolors`.
This means when either the hash has a path of `#bgcolors/red` or `#colors/red`,
then the red background will be displayed.
