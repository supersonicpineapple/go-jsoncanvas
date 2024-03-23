package util

import (
	"crypto/rand"
	"encoding/hex"
)

func NewID() string {
	var b [8]byte

	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b[:])
}
