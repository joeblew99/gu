package design

// Definition  defines a base level definition structure which implements the Def interface.
type Definition struct {
	dependents []Root
	Name       string
	Dsl        DSL
	deffered   bool
}

// Defered returns true/false if the giving definition is supposed to be
// executed later at the end of previous definition.
func (d *Definition) Defered() bool {
	return d.deffered
}

// DSL returns the given function which generates the provided contents for this definition.
func (d *Definition) DSL() DSL {
	return d.dsl
}

func (d *Definition) DependsOn() []Root {
	return d.dependents
}

// Context returns the name representing this given definition.
func (d *Definition) Context() string {
	return d.Name
}

//==============================================================================

// ResourceDefinition defines a high-level definition for managing resources for
// which other definitions build from.
type ResourceDefinition struct {
	 def *Definition
	 grp *gu.RenderGrou

	 bodyMarkup gutrees.Markup
	 headerMarkup gutrees.Markup

	 views []ViewDef
	 markups []MarkupDef
}

