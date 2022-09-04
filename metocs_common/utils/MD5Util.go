package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func Md5(length int) string {
	h := md5.New()
	h.Write([]byte(GetRandomString(length)))
	return hex.EncodeToString(h.Sum(nil))
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
