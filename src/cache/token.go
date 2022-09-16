package cache

import (
	"XDSEC2022-Backend/src/config"
	"XDSEC2022-Backend/src/logger"
	"strconv"
	"time"
)

var tokenCache RedisClient

func init() {
	Register("token", &tokenCache)
}

func ValidateToken(token string) bool {
	result, err := tokenCache.Client.Exists(token).Result()
	if err != nil {
		return false
	}
	return result == 1
}

func PermitToken(token string, userID uint) error {
	err := tokenCache.Client.Set(token, strconv.FormatUint(uint64(userID), 10), time.Duration(config.TokenConfig.ExpiresTime)*time.Second).Err()
	if err != nil {
		return err
	}
	return tokenCache.Client.SAdd(strconv.FormatUint(uint64(userID), 10), token).Err()
}

func ExpireToken(token string) error {
	userID, err := tokenCache.Client.Get(token).Result()
	if err != nil {
		return err
	}
	err = tokenCache.Client.Del(token).Err()
	if err != nil {
		return err
	}
	tokenCache.Client.SRem(userID, token)
	return err
}

func ExpireAllTokenOfUser(userID uint) (err error) {
	userTokens := tokenCache.Client.SMembers(strconv.FormatUint(uint64(userID), 10)).Val()
	for _, token := range userTokens {
		err = ExpireToken(token)
		if err != nil {
			logger.WarnFmt("expire one token of user %u failed: %s", userID, err.Error())
		} else {
			logger.InfoFmt("expired one token of user %u.", userID)
		}
	}
	return err
}
