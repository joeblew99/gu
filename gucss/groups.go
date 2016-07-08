package gucss

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
)

//==============================================================================

// minifyOption defines a variable where all the option to minify css rendering
// is turned on.
var minifyOption = struct {
	minify bool
	rl     sync.RWMutex
}{}

// SetMinitfy sets the current value of the minify option.
func SetMinitfy(state bool) {
	minifyOption.rl.Lock()
	defer minifyOption.rl.Unlock()
	minifyOption.minify = state
}

func getMinitfy() bool {
	minifyOption.rl.RLock()
	defer minifyOption.rl.RUnlock()
	return minifyOption.minify
}

//==============================================================================

// CSSRender defines a package level handle for access the base css parser.
var CSSRender render

type render struct{}

// cssMediaFormat defines the format for each css property.
const cssMediaFormat = "(%+s: %+s)"

// RenderMediaAttribute renders the attribute format for a css key: value pair.
func (r render) RenderMediaAttribute(key string, val interface{}, dst io.Writer) {
	switch val.(type) {
	case PropertyRender:
		dst.Write([]byte(fmt.Sprintf(cssMediaFormat, key, val.(PropertyRender).Render())))
		return
	default:
		dst.Write([]byte(fmt.Sprintf(cssMediaFormat, key, val)))
		return
	}
}

// cssPropertyFormat defines the format for each css property.
const cssPropertyFormat = "\t%+s: %+s;\n"
const cssPropertyFormatMinified = "%+s: %+s;"

// RenderAttribute renders the attribute format for a css key: value pair.
func (r render) RenderAttribute(key string, val interface{}, dst io.Writer) {
	var format string

	if getMinitfy() {
		format = cssPropertyFormatMinified
	} else {
		format = cssPropertyFormat
	}

	switch val.(type) {
	case PropertyRender:
		dst.Write([]byte(fmt.Sprintf(format, key, val.(PropertyRender).Render())))
		return
	case int, int64:
		dst.Write([]byte(fmt.Sprintf(format, key, fmt.Sprintf("%d", val.(int)))))
		return
	case float32, float64:
		dst.Write([]byte(fmt.Sprintf(format, key, fmt.Sprintf("%.6f", val.(float64)))))
		return
	default:
		dst.Write([]byte(fmt.Sprintf(format, key, val)))
		return
	}
}

// cssPropertyBracket defines the bracket format for the css bracket.
const cssPropertyBracket = "%+s {\n%+s\n}\n"

// cssPropertyBracketMinify defines the bracket format for the css bracket.
const cssPropertyBracketMinify = "%+s {%+s}"

// Render calls the giving render function for the giving properties and selector
// into the provided destination writer.
func (r render) Render(sel string, p Properties, dst io.Writer) {
	var ws bytes.Buffer

	var format string
	var bracketformat string

	if getMinitfy() {
		format = cssPropertyFormatMinified
		bracketformat = cssPropertyBracketMinify
	} else {
		bracketformat = cssPropertyBracket
		format = cssPropertyFormat
	}

	for key, val := range p {
		switch val.(type) {
		case PropertyRender:
			ws.Write([]byte(fmt.Sprintf(format, key, val.(PropertyRender).Render())))
			continue
		default:
			ws.Write([]byte(fmt.Sprintf(format, key, val)))
			continue
		}
	}

	dst.Write([]byte(fmt.Sprintf(bracketformat, sel, ws.Bytes())))
}

//==============================================================================

// MediaRender defines an interface for a MediaQuery rendering structure.
type MediaRender interface {
	Render(Group, io.Writer)
}

// Media defines a structure which constructs the media query tree for
// gucss.
type Media struct {
	query  string
	device string
}

// NewMedia returns a new Media instance with the provided group.
func NewMedia(query string, device string) *Media {
	mn := Media{query: query, device: device}
	return &mn
}

// Render outputs the content of the media into the provided writer.
func (m *Media) Render(g Group, dst io.Writer) {
	var b bytes.Buffer
	g.Render(&b)

	var indent bool
	var format string

	if getMinitfy() {
		format = cssPropertyBracketMinify
	} else {
		format = cssPropertyBracket
		indent = true
	}

	if indent {
		var parts [][]byte

		for _, item := range bytes.Split(b.Bytes(), []byte("\n")) {
			parts = append(parts, append([]byte("   "), item...))
		}

		b.Reset()
		b.Write(bytes.Join(parts, []byte("\n")))
	}

	query := fmt.Sprintf("@media %s %s", m.device, m.query)
	qm := fmt.Sprintf(format, query, b.Bytes())

	dst.Write([]byte(qm))
}

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

