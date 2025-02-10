package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

func GenerateRandomState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return base64.URLEncoding.EncodeToString([]byte(time.Now().String()))
	}

	return base64.URLEncoding.EncodeToString(b)
}
