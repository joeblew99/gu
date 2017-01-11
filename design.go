package gu

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu/css"
	"github.com/gu-io/gu/dispatch"
	"github.com/gu-io/gu/shell"
	"github.com/gu-io/gu/shell/cache"
	"github.com/gu-io/gu/shell/fetch"
	"github.com/gu-io/gu/trees"
)

//==============================================================================

// DSL defines a function type which is used to generate the contents
// of a Def(Definition).
type DSL func(shell.Fetch, shell.Cache)

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
	Render(...*ResourceDefinition)
	BindEvent(*trees.Event, *js.Object)
	RenderUpdate(Renderable, string, bool)
	TriggerBindEvent(*js.Object, *js.Object, *trees.Event)
}

// Resource creates a new resource addding into the resource lists for the root.
func Resource(dsl interface{}) int {
	var udsl DSL

	switch dslo := dsl.(type) {
	case func(shell.Fetch, shell.Cache):
		udsl = dslo
	case func(shell.Fetch):
		udsl = func(f shell.Fetch, _ shell.Cache) {
			dslo(f)
		}
	case func():
		udsl = func(f shell.Fetch, _ shell.Cache) {
			dslo()
		}
	default:
		panic("type not supported for Resource argument")
	}

	var index int

	collection.cl.Lock()
	{
		collection.collection = append(collection.collection, udsl)
		index = len(collection.collection)
	}
	collection.cl.Unlock()

	return index
}

// Resources defines a structure which contains the fully embodiement of different resources.
// It contains a stack which will be the current rendering resources for the current match paths.
// It also contains a nStack which contains resource which must be rendered regardless of the
// current resources available.
type Resources struct {
	lastPath  dispatch.Path
	Resources []*ResourceDefinition
	renderer  ResourceRenderer
	cache     shell.Cache
	fetch     shell.Fetch
	manifests []shell.AppManifest
}

// New creates a new instance of a Resources struct which expects a unqiue name
// and and a slice of renderers if desired.
func New(appName string, renderer ...ResourceRenderer) *Resources {
	var rn ResourceRenderer
	if len(renderer) > 0 {
		rn = renderer[0]
	}

	var res Resources
	res.renderer = rn
	res.cache = cache.New(appName)
	res.fetch = fetch.New(res.cache)

	return &res
}

// GetIndex returns the resource for this giving index.
func (rs *Resources) GetIndex(index int) *ResourceDefinition {
	return rs.Resources[index]
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

	var dslList []DSL

	collection.cl.Lock()
	{
		dslList = collection.collection
		collection.root = rs
		collection.collection = nil
	}
	collection.cl.Unlock()

	// Attempt to retrieve a manifest.json from the backend.
	if _, manifestResponse, err := rs.fetch.Get("manifest.json", shell.DefaultStrategy); err == nil {

		// We sucessfully retrieved response.
		var appManifest []shell.AppManifest

		if err := json.Unmarshal(manifestResponse.Body, &appManifest); err != nil {
			errorMsg := fmt.Sprintf("Failed to load shell.AppManifest, resource loading is unavailable")
			panic(errorMsg)
		}

		rs.manifests = appManifest
	}

	for _, dsl := range dslList {
		res := newResource(rs, dsl)
		res.Init()
	}

	collection.cl.Lock()
	{
		collection.root = nil
	}
	collection.cl.Unlock()

	if detect.IsBrowser() && rs.renderer != nil {
		du := rs.Resolve(dispatch.GetLocationHashAsPath())
		rs.renderer.Render(du...)
	}

	return rs
}

