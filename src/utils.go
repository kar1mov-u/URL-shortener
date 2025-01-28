package src

import (
	"crypto/sha256"
	"encoding/base64"
)

func Hashing(url string) string {
	hash := sha256.Sum256([]byte(url))
	short := base64.URLEncoding.EncodeToString(hash[:])
	return short[:8]
}
