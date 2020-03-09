package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(source string) string {
	h := sha1.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