// Resolve returns the giving definitions that matches the provided path.
func (rs *Resources) Resolve(path dispatch.Path) []*ResourceDefinition {
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
func (rs *Resources) ResolveWith(path string) []*ResourceDefinition {
	var first, last, any []*ResourceDefinition

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

			continue
		}

		if resource.active {
			resource.Unmount()
			resource.active = false
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

	return rs.Resources[rlen-1]
}

// Render creates a complete markup definition using the giving set of Resource
// Definition by applying the giving path.
func (rs *Resources) Render(path string) *trees.Markup {
	return rs.RenderPath(dispatch.UseLocation(path))
}

// RenderWithScript creates a complete markup definition using the giving set of Resource
// Definition by applying the giving path.
func (rs *Resources) RenderWithScript(path string, script string) *trees.Markup {
	return rs.RenderPathWithScript(dispatch.UseLocation(path), script)
}

// RenderPath creates a complete markup definition using the giving set of Resource
// Definition by applying the giving dispatch.Path.
func (rs *Resources) RenderPath(path dispatch.Path) *trees.Markup {
	result := rs.Resolve(path)
	return rs.render(nil, result...)
}

// RenderPathWithScript creates a complete markup definition using the giving set of Resource
// Definition by applying the giving dispatch.Path, and adding the script path as a script tag.
func (rs *Resources) RenderPathWithScript(path dispatch.Path, script string) *trees.Markup {
	result := rs.Resolve(path)

	if script != "" {
		src := trees.NewAttr("src", script)
		srctype := trees.NewAttr("gu-script-root", "true")
		scriptElem := trees.NewMarkup("script", false)

		src.Apply(scriptElem)
		srctype.Apply(scriptElem)

		return rs.render([]*trees.Markup{scriptElem}, result...)
	}

	return rs.render(nil, result...)
}

// loadGlobalResources loads the provided Resources specified through the
// manifests
func (rs *Resources) loadGlobalResources() {

}

// render performs the needed work of collecting the giving markups and
// creating the complete html rendering.
func (rs *Resources) render(attachElems []*trees.Markup, rsx ...*ResourceDefinition) *trees.Markup {
	var head = trees.NewMarkup("head", false)
	var body = trees.NewMarkup("body", false)

	for _, rx := range rsx {

		for _, item := range rx.Links {
			item.Render().Apply(head)
		}

		for _, item := range rx.Renderables {
			markup := item.View.Render()

			if item.Targets != "" {
				for _, target := range trees.Query.QueryAll(body, item.Targets) {
					markup.Apply(target)
				}

				continue
			}

			markup.Apply(body)
		}

		// Render all deferred renderables into the body markup
		for _, item := range rx.DRenderables {
			markup := item.View.Render()

			if item.Targets != "" {
				for _, target := range trees.Query.QueryAll(body, item.Targets) {
					markup.Apply(target)
				}

				continue
			}

			markup.Apply(body)
		}

		for _, item := range rx.DeferLinks {
			item.Render().Apply(body)
		}
	}

	for _, item := range attachElems {
		item.Apply(body)
	}

	htmlMarkup := trees.NewMarkup("html", false)
	head.Apply(htmlMarkup)
	body.Apply(htmlMarkup)

	return htmlMarkup
}

// ResourceDefinition defines a high-level definition for managing resources for
// which other definitions build from.
type ResourceDefinition struct {
	Dsl    DSL
	active bool
	uuid   string

	ViewHooks    []ViewHooks
	Views        []RenderView
	Links        []StaticView
	DeferLinks   []StaticView
	Renderables  []targetRenderable
	DRenderables []targetRenderable

	Order    RenderingOrder
	Manager  *RouteManager
	Renderer ResourceRenderer
	Resolver dispatch.Resolver
	Root     *Resources
}

// ResourceViewUpdate defines a view update notification which contains the name of the
// view to be notified for an update.
type ResourceViewUpdate struct {
	View     Renderable
	Resource string
	Target   string
}

// newResource creates a new ResourceDefinition instance and adds the
// new resource into the root ResourceDefinition list.
func newResource(root *Resources, dsl DSL) *ResourceDefinition {
	var rs ResourceDefinition
	rs.Dsl = dsl
	rs.Order = Any
	rs.Root = root
	rs.uuid = NewKey()
	rs.Manager = NewRouteManager()
	rs.Resolver = dispatch.NewResolver("*")

	rsp := &rs
	root.Resources = append(root.Resources, rsp)

	dispatch.Subscribe(func(rv *ResourceViewUpdate) {

		if rv.Resource != rsp.uuid {
			return
		}

		if !rsp.active {
			return
		}

		rs.Root.renderer.RenderUpdate(rv.View, rv.Target, true)
	})

	return rsp
}

// Unmount propagates to all views which define the ViewHooks
func (rd *ResourceDefinition) Unmount() {
	for _, view := range rd.ViewHooks {
		mount, rendered, unmount := view.Hooks()
		unmount.Publish()

		// Reset mount and rendering subscriptions since we are removing this from
		// the DOM.
		mount.Reset()
		rendered.Reset()
	}
}

// UUID returns the uuid associated with the ResourceDefinition.
func (rd *ResourceDefinition) UUID() string {
	return rd.uuid
}

// Init initialize runs the underline DSL for the operation, loads all resolvers
// and connects all internal views for immediate usage.
func (rd *ResourceDefinition) Init() {
	rd.Dsl(rd.Root.fetch, rd.Root.cache)
	rd.Resolver.Flush()

	for _, level := range rd.Manager.Levels {
		rd.Resolver.Register(level)
	}

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

// DoOrder defines a high level function which sets/resets the RenderingOrder of the current ResourceDefinition.
func DoOrder(mode RenderingOrder) {
	getResources().MustCurrentResource().Order = mode
}

//==============================================================================

// UseRoute sets the giving route for the currently used resource of the giving resource root.
func UseRoute(path string) {
	getResources().MustCurrentResource().Resolver = dispatch.NewResolver(path)
}

// DoLocalRouter returns a RouteApplier which can be used to localize the
// internal routes for the base resource router and allow usage with markup
// and views.
func DoLocalRouter(basePath string, mx ...trees.SwitchMorpher) RouteApplier {
	var mo trees.SwitchMorpher

	if len(mx) != 0 {
		mo = mx[0]
	} else {
		mo = &trees.RemoveMorpher{}
	}

	return getResources().MustCurrentResource().Manager.Level(basePath, mo)
}

// DoRouteLevel takes a giving route applier returning the last route created
// from the route levels of the applier.
func DoRouteLevel(level RouteApplier, routes ...string) RouteApplier {
	if len(routes) == 0 {
		return level
	}

	var last RouteApplier

	for _, route := range routes {
		if last == nil {
			last = level.N(route)
			continue
		}

		last = last.N(route)
	}

	return last
}

// DoRoute allows definition of a route level to applying all the giving
// routes path returning the last applier.
func DoRoute(base string, routes ...string) RouteApplier {
	return DoRouteLevel(DoLocalRouter(base), routes...)
}

//==============================================================================

// Viewable defines a generic interface for a generic return type. It exists to
// give symantic representation in the areas it is used to express the expected
// returned to be one of a Viewable souce in the context of the gu library.
type Viewable interface{}

type targetRenderable struct {
	View    Renderable
	Targets string
}

// DoMarkup returns a new instance of a provided value which either is a function
// which returns a needed trees.Markup or a trees.Markup or slice of trees.Markup
// itself.
func DoMarkup(markup Viewable, targets string, deferRender bool, targetAlreadyInDom bool) {
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
		markupFn = trees.ParseTree(mo)
	default:
		panic("Unknown markup processable type")
	}

	current := getResources().MustCurrentResource()

	for _, markup := range markupFn {
		static := Static(markup)
		static.Morph = true

		trees.NewAttr("resource-id", current.UUID()).Apply(static.Content)

		if deferRender {
			current.DRenderables = append(current.DRenderables, targetRenderable{
				Targets: targets,
				View:    static,
			})

			continue
		}

		if !deferRender && targetAlreadyInDom {
			current.Renderables = append(current.Renderables, targetRenderable{
				Targets: targets,
				View:    static,
			})

			continue
		}

		if !deferRender && !targetAlreadyInDom && targets != "" {
			current.DRenderables = append(current.DRenderables, targetRenderable{
				View:    static,
				Targets: targets,
			})

			continue
		}

		current.Renderables = append(current.Renderables, targetRenderable{
			View:    static,
			Targets: targets,
		})

	}
}

// DoHead adds the provided markup as part of the children to the head tag.
func DoHead(markup Viewable, deferRender bool) {
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
		markupFn = trees.ParseTree(mo)
	default:
		panic("Unknown markup processable type")
	}

	current := getResources().MustCurrentResource()

	for _, item := range markupFn {
		var static StaticView
		static.Morph = true
		static.Content = item

		trees.NewAttr("resource-id", current.UUID()).Apply(static.Content)

		if deferRender {
			current.DeferLinks = append(current.DeferLinks, static)
			continue
		}

		current.Links = append(current.Links, static)
	}
}