// mediaGroup defines a struct to store a pair of Group and MediaRender
type mediaGroup struct {
	Root  Group
	Media MediaRender
}

// Render implements the Render interface for the baseGroup and returns the
// property title and content for this group.
func (mg mediaGroup) Render(dst io.Writer) {
	mg.Media.Render(mg.Root, dst)
}

//==============================================================================

// baseGroup defines the base line group structure which others compose. It contains
// the base structure needed to meet the Group interface.
type baseGroup struct {
	base *NGroup

	sel      string
	selector string

	root bool

	kids  []Render
	media []mediaGroup

	properties Properties
	prw        sync.RWMutex
}

// NewbaseGroup return a new instance of a baseGroup.
func newbaseGroup(sel string, selector string, p Properties, root Group) *baseGroup {
	bg := &baseGroup{
		sel:        sel,
		selector:   selector,
		properties: p,
	}

	if root != nil {
		bg.base = root.Tree().Append(bg)
	} else {
		bg.base = &NGroup{grp: bg}
	}

	return bg
}

// IsRoot returns true/false if the giving Group is a root.
func (bg *baseGroup) IsRoot() bool {
	return bg.root
}

// Tree returns the internal group chain tree.
func (bg *baseGroup) Tree() *NGroup {
	return bg.base
}

// Render implements the Render interface for the baseGroup and returns the
// property title and content for this group.
func (bg *baseGroup) Render(dst io.Writer) {
	if !bg.root {
		bg.prw.RLock()
		CSSRender.Render(bg.Selector(), bg.properties, dst)
		bg.prw.RUnlock()
	}

	// Wrap up the children tree and build the appropriate content
	for _, kid := range bg.kids {
		kid.Render(dst)
	}

	// Wrap up the media tree and build the appropriate content
	for _, media := range bg.media {
		media.Render(dst)
	}
}

