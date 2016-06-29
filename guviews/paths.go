package guviews

import "github.com/influx6/gu/gudispatch"

// FailNormal is a fail function type that accepts no arguments.
type FailNormal func()

// FailPath is a fail function type that accepts a gudispatch.Path.
type FailPath func(gudispatch.Path)

// AttachURL attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the full URL(Path+Hash).
// failFn must either be a FailNormal, FailPath or nil.
func AttachURL(pattern string, v Views, failFn interface{}) {
	var failCb FailPath

	switch failFn.(type) {
	case FailNormal:
		failCb = func(_ gudispatch.Path) { failCb.(FailNormal)() }
	case FailPath:
		failCb = failFn.(FailPath)
	}

	gudispatch.AttachURL(pattern, func(p gudispatch.Path) {
		v.Path(p)
	}, failCb)
}

// AttachHash attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the URL hash.
// failFn must either be a FailNormal, FailPath or nil.
func AttachHash(pattern string, v Views, failFn interface{}) {
	var failCb FailPath

	switch failFn.(type) {
	case FailNormal:
		failCb = func(_ gudispatch.Path) { failCb.(FailNormal)() }
	case FailPath:
		failCb = failFn.(FailPath)
	}

	gudispatch.AttachHash(pattern, func(p gudispatch.Path) {
		v.Path(p)
	}, failCb)
}

// ResolveSubscriber defines a function type for a Resolver subcriber.
type ResolveSubscriber func(gudispatch.Path)

// Resolver defines an interface for a type that resolves a
// provided instance of a gudispatch.Path.
type Resolver interface {
	Subscribe(ResolveSubscriber)
	Register(Resolver)
	Resolve(gudispatch.Path)
}

// NewResolver returns a new instance of a structure that matches
// the Resolver interface.
func NewResolver(path string) Resolver {
	var br basicResolver

	if path != "" {
		br.matcher = gudispatch.URIMatcher(path)
	}

	return &br
}

// basicResolver defines a struct that implements
type basicResolver struct {
	matcher  pattern.URIPattern
	subs     []ResolveSubscriber
	children []Resolver
}

// Register adds a resolver into the list which will get triggerd
// when this resolver gets triggered, they will recieve a new Path
// made out of the remaining path from the Path received by this.
func (b *basicResolver) Register(r Resolver) {
	b.children = append(b.children, r)
}

// Resolve takes a `gudispatch.Path` instance, matches the content
// against it's own matcher, if its a match, it calls its subscribers
// then takes the remaining path left after removing its own matching
// piece and sends that off to its children, if there exists any remaining
// path that is.
func (b *basicResolver) Resolve(path gudispatch.Path) {
	if b.matcher == nil {

		// Notify the subscribers.
		for _, sub := range b.subs {
			sub(path)
		}

		// Notify the kids with what is left in the Path.
		for _, child := range b.children {
			child.Resolve(path)
		}

		return
	}

	params, rem, ok := b.matcher.Validate(path.Rem)
	if !ok {
		return
	}

	var hash, path string

	hashIndex := strings.IndexOf(path.Rem, "#")
	if hashIndex != -1 {
		hash = path.Rem[hashIndex:]
		path = path.Rem[:hashIndex]
	}

	newPath := gudispatch.Path{
		PathDirectie: gudispatch.PathDirective{
			Host:     path.Host,
			Hash:     hash,
			Path:     path,
			Sequence: path.Sequence,
		},
		Rem:    rem,
		Params: params,
	}

	// Notify the subscribers.
	for _, sub := range b.subs {
		sub(newPath)
	}

	// Notify the kids with what is left in the Path.
	for _, child := range b.children {
		child.Resolve(newPath)
	}
}

// Subscribe adds a function to the subscription list for this
// resolver.
func (b *basicResolver) Subscribe(sub ResolveSubscriber) {
	b.subs = append(b.subs, sub)
}
