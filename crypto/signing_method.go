package crypto

import (
	"crypto"
)

/*
SigningMethod provides the interfaceo for signing tokens
*/
type SigningMethod interface {
	// Alg returns the name of the algorithm being used to sign a token
	Alg() string

	// Verify accepts the raw content, the signature, and the key used to sign
	// and returns any errors found while validating the signature and content
	Verify(raw []byte, sig Signature, key interface{}) error

	// Sign returns a Signature for the raw content and as any errors that
	// occurred during the signing process
	Sign(raw []byte, key interface{}) (Signature, error)

	// Hasher returns the hashing algorithm itself
	Hasher() crypto.Hash
}