// DoStyle adds a element which generates a <style> tag.
func DoStyle(styles interface{}, bind interface{}, deferRender bool) {
	var rs *css.Rule

	switch so := styles.(type) {
	case string:
		rs = css.New(so)
	case *css.Rule:
		rs = so
	default:
		panic("Invalid Acceptable type for css: Only string or *css.Rule")
	}

	current := getResources().MustCurrentResource()

	var static StaticView
	static.Morph = true
	static.Content = trees.CSSStylesheet(rs, bind)

	trees.NewAttr("resource-id", current.UUID()).Apply(static.Content)

	if deferRender {
		current.DeferLinks = append(current.DeferLinks, static)
		return
	}

	current.Links = append(current.Links, static)
}

//==============================================================================

// DoView creates a RenderView and applies it to the provided resource.
func DoView(vrs Viewable, targets string, deferRender bool, targetAlreadyInDom bool) RenderView {
	var rs Renderables

	switch vwo := vrs.(type) {
	case Renderables:
		rs = vwo
		break
	case Renderable:
		rs = Renderables{vwo}
		break
	case func() Renderable:
		rs = Renderables{vwo()}
		break
	case func() Renderables:
		rs = vwo()
		break
	default:
		panic("View must either recieve a function that returns Renderables or Renderables themselves")
	}

	master := getResources()
	current := master.MustCurrentResource()
	view := CustomView("section", rs...)

	if vh, ok := view.(ViewHooks); ok {
		current.ViewHooks = append(current.ViewHooks, vh)
	}

	if fv, ok := view.(Fetchable); ok {
		fv.UseFetch(master.fetch, master.cache)
	}

	if rvw, ok := view.(Reactive); ok {
		rvw.React(func() {
			dispatch.Dispatch(&ResourceViewUpdate{
				View:     view,
				Target:   targets,
				Resource: current.UUID(),
			})
		})
	}

	current.Views = append(current.Views, view)

	if deferRender {
		current.DRenderables = append(current.DRenderables, targetRenderable{
			View:    view,
			Targets: targets,
		})

		return view
	}

	if !deferRender && targetAlreadyInDom {
		current.Renderables = append(current.Renderables, targetRenderable{
			View:    view,
			Targets: targets,
		})

		return view
	}

	if !deferRender && !targetAlreadyInDom && targets != "" {
		current.DRenderables = append(current.DRenderables, targetRenderable{
			View:    view,
			Targets: targets,
		})

		return view
	}

	current.Renderables = append(current.Renderables, targetRenderable{
		View:    view,
		Targets: targets,
	})

	return view
}

