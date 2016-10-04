package apidesign

import (
	. "github.com/influx6/gu/gutrees"
)

type (
	// DSL defines a function type which is used to generate the contents of a Def(Definition).
	DSL func()

	// DefType provides a easy int type associated with a giving definition.
	DefType int

	// Def defines an interface for all gu  Definitions structure which embodies different structure types.
	Def interface {
		Context() string
		Finalize() error
	}

	// DefSet defines a slice of Def(Definition) slice.
	DefSet []Def

	// Root defines a structure which handles the sub definition within it self.
	Root interface {
		DependsOn() []Root
		IteratorSet(set DefSet)
	}
)
