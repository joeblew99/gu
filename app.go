package gu

import (
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/gu-io/gu/router"
	"github.com/gu-io/gu/shell"
	"github.com/gu-io/gu/shell/cache"
	"github.com/gu-io/gu/shell/fetch"
	"github.com/gu-io/gu/trees"
	"github.com/gu-io/gu/trees/elems"
	"github.com/influx6/faux/reflection"
)

// Mode defines a type to represent the mode which the library works under.
type Mode int

const (
	// ProductionMode defines production mode for which defines how the context
	// which it behaves as regards certain features.
	ProductionMode Mode = iota

	// DevelopmentMode defines production mode for which defines how the context
	// which it behaves as regards certain features.
	DevelopmentMode
)

// RenderingOrder defines a type used to define the order which rendering is to be done for a resource.
type RenderingOrder int

const (
	// FirstOrder defines that rendering be first in order.
	FirstOrder RenderingOrder = iota

	// AnyOrder defines that rendering be middle in order.
	AnyOrder

	// LastOrder defines that rendering be last in order.
	LastOrder
)

// type ResourceRenderer interface {
// 	Render(...*ResourceDefinition)
// 	Update(Renderable, string, bool)
// }

// Driver defines an interface which provides an Interface to render Apps nd views
// to a display system. eg. GopherJS, wkWebview.
type Driver interface {
}

// Resource defines any set of rendering links, scripts, styles needed by a view.
type Resource struct {
	Manifest shell.AppManifest
	body     []*trees.Markup
	head     []*trees.Markup
}

// AppAttr defines a struct for
type AppAttr struct {
	Mode           Mode   `json:"mode"`
	Name           string `json:"name"`
	Manifests      string `json:"manifests"`
	EnableManifest bool   `json:"ignore_manifest"`
	ListenForFetch bool   `json:"listen_for_fetch"`
}

// NApp defines a struct which encapsulates all the core view management functions
// for views.
type NApp struct {
	uuid   string
	attr   AppAttr
	driver Driver
	cache  shell.Cache
	fetch  shell.Fetch

	local []shell.AppManifest
	views []NView

	globalResources []Resource

	tree *trees.Markup
}

// App creates a new app structure to rendering gu components.
func App(attr AppAttr) *NApp {
	var app NApp
	app.attr = attr
	app.uuid = NewKey()
	app.cache = cache.New(attr.Name)
	app.fetch = fetch.New(app.cache)

	// If we are in development mode empty the cache and reset for new use.
	if attr.Mode == DevelopmentMode {
		if err := rs.cache.Empty(); err != nil {
			fmt.Printf("Failed to clear internal cache for %q in development mode: %q\n", rs.cacheName, err.Error())
		}
	}

	// Attempt to retrieve a manifest.json from the backend.
	if attr.EnableManifest && attr.Manifests != "" {
		var appm []*shell.AppManifest

		if err := json.Unmarshal(manifestResponse.Body, &appm); err != nil {
			panic(fmt.Sprintf("Failed to load shell.AppManifest, resource loading is unavailable: %q", err.Error()))
		} else {

			for _, mani := range appm {
				for _, attr := range mani.Manifests {
					if attr.IsGlobal {
						shell.RegisterManifest(attr)
					}
				}

				if mani.GlobalScope {
					app.globalResources = append(app.globalResources, Resource{
						Manifest: mani,
					})

					continue
				}

				app.local = append(app.local, mani)
			}
		}
	}

	return &app
}

// Render returns the giving rendered tree of the app respective of the path
// found.
func (app *NApp) Render(es interface{}) *trees.Markup {
	var pe notifications.PushEvent

	switch es := es.(type) {
	case string:
		pe = router.NewPushEvent(es, true)
	case router.PushEvent:
		pe = es
	}

	var head = trees.NewMarkup("head", false)
	var body = trees.NewMarkup("body", false)

	for _, view := range app.PushViews(pe) {

	}

	html := trees.NewMarkup("html", false)
	head.Apply(html)
	body.Apply(html)
	return htmlMarkup
}

// PushViews returns a slice of  views that match and pass the provided path.
func (app *NApp) PushViews(event router.PushEvent) []NView {
	var active []NView

	for _, view := range app.views {
		if !view.router.Test(event.String()) {
			view.disableView()
			continue
		}

		active = append(active, view)
		view.enableView(event)
	}

	return active
}

