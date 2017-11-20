package crypto

import (
	"encoding/base64"
	"encoding/json"
)

/*
Signature represents a JWS signature
*/
type Signature []byte

/*
MarshalJSON implements json.Marshaler
*/
func (s Signature) MarshalJSON() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if nil != err {
		return nil, err
	}
	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(jsonBytes))+2)
	buf[0] = '"'
	base64.RawURLEncoding.Encode(buf[1:], jsonBytes)
	buf[len(buf)-1] = '"'
	return buf, nil
}

/*
FromBase64 implements jose.Decoder
*/
func (s Signature) FromBase64(bytes []byte) error {
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(bytes)))
	base64.RawURLEncoding.Decode(bytes, buf)
	return s.UnmarshalJSON(buf)
}

/*
ToBase64 implements jose.Encoder
*/
func (s Signature) ToBase64() ([]byte, error) {
	bytes, err := s.MarshalJSON()
	if nil != err {
		return nil, err
	}
	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(bytes)))
	base64.RawURLEncoding.Encode(buf, bytes)
	return buf, nil
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (s *Signature) UnmarshalJSON(bytes []byte) error {
	if len(bytes) > 1 && bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(bytes)))
	_, err := base64.RawURLEncoding.Decode(buf, bytes)
	if nil != err {
		return err
	}
	*s = Signature(buf)
	return nil
}

//var (
//	_ json.Marshaler   = (Signature)(nil)
//	_ json.Unmarshaler = (*Signature)(nil)
//	_ jose.Encoder     = (Signature)(nil)
//)
