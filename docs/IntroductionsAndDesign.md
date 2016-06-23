# Introduction To Gu
Gu is a library built to provide a rendering system which encapsulates a diff
algorithm that ensures only the needed parts of your views that changed get
re-rendered. It's main goal is to reduce your thinking load by taking care of
the rendering needs of your application which allows you to concentrate on just
your architecture and structure.

Gu adopts different ideas from different projects to meet its need to provide
both a rendering system capable of working either on the backend or frontend
with ease.

## Gu By Design
Gu in its design was never created to be complex and in it's very nature reduces
 the adoption process needed to get started with it. Below are all the design principle that is need to use Gu effectively.

### Renderable
Renderables are the core and key to using Gu, all views and rendering require that this interface be followed, because Gu takes the approach of releaving the developer of the worries about rendering, it only ever requires you to provide a snapshot of your application view using the [GuTrees](./gutrees) sublibrary which contains a large set of custom structures for different DOM elements.

```go
type Renderable interface {
	Render() gutrees.Markup
}
```

Once this interface is met by any `struct` then it's content can be rendered and updated at the most optimize mode possible. This allows us to write a custom rendering markup for any structure we wish.  This is in all of `Gu` where such requirement is needed, beyond this interface the developer is free to structure their application as they please.

- For example, making an array a Renderable item:

```go

type List []string

func (l List) Render() gutrees.Markup {
  root := elems.UnorderedList()

  for _, item := range l {
    elems.ListItem(elems.Text(item)).Apply(root)
  }

  return root
}
```



### ReactiveRenderable
`ReactiveRenderable` is more of a simple mechanism that allows views or other `Renderables` to react to changes from other `Renderables` by implementing the ReactiveRenderable interface.

Ofcourse, it's not magic and still requires the developer to either build such a methods into their custom component, or compose one provided by `gu`.
Views will automatically register themselves for updates with `ReactiveRenderable` components, which allows them to update the frontend when changes occur in all the `Renderable` passed in.

```go

// ReactiveRenderable defines an interface of a Renderable capable of
// notifying subscribe functions of changes to allow them to React.
type ReactiveRenderable interface {
	Renderable
	Subscribe(func())
}

```

The `Reactive` interface defines a Reactor, that can be leverage and which has a concrete implementation in `gu`. It can be created by calling the `guviews.NewReactive()` function. You can embed this into your component which turns it into a `ReactiveRenderable`.

```go
// Reactive extends the ReactiveRenderable by exposing a Publish method
// which allows calling the update notifications list of a ReactiveRenderable.
type Reactive interface {
	Subscribe(func())
	Publish()
}

```

- For example, making a reactive Menu list.

```go

type Menu struct{
  guviews.Reactive
  links []string
}

func NewMenu(links []string) *Menu{
  return &Menu{
    Reactive: guviews.NewReactive(),
    links: links,
  }
}

func (m *Menu) Add(link string) {
  m.links = append(m.links,link)
  m.Publish()
}

func (m *Menu) Render() gutrees.Markup {
  root := elems.UnorderedList()

  for _, item := range m.links {
    elems.ListItem(elems.Text(item)).Apply(root)
  }

  return root
}
```


