package design

import (
	"sync"

	"github.com/influx6/gu"
	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/events"
	"github.com/influx6/gu/trees"
)

// Root defines a package level
var rootRes = struct {
	ml   sync.Mutex
	root *Resources
}{}

// UseResources allows setting the giving root by which resources functions are executed against.
func UseResources(rs *Resources) {
	rootRes.ml.Lock()
	rootRes.root = rs
	rootRes.ml.Unlock()
}

// GetCurrentResources returns the current resources being used by the design functions.
func GetCurrentResources() *Resources {
	rootRes.ml.Lock()
	rs := rootRes.root
	rootRes.ml.Unlock()

	if rs != nil {
		return rs
	}

	panic("No Resource root has been set")
}

// Viewable defines a generic interface for a generic return type. It exists to
// give symantic representation in the areas it is used to express the expected
// returned to be one of a Viewable souce in the context of the gu library.
type Viewable interface{}

// ResourceRenderer  defines an interface for a resource rendering structure which handles rendering
// of a giving resource.
type ResourceRenderer interface {
	Render(dispatch.Path, ResourceDefinition)
}

// Resources defines a structure which contains the fully embodiement of different resources.
// It contains a stack which will be the current rendering resources for the current match paths.
// It also contains a nStack which contains resource which must be rendered regardless of the
// current resources available.
type Resources struct {
	Resources []ResourceDefinition
}

// New creates a new instance of a Resources struct and registers it as the currently used resources
// root.
func New() *Resources {
	var res Resources
	UseResources(&res)
	return &res
}

// Init intializes all attached resources under its giving resource list.
func (rs *Resources) Init() {
	for _, res := range rs.Resources {
		res.Init()
	}
}

// Resolve resolves the giving path by testing all available resource definitions to be
// rendered and returns a slice of all Resources sorted against the order which they should be
// rendered.
func (rs *Resources) Resolve(path string) []ResourceDefinition {
	var first, last, any []ResourceDefinition

	for _, resource := range rs.Resources {
		if _, _, passed := resource.Resolver.Test(path); passed {
			switch resource.Order {
			case First:
				first = append(first, resource)
			case Last:
				last = append(last, resource)
			case Any:
				any = append(any, resource)
			}
		}
	}

	return append(append(first, any...), last...)
}

// MustCurrentResource will panic if no resource exists currently in this resource root.
func (rs *Resources) MustCurrentResource() *ResourceDefinition {
	res := rs.CurrentResource()
	if res != nil {
		return res
	}

	panic("No resource allocated in current resources root")
}

// CurrentResource returns the current available resource for this root.
func (rs *Resources) CurrentResource() *ResourceDefinition {
	rlen := len(rs.Resources)
	if rlen == 0 {
		return nil
	}

	return &(rs.Resources[rlen-1])
}

// DSL defines a function type which is used to generate the contents of a Def(Definition).
type DSL func()

// ResourceDefinition defines a high-level definition for managing resources for
// which other definitions build from.
type ResourceDefinition struct {
	Dsl  DSL
	uuid string

	Views      []targetViews
	Markups    []targetMarkup
	Links      []gu.StaticView
	DeferLinks []gu.StaticView

	Order    RenderingOrder
	Resolver dispatch.Resolver
}

// Resource creates a new resource addding into the resource lists for the root.
func Resource(dsl DSL) *ResourceDefinition {
	var rs ResourceDefinition
	rs.Order = Any
	rs.Dsl = dsl
	rs.uuid = gu.NewKey()

	root := GetCurrentResources()
	root.Resources = append(root.Resources, rs)
	return &rs
}

// UUID returns the uuid associated with the ResourceDefinition.
func (rd *ResourceDefinition) UUID() string {
	return rd.uuid
}

// Initialize runs the underline DSL for the
func (rd *ResourceDefinition) Init() {
	rd.Dsl()

	// If no resolver is provided then use a all path resolver.
	if rd.Resolver == nil {
		rd.Resolver = dispatch.NewResolver("*")
	}

	rd.Resolver.Flush()

	for _, view := range rd.Views {
		rd.Resolver.ResolvedPassed(view.View.Resolve)
	}
}

//==============================================================================

// RenderingOrder defines a type used to define the order which rendering is to be done for a resource.
type RenderingOrder int

const (
	// First defines that a resouce should be among the first rendered.
	First RenderingOrder = iota + 1

	// Last defines that a resouce should be among the last rendered.
	Last

	// Any defines that a resouce should be rendered in what ever position it appears.
	Any
)

// Order defines a high level function which sets/resets the RenderingOrder of the current ResourceDefinition.
func Order(mode RenderingOrder) {
	GetCurrentResources().MustCurrentResource().Order = mode
}

//==============================================================================

