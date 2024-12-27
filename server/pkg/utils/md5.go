package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str []byte) string {
	hash := md5.Sum(str)
	return hex.EncodeToString(hash[:])
}
