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

// CSSRender defines a package level handle for access the base css parser.
var CSSRender render

type render struct{}

// CSSMediaFormat defines the format for each css property.
const CSSMediaFormat = "(%+s: %+s)"

// RenderMediaAttribute renders the attribute format for a css key: value pair.
func (r render) RenderMediaAttribute(key string, val interface{}, dst io.Writer) {
	switch val.(type) {
	case PropertyRender:
		dst.Write([]byte(fmt.Sprintf(CSSMediaFormat, key, val.(PropertyRender).Render())))
		return
	default:
		dst.Write([]byte(fmt.Sprintf(CSSMediaFormat, key, val)))
		return
	}
}

// CSSPropertyFormat defines the format for each css property.
const CSSPropertyFormat = "\t%+s: %+s;\n"

// RenderAttribute renders the attribute format for a css key: value pair.
func (r render) RenderAttribute(key string, val interface{}, dst io.Writer) {
	switch val.(type) {
	case PropertyRender:
		dst.Write([]byte(fmt.Sprintf(CSSPropertyFormat, key, val.(PropertyRender).Render())))
		return
	default:
		dst.Write([]byte(fmt.Sprintf(CSSPropertyFormat, key, val)))
		return
	}
}

// CSSPropertyBracket defines the bracket format for the css bracket.
const CSSPropertyBracket = "%+s {\n%+s}\n\n"

// Render calls the giving render function for the giving properties and selector
// into the provided destination writer.
func (r render) Render(sel string, p Properties, dst io.Writer) {
	var ws bytes.Buffer

	for key, val := range p {
		switch val.(type) {
		case PropertyRender:
			ws.Write([]byte(fmt.Sprintf(CSSPropertyFormat, key, val.(PropertyRender).Render())))
			continue
		default:
			ws.Write([]byte(fmt.Sprintf(CSSPropertyFormat, key, val)))
			continue
		}
	}

	dst.Write([]byte(fmt.Sprintf(CSSPropertyBracket, sel, ws.Bytes())))
}

//==============================================================================

// Media defines a structure which constructs the media query tree for
// gucss.
type Media struct {
	roots  []Group
	query  string
	device string
}

// NewMedia returns a new Media instance with the provided group.
func NewMedia(query string, device string, g Group) *Media {
	mn := Media{query: query, device: device}

	if g != nil {
		mn.roots = append(mn.roots, g)
	}

	return &mn
}

// Add adds a new group into the list of groups to be wrapped by the query.
func (m *Media) Add(g Group) {
	m.roots = append(m.roots, g)
}

// Render outputs the content of the media into the provided writer.
func (m *Media) Render(dst io.Writer) {
	for _, g := range m.roots {
		var b bytes.Buffer
		m.root.Render(&b)

		query := fmt.Sprintf("@media %s %s", m.device, m.query)
		qm := fmt.Sprintf(CSSPropertyBracket, query, b.Bytes())

		dst.Write([]byte(qm))
	}
}

//==============================================================================

// BaseGroup defines the base line group structure which others compose. It contains
// the base structure needed to meet the Group interface.
type BaseGroup struct {
	base *NGroup

	sel       string
	seperator string

	root bool

	kids []Render

	properties Properties
	prw        sync.RWMutex
}

// NewBaseGroup return a new instance of a BaseGroup.
func NewBaseGroup(sel string, p Properties, root Group) *BaseGroup {
	bg := &BaseGroup{
		sel:        sel,
		properties: p,
		seperator:  " ",
	}

	if root != nil {
		bg.base = root.Tree().Append(bg)
	} else {
		bg.base = &NGroup{grp: bg}
	}

	return bg
}

// Tree returns the internal group chain tree.
func (bg *BaseGroup) Tree() *NGroup {
	return bg.base
}

// Render implements the Render interface for the BaseGroup and returns the
// property title and content for this group.
func (bg *BaseGroup) Render(dst io.Writer) {
	if !bg.root {
		bg.prw.RLock()
		CSSRender.Render(bg.Selector(), bg.properties, dst)
		bg.prw.RUnlock()
	}

	bg.renderKids(dst)
}

