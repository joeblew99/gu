package apidesign

// Definition  defines a base level definition structure which implements the Def interface.
type Definition struct {
	Dependents []Root
	Name       string
	dsl        DSL
}

// NewDefinition returns a new instance of the Definition type.
func NewDefinition(definitionName string, dsl DSL) *Definition {
	var df Definition
	df.Name = definitionName
	df.dsl = dsl

	return &df
}

// DSL returns the given function which generates the provided contents for this definition.
func (d *Definition) DSL() DSL {
	return d.dsl
}

func (d *Definition) DependsOn() Root {
	return d.Dependents
}

// Context returns the name representing this given definition.
func (d *Definition) Context() string {
	return d.Name
}

// MarkupDefinition defines a type for markup types which generates the markup for elements.
type MarkupDefinition struct {
	*Definition
	target string
}

// NewMarkup returns a new instance of a giving markup definition
func NewMarkup(markup interface{}, target ...string) MarkupDefinition {
	var md MarkupDefinition

	var dsl DSL

	switch realMarkup := markup.(type) {
	case string:

	}

	return &md
}

// ViewDefinition represents the creation of gu views using the contents of this definitions dsl.
type ViewDefinition struct {
	*Definition
	markups []MarkupDefinition
}

// NewView returns an new instance of a ViewDefinition type.
func View(viewName string, dsl DSL) ViewDefinition {
	vd := &ViewDefinition{
		Definition: NewDefinition(viewName, dsl),
	}

	return vd
}
