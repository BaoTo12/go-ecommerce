package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashEmail(email string) string {
	hash := sha256.New()
	hash.Write([]byte(email))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
