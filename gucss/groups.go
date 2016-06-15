package gucss

import (
	"io"
	"sync"
)

//==============================================================================

// NGroup defines a core structure for the internal groups that connects multiple
// groups together.
type NGroup struct {
	prev *NGroup
	grp  Group
}

// Root returns the root by walking the previous tree else return itself.
func (n *NGroup) Root() *NGroup {
	if n.prev != nil {
		return n.prev.Root()
	}

	return n
}

// NthRoot returns the root after a specific amount of steps backward the parent
// tree.
func (n *NGroup) NthRoot(ind int) *NGroup {
	if ind > 0 && n.prev != nil {
		return n.prev.NthRoot(ind - 1)
	}

	return n
}

// Append returns a new NGroup where this becomes the previous link.
func (n *NGroup) Append(g Group) *NGroup {
	return &NGroup{
		prev: n,
		grp:  g,
	}
}

//==============================================================================

// BaseGroup defines the base line group structure which others compose. It contains
// the base structure needed to meet the Group interface.
type BaseGroup struct {
	sel  string
	base *NGroup

	kids    []Group
	extends []Group

	properties Properties
	prw        sync.RWMutex
}

// NewBaseGroup return a new instance of a BaseGroup.
func NewBaseGroup(sel string, p Properties, root Group) *BaseGroup {
	bg := BaseGroup{
		sel:        sel,
		properties: p,
	}

	bg.base = &NGroup{grp: &bg}

	return &bg
}

// CSS flushes the contents of the group and its sub-groups into a css formatted
// styles data into the provided writer.
func (bg *BaseGroup) CSS(dst io.Writer) {

}

// Add augments the giving properties into the BaseGroup properties map.
func (bg *BaseGroup) Add(p Properties) {
	bg.prw.Lock()
	defer bg.prw.Unlock()

	for key, val := range p {
		bg.properties[key] = val
	}
}

// Extend returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector , childSelector {...}
func (bg *BaseGroup) Extend(sel string, p Properties) {
	root := bg.base.Root()
	if root.grp != bg {
		root.grp.Extend(sel, p)
		return
	}

}

// PreSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector ~ childSelector {...}
func (bg *BaseGroup) PreSibling(sel string, p Properties) Group {

	return nil
}

// PostSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector + childSelector {...}
func (bg *BaseGroup) PostSibling(sel string, p Properties) Group {

	return nil
}

// Within returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector  childSelector {...}
func (bg *BaseGroup) Within(sel string, p Properties) Group {

	return nil
}

// Child returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector > childSelector {...}
func (bg *BaseGroup) Child(sel string, p Properties) Group {

	return nil
}

// Sel returns the selector for the giving BaseGroup.
func (bg *BaseGroup) Sel() string {
	return bg.sel
}

// Root returns the main root of the BaseGroup parent chain.
func (bg *BaseGroup) Root() Group {
	return bg.base.Root().grp
}

// NthParent returns the Group after the giving count if the count is
// greater than zero else returns main root of all parents chain.
func (bg *BaseGroup) NthParent(back int) Group {
	var root *NGroup

	if back > 0 {
		root = bg.base.NthRoot(back)
	} else {
		root = bg.base.Root()
	}

	return root.grp
}

//==============================================================================

// ChildGroup defines a group for child selector properties.
type ExtendGroup struct {
	*BaseGroup
}

// CSS flushes the contents of the group and its sub-groups into a full css
// formatted styles.
func (bg *ChildGroup) CSS(dst io.Writer) {

}

//==============================================================================

// SiblingGroup defines the grouping of sibling properties.
type SiblingGroup struct {
	*BaseGroup
}

// CSS flushes the contents of the group and its sub-groups into a full css
// formatted styles.
func (bg *SiblingGroup) CSS(dst io.Writer) {

}

//==============================================================================

// WithinGroup defines a group for within properties.
type WithinGroup struct {
	*BaseGroup
}

// CSS flushes the contents of the group and its sub-groups into a full css
// formatted styles.
func (bg *WithinGroup) CSS(dst io.Writer) {

}

//==============================================================================

// ChildGroup defines a group for child selector properties.
type ChildGroup struct{}

// CSS flushes the contents of the group and its sub-groups into a full css
// formatted styles.
func (bg *ChildGroup) CSS(dst io.Writer) {

}

//==============================================================================
