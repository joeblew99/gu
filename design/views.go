package design

// ViewDefinition represents the creation of gu views using the contents of this definitions dsl.
type ViewDefinition struct {
	*Definition
	defered bool
	markups  []MarkupDefinition
}

// NewView returns an new instance of a ViewDefinition type.
func View(viewName string, deffer bool, dsl DSL) *ViewDefinition {
	vd := &ViewDefinition{
		Definition: NewDefinition(viewName, dsl, deffer),
		deffered:   deffer,
	}

	return vd
}
