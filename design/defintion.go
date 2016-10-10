package design

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

