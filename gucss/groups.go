package gucss

// NGroup defines a core structure for the internal groups that connects multiple
// groups together.
type NGroup struct {
	prev  *NGroup
	child []NGroup
}

// BaseGroup defines the base line group structure which others compose. It contains
// the base structure needed to meet the Group interface.
type BaseGroup struct{}

// SiblingGroup defines the grouping of sibling properties.
type SiblingGroup struct{}

// WithinGroup defines a group for within properties.
type WithinGroup struct{}

// ChildGroup defines a group for child selector properties.
type ChildGroup struct{}
