package guviews

import "github.com/influx6/gu/gudispatch"

// AttachURL attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the full URL(Path+Hash).
// failFn must either be a FailNormal, FailPath or nil.
func AttachURL(pattern string, v Views, failFn func(gudispatch.Path)) {
	gudispatch.ResolveAttachURL(pattern, v.Path, failFn)
}

// AttachHash attaches the view to the provided Route pattern,
// Using the internal route pattern, it matches all route changes
// and checks against the URL hash.
// failFn must either be a FailNormal, FailPath or nil.
func AttachHash(pattern string, v Views, failFn func(gudispatch.Path)) {
	gudispatch.ResolveAttachHash(pattern, v.Path, failFn)
}
