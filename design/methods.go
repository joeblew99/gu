package design

import (
	"github.com/influx6/gu"
	"github.com/influx6/gu/gutrees"
)

var rootResource = struct {
	res []ResourceDefinition
}{}

func currentResource() *ResourceDefinition {
	if len(rootResource.res) == 0 {
		panic("Require resource to exists first before call")
	}

	return &(rootResource.res[len(rootResource.res)-1])
}

// ResourceDefinition defines a high-level definition for managing resources for
// which other definitions build from.
type ResourceDefinition struct {
	dsl DSL

	grp *gu.RenderGroup

	bodyMarkup   gutrees.Markup
	headerMarkup gutrees.Markup

	views   []ViewDef
	markups []MarkupDef
	links   []MarkupDef
	routes  []RouteDef
}

// Resource creates a new resource addding into the resource lists for the root.
func Resource(dsl DSL) *ResourceDefinition {
	var rs ResourceDefinition
	rs.dsl = dsl

	rootResource.res = append(rootResource.res, rs)
	return &rs
}

// MarkupDef defines the structure which stores markup definitions for the
// generating markups.
type MarkupDef struct {
	deffer   bool
	targets  []string
	children []gutrees.Markup
}

// Context returns the context string for this giving structure.
func (MarkupDef) Context() string {
	return "Markup"
}

// Markup returns a new instance of a provided value which either is a function
// which returns a needed gutrees.Markup or a gutrees.Markup or slice of gutrees.Markup
// itself.
func Markup(markup gu.Viewable, target interface{}, defered bool) {
	var markupFn []gutrees.Markup

	switch mo := markup.(type) {
	case func() []gutrees.Markup:
		markupFn = mo()
	case []gutrees.Markup:
		markupFn = mo
	case func() gutrees.Markup:
		markupFn = []gutrees.Markup{mo()}
	case gutrees.Markup:
		markupFn = []gutrees.Markup{mo}
	case string:
		mp, err := gutrees.ParseTree(mo)
		if err != nil {
			panic("Unable to parse markup string: " + err.Error())
		}

		markupFn = mp
	default:
		panic("Unknown markup processable type")
	}

	var targets []string

	switch to := target.(type) {
	case string:
		targets = []string{to}
	case []string:
		targets = to
	default:
		panic("targets only allowed to be a string or a slice of string")
	}

	mp := MarkupDef{
		deffer:   defered,
		targets:  targets,
		children: markupFn,
	}

	current := currentResource()
	current.markups = append(current.markups, mp)
}

//==============================================================================

// ViewDef defines the structure which stores view creation data for the
// generating markups.
type ViewDef struct {
	deffer bool
	target string
	rs     gu.Renderables
	rv     gu.RenderView
}

// Context returns the context string for this giving structure.
func (ViewDef) Context() string {
	return "View"
}

func View(vrs gu.Viewable, target string, deffer bool) gu.RenderView {
	var vws ViewDef
	vws.target = target
	vws.deffer = deffer

	var rs gu.Renderables

	switch vwo := vrs.(type) {
	case gu.Renderables:
		rs = vwo
	case gu.Renderable:
		rs = gu.Renderables{vwo}
	case func() gu.Renderable:
		rs = gu.Renderables{vwo()}
	case func() gu.Renderables:
		rs = vwo()
	default:
		panic("View must either recieve a function that returns Renderables or Renderables themselves")
	}

	current := currentResource()
	vws.rv = current.grp.View(rs...)
	current.views = append(current.views, mp)

	return vws.rv
}

//==============================================================================

// RouteDef defines a structure for connecting a giving route with a provided
// action.
type RouteDef struct {
	deffer bool
	tag    string
	route  string
}

// Context returns the context string for this giving structure.
func (RouteDef) Context() string {
	return "Route"
}

//==============================================================================

// LinkDef defines the structure which stores link creation data for the
// generating markups.
type LinkDef struct {
	deffer bool
	tag    string
	elem   gutrees.Markup
}

// Context returns the context string for this giving structure.
func (LinkDef) Context() string {
	return "Link"
}

// Title adds a element which generates a <title> tag.
func Title(title string) {
	ml := mLink("title", false)
	gutrees.NewText(title).Apply(ml.elem)
}

// Link adds a element which generates a <link> tag.
func Link(url string, mtype string, defered bool) {
	ml := mLink("link", defered)
	gutrees.NewAttr("href", url).Apply(ml.elem)
	gutrees.NewAttr("type", mtype).Apply(ml.elem)
}

// CSS adds a element which generates a <style> tag.
func CSS(src string, defered bool) {
	ml := mLink("link", defered)
	gutrees.NewAttr("href", src).Apply(ml.elem)
	gutrees.NewAttr("rel", "stylesheet").Apply(ml.elem)
	gutrees.NewAttr("type", "text/css").Apply(ml.elem)
}

// Style adds a element which generates a <style> tag.
func Style(src string, defered bool) {
	ml := mLink("style", defered)
	gutrees.NewAttr("src", src).Apply(ml.elem)
	gutrees.NewAttr("type", "text/css").Apply(ml.elem)
}

// Script adds a element which generates a <style> tag.
func Script(src string, mtype string, defered bool) {
	ml := mLink("script", defered)
	gutrees.NewAttr("src", src).Apply(ml.elem)
	gutrees.NewAttr("type", mtype).Apply(ml.elem)
}

// Meta adds a element which generates a <style> tag.
func Meta(props map[string]string) {
	ml := mLink("meta", false)
	for name, val := range props {
		gutrees.NewAttr(name, val).Apply(ml.elem)
	}
}

// mLink adds tagName with the provided value into the header bar for the
// page content.
func mLink(tag string, deffer bool) *LinkDef {
	var ld LinkDef
	ld.tag = "link"
	ld.deffer = deffer
	ml.elem = gutrees.NewElement(tag, false)

	current := currentResource()
	current.links = append(current.links, ld)
	return &ld
}
