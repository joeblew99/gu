// Package gu provides a UI framework for Go.
package gu

import (
	"fmt"
	"html/template"
	"sync"
	"sync/atomic"

	"github.com/gu-io/gu/notifications"
	"github.com/gu-io/gu/router"
	"github.com/gu-io/gu/shell"
	"github.com/gu-io/gu/trees"
)

// countKeeper handles management of the keys being generating. Guards the incrementation
// using a mutex.
var countKeeper = struct {
	ml        sync.Mutex
	baseKey   string
	baseCount int
}{
	baseKey: "3bce4931-6c75-41ab-afe0-2ec108a30860",
}

// NewKey returns a new string key which is path of the incremental key which once initializes
// constantly increases.
func NewKey() string {
	countKeeper.ml.Lock()
	countKeeper.baseCount++
	countKeeper.ml.Unlock()

	return fmt.Sprintf("%d-%s", countKeeper.baseCount, countKeeper.baseKey)
}

//==============================================================================

// Identity defines an interface which expoese the identity of a giving object.
type Identity interface {
	UUID() string
}

// Services defines a struct which exposes certain fields to be accessible to
// others.
type Services struct {
	AppUUID       string
	Driver        Driver
	Router        router.Resolver
	Fetch         shell.Fetch
	Cache         shell.Cache
	Mounted       Subscriptions
	Rendered      Subscriptions
	Updated       Subscriptions
	Unmounted     Subscriptions
	Notifications *notifications.AppNotification
}

// RegisterService provides an interface which registers the provided fetcher,
// caching and routing system for a component. This will be called before
// any setup of the components structure to allow users set the system they needed
// running.
type RegisterService interface {
	RegisterService(Services)
}

// Renderable provides a interface for a renderable type.
type Renderable interface {
	Render() *trees.Markup
}

// Renderables defines a lists of Renderable structures.
type Renderables []Renderable

// MarkupRenderer provides a interface for a types capable of rendering dom markup.
type MarkupRenderer interface {
	Renderable
	RenderHTML() template.HTML
}

// Properties defines a type which exposes a single method to retrieve values
// from.
type Properties interface {
	Get(string) interface{}
}

// Fetchable exposes a interface which recieves a shell.Fetch and shell.Cache instances
// through a method.
type Fetchable interface {
	UseFetch(shell.Fetch, shell.Cache)
}

// FetchableRenderer defines a Renderer which expects to recieve instances of
// the shell.Fetch and shell.Cache interface through the UseFetch method.
type FetchableRenderer interface {
	Fetchable
	Renderable
}

// Reactor defines an interface for functions subscribing for
// notifications to react.
type Reactor interface {
	React(func())
}

// Reactive extends the ReactiveRenderable by exposing a Publish method
// which allows calling the update notifications list of a ReactiveRenderable.
type Reactive interface {
	Reactor
	Publish()
}

//==============================================================================

// RenderCommand defines a struct to hold a giving command for the rendering
// of a App or View using the JSON format.
type RenderCommand struct {
	Command string   `json:"Command"`
	App     AppJSON  `json:"App,omitempty"`
	View    ViewJSON `json:"View,omitempty"`
}

// AppRenderCommand returns a new RenderCommand for rendering a app.
func AppRenderCommand(app *NApp, route interface{}) RenderCommand {
	return RenderCommand{
		Command: "RenderApp",
		App:     app.RenderJSON(route),
	}
}

// ViewRenderCommand returns a new RenderCommand for rendering a view.
func ViewRenderCommand(view *NView) RenderCommand {
	return RenderCommand{
		Command: "RenderView",
		View:    view.RenderJSON(),
	}
}

//==============================================================================

// NewReactive returns an instance of a Reactive struct.
func NewReactive() Reactive {
	var rc Subscription
	return &rc
}

// Subscriptions exposes an interface which combines a Reactive type and a clear
// function to dispose of subscribers.
type Subscriptions interface {
	Reactive
	Clear()
	Reset()
	Used() bool
}

// Subscription defines a baseline structure that can be composed into
// any struct to provide a reactive view.
type Subscription struct {
	subs           []func()
	totalPublished int64
}

// NewSubscriptions returns an instance of a Subscription pointer.
func NewSubscriptions() *Subscription {
	return &Subscription{}
}

// React adds a function into the subscription list for this reactor.
func (r *Subscription) React(sub func()) {
	r.subs = append(r.subs, sub)
}

// Clear destroys all subscribers in the lists.
func (r *Subscription) Clear() {
	r.subs = nil
}

// Reset resets the subscription has unused.
func (r *Subscription) Reset() {
	atomic.StoreInt64(&r.totalPublished, 0)
}

// Used returns true/false if the subscription has been called to publish.
func (r *Subscription) Used() bool {
	return atomic.LoadInt64(&r.totalPublished) > 0
}

// Publish runs a through the subscription list and calls the registerd functions.
func (r *Subscription) Publish() {
	atomic.AddInt64(&r.totalPublished, 1)

	for _, sub := range r.subs {
		sub()
	}
}

//==============================================================================
