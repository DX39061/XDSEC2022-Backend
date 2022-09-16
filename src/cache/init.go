package cache

import (
	"XDSEC2022-Backend/src/config"
	"XDSEC2022-Backend/src/logger"
	"github.com/go-redis/redis"
	"sort"
	"strconv"
)

type RedisClient struct {
	Client *redis.Client
}

var caches = make(map[string]*RedisClient)

func Register(name string, cache *RedisClient) {
	caches[name] = cache
}

func Initialize() error {
	logger.Info("Initializing cache...")
	redisCfg := config.CacheConfig
	dbCount := 0
	var keys []string
	for k := range caches {
		keys = append(keys, k)
	}
	// sort to ensure consistent order
	sort.Strings(keys)
	for _, i := range keys {
		caches[i].Client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Host + ":" + strconv.Itoa(redisCfg.Port),
			Password: redisCfg.Password,
			DB:       dbCount,
		})
		_, err := caches[i].Client.Ping().Result()
		if err != nil {
			return err
		}
		dbCount++
	}

	logger.Info("Cache server initialized.")
	return nil
}
