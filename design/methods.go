package design

import (
  "github.com/influx6/gu"
  "github.com/influx6/gu/gutrees"
)

var (
   rootResource struct{
      res []ResourceDefinition
    }{}
)

func currentResource() *ResourceDefinition {
  if len(rootResource.res) == 0 {
    panic("Require resource to exists first before call")
  }

  return rootResource.res[len(rootResource.res) - 1]
}

// Resource creates a new resource addding into the resource lists for the root.
func Resource(name string,dsl DSL) *ResourceDefinition {
  var rs ResourceDefinition
  rs.def  = &Definition{
    Name: name,
    Dsl: dsl,
  }

  rootResources.res = append(rootResources.res, rs)
  reurn &rs
}

// MarkupDef defines the structure which stores markup definitions for the 
// generating markups.
type MarkupDef struct{
  deffer bool
  targets []string
  children  []gutrees.Markup
}

// Context returns the context string for this giving structure.
func (MarkupDef) Context() string{
  return "Markup"
}

// Markup returns a new instance of a provided value which either is a function 
// which returns a needed gutrees.Markup or a gutrees.Markup or slice of gutrees.Markup
// itself.
func Markup(markup gu.Viewable, target interface{}, defered bool)  {
  var markupFn []gutrees.Markup 

  switch mo := markup.(type){
  case func() []gutrees.Markup:
    markupFn = mo()
  case []gutrees.Markup:
    markupFn = mo
  case func() gutrees.Markup:
    markupFn = []gutrees.Markup {mo()}
  case gutrees.Markup:
    markupFn = []gutrees.Markup{mo}
  case string:
    mp, err := gutrees.ParseTree(markup)
    if err != nil {
      panic("Unable to parse markup string: "+ err.Error())
    }

    markupFn = mp
  default:
    panic("Unknown markup processable type")
  }

  var targets []string

  switch to := target.(type){
  case string:
    targets = []string{to}
  case []string:
    targets = to
  default:
    panic("targets only allowed to be a string or a slice of string")
  }

  mp := MarkupDef{
    deffer: defered,
    targets: targets,
    children: markupFn,
  }

  current := currentResource()
  current.markups = append(current.markups, mp)
}

// ViewDef defines the structure which stores view creation data for the 
// generating markups.
type ViewDef struct{
  deffer bool
  target string
  rs gu.Renderables
  rv gu.RenderView
}

// Context returns the context string for this giving structure.
func (ViewDef) Context() string{
  return "Markup"
}

func View(vrs gu.Viewable, target string, deffer bool) gu.RenderView {
  var vws ViewDef
  vws.target = target
  vws.deffer = deffer


   var rs gu.Renderables

   switch vwo := vrs.(type){
   case gu.Renderables:
    rs = vwo
   case gu.Renderable:
    rs = gu.Renderables{vwo}
  case func() Renderable:
    rs = gu.Renderables{vwo()}
  case func() Renderables:
    rs = vwo()
  default:
    panic("View must either recieve a function that returns Renderables or Renderables themselves")
   }


  current := currentResource()
  vws.rv = current.grp.View(rs...)
  current.views = append(current.views, mp)
  
  return vws.rv
}