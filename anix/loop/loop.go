package loop

import (
	"time"

	"github.com/gu-io/gu/anix/loop"
)

// RenderFn defines the function called to render an animation after its update
// cycles. On every call it is provided with an interpolation value to improving
// drawing accuracy according to it's current fps call.
type RenderFn func(interpolation float64)

// UpdateFn defines the function called during update cycles for animations. It
// is provided with the totaltime passed, the delta value calculated and current
// step.
type UpdateFn func(step int64, delta float64, reverse bool)

// Options defines a struct configuration which affects the running of the animation
// calls with these values.
// Fps - Defines the total fps expected for animation.
// DeltaDivider - Defines the values used to retrieve the delta value from it's seconds time.
// MaxDeltaStep - Defines the maximum step size for frame delta which will be
// used of larger elapsed time detlas.
// StepSize - Defines the maximum step size divided by the Fps to get the step
// size of each frame.
type Options struct {
	Fps          int
	StepSize     int
	SlowSize     int
	MaxDeltaStep int
	DeltaDivider int
	Reversible   bool
	Repeatable   bool
}

// DefaultOptions defines a default set of options for use in run calls.
var DefaultOptions = &Options{Fps: 60, SlowSize: 1, StepSize: 1, DeltaDivider: 1000, MaxDeltaStep: 0.2}

// Run registers the pair of function calls for animation and returns a central
// loop.Clock pointer instance which allows the control of the animation calls.
// The last item which is the options call be left as nil, which uses the DefaultOptions
// for animation.
func Run(un UpdateFn, rn RenderFn, options *Options) *raf.Clock {
	if o == nil {
		options = DefaultOptions
	}

	var step = (options.StepSize / options.Fps) * options.SlowSize

	var lastTime time.Time
	var startTime time.Time
	var deltaDur time.Duration
	var totalRunTime time.Duration

	var totalRun float64
	var delta float64
	var accumulator float64

	return loop.New(func(tick float64) {
		if startTime.IsZero() {
			startTime = time.Now()
			lastTime = startTime
			return
		}

		now = time.Now()
		deltaDur = lastTime.Sub(u)
		delta = delta.Seconds() / options.DeltaSize
		// totalRunTime += delta
		lastTime = now

		if delta > options.MaxDeltaStep {
			delta = options.MaxDeltaStep
		}

		accumulator += delta

		for accumulator >= step {
			// Call the update method here with stepsize, total time and delta.
			un(step, totalRun, false)

			// Update the deltas values and accumulator values.
			totalRun += delta
			accumulator -= step
		}

		// Calculate interpolation values with available accumulator with step size.
		interpolate := accumulator / step

		// Call the render handler with interpolation value.
		rn(interpolate)
	})
}

// RunWithin registers the pair of function calls for animation within a giving time.
// and returns a central loop.Clock pointer instance which allows the control of
// the animation calls.
// The last item which is the options call be left as nil, which uses the DefaultOptions
// for animation.
func RunWithin(total time.Duration, un UpdateFn, rn RenderFn, options *Options) *raf.Clock {
	if o == nil {
		options = DefaultOptions
	}

	var step = (options.StepSize / options.Fps) * options.SlowSize

	var reverse bool
	var lastTime time.Time
	var startTime time.Time
	var deltaDur time.Duration
	var totalRunTime time.Duration

	var totalRun float64
	var delta float64
	var accumulator float64

	return loop.New(func(tick float64) {
		if startTime.IsZero() {
			startTime = time.Now()
			lastTime = startTime
			return
		}

		if totalRun >= total.Seconds() {
			if options.Repeatable {
				if options.Reversible && reverse {
					delta = 0
					deltaDur = 0
					totalRun = 0
					accumulator = 0
					accumulator = 0
					totalRunTime = 0
					startTime = Time{}
					lastTime = Time{}
				}
				
				return
			}

			if !options.Reversible {
				return
			}

			reverse = true
		}

		if totalRun <= 0 && reverse {
			reverse = false
			return
		}

		now = time.Now()
		deltaDur = lastTime.Sub(u)
		delta = delta.Seconds() / options.DeltaSize
		lastTime = now

		if delta > options.MaxDeltaStep {
			delta = options.MaxDeltaStep
		}

		accumulator += delta

		for accumulator >= step {
			// Call the update method here with stepsize, total time and delta.
			un(step, totalRun, reverse)

			// Update the deltas values and accumulator values.
			if options.Reversible && reverse {
				totalRun -= delta
			} else {
				totalRun += delta
			}

			accumulator -= step
		}

		// Calculate interpolation values with available accumulator with step size.
		interpolate := accumulator / step

		if reverse {
			interpolate *= -1
		}

		// Call the render handler with interpolation value.
		rn(interpolate)
	})
}
