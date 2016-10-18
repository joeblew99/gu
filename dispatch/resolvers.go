package dispatch

import (
	"errors"
	"strings"

	"github.com/influx6/faux/pattern"
)

// ResolveSubscriber defines a function type for a Resolver subcriber.
type ResolveSubscriber func(Path)

// Resolvable defines an interface for a resolvable type object.
type Resolvable interface {
	Resolve(Path)
}

// Resolver defines an interface for a type that resolves a
// provided instance of a dispatch.Path.
type Resolver interface {
	Resolvable

	Flush()
	Register(Resolver)
	ResolvedPassed(ResolveSubscriber)
	ResolvedFailed(ResolveSubscriber)
	Test(string) (map[string]string, string, bool)
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
	fails    []ResolveSubscriber
	children []Resolver
}

// Flush resets the subscriptions and children lists to empty.
func (b *basicResolver) Flush() {
	b.subs = nil
	b.fails = nil
	b.children = nil
}

// Register adds a resolver into the list which will get triggerd
// when this resolver gets triggered, they will recieve a new Path
// made out of the remaining path from the Path received by this.
func (b *basicResolver) Register(r Resolver) {
	b.children = append(b.children, r)
}

// Test evaluates the giving path returning a map of parameters, a string of path left unmatch if
// possible(incase the resolver allows extra path lines within its pattern using the /*) and a boolean
// indicating when the path was indeed matched succesfully or if it was a total failure.
func (b *basicResolver) Test(path string) (map[string]string, string, bool) {
	return b.matcher.Validate(path)
}

// Resolve takes a `dispatch.Path` instance, matches the content
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

		// Notify the fail subscribers.
		for _, sub := range b.fails {
			sub(path)
		}

		return
	}

	// Copy over the parameter left from the previous path.
	for key, val := range path.Params {
		params[key] = val
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

// ResolvedFailed adds a function to the failed subscription list for this
// resolver.
func (b *basicResolver) ResolvedFailed(sub ResolveSubscriber) {
	b.fails = append(b.fails, sub)
}

// ResolvedPassed adds a function to the subscription list for this
// resolver.
func (b *basicResolver) ResolvedPassed(sub ResolveSubscriber) {
	b.subs = append(b.subs, sub)
}

//==============================================================================

// ResolvePath returns a new path created from match the giving matcher if
// valid else an non-nil error. It returns the remaining path of the giving
// path as the new path.
func ResolvePath(matcher pattern.URIMatcher, path Path) (Path, error) {
	params, rem, ok := matcher.Validate(path.Rem)
	if !ok {
		return Path{}, errors.New("Path does not match")
	}

	var hash, npath string

	hashIndex := strings.IndexRune(path.Rem, '#')
	if hashIndex != -1 {
		hash = path.Rem[hashIndex:]
		npath = path.Rem[:hashIndex]
	}

	return Path{
		Rem:    rem,
		Params: params,
		PathDirective: PathDirective{
			Host:     path.Host,
			Hash:     hash,
			Path:     npath,
			Sequence: path.Sequence,
		},
	}, nil
}
