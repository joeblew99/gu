package gutrees

import "github.com/influx6/gu/gudispatch"

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
	Flip(interface{})
}

// Routing defines a gudispatch.Resolve structure that allows morphing the output
// of a markup based on a giving route.
type Routing struct {
	gudispatch.Resolver
	morpher SwitchMorpher
}

// Routing returns a new instance of a Routing struct.
func NewRouting(path string, morpher SwitchMorpher) *Routing {
	var rs Routing
	rs.morpher = morpher
	rs.Resolver = gudispatch.NewResolver(path)
	rs.Resolver.ResolvedPassed(func(p gudispatch.Path) { morpher.Flip(p) })
	rs.Resolver.ResolvedFailed(func(p gudispatch.Path) { morpher.Flip(p) })
	return &rs
}

// Morph implements the Morpher interface providing the routing with the ability.
// It lets routing morph markup passed into it.
func (r *Routing) Morph(mr Markup) Markup {
	return r.morpher.Morph(mr)
}

// Apply adds this routing as a morpher into the provided markup.
func (r *Routing) Apply(mr Markup) {
	mr.AddMorpher(r)
	mr.AddResolver(r)
}
