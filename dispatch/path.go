package dispatch

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-humble/detect"
	"github.com/gopherjs/gopherjs/js"
	"github.com/influx6/faux/pattern"
)

//==============================================================================

// PathDirective represent the current path and hash values.
type PathDirective struct {
	Host     string
	Hash     string
	Path     string
	Sequence string
}

// String returns the hash and path.
func (p PathDirective) String() string {
	return fmt.Sprintf("%s%s", pattern.TrimEndSlashe(p.Path), p.Hash)
}

//==============================================================================

// MakePath returns PathDirective based on the oath string provided.
func MakePath(path string) (PathDirective, error) {
	ups, err := url.Parse(path)
	if err != nil {
		return PathDirective{}, err
	}

	return PathDirective{
		Path: ups.Path,
		Host: ups.Host,
		Hash: ups.Fragment,
	}, nil
}

//==============================================================================

// WrapHandler returns a function of type PathHandler by wrapping the
// provided no argument function in one.
func WrapHandler(fx func()) func(Path) {
	return func(_ Path) {
		fx()
	}
}

// Path defines a representation of a location path matching a specific sequence.
type Path struct {
	PathDirective
	Rem    string
	Params map[string]string
}

// GetLocationPath returns the current Path associated with the current location.
// Using the complete path has the Remaining path value.
func GetLocationPath() Path {
	directive := GetLocationDirective()

	return Path{
		Rem:           directive.String(),
		PathDirective: directive,
		Params:        make(map[string]string),
	}
}

// UseLocation returns the current Path associated with the provided path.
// Using the complete path has the Remaining path value.
func UseLocation(path string) Path {
	directive, err := MakePath(path)
	if err != nil {
		return Path{}
	}

	return Path{
		Rem:           directive.String(),
		PathDirective: directive,
		Params:        make(map[string]string),
	}
}

// UseLocationHash returns the current Path associated with the provided path.
// Using the hash path has the Remaining path value.
func UseLocationHash(path string) Path {
	directive, err := MakePath(path)
	if err != nil {
		return Path{}
	}

	return Path{
		Rem:           directive.Hash,
		PathDirective: directive,
		Params:        make(map[string]string),
	}
}

// UseDirective returns a Path which wraps the provided directive using
// the full path including the hash string.
func UseDirective(directive PathDirective) Path {
	return Path{
		Rem:           directive.String(),
		PathDirective: directive,
		Params:        make(map[string]string),
	}
}

// UseHashDirective returns a Path which wraps the provided directive using
// the full path including the hash string.
func UseHashDirective(directive PathDirective) Path {
	return Path{
		Rem:           directive.Hash,
		PathDirective: directive,
		Params:        make(map[string]string),
	}
}

// GetLocationHashAsPath returns the current Path associated with the current location.
// Using the current hash has the Remaining path value.
func GetLocationHashAsPath() Path {
	directive := GetLocationDirective()

	return Path{
		Rem:           directive.Hash,
		PathDirective: directive,
		Params:        make(map[string]string),
	}
}

//==============================================================================

// AttachURLTo takes the giving pattern, matches it against changes provided by
// the current PathObserver, if the full URL(i.e Path+Hash) matches then fires
// the provided function.
func AttachURLTo(pattern string, fx func(Path), fail func(Path)) {
	matcher := URIMatcher(pattern)

	Subscribe(func(p PathDirective) {
		if params, rem, found := matcher.Validate(p.String()); found {
			fx(Path{
				Params:        params,
				Rem:           rem,
				PathDirective: p,
			})
			return
		}

		if fail != nil {
			fail(Path{PathDirective: p})
		}
	})

	// Follow the current location to see if we should be triggered.
	Follow(GetLocation())
}

// AttachHashTo takes the giving pattern, matches it against changes provided by
// the current PathObserver, if the URL hash matches then fires
// the provided function.
func AttachHashTo(pattern string, fx func(Path), fail func(Path)) {
	matcher := URIMatcher(pattern)

	Subscribe(func(p PathDirective) {
		if params, rem, found := matcher.Validate(p.Hash); found {
			fx(Path{
				Params:        params,
				Rem:           rem,
				PathDirective: p,
			})
			return
		}

		if fail != nil {
			fail(Path{PathDirective: p})
		}
	})

	// Follow the current location to see if we should be triggered.
	Follow(GetLocation())
}

//==============================================================================

// ResolveAttachURL takes the giving pattern, matches it against changes provided by
// the current PathObserver, if the full URL(i.e Path+Hash) matches then fires
// the provided function.
func ResolveAttachURL(pattern string, fx func(Path), fail func(Path)) Resolver {
	resolver := NewResolver(pattern)
	resolver.ResolvedPassed(fx)
	resolver.ResolvedFailed(fail)

	Subscribe(func(p PathDirective) {
		resolver.Resolve(Path{
			Rem:           p.String(),
			PathDirective: p,
		})
	})

	// Follow the current location to see if we should be triggered.
	Follow(GetLocation())
	return resolver
}

