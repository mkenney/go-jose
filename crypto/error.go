package crypto

import (
	"errors"
)

var (
	// ErrInvalidKey is returned when the provided key is invalid
	ErrInvalidKey = errors.New("key is invalid")

	// ErrInvalidSignature is returned when the provided signature is invalid
	ErrInvalidSignature = errors.New("signature is invalid")
)
