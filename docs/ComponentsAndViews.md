# Components And Subcomponents
Components in `Gu` are simple structures that satisfy the requirement of the
`Renderable` interface. This can be a single `struct` or a composite struct that
combines multiple `Renderable` structures to make up it's rendering either as a
flat object with internal arrays of components, or where others are composed together
into a new Renderable. This simplicity in thinking provides a sensible organization
approach for the developer has he/she plans out the structures of their application.

### Example:

To demonstrate how a component with Subcomponents can be built together,
we will sample a demonstration for a dashboard display that allows us to
details different metrics. Where the main component is a dashboard with
stats which are its Subcomponents.


```go

// StatItem defines a component which defines a stat item on a dashboard.
type StatItem struct {
	guviews.Reactive
	Name  string
	Value int
}

// Set sets the value of the StatItem to the provided value.
func (s *StatItem) Set(d int) {
	s.Value = d
	s.Publish()
}

// NewStatItem returns a new instance of a StatItem.
func NewStatItem(name string, val int) *StatItem {
	return &StatItem{
		Reactive: guviews.NewReactive(),
		Name:     name,
		Value:    val,
	}
}

// Render returns the markup for a StatItem.
func (s *StatItem) Render() gutrees.Markup {
	return elems.Section(
		attrs.Class("stat"),
		elems.Label(
			attrs.Class("statDetail header"),
			elems.Header3(elems.Text(s.Name)),
		),
		elems.Label(
			attrs.Class("statDetail value"),
			elems.Text(s.Value),
		),
	)
}

// Dashboard defines a dashboard component which displays different stats.
type Dashboard struct {
	guviews.Reactive
	stats []*StatItem
}

// NewDashboard returns a new instance of a Dashboard.
func NewDashboard() *Dashboard {
	return &Dashboard{Reactive: guviews.NewReactive()}
}

// Add adds a new stat into the list of stats displayed by the dashboard
func (d *Dashboard) Add(stat *StatItem) {
	d.stats = append(d.stats, stat)
	stat.Subscribe(d.Publish)
}

// Render returns the markup for the Dashboard component.
func (d *Dashboard) Render() gutrees.Markup {
	root := elems.Div(attrs.Class("dashboard"))

	for _, stat := range d.stats {
		stat.Render().Apply(root)
	}

	return root
}

```

Breaking the code into more readable chunk, let me explain what each piece does.

- The Components

The two main components within our example is the `StatItem` and the `Dashboard`, where the `StatItem` defines the stat items to be displayed on the dashboard and their representation for the page. Each distinctively embeds the `guviews.Reactive` interface and
assigns the `guviews.NewReactive()` value in their `New` functions. This ensures both are reactive and can be listened to for changes. But one difference is the `Dashboard` binds to the `StatItem` it has and this is an excellent example on how other `Renderables` can listen to each other.

```go

// StatItem defines a component which defines a stat item on a dashboard.
type StatItem struct {
	guviews.Reactive
	Name  string
	Value interface{}
}

// Dashboard defines a dashboard component which displays different stats.
type Dashboard struct {
	guviews.Reactive
	stats []StatItem
}
```  

- The Rendering

The two components define their own rendering markup. This also showcases the idea, Root components need not contain all the markups, you should always allow each component to encapsulate it's own rendering, these way each just extends it's root and can be used in other places as needed. The most important partt is that their `Render()` methods return a `gutrees.Markup`. This is the DOM language understood by `gu` and the frontend will be constructed using the returned markup.

```go

// Render returns the markup for a StatItem.
func (s *StatItem) Render() gutrees.Markup {
	return elems.Section(
		attrs.Class("stat"),
		elems.Label(
			attrs.Class("statDetail header"),
			elems.Header3(elems.Text(s.Name)),
		),
		elems.Label(
			attrs.Class("statDetail value"),
			elems.Text(s.Value),
		),
	)
}

// Render returns the markup for the Dashboard component.
func (d *Dashboard) Render() gutrees.Markup {
	root := elems.Div(attrs.Class("dashboard"))

	for _, stat := range d.stats {
		stat.Render().Apply(root)
	}

	return root
}
```  


# Views And SubViews
Views are the core mechanism by which components get exposed to the DOM. They manage the rendering and update of the DOM by listening for change notifications. Due to the need for simplicity, this library ensures to only provided the minimal
needed tools to attain the best result, that is, it does not have an opinionated approach for how developers should use it but rather suggest best approach that bring best results.

*Tip: A view can be told to re-render itself by sending a update signal to it: `gudispatch.Dispatch(ViewUpdate{ID:"view custom-id or view UUID"})`*

We can turn the dashboard in the above section into a view:

