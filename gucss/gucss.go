package gucss

import "io"

//==============================================================================

// PropertyRender defines an interface for a custom css property.
type PropertyRender interface {
	Render() string
}

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

// // Writer defines an interface for a structure which reads out its content
// // as css style into a passed in io.Writer.
// type Writer interface {
// 	CSS(io.Writer)
// }

// Render defines an interface for rendering the title/ID and the content of
// a structure.
type Render interface {
	Render(io.Writer)
}

//==============================================================================

// GroupTree defines an interface that returns the internal chain tree.
type GroupTree interface {
	Tree() *NGroup
}

// Group defines a interface for a grouping of properties under a specific
// boundaries.
type Group interface {
	GroupTree
	Render

	Add(Properties)

	Root() Group
	NthParent(int) Group

	NS(string, Properties) Group
	Child(string, Properties) Group
	Within(string, Properties) Group
	PreSibling(string, Properties) Group
	PostSibling(string, Properties) Group
	Extend(string, Properties) Group
	Media(device string, query string) Group

	Sel() string
	Selector() string
	IsRoot() bool
}

// NewRoot returns a clean root which has no tree and can provide
// a means of creating a clean stylesheet slate from. It acts as the
// root before '*' within a css stylesheet.
func NewRoot() Group {
	bg := newbaseGroup("", "", nil, nil)
	bg.root = true
	return bg
}

// Copy copies the properties provided into the giving groups set thereby creating
// this style of css properties.
//
// Format:
// 1. If all the groups were of the same root
// 	.item box, .guz div, #deck{...}
//
// 2. If each group has a different root
// 	.item box{...}
//  .guz div{...}
//  #deck{...}
func Copy(p Properties, groups ...Group) {
	for _, group := range groups {
		group.Root().Extend(group.Sel(), p)
	}
}

// CopyNS copies the properties provided into the giving groups agumented by the
// namespace provided, thereby creating this style of css properties.
//
// Format:
//   1. When the groups are of the same root
//	  .item box{NS}, .guz div{NS}, #deck{NS} {...}
//
//      Example: .item box:hover, .guz div:hover, #deck:hover{...}
//
//   2. When the groups are of the different root
//	  .item box{NS}
//		.guz div{NS}
//		#deck{NS} {...}
//
//     Example: .item box:hover, .guz div:hover, #deck:hover{...}
func CopyNS(ns string, p Properties, groups ...Group) {
	for _, group := range groups {
		group.Root().NS(group.Sel(), p)
	}
}
