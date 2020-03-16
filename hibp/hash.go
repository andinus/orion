package hibp

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

// GetHsh takes a string as an input & returns SHA-1 Hash
func GetHsh(pass string) string {
	alg := sha1.New()
	alg.Write([]byte(pass))

	return strings.ToUpper(hex.EncodeToString(alg.Sum(nil)))
}
