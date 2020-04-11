package hibp

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// GetSHA1Hash takes a string as an input & returns SHA-1 Hash
func GetSHA1Hash(pass string) string {
	alg := sha1.New()
	alg.Write([]byte(pass))

	return strings.ToUpper(hex.EncodeToString(alg.Sum(nil)))
}