// Resources return the giving resource headers which relate with the
// view.
func (app *NApp) Resources() ([]*trees.Markup, []*trees.Markup) {
	var head, body []*trees.Markup

	for _, def := range app.global {
		if def.body != nil || def.head != nil {
			head = append(head, def.head...)
			body = append(body, def.body...)
			continue
		}

		if def.Manifest.Manifests == nil {
			continue
		}

		for _, manifest := range def.Manifests {
			if !manifest.Init {
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
				def.head = append(def.head, markup)
				head = append(head, markup)
				continue
			}

			def.body = append(def.body, markup)
			body = append(body, markup)
		}
	}

	return head, body
}

// UUID returns the uuid specific to the giving view.
func (app *NApp) UUID() string {
	return app.uuid
}

// ViewAttr defines a structure to define a option values for setting up the appropriate
// settings for the view.
type ViewAttr struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Route  string `json:"route"`
	Target string `json:"target"`
}

// RenderableData defines a struct which contains the name of a giving renderable
// and it's package.
type RenderableData struct {
	Name string
	Pkg  string
}

// NView defines a structure to encapsulates all rendering component for a given
// view.
type NView struct {
	uuid   string
	active bool
	cache  shell.Cache
	fetch  shell.Fetch
	title  *trees.Markup
	router router.Resolver
	attr   ViewAttr

	renderingData []RenderableData
	local         []shell.AppManifest

	localResources []Resource

	beginComponents []Component
	anyComponents   []Component
	lastComponents  []Component
}

// UUID returns the uuid specific to the giving view.
func (v *NView) UUID() string {
	return v.uuid
}

// Render returns the markup for the giving views.
func (v *NView) Render() *trees.Markup {
	return nil
}

// Resources return the giving resource headers which relate with the
// view.
func (v *NView) Resources() ([]*trees.Markup, []*trees.Markup) {
	var head, body []*trees.Markup

	for _, def := range v.localResources {
		if def.body != nil || def.head != nil {
			head = append(head, def.head...)
			body = append(body, def.body...)
			continue
		}

		if def.Manifest.Manifests == nil {
			continue
		}

		for _, manifest := range def.Manifests {
			if !manifest.Init {
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
				def.head = append(def.head, markup)
				head = append(head, markup)
				continue
			}

			def.body = append(def.body, markup)
			body = append(body, markup)
		}
	}

	return head, body
}

// ComponentAttr defines a structure to define a component and its appropriate settings.
type ComponentAttr struct {
	Order     RenderingOrder `json:"order"`
	Tag       string         `json:"tag"`
	Target    string         `json:"target"`
	Route     string         `json:"route"`
	Base      interface{}    `json:"base"`
	Relations []string       `json:"relations"`
}

// Component defines a struct which
type Component struct {
	Reactive
	uuid      string
	attr      ComponentAttr
	Rendering Renderable

	Mounted   Subscriptions
	Unmounted Subscriptions
	Rendered  Subscriptions
	live      *trees.Markup
}

// UUID returns the identification for the giving component.
func (c Component) UUID() string {
	return c.uuid
}

// Render returns the markup corresponding to the internal Renderable.
func (c *Component) Render() *trees.Markup {
	newTree := c.Rendering.Render()
	newTree.SwapUID(c.uuid)

	if c.live != nil {
		live = c.live
		live.EachEvent(func(e *trees.Event, _ *trees.Markup) {
			if e.Handle != nil {
				e.Handle.End()
			}
		})

		newTree.Reconcile(live)
		live.Empty()
	}

	c.live = newTree

	return c.live.ApplyMorphers()
}