// Add augments the giving properties into the BaseGroup properties map.
func (bg *BaseGroup) Add(p Properties) {
	bg.prw.Lock()
	defer bg.prw.Unlock()

	if bg.properties == nil {
		return
	}

	for key, val := range p {
		bg.properties[key] = val
	}
}

//==============================================================================

// Extend returns a new Group based off the selector of the previous using
// the format:
//
//    rootSel parentSelector , rootSel sel {...}
func (bg *BaseGroup) Extend(sel string, p Properties) Group {
	var sig string

	if !bg.root {
		sig = ", "
	}

	newSel := bg.NthParent(1).Selector() + " " + sel
	kid := newSignGroup(sig, newSel, p, bg)
	bg.kids = append(bg.kids, kid)
	return bg
}

// PreSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector ~ childSelector {...}
func (bg *BaseGroup) PreSibling(sel string, p Properties) Group {
	var sig string

	if !bg.root {
		sig = " ~ "
	}

	kid := newSignGroup(sig, sel, p, bg)
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
	var sig string

	if !bg.root {
		sig = ns
	}

	kid := newSignGroup(sig, " ", p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// PostSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector + childSelector {...}
func (bg *BaseGroup) PostSibling(sel string, p Properties) Group {
	var sig string

	if !bg.root {
		sig = " + "
	}

	kid := newSignGroup(sig, sel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Within returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector  childSelector {...}
func (bg *BaseGroup) Within(sel string, p Properties) Group {
	var sig string

	if !bg.root {
		sig = " "
	}

	kid := newSignGroup(sig, sel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Child returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector > childSelector {...}
func (bg *BaseGroup) Child(sel string, p Properties) Group {
	var sig string

	if !bg.root {
		sig = " > "
	}

	kid := newSignGroup(sig, sel, p, bg)
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
	root := bg.NthParent(1)
	if root != nil && root != bg {
		sel := root.Selector()

		if sel != "" {
			return root.Selector() + bg.seperator + bg.sel
		}
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

// renderKids renders the children of the BaseGroup.
func (bg *BaseGroup) renderKids(dst io.Writer) {

	// Wrap up the children tree and build the appropriate content
	for _, kid := range bg.kids {
		kid.Render(dst)
	}
}

//==============================================================================

// newSignGroup returns a new instance of a group with a moded sign.
// The format:
//
//    parentSelector {{signature}} childSelector {...}
func newSignGroup(sign string, sel string, p Properties, root Group) Group {
	bg := NewBaseGroup(sel, p, root)
	bg.seperator = sign
	return bg
}

//==============================================================================

// Child defines a group for child selector properties.
// The format:
//
//    parentSelector > childSelector {...}
func Child(sel string, p Properties, root Group) Group {
	return root.Child(sel, p)
}

//==============================================================================

// NS defines a group for child selector properties.
// The format:
//
//    parentSelector{{NS}} {...}
func NS(ns string, p Properties, root Group) Group {
	return root.NS(ns, p)
}

//==============================================================================

// Extend defines a group for child selector properties.
// The format:
//
//    parentSelector , childSelector {...}
func Extend(sel string, p Properties, root Group) Group {
	return root.Extend(sel, p)
}

//==============================================================================

// PreSibling defines the grouping of sibling properties.
// The format:
//
//    Pre: parentSelector ~ childSelector {...}
func PreSibling(sel string, p Properties, root Group) Group {
	return root.PreSibling(sel, p)
}

// PostSibling defines the grouping of sibling properties.
// The format:
//
//    Post: parentSelector + childSelector {...}
func PostSibling(sel string, p Properties, root Group) Group {
	return root.PostSibling(sel, p)
}

//==============================================================================

// Within defines a group for within properties.
// The format:
//
//    parentSelector  childSelector {...}
func Within(sel string, p Properties, root Group) Group {
	return root.Within(sel, p)
}

//==============================================================================