### Markup
Gu uses a markup library called [Gutrees](./gutrees) inspired by [Richard Musiol's](https://github.com/neelance) work on a demo library for a DOM-like markup library written in native Go, using native Go.
By grafting such an approach, `Gu` simplified the markup approach where rather than depending on `go templates`, the developer gets the benefit of a native system where the DOM markup are expressed within the Go syntax, thereby reducing the cognitive gap and allow expressiveness with the full power of the language.

`GuTrees` provides both a markup approach which embeds event registration and inline style addition that fits well using the native language syntax.

#### Gutrees Subpackage
- [gutrees/Elems](./gutrees/elems)
  Contains the functions which create the needed markup objects for different HTML elements.

- [gutrees/Attrs](./gutrees/attrs)
  Contains the functions which creates a limited set of attribute types for the markup elements.

- [gutrees/Styles](./gutrees/styles)
  Contains the functions which creates a limited set of inline-style types for the markup elements.

- [gutrees/Events](./gutrees/events)
  Contains the functions which create the needed event objects for different DOM events types and is used by `Gutrees` to bind events for its markup.

Gutrees sets up the event functions within

```go
	import "github.com/influx6/gu/gutrees"
	import "github.com/influx6/gu/gutrees/attrs"
	import "github.com/influx6/gu/gutrees/styles"
	import "github.com/influx6/gu/gutrees/elems"
	import "github.com/influx6/gu/gutrees/events"

	div := elems.Div(
	 styles.Width(400),
	 attrs.ID("main-div"),
	 gutrees.NewAttribute("data-warp","warp-div"),
	 gutrees.NewStyle("height","500px"),
	 events.Click("",func (ev guevents.Event,root gutrees.Markup){
	   fmt.Println("Am clicked!")
	 },
	)
```

### Styles
Gu provides a optional css generating library called [GucSS](./gucss), inspired by [GuTrees](./gutrees).
It provides a Go native, but yet simple approach to building css styles which can be rendered and attached to inline style tags.

```go
	import "github.com/influx6/gu/gutrees"
	import "github.com/influx6/gu/gutrees/gucss"

	root := gucss.NewRoot()

	root.Within("*", gucss.Properties{
		"margin":  "0px",
		"padding": "0px",
	})

	root.Within("html", gucss.Properties{
		"width": "100%",
	}).Within("body", gucss.Properties{
		"width":     "100%",
		"font-size": "1em",
	})

	root.Child("div", gucss.Properties{}).Within("ul", gucss.Properties{
		"list-style-type": "none",
	}).Extend("ol", gucss.Properties{
		"list-style-type": "none",
	}).PreSibling("li", gucss.Properties{
		"padding": "10px",
	}).PostSibling("span", gucss.Properties{
		"padding": "3px",
	}).NthParent(1).Child("a", gucss.Properties{
		"color": "#ccc",
	}).NthParent(2).NS(":hover", gucss.Properties{
		"padding": "30px",
	})

	var b bytes.Buffer
	gucss.NewMedia("screen", "(width: 30em) and (height: 40em)").Render(root, &b)
```

### Views
The views are the core of `gu`. They are the key to how the `Renderables` are rendered, and they basically manage the cycle of rendering and update.

*I personally believe when working with Views, one should only have them attached to a root `Renderable`, which encapsulates others. This leads to a far more simpler design
and architecture rather than having multiple views managing
parts of a single `Renderable`.*

Views are by default not reactive ofcourse this can be changed with by meeting the `ReactiveRenderable`, but they do watch for notifications through `gudispatch`, that tell them to update.

Views can easily be created by supplying to them the needed `Renderable` or if passed a set of `Renderables` then they will be rendered into a div before getting attached to the DOM target.

```go
type List []string

func (l List) Render() gutrees.Markup {
  root := elems.UnorderedList()

  for _, item := range l {
    elems.ListItem(elems.Text(item)).Apply(root)
  }

  return root
}

listView := guviews.New(List([]string{"Wombat", "Rabbits"}))
```

#### Views and Path Changes
To ensure developers can react with the URL changes, `gu` includes the ability to watch the changes of the browsers history. When a view is assigned a given path, changes to the browsers history will automatically hides or shows the view as needed.

```go
	import "github.com/influx6/gu/guviews"

	guviews.AttachView(view, "/buckler")
```

#### Views Initialization
Gu includes a optional initalization system which lets you register a `Renderable` returning function which can then be called to be created on demand. This is nice if you wish to centralize your view initialization code and call it on the server or client. Initializing this way allows you to maintain an important detail with `Views`, that is, its `ID`.
Views have a unique `ID` generated by default or assignable if desired. This lets views rendering equally well when rendered from the server and allowed to take over on the client. More on this in [Initializing and Server Side Rendering](./docs/InitializationsAndServerSide.md)

*Ofcourse this is not enfored and you are free to set things as you desire.*


```go
type List []string

func (l List) Render() gutrees.Markup {
  root := elems.UnorderedList()

  for _, item := range l {
    elems.ListItem(elems.Text(item)).Apply(root)
  }

  return root
}

guviews.Register("app/list", func (items []string) guviews.Renderable {
 return List(items)
})

guviews.Create({
  Name: "app/list",
  Elem: "body",
  ID: "32-223-323232-453435",
  Param: []string{"Wombat", "Chilly"},
})
```

By using this approach we can ensure important details like the `Views ID` is maintained regardless of if the view gets rendered at the backend then initialized on the front-end to take control or rendered from the frontend directly.


### Conclusion
The parts that make up the `gu` library are very simple in their design and approach, and the next set of links presents more practically use cases of the `gu` library to build your  components.

  - **[Components and Views](./docs/ComponentsAndViews.md)**
  - **[Components and Subcomponents](./docs/ComponentsAndSubComponents.md)**
  - **[Components and Styles](./docs/ComponentsAndStyles.md)**
  - **[Reactivity and Notifications](./docs/ReactivityAndNotifications.md)**
  - **[Initializing and Server Side Rendering](./docs/InitializationsAndServerSide.md)**
  - **[Limitations](./docs/Limitations.md)**
