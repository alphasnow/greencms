// @author AlphaSnow

package xcache

import (
	"context"
	"fmt"
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	redis_store "github.com/eko/gocache/store/redis/v4"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

func ProvideCache(conf *viper.Viper) (*cache.Cache[string], error) {
	var s store.StoreInterface
	var err error
	switch conf.GetString("cache.default") {
	case "redis":
		s, err = newRedisCache(conf)
		break
	case "memory":
		s, err = newMemoryCache(conf)
		break
	default:
		s, err = nil, fmt.Errorf("invalid value for cache.default: %s", conf.GetString("cache.default"))
	}

	if err != nil {
		return nil, err
	}
	return cache.New[string](s), nil
}

func newRedisCache(conf *viper.Viper) (store.StoreInterface, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("cache.stores.redis.addr"),
		Password: conf.GetString("cache.stores.redis.password"),
		DB:       conf.GetInt("cache.stores.redis.db"),
		// 默认0不限制
		MaxIdleConns: 10,
		//ConnMaxIdleTime: 30 * time.Minute,
		// 限制最大连接数, 控制连接占用资源
		//MaxActiveConns: 100,
		//ConnMaxLifetime: 1*time.Hour,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return redis_store.NewRedis(client), nil
}

func newMemoryCache(conf *viper.Viper) (store.StoreInterface, error) {
	cfg := &ristretto.Config{
		NumCounters: conf.GetInt64("cache.stores.memory.num_counters"),
		MaxCost:     conf.GetInt64("cache.stores.memory.max_cost"),
		BufferItems: conf.GetInt64("cache.stores.memory.buffer_items"),
	}
	ristrettoCache, err := ristretto.NewCache(cfg)
	if err != nil {
		return nil, err
	}
	ristrettoStore := ristretto_store.NewRistretto(ristrettoCache)
	return ristrettoStore, nil
}

// newBigCacheMemoryCache 废弃
// 只能接收 cache.New[[]byte](bigcacheStore)
//func newBigCacheMemoryCache(conf *viper.Viper) (store.StoreInterface, error) {
//	ctx := context.Background()
//	//expire := conf.GetInt("cache.stores.memory.expire")
//	//expireTime := time.Duration(expire) * time.Minute
//	bigcacheClient, err := bigcache.New(ctx, bigcache.DefaultConfig(10*time.Minute))
//	if err != nil {
//		return nil, err
//	}
//	return bigcache_store.NewBigcache(bigcacheClient), nil
//}
