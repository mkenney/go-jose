package jose

import (
	"crypto"

	jose_crypto "github.com/mkenney/go-jose/crypto"
)

/*
Decoder is satisfied if the type construct its data from a base64 string
*/
type Decoder interface {
	FromBase64() ([]byte, error)
}

/*
Encoder is satisfied if the type can encode it's data into a base64 string
*/
type Encoder interface {
	ToBase64() ([]byte, error)
}

/*
Signer provides the interfaceo for signing tokens
*/
type Signer interface {
	// Alg returns the name of the algorithm being used to sign a token
	Alg() string

	// Hasher returns the hashing algorithm itself
	Hasher() crypto.Hash

	// Sign returns a Signature for the raw content and as any errors that
	// occurred during the signing process
	Sign(raw []byte, key interface{}) (jose_crypto.Signature, error)

	// Verify accepts the raw content, the signature, and the key used to sign
	// and returns any errors found while validating the signature and content
	Verify(raw []byte, sig jose_crypto.Signature, key interface{}) error
}
