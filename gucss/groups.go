package gucss

import (
	"bytes"
	"fmt"
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

//==============================================================================

const (
	// CSSPropertyBracket defines the bracket format for the css bracket.
	CSSPropertyBracket = "\n %+s {\n%+s\n}\n"

	// CSSPropertyFormat defines the format for each css property.
	CSSPropertyFormat = "\t\t%+s: %+s\n"
)

// Render implements the Render interface for the BaseGroup and returns the
// property title and content for this group.
func (bg *BaseGroup) Render() []byte {
	var ws bytes.Buffer

	for key, val := range bg.properties {
		switch val.(type) {
		case PropertyRender:
			ws.Write([]byte(fmt.Sprintf(CSSPropertyFormat, key, val.(PropertyRender).Render())))
			continue
		default:
			ws.Write([]byte(fmt.Sprintf(CSSPropertyFormat, key, val)))
			continue
		}
	}

	return ws.Bytes()
}

//==============================================================================

// CSS flushes the contents of the group and its sub-groups into a css formatted
// styles data into the provided writer.
func (bg *BaseGroup) CSS(dst io.Writer) {
	var writer bytes.Buffer

	// Map out the properties into proper css style processing.
	writer.Write([]byte(fmt.Sprintf(CSSPropertyBracket, bg.Selector(), bg.Render())))

	// Wrap up the extends list and pull down into appropriate content.
	for _, extend := range bg.extends {
		sel := bg.Selector() + extend.Selector()
		writer.Write([]byte(fmt.Sprintf(CSSPropertyBracket, sel, extend.Render())))
	}

	// Wrap up the children tree and build the appropriate content
	for _, kid := range bg.kids {
		kid.CSS(&writer)
	}
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

	kid := NewSignatureGroup(",  ", sel, p, bg)
	bg.extends = append(bg.extends, kid)
}

// PreSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector ~ childSelector {...}
func (bg *BaseGroup) PreSibling(sel string, p Properties) Group {
	kid := NewSignatureGroup(" ~ ", sel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// NS returns a group where the current selector gets attached the namespace value
// provided.
// The format:
//
//    parentSelector{{NS}} {}
//       eg
//        if NS = ":hover"
//         then 'div' becomes 'div:hover'
func (bg *BaseGroup) NS(ns string, p Properties) Group {
	kid := NewSignatureGroup(ns, bg.sel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// PostSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector + childSelector {...}
func (bg *BaseGroup) PostSibling(sel string, p Properties) Group {
	kid := NewSignatureGroup(" + ", sel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Within returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector  childSelector {...}
func (bg *BaseGroup) Within(sel string, p Properties) Group {
	kid := NewSignatureGroup(" ", sel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Child returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector > childSelector {...}
func (bg *BaseGroup) Child(sel string, p Properties) Group {
	kid := NewSignatureGroup(" > ", sel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Sel returns the selector for the giving BaseGroup.
func (bg *BaseGroup) Sel() string {
	return bg.sel
}

// Selector returns the selector with the appropriate sign for the type
// for the giving BaseGroup.
func (bg *BaseGroup) Selector() string {
	root := bg.Root()
	if root != nil && root != bg {
		return root.Selector() + " " + bg.sel
	}

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

// SignatureGroup defines a group for a giving signature selector properties.
// The format:
//
//    parentSelector {{signature}} childSelector {...}
type SignatureGroup struct {
	*BaseGroup
	signature string
}

// NewSignatureGroup returns a new instance of a signature group
func NewSignatureGroup(sign, sel string, p Properties, root Group) *SignatureGroup {
	sm := SignatureGroup{
		signature: sign,
		BaseGroup: NewBaseGroup(sel, p, root),
	}

	return &sm
}

// Selector returns the selector with the appropriate sign for the type
// for the giving BaseGroup.
func (s *SignatureGroup) Selector() string {
	return s.BaseGroup.Selector() + s.signature
}

//==============================================================================

// Child defines a group for child selector properties.
// The format:
//
//    parentSelector > childSelector {...}
func Child(sel string, p Properties, root Group) Group {
	return NewSignatureGroup(" > ", sel, p, root)
}

//==============================================================================

// NS defines a group for child selector properties.
// The format:
//
//    parentSelector{{NS}}childSelector {...}
func NS(ns string, sel string, p Properties, root Group) Group {
	return NewSignatureGroup(ns, sel, p, root)
}

//==============================================================================

// Extend defines a group for child selector properties.
// The format:
//
//    parentSelector , childSelector {...}
func Extend(sel string, p Properties, root Group) Group {
	return NewSignatureGroup(", ", sel, p, root)
}

//==============================================================================

// PreSibling defines the grouping of sibling properties.
// The format:
//
//    Pre: parentSelector ~ childSelector {...}
func PreSibling(sel string, p Properties, root Group) Group {
	return NewSignatureGroup(" ~ ", sel, p, root)
}

// PostSibling defines the grouping of sibling properties.
// The format:
//
//    Post: parentSelector + childSelector {...}
func PostSibling(sel string, p Properties, root Group) Group {
	return NewSignatureGroup(" + ", sel, p, root)
}

//==============================================================================

// Within defines a group for within properties.
// The format:
//
//    parentSelector  childSelector {...}
func Within(sel string, p Properties, root Group) Group {
	return NewSignatureGroup(" ", sel, p, root)
}

//==============================================================================