```go

func wave(factor float64, val int) float64 {
	vm := math.Max(float64(val), 2.5)
	return vm + (math.Sin(math.Tan(factor)*math.Pow(math.Sin(factor), 10)) * vm)
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
	doc := dom.GetWindow().Document()

	css := gucss.NewRoot()

	css.Child("*", gucss.Properties{
		"padding": "0px",
		"margin":  "0px",
	})

	gucss.Extend(css, gucss.Properties{
		"width": "100%",
	}, "html", "html body", ".dashboard")

	css.Child(".dashboard", gucss.Properties{
		"height": "400px",
	}).Child(".stat", gucss.Properties{
		"display":   "inline-block",
		"width":     "30%",
		"color":     "#000",
		"border":    "1px dotted rgba(20,20,20,0.5)",
		"padding":   "30px",
		"font-size": "3em",
	})

	makeCSS(doc, css)

	dashboard := NewDashboard()

	visitors := NewStatItem("Visitor's Count", 200)
	clickRate := NewStatItem("Click Rates", 100)
	bounceRate := NewStatItem("Bounce Rate", 50)

	dashboard.Add(visitors)
	dashboard.Add(clickRate)
	dashboard.Add(bounceRate)

	dashboardView := guviews.New(dashboard)
	dashboardView.Mount(doc.QuerySelector("body").Underlying())

	factor := 0.05

	go func() {
		for {
			<-time.After(300 * time.Millisecond)
			visitors.Set(int(wave(factor, visitors.Value)))
			clickRate.Set(int(wave(factor, clickRate.Value)))
			bounceRate.Set(int(wave(factor, bounceRate.Value)))
			factor += 0.1
		}
	}()

}
```

Looking deeper into the example code above, we can see a few things happening, which do play an interesting role in how our demo works and presents itself.


- The `wave` function takes it's input and returns a value
in a periodic up and down fashion, these allows us abit of
near interesting datapoint for our stats.

```go
func wave(factor float64, val int) float64 {
	vm := math.Max(float64(val), 2.5)
	return vm + (math.Sin(math.Tan(factor)*math.Pow(math.Sin(factor), 10)) * vm)
}
```

- The `makeCSS` provides a custom function that takes the internal css library `gucss.Render` and takes its output into a style tag created and added into the DOM.

```go
// css takes a document and a gucss.Render, creating the need dom element
// with the content of the rendered css stylesheet.
func makeCSS(doc dom.Document, css gucss.Render) {
	styl := doc.CreateElement("style")
	doc.QuerySelector("head").AppendChild(styl)

	var data bytes.Buffer
	css.Render(&data)

	styl.SetTextContent(string(data.Bytes()))
}
```

-  The `main` function contains the real core meat of our demo. It wraps the css styles loaded as a style tags and
creates a new dashboard with 3 stat items added.

Creating a view is as simple as calling the `guviews.New` or `guviews.NewWithID` which has its own usage especially when concerning server-sided rendering.

```go
	dashboardView := guviews.New(dashboard)
	dashboardView.Mount(doc.QuerySelector("body").Underlying())

```
We mount the view on the body of the page it gets loaded with, although we use the [DOM](honnef.co/go/js/dom) package by [Dominik Honnef](https://github.com/dominikh), which provides a more friendly API to interact with the DOM, yet views use the underline `js.Object`, which is why that is being passed into the `Mount` function of the view.

*Now our view should be rendered on the DOM and ready to be updated at anytime*

- The update sequence is very simple for the demo, has it updates the values of each stat item on the dashboard as within the code.

```go
	factor := 0.05

	go func() {
		for {
			<-time.After(300 * time.Millisecond)
			visitors.Set(int(wave(factor, visitors.Value)))
			clickRate.Set(int(wave(factor, clickRate.Value)))
			bounceRate.Set(int(wave(factor, bounceRate.Value)))
			factor += 0.1
		}
	}()
```

Using the wave function, we take an incremental factor and current value of each stat, to update each with their next value. Now the beauty of this is the fact the DOM updates, with minimal effort, these is because we had embedded the `guviews.Reactive` into both the `Dashboard`, which subscribes to the `guviews.Reactive` within all its `StatItem`, with such a clean, bottom-up approach, we get the benefit of reactivity with elements and their parents directly to their respective views.

*Ofcourse, the parent component must subscribe to their children adequately to have this same effect, as within the code which puts that responsibility on the developer*


## SubViews
I thought hard and long on the concepts of subviews after stripping the need that subcomponents should have their own views within their parent components. Naturally there are cases where a one view actually depends on another for it's update/re-render signal or even hides itself in relation to another.
Due to this two major needs, views are allowed to bind to one another, where an update operation on one view will cascade to another view or sets of views, moreover views can also bind not just for update but also for display status where if another view hides itself, this cascades to the others to hide themselves as well.

Accomplishing this is rather simple, as demonstrated with the below pseudo-code:

```go

	dashboard := NewDashboard()
	dashboard.Add(NewStatItem("Clicks",100))
	dashView := guviews.New(dashboard)

	gameboard := NewDashboard()
	gameboard.Add(NewStatItem("Wins",400))
	gameView := guviews.New(gameboard)

	// To have dashView tell gameView to update, we use the `Bind` method.
	gameView.Bind(dashView)

	// To have dashView tell gameView to hide/show itself, we use the `Sync` method.
	gameView.Sync(dashView)

	// To have dashView tell gameView to both update or hide/show itself, we use the `Sync` method.
	gameView.BindSync(dashView)

```

## Views and Browser History
