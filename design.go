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

// RenderingOrder defines a type used to define the order which rendering is to be done for a resource.
type RenderingOrder int

// Mode defines a type to represent the mode which the library works under.
type Mode int

const (
	// ProductionMode defines production mode for which defines how the context
	// which it behaves as regards certain features.
	ProductionMode Mode = iota + 30

	// DevelopmentMode defines production mode for which defines how the context
	// which it behaves as regards certain features.
	DevelopmentMode
)

const (

	// First defines that a resouce should be among the first rendered.
	First RenderingOrder = iota + 1

	// Last defines that a resouce should be among the last rendered.
	Last

	// Any defines that a resouce should be rendered in what ever position it appears.
	Any
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

//==============================================================================

// Options defines a set of configurations which the resource manager uses
// in the way it operates and works.
//
// ManifestURI:
//		 Sets the path to use when request the manifestURI which will be the definite
//		 url used to retrieve the manifest.
//
// Mode:
//		 Sets the mode by which gu works most specifically to the wiping of the
//		 browser cache which loads all resources into the cache default to have
//		 true offline working of app
//
// IgnoreManifest:
//		 Turns of the manifest capabilities in gu, hence leaving the
// 		 loading of resources to the user.
type Options struct {
	Mode           Mode   `json:"mode"`
	ManifestURI    string `json:"manifest_path"`
	IgnoreManifest bool   `json:"ignore_manifest"`
}

// Resources defines a structure which contains the fully embodiement of different resources.
// It contains a stack which will be the current rendering resources for the current match paths.
// It also contains a nStack which contains resource which must be rendered regardless of the
// current resources available.
type Resources struct {
	cacheName    string
	options      *Options
	cache        shell.Cache
	fetch        shell.Fetch
	lastPath     dispatch.Path
	renderer     ResourceRenderer
	manifests    []*shell.AppManifest
	gmanifests   []*shell.AppManifest
	Resources    []*ResourceDefinition
	headResource []*trees.Markup
	bodyResource []*trees.Markup
	initOnce     sync.Once
}

// ResourceRenderer  defines an interface for a resource rendering structure which handles rendering
// of a giving resource.
type ResourceRenderer interface {
	Render(...*ResourceDefinition)
	BindEvent(*trees.Event, *js.Object)
	RenderUpdate(Renderable, string, bool)
	TriggerBindEvent(*js.Object, *js.Object, *trees.Event)
}

// New creates a new instance of a Resources struct which expects a unqiue name
// and and a slice of renderers if desired.
func New(appName string, op *Options, renderer ...ResourceRenderer) *Resources {
	var rn ResourceRenderer
	if len(renderer) > 0 {
		rn = renderer[0]
	}

	var res Resources
	res.cacheName = appName
	res.options = op
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
func (rs *Resources) Init(useHashOnly bool) *Resources {
	rs.initOnce.Do(func() {
		rs.initResource(useHashOnly)
	})

	return rs
}

// initResource holds the initialization calls for the resource manager.
func (rs *Resources) initResource(watchHash bool) {
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

	// If we are in development mode empty the cache and reset for new use.
	fmt.Printf("Mode: %q -> %q", rs.options.Mode, DevelopmentMode)
	if rs.options.Mode == DevelopmentMode {
		fmt.Printf("Empty keys: %q", rs.cacheName)
		if err := rs.cache.Empty(); err != nil {
			fmt.Printf("Failed to clear internal cache for %q in development mode: %q\n", rs.cacheName, err.Error())
		}
	}

	// Attempt to retrieve a manifest.json from the backend.
	if !rs.options.IgnoreManifest {

		_, manifestResponse, err := rs.fetch.Get(rs.options.ManifestURI, shell.DefaultStrategy)
		if err != nil {
			fmt.Printf("Failed to load shell.AppManifest, resource loading is unavailable: %q\n", err.Error())
		} else {
			var appm []*shell.AppManifest

			if err := json.Unmarshal(manifestResponse.Body, &appm); err != nil {
				panic(fmt.Sprintf("Failed to load shell.AppManifest, resource loading is unavailable: %q", err.Error()))
			} else {

				var gmanifests []*shell.AppManifest
				var manifests []*shell.AppManifest

				for _, mani := range appm {
					for _, attr := range mani.Manifests {
						if attr.IsGlobal {
							shell.RegisterManifest(attr)
						}
					}

					if mani.GlobalScope {
						gmanifests = append(gmanifests, mani)
						continue
					}

					manifests = append(manifests, mani)
				}

				rs.manifests = manifests
				rs.gmanifests = gmanifests
				appm = nil
			}
		}
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
		scriptElem := trees.NewMarkup("script", false)
		trees.NewAttr("src", script).Apply(scriptElem)
		trees.NewAttr("gu-resource", "true").Apply(scriptElem)
		trees.NewAttr("gu-resource-root", "true").Apply(scriptElem)

		return rs.render([]*trees.Markup{scriptElem}, result...)
	}

	return rs.render(nil, result...)
}

// GenerateResources returns the markups based on specific resources which should be loaded
// along with the resource definition.
func (rs *Resources) GenerateResources() ([]*trees.Markup, []*trees.Markup) {
	if rs.headResource != nil && rs.bodyResource != nil {
		return rs.headResource, rs.bodyResource
	}

	var head, body []*trees.Markup

	for _, def := range rs.gmanifests {
		for _, manifest := range def.Manifests {
			if !manifest.InitStartup {
				continue
			}

			hook, err := shell.Get(manifest.HookName)
			if err != nil {
				fmt.Printf("Hook[%q] does not exists: Resource[%q] unable to install\n", manifest.HookName, manifest.Name)
				continue
			}

			markup, toHead, err := hook.Fetch(rs.fetch, manifest)
			if err != nil {
				fmt.Printf("Hook[%q] failed to retrieve Resource {Name: %q, Path: %q}\n", manifest.HookName, manifest.Name, manifest.Path)
				continue
			}

			trees.NewAttr("gu-resource", "true").Apply(markup)
			trees.NewAttr("gu-resource-from", manifest.Path).Apply(markup)
			trees.NewAttr("gu-resource-name", manifest.Name).Apply(markup)
			trees.NewAttr("gu-resource-id", manifest.ID).Apply(markup)

			if toHead {
				head = append(head, markup)
				continue
			}

			body = append(body, markup)
		}
	}

	rs.headResource = head
	rs.bodyResource = body

	return head, body
}

// render performs the needed work of collecting the giving markups and
// creating the complete html rendering.
func (rs *Resources) render(attachElems []*trees.Markup, rsx ...*ResourceDefinition) *trees.Markup {
	var head = trees.NewMarkup("head", false)
	var body = trees.NewMarkup("body", false)

	for _, rx := range rsx {
		toHead, toBody := rx.Resources()

		head.AddChild(toHead...)
		body.AddChild(toBody...)

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

//==============================================================================

type targetRenderable struct {
	View    Renderable
	Targets string
}

// ResourceDefinition defines a high-level definition for managing resources for
// which other definitions build from.
type ResourceDefinition struct {
	Dsl    DSL
	active bool
	uuid   string

	headResource []*trees.Markup
	bodyResource []*trees.Markup

	ViewHooks      []ViewHooks
	Views          []RenderView
	Links          []StaticView
	DeferLinks     []StaticView
	Renderables    []targetRenderable
	DRenderables   []targetRenderable
	RenderableData []RenderableData
	Relations      []*shell.AppManifest

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

// Resources returns the markups based on specific resources which should be loaded
// along with the resource definition.
func (rd *ResourceDefinition) Resources() ([]*trees.Markup, []*trees.Markup) {
	if rd.headResource != nil && rd.bodyResource != nil {
		return rd.headResource, rd.bodyResource
	}

	var head, body []*trees.Markup

	for _, def := range rd.Relations {
		for _, manifest := range def.Manifests {
			if !manifest.InitStartup {
				continue
			}

			hook, err := shell.Get(manifest.HookName)
			if err != nil {
				fmt.Printf("Hook[%q] does not exists: Resource[%q] unable to install\n", manifest.HookName, manifest.Name)
				continue
			}

			markup, toHead, err := hook.Fetch(rd.Root.fetch, manifest)
			if err != nil {
				fmt.Printf("Hook[%q] failed to retrieve Resource {Name: %q, Path: %q}\n", manifest.HookName, manifest.Name, manifest.Path)
				continue
			}

			trees.NewAttr("gu-resource", "true").Apply(markup)
			trees.NewAttr("gu-resource-from", manifest.Path).Apply(markup)
			trees.NewAttr("gu-resource-name", manifest.Name).Apply(markup)
			trees.NewAttr("gu-resource-id", manifest.ID).Apply(markup)

			if toHead {
				head = append(head, markup)
				continue
			}

			body = append(body, markup)
		}
	}

	// Cache the results so we dont re-render.
	rd.headResource = head
	rd.bodyResource = body

	return head, body
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

// hasRenderable returns true/false if a giving dependencies has been identified
// and is in the views dependencies list.
func (rd *ResourceDefinition) hasRenderable(name string) bool {
	for _, rd := range rd.RenderableData {
		if rd.Name == name {
			return true
		}
	}

	return false
}

// Init initialize runs the underline DSL for the operation, loads all resolvers
// and connects all internal views for immediate usage.
func (rd *ResourceDefinition) Init() {
	rd.Dsl(rd.Root.fetch, rd.Root.cache)
	rd.Resolver.Flush()

	{
		// Search root manifests lists component relation using the resource
		// RenderableData.
		for _, relation := range rd.RenderableData {

			if app := shell.FindByRelation(rd.Root.manifests, relation.Name); app != nil {
				if rd.hasRelation(app) {
					continue
				}

				rd.Relations = append(rd.Relations, app)
				rd.initRelation(app)
			}
		}
	}

	for _, level := range rd.Manager.Levels {
		rd.Resolver.Register(level)
	}

	for _, view := range rd.Views {
		rd.Resolver.ResolvedPassed(view.Resolve)
	}
}

// initRelation walks down the provided app relation adding the giving AppManifest
// which connect with this if not already in the list.
func (rd *ResourceDefinition) initRelation(app *shell.AppManifest) {
	for _, relation := range app.Relation.Composites {
		// fmt.Printf("Walking composite Relation to Manifest: %#v {to}-> %#v\n", relation, app)

		if related := shell.FindByRelation(rd.Root.manifests, relation); related != nil {
			// fmt.Printf("Found Composite Relation to Manifest: %#v -> %#v\n", relation, related)

			if !rd.hasRelation(related) {
				// fmt.Printf("Adding Composite Relation to Manifest: %#v -> %#v\n", relation, related)
				rd.Relations = append(rd.Relations, related)
				rd.initRelation(related)
			}
		}
	}

	for _, field := range app.Relation.FieldTypes {
		// fmt.Printf("Walking composite Relation to Manifest: %#v {to}-> %#v\n", field, app)

		if related := shell.FindByRelation(rd.Root.manifests, field); related != nil {
			// fmt.Printf("Found Field Relation to Manifest: %#v -> %#v\n", field, related)

			if !rd.hasRelation(related) {
				// fmt.Printf("Adding Composite Relation to Manifest: %#v -> %#v\n", field, related)
				rd.Relations = append(rd.Relations, related)
				rd.initRelation(related)
			}
		}
	}
}

// hasRelation returns true/false if the giving relation exists in the
// ResourceDefinition.
func (rd *ResourceDefinition) hasRelation(app *shell.AppManifest) bool {
	for _, rel := range rd.Relations {
		if rel == app {
			return true
		}
	}

	return false
}

// Head adds the provided markup as part of the children to the head tag.
func Head(markup interface{}, deferRender bool) {
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

// Style adds a element which generates a <style> tag.
func Style(styles interface{}, bind interface{}, deferRender bool) {
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

// View generates th necessary view for the specific type of data passed in
// to the function.
func View(markup interface{}, targets string, ops ...bool) Identity {
	var deferRender bool
	var targetAlreadyInDom bool

	if len(ops) == 1 {
		deferRender = ops[0]
	}

	if len(ops) > 1 {
		deferRender = ops[0]
		targetAlreadyInDom = ops[0]
	}

	switch mo := markup.(type) {
	case func() *trees.Markup:
		return doMarkup(mo(), targets, deferRender, targetAlreadyInDom)
	case *trees.Markup:
		return doMarkup(mo, targets, deferRender, targetAlreadyInDom)
	case string:
		parseTree := trees.ParseTree(mo)
		if len(parseTree) == 1 {
			return doMarkup(parseTree[0], targets, deferRender, targetAlreadyInDom)
		}

		div := trees.NewMarkup("section", false)
		div.AddChild(parseTree...)

		return doMarkup(div, targets, deferRender, targetAlreadyInDom)
	case Renderables:
		return doRenderView(mo, targets, deferRender, targetAlreadyInDom)
	case Renderable:
		return doRenderView([]Renderable{mo}, targets, deferRender, targetAlreadyInDom)
	case func() Renderable:
		return doRenderView(Renderables{mo()}, targets, deferRender, targetAlreadyInDom)
	case func() Renderables:
		return doRenderView(mo(), targets, deferRender, targetAlreadyInDom)
	default:
		panic(`
	Unknown markup or view processable type

		Accepted Markup Arguments:
			-	*trees.Markup
			- func() *trees.Markup

		Accepted View Arguments:
			-	[]Renderable
			-	Renderable
			-	func() []Renderable
			-	func() Renderable

			`)
	}
}

// doMarkup returns a new instance of a provided value which either is a function
// which returns a needed trees.Markup or a trees.Markup or slice of trees.Markup
// itself.
func doMarkup(markup *trees.Markup, targets string, deferRender bool, targetAlreadyInDom bool) Identity {
	current := getResources().MustCurrentResource()

	static := Static(markup)
	static.Morph = true
	trees.NewAttr("resource-id", current.UUID()).Apply(static.Content)

	if deferRender {
		current.DRenderables = append(current.DRenderables, targetRenderable{
			Targets: targets,
			View:    static,
		})

		return static
	}

	if !deferRender && targetAlreadyInDom {
		current.Renderables = append(current.Renderables, targetRenderable{
			Targets: targets,
			View:    static,
		})

		return static
	}

	if !deferRender && !targetAlreadyInDom && targets != "" {
		current.DRenderables = append(current.DRenderables, targetRenderable{
			View:    static,
			Targets: targets,
		})

		return static
	}

	current.Renderables = append(current.Renderables, targetRenderable{
		View:    static,
		Targets: targets,
	})

	return static
}

// doRenderView creates a RenderView and applies it to the provided resource.
func doRenderView(rs Renderables, targets string, deferRender bool, targetAlreadyInDom bool) Identity {
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

	for _, rd := range view.Dependencies() {
		if current.hasRenderable(rd.Name) {
			continue
		}

		current.RenderableData = append(current.RenderableData, rd)
	}

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

// Title adds a element which generates a <title> tag.
func Title(title string) {
	ml := mLink("title", false)
	trees.NewText(title, nil).Apply(ml.Content)
}

// Link adds a element which generates a <link> tag.
func Link(url string, mtype string, defered bool) {
	ml := mLink("link", defered)
	trees.NewAttr("href", url).Apply(ml.Content)
	trees.NewAttr("rel", mtype).Apply(ml.Content)
}

// CSS adds a element which generates a <style> tag.
func CSS(src string, defered bool) {
	ml := mLink("link", defered)
	trees.NewAttr("href", src).Apply(ml.Content)
	trees.NewAttr("rel", "stylesheet").Apply(ml.Content)
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
func mLink(tag string, deffer bool) *StaticView {
	static := Static(trees.NewMarkup(tag, false))
	current := getResources().MustCurrentResource()

	if deffer {
		current.DeferLinks = append(current.DeferLinks, *static)
	} else {
		current.Links = append(current.Links, *static)
	}

	return static
}

// Order defines a high level function which sets/resets the RenderingOrder of
// the current ResourceDefinition.
func Order(mode RenderingOrder) {
	getResources().MustCurrentResource().Order = mode
}

// GlobalRoute sets the giving root route for the currently used resource of the
// giving resource root.
func GlobalRoute(path string) {
	getResources().MustCurrentResource().Resolver = dispatch.NewResolver(path)
}

// Route allows definition of a route level to applying all the giving
// routes path returning the last applier.
func Route(base string, routes ...string) RouteApplier {
	return RouteLevel(LocalRouter(base), routes...)
}

// LocalRouter returns a RouteApplier which can be used to localize the
// internal routes for the base resource router and allow usage with markup
// and views.
func LocalRouter(basePath string, mx ...trees.SwitchMorpher) RouteApplier {
	var mo trees.SwitchMorpher

	if len(mx) != 0 {
		mo = mx[0]
	} else {
		mo = &trees.RemoveMorpher{}
	}

	return getResources().MustCurrentResource().Manager.Level(basePath, mo)
}

// RouteLevel takes a giving route applier returning the last route created
// from the route levels of the applier.
func RouteLevel(level RouteApplier, routes ...string) RouteApplier {
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