// Add augments the giving properties into the baseGroup properties map.
func (bg *baseGroup) Add(p Properties) {
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

// Media adds a returns a new root Group which creates a media query section
// based on the properties added to the returned group.
func (bg *baseGroup) Media(device string, query string) Group {
	mg := mediaGroup{
		Root:  NewRoot(),
		Media: NewMedia(query, device),
	}

	bg.media = append(bg.media, mg)

	return mg.Root
}

// Extend returns a new Group based off the selector of the previous using
// the format:
//
//    rootSel parentSelector , rootSel sel {...}
func (bg *baseGroup) Extend(sel string, p Properties) Group {
	newSel := sel
	root := bg.NthParent(1)

	if !bg.root {
		newSel = bg.Selector() + ", " + root.Selector() + " " + sel
	}

	kid := newbaseGroup(sel, newSel, p, bg)
	bg.kids = append(bg.kids, kid)
	return bg
}

// PreSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector ~ childSelector {...}
func (bg *baseGroup) PreSibling(sel string, p Properties) Group {
	newSel := sel

	if !bg.root {
		newSel = bg.Selector() + " ~ " + sel
	}

	kid := newbaseGroup(sel, newSel, p, bg)
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
func (bg *baseGroup) NS(ns string, p Properties) Group {
	if bg.root {
		return bg
	}

	newSel := bg.selector + ns

	kid := newbaseGroup(bg.selector, newSel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// PostSibling returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector + childSelector {...}
func (bg *baseGroup) PostSibling(sel string, p Properties) Group {
	newSel := sel

	if !bg.root {
		newSel = bg.Selector() + " + " + sel
	}

	kid := newbaseGroup(sel, newSel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Within returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector  childSelector {...}
func (bg *baseGroup) Within(sel string, p Properties) Group {
	newSel := sel

	if !bg.root {
		newSel = bg.Selector() + " " + sel
	}

	kid := newbaseGroup(sel, newSel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Child returns a new Group based off the selector of the previous using
// the format:
//
//    parentSelector > childSelector {...}
func (bg *baseGroup) Child(sel string, p Properties) Group {
	newSel := sel

	if !bg.root {
		newSel = bg.Selector() + " > " + sel
	}

	kid := newbaseGroup(sel, newSel, p, bg)
	bg.kids = append(bg.kids, kid)
	return kid
}

// Sel returns the selector for the giving baseGroup.
func (bg *baseGroup) Sel() string {
	return bg.sel
}

// Selector returns the selector with the appropriate sign for the type
// for the giving baseGroup.
func (bg *baseGroup) Selector() string {
	return bg.selector
}

// Root returns the main root of the baseGroup parent chain.
func (bg *baseGroup) Root() Group {
	return bg.base.Root().grp
}

// NthParent returns the Group after the giving count if the count is
// greater than zero else returns main root of all parents chain.
func (bg *baseGroup) NthParent(back int) Group {
	var root *NGroup

	if back > 0 {
		root = bg.base.NthRoot(back)
	} else {
		root = bg.base.Root()
	}

	return root.grp
}

// addGroup adds the group into thhe BaseGroup.
func (bg *baseGroup) addGroup(g Group) {
	bg.kids = append(bg.kids, g)
}

//==============================================================================

// Child defines a group for a set of child selector with given properties.
// The format:
//
//    parentSelector > childSelector {...}
func Child(root Group, p Properties, sels ...string) {
	var oldSel string

	if !root.IsRoot() {
		oldSel = root.Selector() + " > "
	}

	var newSel []string

	for _, sel := range sels {
		newSel = append(newSel, oldSel+sel)
	}

	sel := strings.Join(newSel, ",")
	gm := newbaseGroup(sel, sel, p, root)

	if bm, ok := root.(*baseGroup); ok {
		bm.addGroup(gm)
	}
}

//==============================================================================

// NS defines a group for child selector properties.
// The format:
//
//    parentSelector{{NS}} {...}
func NS(root Group, p Properties, sels ...string) {
	var oldSel string

	if !root.IsRoot() {
		oldSel = root.Selector()
	}

	var newSel []string

	for _, sel := range sels {
		newSel = append(newSel, oldSel+sel)
	}

	sel := strings.Join(newSel, ",")
	gm := newbaseGroup(sel, sel, p, root)

	if bm, ok := root.(*baseGroup); ok {
		bm.addGroup(gm)
	}

}

//==============================================================================

// Extend defines a group for child selector properties.
// The format:
//
//    parentSelector , childSelector {...}
func Extend(root Group, p Properties, sels ...string) {
	var oldSel string

	if !root.IsRoot() {
		oldSel = root.NthParent(1).Selector()
	}

	var newSel []string

	for _, sel := range sels {
		if oldSel != "" {
			sel = oldSel + " " + sel
		}

		if strings.TrimSpace(sel) == "" {
			continue
		}

		newSel = append(newSel, sel)
	}

	sel := strings.Join(newSel, ", ")

	if !root.IsRoot() {
		sel = root.Selector() + ", " + sel
	}

	gm := newbaseGroup(sel, sel, p, root)

	if bm, ok := root.(*baseGroup); ok {
		bm.addGroup(gm)
	}
}

//==============================================================================

// PreSibling defines the grouping of sibling properties.
// The format:
//
//    Pre: parentSelector ~ childSelector {...}
func PreSibling(root Group, p Properties, sels ...string) {
	var oldSel string

	if !root.IsRoot() {
		oldSel = root.Selector() + " ~ "
	}

	var newSel []string

	for _, sel := range sels {
		newSel = append(newSel, oldSel+sel)
	}

	sel := strings.Join(newSel, ",")
	gm := newbaseGroup(sel, sel, p, root)

	if bm, ok := root.(*baseGroup); ok {
		bm.addGroup(gm)
	}
}

// PostSibling defines the grouping of sibling properties.
// The format:
//
//    Post: parentSelector + childSelector {...}
func PostSibling(root Group, p Properties, sels ...string) {
	var oldSel string

	if !root.IsRoot() {
		oldSel = root.Selector() + " + "
	}

	var newSel []string

	for _, sel := range sels {
		newSel = append(newSel, oldSel+sel)
	}

	sel := strings.Join(newSel, ",")
	gm := newbaseGroup(sel, sel, p, root)

	if bm, ok := root.(*baseGroup); ok {
		bm.addGroup(gm)
	}
}

//==============================================================================

// Within defines a group for within properties.
// The format:
//
//    parentSelector  childSelector {...}
func Within(root Group, p Properties, sels ...string) {
	var oldSel string

	if !root.IsRoot() {
		oldSel = root.Selector() + "  "
	}

	var newSel []string

	for _, sel := range sels {
		newSel = append(newSel, oldSel+sel)
	}

	sel := strings.Join(newSel, ",")
	gm := newbaseGroup(sel, sel, p, root)

	if bm, ok := root.(*baseGroup); ok {
		bm.addGroup(gm)
	}
}

//==============================================================================
