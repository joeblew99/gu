package loop

import (
  "github.com/gu-io/gu/anix/loop"
)

// RenderFn defines the function called to render an animation after its update
// cycles. On every call it is provided with an interpolation value to improving
// drawing accuracy according to it's current fps call.
type RenderFn func(interpolation float64)

// UpdateFn defines the function called during update cycles for animations. It
// is provided with the totaltime passed, the delta value calculated and current
// step.
type UpdateFn func(totalTime int64, delta float64)

// Options defines a struct configuration which affects the running of the animation
// calls with these values.
type Options struct{
  Fps int64
}

// Run registers the pair of function calls for animation and returns a central
// loop.Clock pointer instance which allows the control of the animation calls.
func Run(un UpdateFn, rn RenderFn, o *Options) *loop.Clock{

  return loop.New(func(tick float64){

  })
}
