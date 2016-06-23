# Components
Components in `Gu` are simple structures that satisfy the requirement of the
`Renderable` interface. This can be a single `struct` or a composite struct that
combines multiple `Renderable` structures to make up it's rendering.
This simple thinking provides a sensible organization approach for the developer
has he/she plans out the structures of their application.

### Example:

To demonstrate how a component can be built together, we will sample a demonstration for a dashboard display that allows us to details different metrics.


```go

// StatItem defines a component which defines a stat item on a dashboard.
type StatItem struct {
	guviews.Reactive
	Name  string
	Value interface{}
}

// NewStatItem returns a new instance of a StatItem.
func NewStatItem(name string, val interface{}) *StatItem{
	return &StatItem{
		Reactive: guviews.NewReactive(),
		Name: name,
		Value: val,
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
	stats []StatItem
}

// NewDashboard returns a new instance of a Dashboard.
func NewDashboard() *Dashboard{
	return &Dashboard{Reactive: guviews.NewReactive()}
}

// Add adds a new stat into the list of stats displayed by the dashboard
func (d Dashboard) Add(stat StatItem) {
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

The two main components within our example is the `StatItem` and the `Dashboard`, where the `StatItem` defines the individual items and their representation for the page. Each distinctively embeds the `guviews.Reactive` interface and
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

```
