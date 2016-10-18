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

	if rs == nil {
		panic("No Resource root has been set")
	}

	return rs
}

// Resources defines a structure which contains the fully embodiement of different resources.
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

// MustCurrentResource will panic if no resource exists currently in this resource root.
func (rs *Resources) MustCurrentResource() *ResourceDefinition {
	res := rs.CurrentResource()
	if res == nil {
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
	dsl DSL

	views    []targetViews
	markups  []targetMarkup
	links    []gu.StaticView
	resolver dispatch.Resolvable
}

// Initialize runs the underline DSL for the
func (rd *ResourceDefinition) Init() {
	rd.dsl()
}

// Resource creates a new resource addding into the resource lists for the root.
func Resource(dsl DSL) *ResourceDefinition {
	var rs ResourceDefinition
	rs.dsl = dsl

	root := GetCurrentResources()
	root.Resources = append(root.Resources, rs)
	return &rs
}

//==============================================================================

// UseRoute sets the giving route for the currently used resource of the giving resource root.
func UseRoute(path string) {
	GetCurrentResources().MustCurrentResource().resolver = dispatch.NewResolver(path)
}

//==============================================================================

type targetMarkup struct {
	View    gu.StaticView
	Targets []string
}

// Markup returns a new instance of a provided value which either is a function
// which returns a needed trees.Markup or a trees.Markup or slice of trees.Markup
// itself.
func Markup(markup gu.Viewable, target interface{}, defered bool) {
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

		current.markups = append(current.markups, targetMarkup{
			Targets: targets,
			View:    static,
		})
	}
}

//==============================================================================

type targetViews struct {
	View    gu.RenderView
	Targets string
}

// View creates a gu.RenderView and applies it to the provided resource.
func View(vrs gu.Viewable, target string) gu.RenderView {
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

	view := gu.CustomView("section", events.NewEventManager(), rs...)

	current := GetCurrentResources().MustCurrentResource()
	current.views = append(current.views, targetViews{
		View:    view,
		Targets: target,
	})

	return view
}

//==============================================================================

// Title adds a element which generates a <title> tag.
func Title(title string) {
	ml := mLink("title", false)
	trees.NewText(title).Apply(ml.Content)
}

// Link adds a element which generates a <link> tag.
func Link(url string, mtype string, defered bool) {
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

// Style adds a element which generates a <style> tag.
func Style(src string, defered bool) {
	ml := mLink("style", defered)
	trees.NewAttr("src", src).Apply(ml.Content)
	trees.NewAttr("type", "text/css").Apply(ml.Content)
}

// Script adds a element which generates a <style> tag.
func Script(src string, mtype string, defered bool) {
	ml := mLink("script", defered)
	trees.NewAttr("src", src).Apply(ml.Content)
	trees.NewAttr("type", mtype).Apply(ml.Content)
}

// Meta adds a element which generates a <style> tag.
func Meta(props map[string]string) {
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
	current.links = append(current.links, static)
	return static
}