// ResolveAttachHash takes the giving pattern, matches it against changes provided by
// the current PathObserver, if the URL hash matches then fires
// the provided function.
func ResolveAttachHash(pattern string, fx func(Path), fail func(Path)) Resolver {
	resolver := NewResolver(pattern)
	resolver.ResolvedPassed(fx)
	resolver.ResolvedFailed(fail)

	Subscribe(func(p PathDirective) {
		resolver.Resolve(Path{
			Rem:           p.Hash,
			PathDirective: p,
		})
	})

	// Follow the current location to see if we should be triggered.
	Follow(GetLocation())
	return resolver
}

//==============================================================================

// PathSequencer provides a function to convert either the path/hash into a
// dot seperated sequence string for use with States.
type PathSequencer func(path string, hash string) string

// HashSequencer provides a PathSequencer that returns the hash part of a url,
// as the path sequence.
func HashSequencer(path, hash string) string {
	cleanHash := strings.Replace(hash, "#", ".", -1)
	return strings.Replace(cleanHash, "/", ".", -1)
}

// URLPathSequencer provides a PathSequencer that returns the path part of a url,
// as the path sequence.
func URLPathSequencer(path, hash string) string {
	return strings.Replace(path, "/", ".", -1)
}

//==============================================================================

//ErrNotSupported is returned when a feature requested is not supported by the environment
var ErrNotSupported = errors.New("Feature not supported")

// PathObserver represent any continouse changing route path by the browser
type PathObserver struct {
	usingHash bool
	sequencer PathSequencer
}

// NewPathObserver returns a new PathObserver instance.
func NewPathObserver(ps PathSequencer) *PathObserver {
	if ps == nil {
		ps = HashSequencer
	}

	return &PathObserver{
		sequencer: ps,
	}
}

//==============================================================================

// GetLocation returns the path and hash of the browsers location api else
// panics if not in a browser.
func GetLocation() (host string, path string, hash string) {
	if !detect.IsBrowser() {
		return
	}

	loc := js.Global.Get("location").String()
	ups, err := url.Parse(loc)
	if err != nil {
		return
	}

	host = ups.Host
	path = ups.Path
	hash = ups.Fragment
	return
}

// GetLocationDirective returns the current location directive.
func GetLocationDirective() PathDirective {
	host, path, hash := GetLocation()

	return PathDirective{
		Host:     host,
		Hash:     hash,
		Path:     path,
		Sequence: URLPathSequencer(path, hash),
	}
}

//==============================================================================

// HashChangePath returns a path observer path changes
func HashChangePath(ps PathSequencer) *PathObserver {
	panicBrowserDetect()
	path := NewPathObserver(ps)
	path.usingHash = true

	js.Global.Set("onhashchange", func() {
		path.Follow(GetLocation())
	})

	return path
}

//==============================================================================

// BrowserSupportsPushState checks if browser supports pushState
func BrowserSupportsPushState() bool {
	if !detect.IsBrowser() {
		return false
	}

	return (js.Global.Get("onpopstate") != js.Undefined) &&
		(js.Global.Get("history") != js.Undefined) &&
		(js.Global.Get("history").Get("pushState") != js.Undefined)
}

//==============================================================================

// PopStatePath returns a path observer path changes
func PopStatePath(ps PathSequencer) (*PathObserver, error) {
	panicBrowserDetect()

	if !BrowserSupportsPushState() {
		return nil, ErrNotSupported
	}

	path := NewPathObserver(ps)

	js.Global.Set("onpopstate", func() {
		path.Follow(GetLocation())
	})

	return path, nil
}

// Follow creates a Pathspec from the hash and path and sends it
func (p *PathObserver) Follow(host, path, hash string) {
	Dispatch(PathDirective{
		Host:     host,
		Hash:     hash,
		Path:     path,
		Sequence: p.sequencer(path, hash),
	})
}

//==============================================================================

// PushDOMState adds a new state the dom push history.
func PushDOMState(path string) {
	if !detect.IsBrowser() {
		return
	}

	// Use the advance pushState feature.
	js.Global.Get("history").Call("pushState", nil, "", path)

	// Set the browsers hash accordinly.
	js.Global.Get("location").Set("hash", path)
}

// SetDOMHash sets the dom location hash.
func SetDOMHash(hash string) {
	panicBrowserDetect()
	js.Global.Get("location").Set("hash", hash)
}

//==============================================================================

// HistoryProvider wraps the PathObserver with methods that allow easy control of
// client location
type HistoryProvider struct {
	*PathObserver
}

// History returns a new PathObserver and depending on browser support will
// either use the popState or HashChange.
func History(ps PathSequencer) *HistoryProvider {
	pop, err := PopStatePath(ps)

	if err != nil {
		pop = HashChangePath(ps)
	}

	return &HistoryProvider{pop}
}

// Go changes the path of the current browser location depending on wether
// its underline observer is hashed based or pushState based,
// it will use SetDOMHash or PushDOMState appropriately.
func (h *HistoryProvider) Go(path string) {
	if h.usingHash {
		SetDOMHash(path)
		return
	}
	PushDOMState(path)
}

//==============================================================================

// panicBrowserDetect panics if the current gc is not a browser based
// one.
func panicBrowserDetect() {
	if !detect.IsBrowser() {
		panic("expected to be used in a dom/browser env")
	}
}
