package trees

import "strings"

//==============================================================================

// Property defines the interface for attributes in trees.
// It provides a apply and RenderAttribute which returns the key
// and value for that attribute.
type Property interface {
	Apply(*Markup)
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

// Attribute define the struct  for attributes
type Attribute struct {
	Name  string
	Value string
	After func(*Markup)
}

// NewAttr returns a new attribute instance.
func NewAttr(name, val string) *Attribute {
	a := Attribute{Name: strings.ToLower(name), Value: val}
	return &a
}

// NewAttrWith returns a new attribute instance with a provided function
// to call to provide a after effect to the markup.
func NewAttrWith(name, val string, after func(*Markup)) *Attribute {
	a := Attribute{Name: strings.ToLower(name), Value: val, After: after}
	return &a
}

// Render returns the key and value for this attribute rendered.
func (a *Attribute) Render() (string, string) {
	return a.Name, a.Value
}

// Apply applies a set change to the giving element attributes list
func (a *Attribute) Apply(e *Markup) {
	if e.allowAttributes {
		e.AddAttribute(a)

		if a.After != nil {
			a.After(e)
		}
	}
}

//Clone replicates the attribute into a unique instance
func (a *Attribute) Clone() Property {
	return &Attribute{Name: a.Name, Value: a.Value, After: a.After}
}

//==============================================================================

// CSSStyle define the style specification for element styles
type CSSStyle struct {
	Name  string
	Value string
}

// NewCSSStyle returns a new style instance
func NewCSSStyle(name, val string) *CSSStyle {
	s := CSSStyle{Name: strings.ToLower(name), Value: val}
	return &s
}

// Render returns the key and value for this style rendered.
func (s *CSSStyle) Render() (string, string) {
	return s.Name, s.Value
}

//Clone replicates the style into a unique instance
func (s *CSSStyle) Clone() Property {
	return &CSSStyle{Name: s.Name, Value: s.Value}
}

// Apply applies a set change to the giving element style list
func (s *CSSStyle) Apply(e *Markup) {
	if e.allowStyles {
		e.AddStyle(s)
	}
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
		name: strings.ToLower(name),
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
func (c *ClassList) Apply(em *Markup) {
	if em.allowStyles {
		em.AddStyle(c)
	}
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
