package gutrees

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// RandString generates a set of random numbers of a set length
func RandString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// Augment adds new markup to an the root if its Element
func Augment(root Markup, m ...Markup) {
	if el, ok := root.(*Element); ok {
		for _, mo := range m {
			mo.Apply(el)
		}
	}
}

// Events defines an interface that returns a list of events.
type Events interface {
	Events() []Event
}

// ReconcileEvents checks through two markup events against each other and if it finds any disparity marks
// event objects as Removed
func ReconcileEvents(e, em Events) {
	oldevents := em.Events()
	newevents := e.Events()

	if len(oldevents) <= 0 && len(newevents) <= 0 {
		return
	}

	if len(newevents) <= 0 && len(oldevents) > 0 {
		for _, ev := range oldevents {
			ev.Meta().Remove()
		}
		return
	}

	checkOut := func(ev Event) bool {
		for _, evs := range newevents {
			if evs.Meta().Type() == ev.Meta().Type() {
				return true
			}
		}
		return false

	}
	//set to equal as the logic will try to assert its falsiness

	// outerfind:
	for _, ev := range oldevents {
		if checkOut(ev) {
			continue
		}

		ev.Meta().Remove()
	}
}

// Styles defines an interface that returns a list of styles.
type Styles interface {
	Styles() []Property
}

// EqualStyles returns true/false if the style values are all equal attribute.
func EqualStyles(e, em Styles) bool {
	old := em.Styles()

	if len(old) < 1 {
		if len(e.Styles()) > 0 {
			return false
		}
		return true
	}

	for _, oa := range old {
		name, val := oa.Render()

		ta, err := GetStyle(e, name)
		if err != nil {
			return false
		}

		_, tvalue := ta.Render()
		if tvalue != val {
			return false
		}
	}

	return true
}

// Attributes defines an interface that returns a list of attributes.
type Attributes interface {
	Attributes() []Property
}

// EqualAttributes returns true/false if the elements and the giving markup have equal attribute
func EqualAttributes(e, em Attributes) bool {
	old := em.Attributes()

	if len(old) < 1 {
		if len(e.Attributes()) > 0 {
			return false
		}
		return true
	}

	for _, oa := range old {
		name, val := oa.Render()

		ta, err := GetAttr(e, name)
		if err != nil {
			return false
		}

		_, tvalue := ta.Render()
		if tvalue != val {
			return false
		}
	}

	return true
}

