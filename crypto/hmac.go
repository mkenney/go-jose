package crypto

import (
	"crypto"
	"crypto/hmac"
	"encoding/json"
)

/*
HMAC implements the HMAC-SHA family of signing methods
*/
type HMAC struct {
	Name string
	Hash crypto.Hash
	_    struct{}
}

/*
Specific instances of signing methods
*/
var (
	// HS256 implements HS256.
	HS256 = &HMAC{
		Name: "HS256",
		Hash: crypto.SHA256,
	}

	// HS384 implements HS384.
	HS384 = &HMAC{
		Name: "HS384",
		Hash: crypto.SHA384,
	}

	// HS512 implements HS512.
	HS512 = &HMAC{
		Name: "HS512",
		Hash: crypto.SHA512,
	}
)

/*
Alg implements the Signature interface
*/
func (sig *HMAC) Alg() string {
	return sig.Name
}

/*
Hasher implements the Signature interface
*/
func (sig *HMAC) Hasher() crypto.Hash {
	return sig.Hash
}

/*
MarshalJSON implements the json.Marshaler interface
*/
func (sig *HMAC) MarshalJSON() ([]byte, error) {
	return []byte(`"` + sig.Alg() + `"`), nil
}

/*
Verify implements the Signature interface
*/
func (sig *HMAC) Verify(raw []byte, signature Signature, key interface{}) error {
	keyBytes, ok := key.([]byte)
	if !ok {
		return ErrInvalidKey
	}

	hasher := hmac.New(sig.Hash.New, keyBytes)
	hasher.Write(raw)
	if hmac.Equal(signature, hasher.Sum(nil)) {
		return nil
	}

	return ErrInvalidSignature
}

/*
Sign implements the Signature interface
*/
func (sig *HMAC) Sign(data []byte, key interface{}) (Signature, error) {
	keyBytes, ok := key.([]byte)
	if !ok {
		return nil, ErrInvalidKey
	}
	hasher := hmac.New(sig.Hash.New, keyBytes)
	hasher.Write(data)
	return Signature(hasher.Sum(nil)), nil
}

var _ json.Marshaler = (*HMAC)(nil)
