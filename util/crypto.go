package util

import (
	"crypto/sha512"
	"encoding/base64"
)

// Hash returns base64 encoded SHA512 salted hash of password as a string.
func Hash(password, salt string) string {
	h := sha512.Sum512([]byte(password + salt))
	return base64.StdEncoding.EncodeToString(h[:])
}
