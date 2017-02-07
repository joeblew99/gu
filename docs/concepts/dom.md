DOM
====
The DOM in Gu is package which provides a similar but functional approach to define
the representation of the expected DOM which must be rendered by the selected [Driver](./drivers.md).

Gu ensures to provide a comprehensive list of functions to generate and create
different tags and attributes to fit the standard HTML/HTML5 DOM nodes and events.

Which then allows an expressive and functional style in creating contents for components,
views and the central app.

Using the DOM packages requires the following packages, with each providing different
portions of the HTML/HTML5 API:

- Trees Package(https://github.com/gu-io/gu/trees)
The `trees` package provides the central root structure for the GU DOM. It is the
baseline package used in the creation of object instance to represent different tags,
events, and attributes types.

```go

import (
  "github.com/gu-io/gu/trees"
)

func main(){
  selfClosing := false

  // Creates a div structure and indicates that it is not a self closing tag.
  div := trees.NewMarkup("div", selfClosing)
  trees.NewAttr("class","articles").Apply(div)
  trees.NewCSSStyle("display", "block").Apply(div)

  // Parses the provided string as childing of the provided div.
  trees.ParseToRoot(div,`
    <article class="article">
      <div>
        <label>Name</Label>
        <label>Story</Label>
    </article>
    <article class="article">
      <div>
        <label>Name</Label>
        <label>Story</Label>
    </article>
    <article class="article">
      <div>
        <label>Name</Label>
        <label>Story</Label>
    </article>
  `)

  div.HTMl() => `
    <div class="articles" style="display:block; ">
      <article class="article">
        <div>
          <label>Name</Label>
          <label>Story</Label>
      </article>
      <article class="article">
        <div>
          <label>Name</Label>
          <label>Story</Label>
      </article>
      <article class="article">
        <div>
          <label>Name</Label>
          <label>Story</Label>
      </article>
    </div>
  `
}
```

Although a trivial example has above, the code does demonstrate how the `trees`
foundational structures help define and construct the content to be generated
effectively and efficiently.

- Elems Package(https://github.com/gu-io/gu/trees/elems)
The `elems` package provides is an auto-generated package which provides a functional
style of calls to describing the structures of the HTML to be rendered and provides
a more cleaner and easier use built on the foundation of the `trees` package.

```go

import (
  "github.com/gu-io/gu/elems"
)

func main(){
		div := elems.Div(
			elems.CSS(`
				${
					width:100%;
					height: 100%;
          background: {{.Color}};
				}
		`, struct{ Color string }{Size: "#ccc"}),
		elems.Header1(elems.Text("Hello")}),
		elems.Span(elems.Text("Click me")))

    div.HTMl() => `
      <div uid="34343440KK32232232">
        <style>
          div[uid="34343440KK32232232"]{
            width: 100%;
            height: 100%;
            background: #ccc;
          }
        </style>
        <h1>Hello</h1>
        <span>Click me</span>
      </div>
    `
}
```

By using a more functional declarative style, constructing complicated markup
directives becomes easier and simpler.

- Property Package(https://github.com/gu-io/gu/trees/property)
The `property` package follows in the style of the `elems` package to provide a functional
and declarative approach in provided attributes and styles to the constructed elements.
The `property` package differentiates attributes and styles by append a suffix of
`Attr` to the name of the property if an attribute and a suffix of `Style` to a style
property.

```go

import (
  "github.com/gu-io/gu/elems"
  "github.com/gu-io/gu/property"
)

func main(){
		div := elems.Div(
      property.ClassAttr("cage", "wrapper"),
      property.DisplayStyle("inline-block"),
			elems.CSS(`
				${
					width:100%;
					height: 100%;
          background: {{.Color}};
				}
		`, struct{ Color string }{Size: "#ccc"}),
		elems.Header1(elems.Text("Hello")}),
		elems.Span(elems.Text("Click me")))

    div.HTMl() => `
      <div uid="34343440KK32232232" class="cage wrapper" style="display:inline-block; ">
        <style>
          div[uid="34343440KK32232232"]{
            width: 100%;
            height: 100%;
            background: #ccc;
          }
        </style>
        <h1>Hello</h1>
        <span>Click me</span>
      </div>
    `
}
```

- Events Package(https://github.com/gu-io/gu/trees/events)
The `events` package follows in the style of the `elems` package to provide a functional
and declarative approach in defining the expected events which must occur on specific
elements. The events are actually bound on the target where the giving DOM is mounted into
but a checked is done to validate the real target of the event and if it matches the
desired elements, this is handled by the driver used. This way we can easily declare and
describe the behaviours we need for when events occur.

```go

import (
  "github.com/gu-io/gu/elems"
  "github.com/gu-io/gu/events"
)

func main(){
		div := elems.Div(
		elems.Span(elems.Text("Click me")),
    events.ClickEvent(func(event trees.EventObject, root *trees.Markup){
      event.PreventDefault()
      // do something.....
    })
  )

    div.HTMl() => `
      <div uid="34343440KK32232232">
        <span>Click me</span>
      </div>
    `
}
```
