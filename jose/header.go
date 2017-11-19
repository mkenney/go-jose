package jose

import (
	"encoding/base64"
	"encoding/json"
)

/*
Header implements a JOSE Header with the addition and methods
*/
type Header struct {
	Claims *Claims
}

/*
Protected headers are base64-encoded after they're marshaled into JSON
*/
type Protected struct {
	Claims *Claims
}

/*
MarshalJSON implements json.Marshaler for Protected.
*/
func (protected Protected) MarshalJSON() ([]byte, error) {
	jsonBytes, err := json.Marshal(protected)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(jsonBytes)))
	base64.RawURLEncoding.Encode(buf, jsonBytes)
	return buf, nil
}

// UnmarshalJSON implements json.Unmarshaler for Protected.
func (protected *Protected) UnmarshalJSON(buf []byte) error {
	var header *Header

	jsonBytes := make([]byte, base64.RawURLEncoding.DecodedLen(len(buf)))
	_, err := base64.RawURLEncoding.Decode(jsonBytes, buf)
	if nil != err {
		return err
	}
	header = &Header{}
	if err := header.Claims.UnmarshalJSON(jsonBytes); err != nil {
		return err
	}
	protected = &Protected{header.Claims}

	return nil
}
