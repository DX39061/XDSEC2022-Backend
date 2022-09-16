package utility

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(v string) error {
	exist, err := PathExists(v)
	if err != nil {
		return err
	}
	if !exist {
		err = os.MkdirAll(v, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return err
}

func GetFileSHA256(filePath string) (string, error) {
	file, err := os.Open(filePath)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Hashing file close error: ", err)
		}
	}(file)
	if err != nil {
		log.Println("Hashing file open error: ", err)
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Println("Hashing file open error: ", err)
	}
	sum := hash.Sum(nil)
	res := strings.ToLower(fmt.Sprintf("%x", sum))
	return res, nil
}
