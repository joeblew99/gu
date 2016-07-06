# Views And Routes
Although the internal routing system for `gu` is built to be an optional addition just as the use of the Reactive interface. It does provide an nice and elegant solution for routing.

Taking inspiration from the idea of delayed routing and insulated routes that are not flat but nest-able regardless of their parent (i.e where a route can be attached to another high level route without loosing functionality), still the approach still leaves the developer to either use or adopt another suitable means.

*Previous versions of the routing system actually were opinionated and had routes flip the view state of views, but the new system allows more control on what should happen when a route actually changes*

## Resolvers
Taking inspiration from Polymer and other approaches, `Resolvers` exists to provide a custom registering system of nest-able routes, where a Component/Renderable can embedded a resolver within itself, then have its view or parent view register to a root route. This way anytime, the root route matches, the resolver gets called to match against the remaining of the path after the view matches its criteria.

This approach will allow any component to be moved or initialized into another without needing to change routing details.

*This routing system is as much an experimental one and will evolve as time goes by*
