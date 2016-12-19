package gu

import (
	"github.com/gu-io/gu/dispatch"
	"github.com/gu-io/gu/trees"
)

// RouteApplier defines a interface which composes trees Applier, Morpher and
// the dispatch.Resolver to handle routing over against a tree markup.
type RouteApplier interface {
	trees.Morpher
	trees.Appliable
	dispatch.Resolver

	Root() RouteApplier
	N(string) RouteApplier
	Next(string, trees.SwitchMorpher) RouteApplier
}

// RouteManager defines a router which handles creation and registering of a
// set of level connecting routes
type RouteManager struct {
	Levels map[string]RouteApplier
}

// NewRoutingManager returns a new instance of  RouteManager.
func NewRouteManager() *RouteManager {
	return &RouteManager{
		Levels: make(map[string]RouteApplier),
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
	if rx, ok := r.Levels[path]; ok {
		return rx
	}

	var root rm
	root.level = make(map[string]RouteApplier)
	root.routing = newRouting(path, morpher)

	r.Levels[path] = &root

	return &root
}

//==============================================================================

type rm struct {
	*routing
	prev  RouteApplier
	next  RouteApplier
	level map[string]RouteApplier
}

// Root returns the giving root if existing of this giving RouteApplier.
func (r *rm) Root() RouteApplier {
	return r.prev
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
	if rx, ok := r.level[path]; ok {
		return rx
	}

	block := newRouting(path, morpher)
	r.Resolver.Register(block.Resolver)

	var nextRoot rm
	nextRoot.level = make(map[string]RouteApplier)
	nextRoot.prev = r
	nextRoot.routing = block

	r.level[path] = &nextRoot

	return &nextRoot
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

	rs.Resolver.ResolvedPassed(func(p dispatch.Path) {
		morpher.Off(p)
	}).ResolvedFailed(func(p dispatch.Path) {
		morpher.On(p)
	})

	morpher.On(nil)

	return &rs
}

// Morph implements the Morpher interface providing the routing with the ability.
// It lets routing morph markup passed into it.
func (r *routing) Morph(mr *trees.Markup) *trees.Markup {
	return r.m.Morph(mr)
}

// Apply adds this routing as a morpher into the provided markup.
func (r *routing) Apply(mr *trees.Markup) {
	mr.AddMorpher(r)
}
