package shell

import (
	"errors"

	"github.com/gu-io/gu/trees"
)

// Hook defines an interface which handles the retrieval and installation of
// a ManifestAttr. It expects to return two values, the markup to be installed
// into the page and a boolean indicating if it should be added into the head.
type Hook interface {
	Fetch(Fetch, ManifestAttr) (res *trees.Markup, addToHeader bool, err error)
}

// hooks provides a global registry for registering hooks.
var hooks = newRegistry()

// Register adds the Hook into the registry, passing the giving struct with the
// provider for usage.
func Register(name string, hook Hook) {
	hooks.Create(name, hook)
}

// Get retrieves the giving provider by the name.
func Get(name string) (Hook, error) {
	return hooks.Get(name)
}

// registry defines a structure for registery for Providers where all providers
// handles the lifecycle process for a asset from installation to removal.
type registry struct {
	hooks map[string]Hook
}

// newRegistry returns a new instance of Registry.
func newRegistry() *registry {
	return &registry{
		hooks: make(map[string]Hook),
	}
}

// Create returns a provider associated with the given key if found and if not, then
// it is created and added into the hook lists.
func (r *registry) Create(name string, hook Hook) {
	r.hooks[name] = hook
}

// ErrHookNotFound is returned when the desired hook is not on the registry.
var ErrHookNotFound = errors.New("Hook not found")

// Get returns a giving provider with the provided key.
func (r *registry) Get(name string) (Hook, error) {
	hl, ok := r.hooks[name]
	if !ok {
		return nil, ErrHookNotFound
	}

	return hl, nil
}
