package design

import (
	"github.com/influx6/gu/dispatch"
	"github.com/influx6/gu/trees"
)

// ResourcesRouter handles rendering giving resource output using the
// provided resources and the resource renderer.
type ResourcesRouter struct {
	core   *Resources
	engine ResourceRenderer
}

// New returns a new instance of the ResourcesRouter to handle a direct
// rendering of the output of
func New(res *Resources, engine ResourceRenderer) *ResourcesRouter {
	rsx := &ResourcesRouter{
		core:   res,
		engine: engine,
	}

	return rsx
}

// Begin intializes the resources located within the router.
// Setting up any required listeners and also takes in a series of boolean
// values but only the first is used, as this signature is used for sake
// of convenience and for a simpler method. This provides a means to delegate
// the router to watch for only hash paths instead of full paths.
func (rsx *ResourcesRouter) Begin(hashOnly ...bool) *ResourcesRouter {
	var watchHash bool

	if len(hashOnly) != 0 {
		watchHash = hashOnly[0]
	}

	rsx.core.Init()
	dispatch.Subscribe(func(pd dispatch.PathDirective) {
		if !watchHash {
			rsx.core.Resolve(dispatch.UseDirective(pd))
			return
		}

		rsx.core.Resolve(dispatch.UseHashDirective(pd))
	})

	return rsx
}
