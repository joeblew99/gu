package gopherjs

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gu-io/gu"
	"github.com/gu-io/gu/router"
	"github.com/gu-io/gu/shell"
	"github.com/gu-io/gu/shell/fetch"
	"github.com/gu-io/gu/trees"
	"honnef.co/go/js/dom"
)

// JSDriver provides a concrete implementation of the Gu.Driver interface.
type JSDriver struct {
	readyHandlers []func()
	window        dom.Window
	doc           dom.Document
}

// NewJSDriver returns a new instance of a js driver.
func NewJSDriver() *JSDriver {
	win := dom.GetWindow()
	driver := &JSDriver{
		window: win,
		doc:    win.Document(),
	}

	DOMReady(driver.init)

	return driver
}

func (driver *JSDriver) init() {
	fmt.Printf("Initializing Driver(%q)\n", driver.Name())
	for _, ready := range driver.readyHandlers {
		ready()
	}
}

// Name returns the name of the driver.
func (driver *JSDriver) Name() string {
	return "Gopherjs(https://github.com/gopherjs/gopherjs)"
}

// OnReady registers the giving handle function to be called when the giving
// driver is ready and loaded.
func (driver *JSDriver) OnReady(handle func()) {
	driver.readyHandlers = append(driver.readyHandlers, handle)
}

// Location returns the current location of the browser.
func (driver *JSDriver) Location() router.PushEvent {
	host, path, hash := GetLocation()
	return router.PushEvent{
		Host: host,
		Path: path,
		Hash: hash,
		Rem:  hash,
	}
}

// Navigate takes the provided route and navigates the rendering system to the
// desired page.
func (driver *JSDriver) Navigate(route router.PushDirectiveEvent) {
	SetLocationByPushEvent(route)
}

// OnRoute registers the NApp instance for route changes and re-rendering.
func (driver *JSDriver) OnRoute(app *gu.NApp) {
	ListenForHistory(func(ev router.PushEvent) {
		app.ActivateRoute(ev)
		driver.Render(app)
	})
}

// Render issues a clean rendering of all content clearing out the current content
// of the browser to the one provided by the appliation.
func (driver *JSDriver) Render(app *gu.NApp) {
	head := driver.doc.QuerySelector("head")
	body := driver.doc.QuerySelector("body")

	// clear all children of head and body if the belong to us.
	for _, item := range head.QuerySelectorAll("[data-gen*='gu']") {
		if !item.HasAttribute("gu-resource-root") {
			item.ParentNode().RemoveChild(item)
		}
	}

	for _, item := range body.QuerySelectorAll("[data-gen*='gu']") {
		if !item.HasAttribute("gu-resource-root") {
			item.ParentNode().RemoveChild(item)
		}
	}

	renderTree := app.Render(nil)
	headTree := trees.Query.Query(renderTree, "head")
	bodyTree := trees.Query.Query(renderTree, "body")

	head.SetInnerHTML(headTree.HTML())
	body.SetInnerHTML(bodyTree.HTML())

	for _, child := range headTree.Children() {
		// TODO: Do we want to bind to the body instead of binding to a closer parent.
		child.EachEvent(func(ev *trees.Event, root *trees.Markup) {
			BindEvent(ev, head.Underlying())
		})
	}

	for _, child := range bodyTree.Children() {
		// TODO: Do we want to bind to the body instead of binding to a closer parent.
		child.EachEvent(func(ev *trees.Event, root *trees.Markup) {
			BindEvent(ev, body.Underlying())
		})
	}
}

// Update updates a giving view portion of a giving App within the designated
// rendering system(browser) provided by the driver.
func (driver *JSDriver) Update(app *gu.NApp, view *gu.NView) {

	attr := view.Attr()

	if attr.Target == gu.HeadTarget {
		head := driver.doc.QuerySelector("head")
		renderTree := view.Render()

		// TODO: Do we wish to call this here?
		// view.Unmounted()

		Patch(CreateFragment(renderTree.HTML()), head.Underlying(), false)

		// TODO: Do we wish to call this here?
		view.Updated()
		return
	}

	body := driver.doc.QuerySelector("body")
	renderTree := view.Render()
	Patch(CreateFragment(renderTree.HTML()), body.Underlying(), false)

	// TODO: Do we wish to call this here?
	view.Updated()
}

// Services returns the Fetcher and Cache associated with the provided cacheName.
// Intercepting requests for usage.
func (driver *JSDriver) Services(cacheName string, intercept bool) (shell.Fetch, shell.Cache) {
	cache := NewCache(cacheName)
	fetch := fetch.New(cache)

	// If we are supposed to intercept all request signals then bind.
	if intercept {
		driver.doc.AddEventListener("fetch", true, func(event dom.Event) {
			request := event.Underlying().Get("request")
			if request == nil || request == js.Undefined {
				return
			}

			// Extract shell.Request from the request object.
			req, err := shell.ObjectToWebRequest(request)
			if err != nil {
				fmt.Printf("Unable to transform request object: %+q -> %q", request.String(), err.Error())
				return
			}

			res, err := cache.GetRequest(req)
			if err != nil {
				if req.Cache == "" {
					req.Cache = shell.LastCacheStrategy
				}

				madeRes, err := fetch.Do(req)
				if err != nil {
					fmt.Printf("Failed to fetch requests: %#v : %#v -> %q", request, req, err.Error())
					return
				}

				resObj := shell.WebResponseToJSResponse(&madeRes)
				event.Underlying().Call("respondWith", resObj)
				fmt.Printf("Request serviced successfuly. {Req: %q, Res: %d}", req.URL, madeRes.Status)
				return
			}

			resObj := shell.WebResponseToJSResponse(&res)
			event.Underlying().Call("respondWith", resObj)
			fmt.Printf("Request serviced successfuly. {Req: %q, Res: %d}", req.URL, res.Status)
			return
		})
	}

	return fetch, cache
}
