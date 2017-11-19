package jose

import (
	"encoding/base64"
	"encoding/json"
)

/*
Claims implements the JWS and JWE Claims type
*/
type Claims map[string]interface{}

/*
Delete removes the specified Claim from the set
*/
func (claims Claims) Delete(key string) {
	delete(claims, key)
}

/*
Get retrieves a specified Claim Value
*/
func (claims Claims) Get(key string) interface{} {
	if claims == nil {
		return nil
	}
	return claims[key]
}

/*
Has returns whether a specified Claim Name exists
*/
func (claims Claims) Has(key string) bool {
	_, ok := claims[key]
	return ok
}

/*
Set sets the specified Claim Value
*/
func (claims Claims) Set(key string, val interface{}) {
	claims[key] = val
}

/*
FromBase64 implements jose.Decoder
This implements RawURLEncodingBase64 -> JSON -> T
*/
func (claims *Claims) FromBase64(bytes []byte) error {
	if bytes == nil {
		return nil
	}
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(bytes)))
	base64.RawURLEncoding.Decode(bytes, buf)
	return claims.UnmarshalJSON(buf)
}

/*
ToBase64 implements jose.Encoder
This implements T -> JSON -> RawURLEncodingBase64
*/
func (claims Claims) ToBase64() ([]byte, error) {
	bytes, err := claims.MarshalJSON()
	if nil != err {
		return nil, err
	}
	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(bytes)))
	base64.RawURLEncoding.Encode(buf, bytes)
	return buf, nil
}

/*
MarshalJSON implements json.Marshaler
*/
func (claims Claims) MarshalJSON() ([]byte, error) {
	if len(claims) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(map[string]interface{}(claims))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (claims *Claims) UnmarshalJSON(bytes []byte) error {
	if bytes == nil {
		return nil
	}
	return json.Unmarshal(bytes, (*map[string]interface{})(claims))
}
