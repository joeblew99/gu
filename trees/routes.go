package trees

import (
	"sync"

	"github.com/influx6/gu/dispatch"
)

// Morpher defines an interface which morphs the giving markup based on
// its current internal state based on some internal condition.
type Morpher interface {
	Morph(Markup) Markup
}

// SwitchMorpher defines an interface which allows switching the behaviour of
// which determines what the markup provided gets morphed into. This allows a nice
// simple binary morph of input based on certain conditions.
type SwitchMorpher interface {
	Morpher
	On(interface{})
	Off(interface{})
}

// RemoveMorpher defines a morpher which sets the giving markup as removed.
type RemoveMorpher struct {
	wl     sync.RWMutex
	remove bool
}

// On switches the state of the morpher out of a remove to be true.
func (r *RemoveMorpher) On(m interface{}) {
	r.wl.Lock()
	r.remove = true
	r.wl.Unlock()
}

// Off switches the state of the morpher to set remove to be true.
func (r *RemoveMorpher) Off(m interface{}) {
	r.wl.Lock()
	r.remove = false
	r.wl.Unlock()
}

// Morph sets the markup received as removed.
func (r *RemoveMorpher) Morph(m Markup) Markup {
	r.wl.RLock()
	{
		if r.remove {
			m.Remove()
		} else {
			m.UnRemove()
		}
	}
	r.wl.RUnlock()

	return m
}

// HideMorpher defines a morpher which sets the giving markup as removed.
type HideMorpher struct {
	wl     sync.RWMutex
	hidden bool
}

// On switches the state of the morpher out of a hidden state to be true.
func (r *HideMorpher) On(m interface{}) {
	r.wl.Lock()
	r.hidden = true
	r.wl.Unlock()
}

// Off switches the state of the morpher to set hidden state to be true.
func (r *HideMorpher) Off(m interface{}) {
	r.wl.Lock()
	r.hidden = false
	r.wl.Unlock()
}

// Morph sets the markup received as removed.
func (r *HideMorpher) Morph(m Markup) Markup {
	r.wl.RLock()
	{
		if r.hidden {
			Hide.Mode(m)
		} else {
			Show.Mode(m)
		}
	}
	r.wl.RUnlock()

	return m
}

//==============================================================================

// Routing defines a dispatch.Resolve structure that allows morphing the output
// of a markup based on a giving route.
type Routing struct {
	dispatch.Resolver
	m Morpher
}

// Route provides a convenient function for creating a route with the default RemovalMorpher.
func Route(path string) *Routing {
	return NewRouting(path, &RemoveMorpher{})
}

// Routing returns a new instance of a Routing struct.
func NewRouting(path string, morpher SwitchMorpher) *Routing {
	var rs Routing
	rs.m = morpher
	rs.Resolver = dispatch.NewResolver(path)
	rs.Resolver.ResolvedPassed(func(p dispatch.Path) { morpher.Off(p) })
	rs.Resolver.ResolvedFailed(func(p dispatch.Path) { morpher.On(p) })
	return &rs
}

// NewFrom creates a router dependent on this Routing, where the internal
// resolvers strips down to the returning Routing resolver.
func (r *Routing) NewFrom(path string, morpher SwitchMorpher) *Routing {
	rsm := NewRouting(path, morpher)
	r.Resolver.Register(rsm)
	return rsm
}

// Morph implements the Morpher interface providing the routing with the ability.
// It lets routing morph markup passed into it.
func (r *Routing) Morph(mr Markup) Markup {
	return r.m.Morph(mr)
}

// Apply adds this routing as a morpher into the provided markup.
func (r *Routing) Apply(mr Markup) {
	mr.AddMorpher(r)
	mr.UseRouter(r)
}

//==============================================================================
