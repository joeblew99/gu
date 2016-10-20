package gu

import (
	"sync"

	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/trees"
)

// RouteApplier defines a interface which composes trees Applier, Morpher and
// the dispatch.Resolver to handle routing over against a tree markup.
type RouteApplier interface {
	trees.Morpher
	trees.Appliable
	dispatch.Resolver
	N(string) RouteApplier
	Next(string, trees.SwitchMorpher) RouteApplier
}

// RouteManager defines a router which handles creation and registering of a
// set of level connecting routes
type RouteManager struct {
	levels map[string]RouteApplier
}

// NewRoutingManager returns a new instance of  RouteManager.
func NewRouteManager() *RouteManager {
	return &RouteManager{
		levels: make(map[string]RouteApplier),
	}
}

// L returns a RouteApplier attached to a giving manager which cache and manages
// the different defined routes.
func (r *RouteManager) L(path string) RouteApplier {
	return r.Level(path, &trees.RemoveMorpher{})
}

// Level creates a new route level using the provided string as the base path
// for its root router.
func (r *RouteManager) Level(path string, morpher trees.SwitchMorpher) RouteApplier {
	if rx, ok := r.levels[path]; ok {
		return rx
	}

	var rs rm
	rs.path = path

	rmi := newRouting(path, morpher)

	rs.next = rmi
	rs.root = rmi
	rs.Resolver = rmi
	rs.linked = make(map[string]*routing)

	r.levels[path] = &rs
	return &rs
}

//==============================================================================

type rm struct {
	dispatch.Resolver
	*routing

	rw     sync.RWMutex
	path   string
	linked map[string]*routing

	root *routing
	next *routing
}

// N returns a new trees.Applier based on the internal routing struct. Which stores
// its modifier
func (r *rm) N(path string) RouteApplier {
	return r.Next(path, &trees.RemoveMorpher{})
}

// Next creates a route which depends on the previous route created. It
// allows multi-level routing where the next depends on the outcome of the
// previous. It uses dispatch.Resolver which shifts the matching paths by passing
// down the remants of path unmatched to its next subscribers.
func (r *rm) Next(path string, morpher trees.SwitchMorpher) RouteApplier {
	if rx, ok := r.linked[path]; ok {
		return rx
	}

	prv := r.next
	r.next = newRouting(path, morpher)
	prv.Register(r.next)

	r.linked[path] = r.next

	return r.next
}

//==============================================================================

// routing defines a dispatch.Resolve structure that allows morphing the output
// of a markup based on a giving route.
type routing struct {
	dispatch.Resolver
	m trees.Morpher
}

// Routing returns a new instance of a Routing struct.
func newRouting(path string, morpher trees.SwitchMorpher) *routing {
	var rs routing
	rs.m = morpher
	rs.Resolver = dispatch.NewResolver(path)
	rs.Resolver.ResolvedPassed(func(p dispatch.Path) { morpher.Off(p) })
	rs.Resolver.ResolvedFailed(func(p dispatch.Path) { morpher.On(p) })
	return &rs
}

// Morph implements the Morpher interface providing the routing with the ability.
// It lets routing morph markup passed into it.
func (r *routing) Morph(mr trees.Markup) trees.Markup {
	return r.m.Morph(mr)
}

// Apply adds this routing as a morpher into the provided markup.
func (r *routing) Apply(mr trees.Markup) {
	mr.AddMorpher(r)
}
