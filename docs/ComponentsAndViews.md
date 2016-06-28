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


# Views
Views are the core mechanism by which components get exposed to the DOM. They manage the rendering and update of the DOM by listening for change notifications. Due to the need for simplicity, this library ensures to only provided the minimal
needed tools to attain the best result, that is, it does not have an opinionated approach for how developers should use it but rather suggest best approach that bring best results.

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

	gu.RenderAsBody(guviews.New(dashboard))

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
