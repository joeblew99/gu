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
ReactiveRenderable is more of a simple mechanism that allows views or other Renderables to react to changes from other Renderables by implementing the ReactiveRenderable interface.
Ofcourse it's not magic and still requires the developer to either build such a methods into their custom component or compose one provided by `gu`. Views will automatically register themselves for updates with `ReactiveRenderable` components which allows them to update themselves when changes occur in the passed in `Renderable` or `Renderables` if more than one.

```go

// ReactiveRenderable defines an interface of a Renderable capable of
// notifying subscribe functions of changes to allow them to React.
type ReactiveRenderable interface {
	Renderable
	Subscribe(func())
}

```

The `Reactive` interface defines a Reactor that can be leverage and which has a concrete implementation in `gu`, which can be retrieved by calling the `guviews.NewReactive()` function.
You can embed this into your components and have them
provided a structure which handles the subscription and publishing needs.

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
Gu uses a markup library called [GuTrees](./gutrees) inspired by [Richard Musiol](https://github.com/neelance) work on a demo library for a dom-like markup library written in go and generated using the go code generation feature.
By grafting such a approach, `Gu` simplified the markup approach where rather than depending on `go templates`, the developer gets the benefit of a simple native based system where the DOM markup could be expressed within the native syntax and idiomatic thinking of the Go language, thereby reducing learning and cognitive gap and allow expressiveness with the full power of the language.

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
Gu uses provides a optional style library called [GuCSS](./gucss) inspired by [GuTrees](./gutrees) which provides a go-centric approach to building css styles which can be rendered and attached to inline style tags, they allow you to organize your styles from being inlined to the specific Markup but more organized as a whole.

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
The views are the core structure within `gu` and they are the key to how the `Renderables` are rendered, basically they manage the core operations and provide you the mechanism to have your structures adequately updated and rendered as needed.

I personally believe when working with `gu` views, one should only have them attached to a root `Renderable`, which encapsulates others. This leads to a far more simpler design
and architecture rather than having multiple views managing
parts of a single `Renderable`.

Views are by default not reactive, they watch for notifications that tell them to update but `gu` does provide
a optional interface `Reactive` that can be implement or using its default implementation to have views notified by their `Renderable` or `Renderables` when to update, ofcourse this
still requires the developer to call the `Reactive.Publish`
method appropriate.

Views can easily be created by supplying to them the needed `Renderable` or list of `Renderables` of which they are to manage.

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
To ensure developers can react with the URL changes, `gu` includes the ability to watch the changes of the browsers path when a view is assigned a given path, which automatically hides or shows the view when the path is visited or navigated from.

```go
	import "github.com/influx6/gu/guviews"

	guviews.AttachView(view, "/buckler")
```

#### Views and Initialization
Initialization of views is has important as the views themselves because `gu` as a library is made to function adequately in its rendering capability regardless of whether its being called on the front-end within a browser or on the backend.

This is very important especially when moving a rendered content from backend to frontend as the views use a unique ID which can either be assigned by the developer or auto-generated when initialized. This ID plays a great role for the view to capture a previous rendered point to either replace with new markup at initialization or update as needed.

Hence the `gu` libraries offers a approach which can be used depending on developers choice to register and initialize views both on the client and server.  



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
The parts that make up the `gu` libraries are very simple in their design and approach and the next documentation rather presents the practically use cases of the `gu` library to build applications.

  - **[Components and Views](./docs/ComponentsAndViews.md)**
  - **[Components and Subcomponents](./docs/ComponentsAndSubComponents.md)**
  - **[Components and Styles](./docs/ComponentsAndStyles.md)**
  - **[Reactivity and Notifications](./docs/ReactivityAndNotifications.md)**
  - **[Initializing and Server Side Rendering](./docs/InitializationsAndServerSide.md)**
  - **[Limitations](./docs/Limitations.md)**
