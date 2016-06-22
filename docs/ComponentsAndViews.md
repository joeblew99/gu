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

# Views
Views are the core mechanism by which components get exposed to the DOM. They manage the rendering and update of the DOM by listening for change notifications. Due to the need for simplicity, this library ensures to only provided the minimal
needed tools to attain the best result, that is, it does not have an opinionated approach for how developers should use it but rather suggest best approach that bring best results.
