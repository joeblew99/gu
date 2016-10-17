package trees

import "strings"

//==============================================================================

// Property defines the interface for attributes in trees.
// It provides a apply and RenderAttribute which returns the key
// and value for that attribute.
type Property interface {
	Apply(Markup)
	Clone() Property
	Render() (string, string)
}

// WhenElse checks if the first is a non-nil and returns else returns the second
// property instead.
func WhenElse(first Property, other Property) Property {
	if first != nil {
		return first
	}

	return other
}

// When checks if the giving state is true and returns the first property else
// returns the second.
func When(state bool, first Property, other Property) Property {
	if state {
		return first
	}

	return other
}

//==============================================================================

// PropertyApplier defines a package level property applier for markup properties.
var PropertyApplier properties

type properties struct{}

// MarkupPropertiesProvider defines a interface for a markup to register
// properties for themselves
type MarkupPropertiesProvider interface {
	MarkupState
	AddAttribute(Property)
	AddStyle(Property)
}

// ApplyAttribute adds the property into the Markup attribute lists
func (properties) ApplyAttribute(ex Markup, p Property) {
	if p == nil {
		return
	}

	if em, ok := ex.(MarkupPropertiesProvider); ok {
		if em.AllowAttributes() {
			em.AddAttribute(p)
		}
	}
}

// ApplyStyle adds the property into the Markup style lists
func (properties) ApplyStyle(ex Markup, p Property) {
	if p == nil {
		return
	}

	if em, ok := ex.(MarkupPropertiesProvider); ok {
		if em.AllowStyles() {
			em.AddStyle(p)
		}
	}
}

//==============================================================================

// Attribute define the struct  for attributes
type Attribute struct {
	Name  string
	Value string
}

// NewAttr returns a new attribute instance
func NewAttr(name, val string) *Attribute {
	a := Attribute{Name: name, Value: val}
	return &a
}

// Render returns the key and value for this attribute rendered.
func (a *Attribute) Render() (string, string) {
	return a.Name, a.Value
}

// Apply applies a set change to the giving element attributes list
func (a *Attribute) Apply(e Markup) {
	PropertyApplier.ApplyAttribute(e, a)
}

//Clone replicates the attribute into a unique instance
func (a *Attribute) Clone() Property {
	return &Attribute{Name: a.Name, Value: a.Value}
}

//==============================================================================

// Style define the style specification for element styles
type Style struct {
	Name  string
	Value string
}

// NewStyle returns a new style instance
func NewStyle(name, val string) *Style {
	s := Style{Name: name, Value: val}
	return &s
}

// Render returns the key and value for this style rendered.
func (s *Style) Render() (string, string) {
	return s.Name, s.Value
}

//Clone replicates the style into a unique instance
func (s *Style) Clone() Property {
	return &Style{Name: s.Name, Value: s.Value}
}

// Apply applies a set change to the giving element style list
func (s *Style) Apply(e Markup) {
	PropertyApplier.ApplyStyle(e, s)
}

//==============================================================================

// ClassList defines the list type for class lists.
type ClassList struct {
	list []string
	name string
}

// NewClassList returns a new ClassList instance.
func NewClassList(name string) *ClassList {
	cl := ClassList{
		name: name,
	}

	return &cl
}

// Add adds a class name into the lists.
func (c *ClassList) Add(class string) {
	c.list = append(c.list, class)
}

// Render returns the key and value for this style rendered.
func (c *ClassList) Render() (string, string) {
	return c.name, strings.Join(c.list, " ")
}

// Apply checks for a class attribute
func (c *ClassList) Apply(em Markup) {
	PropertyApplier.ApplyAttribute(em, c)
}

// Clone replicates the lists of classnames.
func (c ClassList) Clone() Property {
	return &ClassList{
		name: c.name,
		list: c.list[:len(c.list)],
	}
}

// Reconcile checks each item against the given lists
// and replaces/add any missing item.
func (c *ClassList) Reconcile(m Property) bool {
	var added bool

	if mc, ok := m.(*ClassList); ok {
		maxlen := len(c.list)

		for ind, val := range mc.list {

			if ind >= maxlen {
				added = true
				c.list = append(c.list, val)
				continue
			}

			if (c.list[ind]) == val {
				continue
			}

			added = true
			c.list[ind] = val
		}
	}

	return added
}
