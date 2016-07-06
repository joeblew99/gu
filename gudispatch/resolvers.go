package gudispatch

import (
	"strings"

	"github.com/influx6/faux/pattern"
)

// ResolveSubscriber defines a function type for a Resolver subcriber.
type ResolveSubscriber func(Path)

// Resolver defines an interface for a type that resolves a
// provided instance of a gudispatch.Path.
type Resolver interface {
	Subscribe(ResolveSubscriber)
	Register(Resolver)
	Resolve(Path)
}

// NewResolver returns a new instance of a structure that matches
// the Resolver interface.
func NewResolver(path string) Resolver {
	var br basicResolver

	if path != "" {
		br.matcher = URIMatcher(path)
	}

	return &br
}

// basicResolver defines a struct that implements
type basicResolver struct {
	matcher  pattern.URIMatcher
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
func (b *basicResolver) Resolve(path Path) {
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

	var hash, npath string

	hashIndex := strings.IndexRune(path.Rem, '#')
	if hashIndex != -1 {
		hash = path.Rem[hashIndex:]
		npath = path.Rem[:hashIndex]
	}

	newPath := Path{
		Rem:    rem,
		Params: params,
		PathDirective: PathDirective{
			Host:     path.Host,
			Hash:     hash,
			Path:     npath,
			Sequence: path.Sequence,
		},
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
