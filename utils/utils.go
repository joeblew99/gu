package utils

import "time"

// TickToHighResTimer provides a method to transform requestAnimationFrame
// clock elapsed time into a appropriate time.Duration
func TickToHighResTimer(ms float64) time.Duration {
	return time.Duration(ms * float64(time.Millisecond))
}
