package generator

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateToken generates a random 32 byte token and encodes it as a base64 string.
func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	key := base64.StdEncoding.EncodeToString(bytes)

	return key, nil
}
