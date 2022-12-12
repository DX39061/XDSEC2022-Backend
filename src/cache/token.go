package cache

import (
	"errors"
	"go.etcd.io/bbolt"
	"strconv"
	"time"
)

func ValidateEmail(email string, code string) bool {
    var val []byte
    err := tokenCache.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte("email"))
        val = b.Get([]byte(email))
        return nil
    })
    if err != nil || val==nil {
        return false
    }
    return string(val) == code
}

func PermitEmail(email string, code string) error {
    err := tokenCache.Update(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte("email"))
        err := b.Put([]byte(email), []byte(code))
        return err
    })
    return err
}

func DeleteEmail(email string) error {
    time.Sleep(5*time.Minute)
    err := tokenCache.Update(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte("email"))
        err := b.Delete([]byte(email))
        return err
    })
    return err
}

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
