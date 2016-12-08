package anix

import (
	"time"
)

// UpdateMux defines a function type which is to be called on the update cycle.
type UpdateMux func(tTime int64, deltaTime float64)

// RenderMux defines a function type for render function calls.
type RenderMux func(interpolate float64)

// CyclerAttr defines values passed to setup the cycle runner.
type CyclerAttr struct {
	Fps      float64
	MaxDelta float64
	Level    int
}

// Cycler defines a type which handles cycling of timer calls providing a gameloop
// calls.
type Cycler struct {
	paused      bool
	fps         float64
	step        float64
	last        time.Time
	start       time.Time
	ts          float64
	accumulator float64
}

// Start initiates the cycler operations to begin running.
func (c *Cycler) Start() {}

// Pause stops the cycle operation, ensure no calls of update/renders are called.
func (c *Cycler) Pause() {}

// Resume sets the cycle to begin animation once again.
func (c *Cycler) Resume() {}

// Stop ends the cycler calls.
func (c *Cycler) Stop() {}

// nextFrame handles update of the internal logic of the cycler.
func (c *Cycler) nextFrame() {
	now := time.now()
	dt := now.Sub(c.last).Seconds() / 1000

	if dt > 0.2 {
		dt = c.step
	}

	c.last = now
	c.accumulator += dt

	for c.accumulator > c.step {
		// do update call.
		this.ts += dt
		c.accumulator -= c.step
	}

	interpolation := c.accumulator / c.step
	// do render calls.

}