//==============================================================================

// DoTitle adds a element which generates a <title> tag.
func DoTitle(title string) {
	ml := mLink("title", false)
	trees.NewText(title).Apply(ml.Content)
}

// DoLink adds a element which generates a <link> tag.
func DoLink(url string, mtype string, defered bool) {
	ml := mLink("link", defered)
	trees.NewAttr("href", url).Apply(ml.Content)
	trees.NewAttr("rel", mtype).Apply(ml.Content)
}

// DoCSS adds a element which generates a <style> tag.
func DoCSS(src string, defered bool) {
	ml := mLink("link", defered)
	trees.NewAttr("href", src).Apply(ml.Content)
	trees.NewAttr("rel", "stylesheet").Apply(ml.Content)
	trees.NewAttr("type", "text/css").Apply(ml.Content)
}

// DoScript adds a element which generates a <style> tag.
func DoScript(src string, mtype string, defered bool) {
	ml := mLink("script", defered)
	trees.NewAttr("src", src).Apply(ml.Content)
	trees.NewAttr("type", mtype).Apply(ml.Content)
}

// DoMeta adds a element which generates a <style> tag.
func DoMeta(props map[string]string) {
	ml := mLink("meta", false)
	for name, val := range props {
		trees.NewAttr(name, val).Apply(ml.Content)
	}
}

// mLink adds tagName with the provided value into the header bar for the
// page content.
func mLink(tag string, deffer bool) StaticView {
	var static StaticView
	static.Morph = true
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
