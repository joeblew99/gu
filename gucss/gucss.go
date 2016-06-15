package gucss

import "io"

//==============================================================================

// Writer defines an interface for a structure which reads out its content
// as css style into a passed in io.Writer.
type Writer interface {
	CSS(io.Writer)
}

//==============================================================================

// Properties defines a map of css property key: value pairs.
type Properties map[string]interface{}

// MergeProperties merges the contents of the src into the content of the
// destination
func MergeProperties(dest, src Properties) {
	for key, val := range src {
		dest[key] = val
	}
}

// CloneProperties clones a Properties map returning the cloned version.
func CloneProperties(m Properties) Properties {
	n := make(Properties)

	for key, val := range m {
		n[key] = val
	}

	return n
}

// CloneWith clones the set of valus and adds the key and Properties to
// the element.
func CloneWith(m []Properties, key string, newVal interface{}) []Properties {
	var mn []Properties

	for _, val := range m {
		cm := CloneProperties(val)
		cm[key] = newVal
		mn = append(mn, cm)
	}

	return mn
}

//==============================================================================

// Group defines a interface for a grouping of properties under a specific
// boundaries.
type Group interface {
	Writer

	Add(Properties)

	Root() Group
	NthParent(int) Group

	Child(string, Properties) Group
	Within(string, Properties) Group
	PreSibling(string, Properties) Group
	PostSibling(string, Properties) Group

	Extend(string, Properties)

	Sel() string
}

// Copy copies the properties provided into the giving groups set thereby creating
// this style of css properties.
//
// Format: .item box, .guz div, #deck{...}
func Copy(p Properties, groups ...Group) {

}

// CopyNS copies the properties provided into the giving groups agumented by the
// namespace provided, thereby creating this style of css properties.
//
// Format: .item box{NS}, .guz div{NS}, #deck{NS} {...}
// Example: .item box:hover, .guz div:hover, #deck:hover{...}
func CopyNS(ns string, p Properties, groups ...Group) {

}

//==============================================================================

// Media defines a media query generator for css properties.
// This produces media query output styles capsulated within them, they are submasters
// of the groups that flows from them and they are always the root for this groups
//
// @media (...) {...}
type Media interface {
	Writer
	Group(string)
}

//==============================================================================
