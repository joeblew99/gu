package guviews

import (
	"html/template"
	"sync"

	"github.com/influx6/gu/gutrees"
)

// Renderable provides a interface for a renderable type.
type Renderable interface {
	Render() gutrees.Markup
}

// Renderables defines a lists of Renderable structures.
type Renderables []Renderable

// MarkupRenderer provides a interface for a types capable of rendering dom markup.
type MarkupRenderer interface {
	Renderable
	RenderHTML() template.HTML
}

// ReactiveRenderable defines an interface of a Renderable capable of
// notifying subscribe functions of changes to allow them to React.
type ReactiveRenderable interface {
	Renderable
	Subscribe(func())
}

// Reactive extends the ReactiveRenderable by exposing a Publish method
// which allows calling the update notifications list of a ReactiveRenderable.
type Reactive interface {
	Subscribe(func())
	Publish()
}

// NewReactive returns an instance of a Reactive struct.
func NewReactive() Reactive {
	var rc reactive
	return &rc
}

//==============================================================================

// reactive defines a baseline structure that can be composed into
// any struct to provide a reactive view.
type reactive struct {
	rw   sync.RWMutex
	subs []func()
}

// Subscribe adds a function into the subscription list for this reactor.
func (r *reactive) Subscribe(sub func()) {
	r.rw.Lock()
	defer r.rw.Unlock()
	r.subs = append(r.subs, sub)
}

// Publish runs a through the subscription list and calls the registerd functions.
func (r *reactive) Publish() {
	r.rw.RLock()
	defer r.rw.RUnlock()
	for _, sub := range r.subs {
		sub()
	}
}
