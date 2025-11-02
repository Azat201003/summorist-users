package passwords

import (
	"golang.org/x/crypto/argon2"
)

const (
	MEMORY      = 64 * 1024
	ITERATIONS  = 3
	PARALLELISM = 3
	SALT        = "chikibambony"
	KEY_LENGTH  = 64
)

func Hash(password string) []byte {
	return argon2.IDKey(
		[]byte(password),
		[]byte(SALT),
		ITERATIONS,
		MEMORY,
		PARALLELISM,
		KEY_LENGTH,
	)
}

func Verify(hash []byte, password string) bool {
	return compareBytes(Hash(password), hash)
}

func compareBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
