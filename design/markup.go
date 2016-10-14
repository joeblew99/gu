package design

// MarkupDefinition defines a type for markup types which generates the markup for elements.
type MarkupDefinition struct {
	*Definition
	target []string
         markup []gutrees.Markup
}

// NewMarkup returns a new instance of a giving markup definition
func NewMarkup(markup interface{}, target ...string) *MarkupDefinition {
	var md MarkupDefinition
        md.target = target

	var dsl DSL


	switch rm := markup.(type) {
          case []gutrees.Markup:
            dsl = func(){
                (&md).markup  = rm 
            }

          case func() gutrees.Markup:
            dsl = func(){
                (&md).markup  = []gutrees{rm()}
            }

          case gutrees.Markup:
            dsl = func(){
                (&md).markup  = []gutrees{rm}
            }

    	case string: 
              dsl = func(){
                tree, err := gutrees.ParseTree(rm)
                if err != nil {
                  panic(tree)
                }

                (&md).markup  = tree
              }

           default:
              panic("Unsupported type unallowed")
	}

        md.Definition = NewDefinition("markup", false, dsl)
	return &md
}

