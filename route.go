package gu

// LevelRouting defines an interface which returns a trees.Applier which adds
// into a trees.Markup structure. This actually return a routing instance which 
// adds itself into the Markup morphers lists and morphs the state of the markup
// tree depending on its current internal state.
type LevelRouting interface{
  dispatch.Resolver
  N(string) trees.Applier
  Next(string, trees.Morpher) trees.Applier
}

// RouteManager defines a router which handles creation and registering of a 
// set of level connecting routes
type RouteManager struct {
  levels map[string]LevelRouting
}

// Level creates a new route level using the provided string as the base path 
// for its root router.
func (r *rm) Level(path string, morpher trees.SwitchMorpher) LevelRouting {
  if rx, ok := r.levels[path]; ok {
    return rx
  }

  var rs rm
  rs.path = path
  rs.Resolver = rs.next = newRouting(path, morpher) 
  rs.linked = make(map[string]*Router)
  
  r.levels[path] = &rs
  return &rs
}

//==============================================================================

type rm {
  dispatch.Resolver

  rw sync.RWMutex
  path string
  linked map[string]*Routing    
  next *Routing
}

// N returns a new trees.Applier based on the internal routing struct. Which stores 
// its modifier 
func (r *rm) N(path string) trees.Applier {
   return r.Next(path, &trees.RemoveMorpher{})
}

// Next creates a route which depends on the previous route created. It 
// allows multi-level routing where the next depends on the outcome of the 
// previous. It uses dispatch.Resolver which shifts the matching paths by passing 
// down the remants of path unmatched to its next subscribers.
func (r *rm) Next(path string, morpher trees.SwitchMorpher) *Routing {
  if rx, ok := r.linked[path]; ok {
    return rx
  }

  prv := r.next
  r.next := newRouting(path, morpher)
  prv.Register(r.next)
  
  r.linked[path] = r.next

  return rsm
}


//==============================================================================

// routing defines a dispatch.Resolve structure that allows morphing the output
// of a markup based on a giving route.
type routing struct {
  dispatch.Resolver
  m trees.Morpher
}

// Routing returns a new instance of a Routing struct.
func newRouting(path string, morpher trees.SwitchMorpher) *Routing {
  var rs routing
  rs.m = morpher
  rs.Resolver = dispatch.NewResolver(path)
  rs.Resolver.ResolvedPassed(func(p dispatch.Path) { morpher.Off(p) })
  rs.Resolver.ResolvedFailed(func(p dispatch.Path) { morpher.On(p) })
  return &rs
}

// Morph implements the Morpher interface providing the routing with the ability.
// It lets routing morph markup passed into it.
func (r *routing) Morph(mr trees.Markup) trees.Markup {
  return r.m.Morph(mr)
}

// Apply adds this routing as a morpher into the provided markup.
func (r *routing) Apply(mr trees.Markup) {
  mr.AddMorpher(r)
}
