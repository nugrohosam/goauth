package infrastructure

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// Ring ..
var Ring *redis.Ring

// CacheConfigRing ..
var CacheConfigRing *cache.Cache

// InitiateRedisCache ..
func InitiateRedisCache(hosts map[string]string) {
	sizeCache := viper.GetInt("cache.redis.size")

	Ring = redis.NewRing(&redis.RingOptions{
		Addrs: hosts,
	})

	CacheConfigRing = cache.New(&cache.Options{
		Redis:      Ring,
		LocalCache: cache.NewTinyLFU(sizeCache, time.Minute),
	})
}

// StoreCacheRedis ..
func StoreCacheRedis(key string, data interface{}) error {
	ctx := context.TODO()

	err := CacheConfigRing.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: data,
		TTL:   time.Hour,
	})

	return err
}

// GetCacheRedis ..
func GetCacheRedis(key string) (interface{}, error) {
	var data string
	ctx := context.TODO()

	if err := CacheConfigRing.Get(ctx, key, data); err == nil {
		return data, nil
	} else {
		return nil, err
	}
}
