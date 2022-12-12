package cache

import (
	"go.etcd.io/bbolt"
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