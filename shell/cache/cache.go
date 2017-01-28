package cache

import (
	"github.com/gu-io/gu/shell"
	"github.com/gu-io/gu/shell/cache/localcache"
	"github.com/gu-io/gu/shell/cache/webcache"
)

// New returns a new cache based on the name provided.
// It attempts to use the new webkit CacheAPI using the
// github.com/gu-io/gu/shell/cache/webcache and if not possible
// resolves to using github.com/gu-io/gu/shell/cache/localcache
// which stores all requests into localstorage.
func New(cacheName string) shell.Cache {
	webCache, err := webcache.New()
	if err != nil {
		return localcache.New(cacheName)
	}

	var cache shell.Cache

	nameCache, err := webCache.Open(cacheName)
	if err != nil {
		if err == webcache.ErrInvalidState {
			cache = localcache.New(cacheName)
			return cache
		}

		panic(err)
	}

	cache = nameCache
	return cache
}
