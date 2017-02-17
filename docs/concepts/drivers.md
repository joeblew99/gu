Drivers
=======

Drivers are the rendering engines for the Gu app system. They provide implementation details for the different endpoint where the app designed will be rendered into.

Examples of Drivers:

-	GopherJS Driver(https://github.com/gu-io/gu/drivers/gopherjs) This provides a driver to handle rendering to the browser and route changes to effectively and with performance render the design package appropriately with the functionality intended.

Drivers are required to meet the Gu `Drivers` interface which then handles coordination of rendering and view updates request from and to the provided app.

```go

// Driver defines an interface which provides an Interface to render Apps and views
// to a display system. eg. GopherJS, wkWebview.
type Driver interface {
	// Method to subscribe to ready state for the driver.
	OnReady(func())

	// Name of the driver.
	Name() string

	// Navigate the Driver to the provided path.
	Navigate(router.PushDirectiveEvent)

	// Current Location of the driver path.
	Location() router.PushEvent

	// OnRoute registers for route change notifications to react accordingly.
	// This does a totally re-write of the whole display.
	OnRoute(*NApp)

	// Render app and it's content.
	Render(*NApp)

	// Update app's view and it's target content.
	Update(*NApp, ...*NView)

	// Fetcher returns a new resource fetcher which can be used to retrieve Resources.
	// Fetcher requires the cacheName and a boolean indicating if it should intercept
	// all requests for resources.
	Fetcher(cacheName string, interceptRequests bool) (shell.Fetch, shell.Cache)
}

```

It's job is to provide an interface by which it's selected platform renders out appropriately the designed application built using GU and to provide the flexibility of not being tide to a specific rendering endpoint. It also exposes means by which, requests for resources can be made through and caches by which response can be sorted to reduce network usage.

Explore the [Driver](https://github.com/gu-io/gu/drivers/) package to see other driver implementations.