// Component adds the provided component into the selected view.
func (v *NView) Component(attr ComponentAttr) {
	var c Component
	c.attr = attr
	c.uuid = NewKey()
	c.Reactive = NewReactive()
	c.mounted = NewSubscriptions()
	c.unmounted = NewSubscriptions()
	c.rendered = NewSubscriptions()

	// Transform the base argument into the acceptable
	// format for the object.
	{
		switch mo := attr.Base.(type) {
		case func() *trees.Markup:
			static := Static(mo())
			static.Morph = true
			c.Rendering = static

		case *trees.Markup:
			static := Static(mo)
			static.Morph = true
			c.Rendering = static
			break

		case string:
			parseTree := trees.ParseTree(mo)
			if len(parseTree) != 1 {
				section := trees.NewMarkup("section", false)
				section.AddChild(parseTree...)

				static := Static(section)
				static.Morph = true
				c.Rendering = static
				break
			}

			static := Static(parseTree[0])
			static.Morph = true
			c.Rendering = static
			break

		case Renderable:
			if renderField, _, err := reflection.StructAndEmbeddedTypeNames(vr); err == nil {
				v.renderingData = append(v.renderingData, RenderableData{
					Name: renderField.TypeName,
					Pkg:  renderField.Pkg,
				})
			}

			c.Rendering = rc
			break

		case func() Renderable:
			rc := mo()
			if renderField, _, err := reflection.StructAndEmbeddedTypeNames(rc); err == nil {
				v.renderingData = append(v.renderingData, RenderableData{
					Name: renderField.TypeName,
					Pkg:  renderField.Pkg,
				})
			}

			c.Rendering = rc
			break

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

	// Connect the views into the rendering reactor if it has one.
	if rc, ok := c.Rendering.(Reactor); ok {
		rc.React(c.Reactive.Publish)
	}

	// Add the component into the right order.
	{
		switch resource.Order {
		case FirstOrder:
			v.beginComponents = append(v.beginComponents, c)
		case LastOrder:
			v.lastComponents = append(v.lastComponents, c)
		case AnyOrder:
			v.anyComponents = append(v.anyComponents, c)
		}
	}

	// Collect necessary app manifest that connect with rendering.
	{
		for _, relation := range v.renderingData {
			if app, err := shell.FindByRelation(v.local, relation.Name); err == nil {
				if v.hasRelation(app) {
					continue
				}

				v.localResources = append(v.localResources, Resource{
					Manifest: app,
				})

				initRelation(views, app, v.local)
			}
		}
	}
}

// Disabled returns true/false if the giving view is disabled.
func (v *NView) Disabled() bool {
	// v.rl.RLock()
	// defer v.rl.RUnlock()

	return state
}

// enableView enables the active state of the view.
func (v *NView) enableView() {
	// v.rl.Lock()
	// {
	v.active = true
	// }
	// v.rl.Unlock()
}

// disableView disables the active state of the view.
func (v *NView) disableView() {
	// v.rl.Lock()
	// {
	v.active = false
	// }
	// v.rl.Unlock()
}

// hasRenderable returns true/false if a giving dependency name exists.
func (v *NView) hasRenderable(name string) bool {
	for _, rd := range v.renderingData {
		if rd.Name == name {
			return true
		}
	}

	return false
}

// hasRelation returns true/false if a giving manifests name exists.
func (v *NView) hasRelation(name string) bool {
	for _, rd := range v.localResources {
		if rd.Manifest.Name == name {
			return true
		}
	}

	return false
}

// View returns a new instance of the view object.
func (app *NApp) View(attr ViewAttr) *View {
	var vw NView
	vw.attr = attr
	vw.uuid = NewKey()

	vw.title = elems.Title(attr.Title)
	vw.router = router.New(attr.Route)
	vw.cache = app.cache
	vw.fetch = app.fetch
	vw.local = app.local

	vw.router.Failed(func(push router.PushEvent) {
		vw.disableView()
	})

	app.views = append(app.views, vw)

	return &vw
}

//==============================================================================

// StaticView defines a MarkupRenderer implementing structure which returns its Content has
// its markup.
type StaticView struct {
	uid      string
	Content  *trees.Markup
	Mounted  Subscriptions
	Rendered Subscriptions
	Morph    bool
}

// Static defines a toplevel function which returns a new instance of a StaticView using the
// provided markup as its content.
func Static(tree *trees.Markup) *StaticView {
	return &StaticView{
		Content: tree,
		uid:     NewKey(),
	}
}

// UUID returns the RenderGroup UUID for identification.
func (s *StaticView) UUID() string {
	return s.uid
}

// Dependencies returns the list of all internal dependencies of the given view.
// It returns the names of the structs and their internals composed values/fields
// to help conditional resource loading.
func (s *StaticView) Dependencies() []RenderableData {
	return nil
}

// Render returns the markup for the static view.
func (s *StaticView) Render() *trees.Markup {
	if s.Morph {
		return s.Content.ApplyMorphers()
	}

	return s.Content
}

// RenderHTML returns the html template version of the StaticView content.
func (s *StaticView) RenderHTML() template.HTML {
	return s.Content.EHTML()
}

//==============================================================================

// initRelation walks down the provided app relation adding the giving AppManifest
// which connect with this if not already in the list.
func initRelation(views *NView, app shell.AppManifest, relations []shell.AppManifests) {
	for _, relation := range app.Relation.Composites {
		if related, err := shell.FindByRelation(relations, relation); err == nil {
			if !views.hasRelation(related) {

				views.localResources = append(views.localResources, Resource{
					Manifest: related,
				})

				initRelation(views, related, relations)
			}
		}
	}

	for _, field := range app.Relation.FieldTypes {
		if related, err := shell.FindByRelation(rd.Root.manifests, field); err == nil {
			if !views.hasRelation(related) {

				views.localResources = append(views.localResources, Resource{
					Manifest: related,
				})

				initRelation(views, related, relations)
			}
		}
	}
}
