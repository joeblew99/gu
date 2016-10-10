package design


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