// GetStyles returns the styles that contain the specified name and if not empty that contains the specified value also, note that strings
// NOTE: string.Contains is used when checking value parameter if present
func GetStyles(e Styles, f, val string) []Property {
	var found []Property
	var styles = e.Styles()

	for _, as := range styles {
		name, value := as.Render()
		if name != f {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		found = append(found, as)
	}

	return found
}

// GetStyle returns the style with the specified tag name
func GetStyle(e Styles, f string) (Property, error) {
	styles := e.Styles()
	for _, as := range styles {
		name, _ := as.Render()
		if name == f {
			return as, nil
		}
	}
	return nil, ErrNotFound
}

// StyleContains returns the styles that contain the specified name and if the val is not empty then
// that contains the specified value also, note that strings
// NOTE: string.Contains is used
func StyleContains(e Styles, f, val string) bool {
	styles := e.Styles()
	for _, as := range styles {
		name, value := as.Render()
		if !strings.Contains(name, f) {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		return true
	}

	return false
}

// GetAttrs returns the attributes that have the specified text within the naming
// convention and if it also contains the set val if not an empty "",
// NOTE: string.Contains is used
func GetAttrs(e Attributes, f, val string) []Property {
	var found []Property

	for _, as := range e.Attributes() {
		name, value := as.Render()
		if name != f {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		found = append(found, as)
	}

	return found
}

// AttrContains returns the attributes that have the specified text within the naming
// convention and if it also contains the set val if not an empty "",
// NOTE: string.Contains is used
func AttrContains(e Attributes, f, val string) bool {
	for _, as := range e.Attributes() {
		name, value := as.Render()
		if !strings.Contains(name, f) {
			continue
		}

		if val != "" {
			if !strings.Contains(value, val) {
				continue
			}
		}

		return true
	}

	return false
}

// GetAttr returns the attribute with the specified tag name
func GetAttr(e Attributes, f string) (Property, error) {
	for _, as := range e.Attributes() {
		name, _ := as.Render()
		if name == f {
			return as, nil
		}
	}
	return nil, ErrNotFound
}

// ReplaceStyle replaces a specific style with the given
// name with the supplied value.
func ReplaceStyle(m Styles, name string, val string) {
	styl, err := GetStyle(m, name)
	if err != nil {
		return
	}

	stylm, ok := styl.(*Style)
	if !ok {
		return
	}

	stylm.Value = val
}

// ReplaceAttribute replaces a specific attribute with the given
// name with the supplied value.
func ReplaceAttribute(m Attributes, name string, val string) {
	attr, err := GetAttr(m, name)
	if err != nil {
		return
	}

	attrm, ok := attr.(*Attribute)
	if !ok {
		return
	}

	attrm.Value = val
}

// ReplaceORAddStyle replaces a specific style with the given
// name with the supplied value if not found it adds a new one
// if found and if the type does not match a *Style then it stops.
func ReplaceORAddStyle(m Properties, name string, val string) {
	styl, err := GetStyle(m, name)
	if err != nil {
		m.AddStyle(NewStyle(name, val))
		return
	}

	stylm, ok := styl.(*Style)
	if !ok {
		return
	}

	stylm.Value = val
}

// ReplaceORAddAttribute replaces a specific attribute with the given
// name with the supplied value if not found it adds a new one
// if found and if the type does not match a *Style then it stops.
func ReplaceORAddAttribute(m Properties, name string, val string) {
	attr, err := GetAttr(m, name)
	if err != nil {
		m.AddAttribute(NewAttr(name, val))
		return
	}

	attrm, ok := attr.(*Attribute)
	if !ok {
		return
	}

	attrm.Value = val
}

//==============================================================================

// MarkupProps defines a custom type that combines the Markup, Styles and
// Attributes interfaces.
type MarkupProps interface {
	Properties
	Children
	Appliable
}

// ElementsUsingStyle returns the children within the element matching the
// stlye restrictions passed.
// NOTE: is uses StyleContains
func ElementsUsingStyle(e MarkupProps, f, val string) []Markup {
	return DeepElementsUsingStyle(e, f, val, 1)
}

// ElementsWithAttr returns the children within the element matching the
// stlye restrictions passed.
// NOTE: is uses AttrContains
func ElementsWithAttr(e MarkupProps, f, val string) []Markup {
	return DeepElementsWithAttr(e, f, val, 1)
}

// DeepElementsUsingStyle returns the children within the element matching the
// style restrictions passed allowing control of search depth
// NOTE: is uses StyleContains
// WARNING: depth must start at 1
func DeepElementsUsingStyle(e MarkupProps, f, val string, depth int) []Markup {
	if depth <= 0 {
		return nil
	}

	var found []Markup

	for _, c := range e.Children() {
		if ch, ok := c.(Markup); ok {
			if StyleContains(ch, f, val) {
				found = append(found, ch)
				cfo := DeepElementsUsingStyle(ch, f, val, depth-1)
				if len(cfo) > 0 {
					found = append(found, cfo...)
				}
			}
		}
	}

	return found
}

// DeepElementsWithAttr returns the children within the element matching the
// attributes restrictions passed allowing control of search depth
// NOTE: is uses Element.AttrContains
// WARNING: depth must start at 1
func DeepElementsWithAttr(e MarkupProps, f, val string, depth int) []Markup {
	if depth <= 0 {
		return nil
	}

	var found []Markup

	for _, c := range e.Children() {
		if ch, ok := c.(Markup); ok {
			if AttrContains(ch, f, val) {
				found = append(found, ch)
				cfo := DeepElementsWithAttr(ch, f, val, depth-1)
				if len(cfo) > 0 {
					found = append(found, cfo...)
				}
			}
		}
	}

	return found
}

// ElementsWithTag returns elements matching the tag type in the parent markup children list
// only without going deeper into children's children lists
func ElementsWithTag(e MarkupProps, f string) []Markup {
	return DeepElementsWithTag(e, f, -1)
}

// DeepElementsWithTag returns elements matching the tag type in the parent markup
// and depending on the depth will walk down other children within the children.
// WARNING: depth must start at 1
func DeepElementsWithTag(e MarkupProps, f string, depth int) []Markup {
	f = strings.TrimSpace(strings.ToLower(f))

	var dp bool
	var found []Markup

	if depth > 0 {
		dp = true
	}

	deepElementsWithTag(e, f, depth, dp, &(found))

	return found
}

func deepElementsWithTag(e MarkupProps, f string, depth int, doDepth bool, res *[]Markup) {
	if doDepth && depth < 0 {
		return
	}

	for _, ch := range e.Children() {
		fmt.Printf("Checking item: %s \n", ch.Name())

		if ch.Name() == f {
			*res = append(*res, ch)

			if doDepth {
				depth = depth - 1
			}

			deepElementsWithTag(ch, f, depth-1, doDepth, res)
		}
	}
}

//==============================================================================
