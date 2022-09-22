package cache

import (
	"errors"
	"go.etcd.io/bbolt"
	"strconv"
)

func ValidateToken(token string) bool {
	var val []byte
	err := tokenCache.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("token"))
		val = b.Get([]byte(token))
		return nil
	})
	if err != nil || val == nil {
		return false
	} else {
		return true
	}
}

func PermitToken(token string, userID uint) error {
	err := tokenCache.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("token"))
		err := b.Put([]byte(token), []byte(strconv.FormatUint(uint64(userID), 10)))
		return err
	})
	return err
}

func ExpireToken(token string) error {
	err := tokenCache.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("token"))
		userID := b.Get([]byte(token))
		if userID == nil {
			return errors.New("token not found")
		} else {
			err := b.Delete([]byte(token))
			return err
		}
	})
	return err
}
