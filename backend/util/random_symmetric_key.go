package util

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomSymmetricKey() string {

	key := make([]byte, 32) // chacha20poly1305.KeySize = 32
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	return hex.EncodeToString(key)

}
