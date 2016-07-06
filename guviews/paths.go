package guviews

import "github.com/influx6/gu/gudispatch"

// FailNormal is a fail function type that accepts no arguments.
type FailNormal func()

// FailPath is a fail function type that accepts a gudispatch.Path.
type FailPath func(gudispatch.Path)

// WrapNormal returns a function of type FailPath by wrapping the
// provided no argument function in one.
func WrapNormal(fx func()) FailPath {
	return func(_ gudispatch.Path) {
		fx()
	}
}

// AttachURL attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the full URL(Path+Hash).
// failFn must either be a FailNormal, FailPath or nil.
func AttachURL(pattern string, v Views, failFn interface{}) {
	var failCb FailPath

	switch failFn.(type) {
	case FailNormal:
		failCb = WrapNormal(failFn.(FailNormal))
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
		failCb = WrapNormal(failFn.(FailNormal))
	case FailPath:
		failCb = failFn.(FailPath)
	}

	gudispatch.AttachHash(pattern, func(p gudispatch.Path) {
		v.Path(p)
	}, failCb)
}
