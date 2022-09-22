package cache

import (
	"XDSEC2022-Backend/src/config"
	"XDSEC2022-Backend/src/logger"
	"go.etcd.io/bbolt"
)

var tokenCache *bbolt.DB

func Initialize() error {
	logger.Info("Initializing cache...")
	cacheCfg := config.CacheConfig
	var err error
	tokenCache, err = bbolt.Open(cacheCfg.Path, 0600, nil)
	if err != nil {
		return err
	}
	err = tokenCache.Update(func(tx *bbolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("token"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	logger.Info("Cache server initialized.")
	return nil
}