// UseRoute sets the giving route for the currently used resource of the giving resource root.
func UseRoute(path string) {
	GetCurrentResources().MustCurrentResource().Resolver = dispatch.NewResolver(path)
}

//==============================================================================

type targetMarkup struct {
	View    gu.StaticView
	Targets []string
}

// Markup returns a new instance of a provided value which either is a function
// which returns a needed trees.Markup or a trees.Markup or slice of trees.Markup
// itself.
func Markup(markup Viewable, target interface{}, defered bool) {
	var markupFn []trees.Markup

	switch mo := markup.(type) {
	case func() []trees.Markup:
		markupFn = mo()
	case []trees.Markup:
		markupFn = mo
	case func() trees.Markup:
		markupFn = []trees.Markup{mo()}
	case trees.Markup:
		markupFn = []trees.Markup{mo}
	case string:
		mp, err := trees.ParseTree(mo)
		if err != nil {
			panic("Unable to parse markup string: " + err.Error())
		}

		markupFn = mp
	default:
		panic("Unknown markup processable type")
	}

	var targets []string

	switch to := target.(type) {
	case string:
		targets = []string{to}
	case []string:
		targets = to
	default:
		panic("targets only allowed to be a string or a slice of string")
	}

	current := GetCurrentResources().MustCurrentResource()
	for _, markup := range markupFn {
		var static gu.StaticView
		static.Content = markup

		current.Markups = append(current.Markups, targetMarkup{
			Targets: targets,
			View:    static,
		})
	}
}

//==============================================================================

// ResourceViewUpdate defines a view update notification which contains the name of the
// view to be notified for an update.
type ResourceViewUpdate struct {
	View     string
	Resource string
}

type targetViews struct {
	View    gu.RenderView
	Targets string
}

// View creates a gu.RenderView and applies it to the provided resource.
func View(vrs Viewable, target string) gu.RenderView {
	var rs gu.Renderables

	switch vwo := vrs.(type) {
	case gu.Renderables:
		rs = vwo
	case gu.Renderable:
		rs = gu.Renderables{vwo}
	case func() gu.Renderable:
		rs = gu.Renderables{vwo()}
	case func() gu.Renderables:
		rs = vwo()
	default:
		panic("View must either recieve a function that returns Renderables or Renderables themselves")
	}

	current := GetCurrentResources().MustCurrentResource()
	view := gu.CustomView("section", events.NewEventManager(), rs...)

	if rvw, ok := view.(gu.Reactive); ok {
		rvw.React(func() {
			dispatch.Dispatch(ResourceViewUpdate{
				View:     view.UUID(),
				Resource: current.UUID(),
			})
		})
	}

	current.Views = append(current.Views, targetViews{
		View:    view,
		Targets: target,
	})

	return view
}

//==============================================================================

// PageTitle adds a element which generates a <title> tag.
func PageTitle(title string) {
	ml := mLink("title", false)
	trees.NewText(title).Apply(ml.Content)
}

// PageLink adds a element which generates a <link> tag.
func PageLink(url string, mtype string, defered bool) {
	ml := mLink("link", defered)
	trees.NewAttr("href", url).Apply(ml.Content)
	trees.NewAttr("type", mtype).Apply(ml.Content)
}

// CSS adds a element which generates a <style> tag.
func CSS(src string, defered bool) {
	ml := mLink("link", defered)
	trees.NewAttr("href", src).Apply(ml.Content)
	trees.NewAttr("rel", "stylesheet").Apply(ml.Content)
	trees.NewAttr("type", "text/css").Apply(ml.Content)
}

// Styles adds a element which generates a <style> tag.
func Styles(src string, defered bool) {
	ml := mLink("style", defered)
	trees.NewAttr("src", src).Apply(ml.Content)
	trees.NewAttr("type", "text/css").Apply(ml.Content)
}

// Scripts adds a element which generates a <style> tag.
func Scripts(src string, mtype string, defered bool) {
	ml := mLink("script", defered)
	trees.NewAttr("src", src).Apply(ml.Content)
	trees.NewAttr("type", mtype).Apply(ml.Content)
}

// Metas adds a element which generates a <style> tag.
func Metas(props map[string]string) {
	ml := mLink("meta", false)
	for name, val := range props {
		trees.NewAttr(name, val).Apply(ml.Content)
	}
}

// mLink adds tagName with the provided value into the header bar for the
// page content.
func mLink(tag string, deffer bool) gu.StaticView {
	var static gu.StaticView
	static.Content = trees.NewElement(tag, false)

	current := GetCurrentResources().MustCurrentResource()

	if deffer {
		current.DeferLinks = append(current.DeferLinks, static)
	} else {
		current.Links = append(current.Links, static)
	}

	return static
}
