package g

import (
	"github.com/eko/gocache/lib/v4/cache"
	"server/pkg/xcache"
)

func Cache() (ca *cache.Cache[string]) {
	var err error
	if globalContainer.Has("cache") {
		ca = globalContainer.Get("cache").(*cache.Cache[string])
	} else {
		ca, err = xcache.ProvideCache(Config())
		if err != nil {
			panic(err)
		}
		globalContainer.Set("cache", ca)
	}
	return ca
}
