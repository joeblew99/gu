package design

import (
	"sync"

	"github.com/influx6/gu"
	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/events"
	"github.com/influx6/gu/trees"
)

//==============================================================================

var collection = struct {
	root       *Resources
	cl         sync.Mutex
	collection []DSL
}{}

func getResources() *Resources {
	collection.cl.Lock()
	{
		if collection.root != nil {
			collection.cl.Unlock()
			return collection.root
		}
	}
	collection.cl.Unlock()

	panic("No Root Resource Manager created.")
}

//==============================================================================

// ResourceRenderer  defines an interface for a resource rendering structure which handles rendering
// of a giving resource.
type ResourceRenderer interface {
	Render(...ResourceDefinition)
	RenderUpdate(gu.Renderable, string)
}

// Resources defines a structure which contains the fully embodiement of different resources.
// It contains a stack which will be the current rendering resources for the current match paths.
// It also contains a nStack which contains resource which must be rendered regardless of the
// current resources available.
type Resources struct {
	Resources []ResourceDefinition
	renderer  ResourceRenderer
	lastPath  dispatch.Path
}

// Resource creates a new resource addding into the resource lists for the root.
func Resource(dsl DSL) {
	collection.cl.Lock()
	{
		collection.collection = append(collection.collection, dsl)
	}
	collection.cl.Unlock()
}

// New creates a new instance of a Resources struct and registers it as the currently used resources
// root.
func New(renderer ...ResourceRenderer) *Resources {
	var rn ResourceRenderer
	if len(renderer) > 0 {
		rn = renderer[0]
	}

	var res Resources
	res.renderer = rn

	collection.cl.Lock()
	{
		collection.root = &res
	}
	collection.cl.Unlock()

	return collection.root
}

// Init intializes all attached resources under its giving resource list.
func (rs *Resources) Init(useHashOnly ...bool) *Resources {
	var watchHash bool

	if len(useHashOnly) != 0 {
		watchHash = useHashOnly[0]
	}

	dispatch.Subscribe(func(pd dispatch.PathDirective) {
		if !watchHash {
			psx := dispatch.UseDirective(pd)

			if psx.String() == rs.lastPath.String() {
				return
			}

			rs.lastPath = psx

			if rs.renderer != nil {
				rs.renderer.Render(rs.Resolve(psx)...)
			}

			return
		}

		psx := dispatch.UseHashDirective(pd)
		if psx.String() == rs.lastPath.String() {
			return
		}

		rs.lastPath = psx

		if rs.renderer != nil {
			rs.renderer.Render(rs.Resolve(psx)...)
		}
	})

	for _, dsl := range collection.collection {
		res := newResource(rs, dsl)
		rs.Resources = append(rs.Resources, *res)
		rs.Init()
	}

	return rs
}

// Resolve returns the giving definitions that matches the provided path.
func (rs *Resources) Resolve(path dispatch.Path) []ResourceDefinition {
	res := rs.ResolveWith(path.Rem)

	for _, rs := range res {
		rs.Resolver.Resolve(path)
		rs.active = true
	}

	return res
}

