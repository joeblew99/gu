package design

// Definition  defines a base level definition structure which implements the Def interface.
type Definition struct {
	dependents []Root
	Name       string
	dsl        DSL
	deffered   bool
}

// NewDefinition returns a new instance of the Definition type.
func NewDefinition(definitionName string, defered bool, dsl DSL) *Definition {
	var df Definition
	df.Name = definitionName
	df.dsl = dsl
	df.deffered = defered

	return &df
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
	RootViews   []ViewDefinition
	RootMarkups []MarkupDefinition

	DRootViews   []ViewDefinition
	DRootMarkups []MarkupDefinition
}

// CurrentView returns the current ViewDefinition being used in the processing of
// the current view.
func (rd *ResourceDefinition) CurrentView() *ViewDefinition {
	if len(rd.RootViews) == 0 {
		return nil
	}

	return &(rd.RootViews[len(rd.RootViews)-1])
}

func (rd *ResourceDefinition) CurrentMarkup() *MarkupDefinition {
	if len(rd.RootViews) == 0 {
		return nil
	}

	return &(rd.RootMarkups[len(rd.RootMarkups)-1])
}
