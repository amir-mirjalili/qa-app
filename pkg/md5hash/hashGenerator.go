package md5hash

import (
	"crypto/md5"
	"encoding/hex"
)

func HashGenerator(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