// ResolveWith resolves the giving path by testing all available resource definitions to be
// rendered and returns a slice of all Resources sorted against the order which they should be
// rendered.
func (rs *Resources) ResolveWith(path string) []ResourceDefinition {
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

// Render creates a complete markup definition using the giving set of Resource
// Definition by applying the giving path.
func (rs *Resources) Render(path dispatch.Path) *trees.Markup {
	return rs.render(rs.Resolve(path)...)
}

// render performs the needed work of collecting the giving markups and
// creating the complete html rendering.
func (rs *Resources) render(rsx ...ResourceDefinition) *trees.Markup {
	var head = trees.NewMarkup("head", false)
	var body = trees.NewMarkup("body", false)

	for _, rx := range rsx {

		for _, item := range rx.Links {
			item.Render().ApplyMorphers().Apply(head)
		}

		for _, item := range rx.Renderables {
			markup := item.View.Render().ApplyMorphers()

			if item.Targets != "" {
				for _, target := range trees.MarkupQueries.QueryAll(body, item.Targets) {
					markup.Apply(target)
				}

				continue
			}

			markup.Apply(body)
		}

		// Render all deferred renderables into the body markup
		for _, item := range rx.DRenderables {
			markup := item.View.Render().ApplyMorphers()

			if item.Targets != "" {
				for _, target := range trees.MarkupQueries.QueryAll(body, item.Targets) {
					markup.Apply(target)
				}

				continue
			}

			markup.Apply(body)
		}

		for _, item := range rx.DeferLinks {
			item.Render().ApplyMorphers().Apply(body)
		}
	}

	htmlMarkup := trees.NewMarkup("html", false)
	head.Apply(htmlMarkup)
	body.Apply(htmlMarkup)

	return htmlMarkup
}

// DSL defines a function type which is used to generate the contents of a Def(Definition).
type DSL func()

// ResourceViewUpdate defines a view update notification which contains the name of the
// view to be notified for an update.
type ResourceViewUpdate struct {
	View     string
	Resource string
}

// newResource creates a new ResourceDefinition instance and adds the
// new resource into the root ResourceDefinition list.
func newResource(root *Resources, dsl DSL) *ResourceDefinition {
	var rs ResourceDefinition
	rs.Dsl = dsl
	rs.Order = Any
	rs.Root = root
	rs.uuid = gu.NewKey()
	rs.Renderer = root.renderer

	root.Resources = append(root.Resources, rs)

	rsp := &rs
	dispatch.Subscribe(func(rv ResourceViewUpdate) {
		if rv.Resource != rsp.uuid {
			return
		}

		if !rsp.active {
			return
		}

	})

	return rsp
}

// ResourceDefinition defines a high-level definition for managing resources for
// which other definitions build from.
type ResourceDefinition struct {
	Dsl    DSL
	active bool
	uuid   string

	Views        []gu.RenderView
	Links        []gu.StaticView
	DeferLinks   []gu.StaticView
	Renderables  []targetRenderable
	DRenderables []targetRenderable

	Order    RenderingOrder
	Renderer ResourceRenderer
	Resolver dispatch.Resolver
	Root     *Resources
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
		rd.Resolver.ResolvedPassed(view.Resolve)
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
	getResources().MustCurrentResource().Order = mode
}

//==============================================================================

// UseRoute sets the giving route for the currently used resource of the giving resource root.
func UseRoute(path string) {
	getResources().MustCurrentResource().Resolver = dispatch.NewResolver(path)
}

//==============================================================================

// Viewable defines a generic interface for a generic return type. It exists to
// give symantic representation in the areas it is used to express the expected
// returned to be one of a Viewable souce in the context of the gu library.
type Viewable interface{}

type targetRenderable struct {
	View    gu.Renderable
	Targets string
}

// Markup returns a new instance of a provided value which either is a function
// which returns a needed trees.Markup or a trees.Markup or slice of trees.Markup
// itself.
func Markup(markup Viewable, targets string, immediateRender ...bool) {
	var immediate bool

	if len(immediateRender) == 0 {
		immediate = immediateRender[0]
	}

	var markupFn []*trees.Markup

	switch mo := markup.(type) {
	case func() []*trees.Markup:
		markupFn = mo()
	case []*trees.Markup:
		markupFn = mo
	case func() *trees.Markup:
		markupFn = []*trees.Markup{mo()}
	case *trees.Markup:
		markupFn = []*trees.Markup{mo}
	case string:
		mp, err := trees.ParseTree(mo)
		if err != nil {
			panic("Unable to parse markup string: " + err.Error())
		}

		markupFn = mp
	default:
		panic("Unknown markup processable type")
	}

	current := getResources().MustCurrentResource()
	for _, markup := range markupFn {
		var static gu.StaticView
		static.Content = markup

		trees.NewAttr("resource-id", current.UUID()).Apply(static.Content)

		if len(targets) != 0 && immediate {
			current.Renderables = append(current.Renderables, targetRenderable{
				Targets: targets,
				View:    &static,
			})

			continue
		}

		current.DRenderables = append(current.DRenderables, targetRenderable{
			Targets: targets,
			View:    &static,
		})
	}
}

//==============================================================================

// View creates a gu.RenderView and applies it to the provided resource.
func View(vrs Viewable, target string, immediateRender ...bool) gu.RenderView {
	var immediate bool

	if len(immediateRender) == 0 {
		immediate = immediateRender[0]
	}

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

	current := getResources().MustCurrentResource()
	view := gu.CustomView("section", events.NewEventManager(), rs...)

	if rvw, ok := view.(gu.Reactive); ok {
		rvw.React(func() {
			dispatch.Dispatch(ResourceViewUpdate{
				View:     view.UUID(),
				Resource: current.UUID(),
			})
		})
	}

	current.Views = append(current.Views, view)

	if target != "" && immediate {
		current.Renderables = append(current.Renderables, targetRenderable{
			Targets: target,
			View:    view,
		})

		return view
	}

	if target != "" {
		current.DRenderables = append(current.DRenderables, targetRenderable{
			View:    view,
			Targets: target,
		})

		return view
	}

	current.Renderables = append(current.Renderables, targetRenderable{
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
	static.Content = trees.NewMarkup(tag, false)

	current := getResources().MustCurrentResource()

	if deffer {
		current.DeferLinks = append(current.DeferLinks, static)
	} else {
		current.Links = append(current.Links, static)
	}

	return static
}

//==============================================================================
