package anix

import "github.com/gu-io/gu/anix/loop"

// Renderable defines a interface which provides an application of giving values
// to a provided target.
type Renderable interface {
	Apply(interface{})
	Render(interpolate float64)
	Update(step int64, delta float64)
}

// FunnelRenderer defines a function which returns a type which Renderable gets
// called to be applied against a target.
type FunnelRenderer func() Renderable

// FunnelTargets defines a function which receives targets to be used for applying
// against a series of pre-provided set of Renderables and returns a pointer of
// a loop.Clock type.
type FunnelTargets func(targets ...interface{}) *raf.Clock

// Animate registers a series of animation builders and callers which generate
// appliers for all provided elements.
func Animate(f ...FunnelRenderer) FunnelTargets {
	return func(targets ...interface{}) *raf.Clock {
		var renderers []Renderable

		for _, fr := range f {
			render := fr()

			func(r Renderable) {
				for _, target := range targets {
					r.Apply(target)
				}
			}(render)

			renderers = append(renderers, render)
		}

		return loop.Run(func(totalRun int64, step int64, delta float64) {
			for _, renderer := range renderers {
				renderer.Update(step, delta)
			}
		}, func(interpolate float64) {
			for _, renderer := range renderers {
				renderer.Apply(interpolate)
			}
		})
	}
}
