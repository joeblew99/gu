package design

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
