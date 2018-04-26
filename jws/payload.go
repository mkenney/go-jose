package jws

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/mkenney/go-jose/jose"
)

/*
Payload represents the JOSE payload data
*/
type Payload struct {
	data        jose.Claims
	ContentType string
}

/*
Delete removes the specified Claim from the set for JSON payloads
*/
func (payload Payload) Delete(key string) error {
	if "application/json" != payload.ContentType {
		return fmt.Errorf("Delete() only applies to the 'application/json' content type")
	}
	return payload.data.Delete(key)
}

/*
Get retrieves a specified Claim Value for JSON payloads
*/
func (payload Payload) Get(key string) (interface{}, error) {
	if "application/json" != payload.ContentType {
		return nil, fmt.Errorf("Get() only applies to the 'application/json' content type")
	}
	return payload.data.Get(key)
}

/*
Has returns whether a specified Claim Name exists for JSON payloads
*/
func (payload Payload) Has(key string) bool {
	if "application/json" != payload.ContentType {
		return false
	}
	return payload.data.Has(key)
}

/*
Set sets the specified Claim Value for JSON payloads
*/
func (payload Payload) Set(key string, val interface{}) error {
	if "application/json" != payload.ContentType {
		return fmt.Errorf("Set() only applies to the 'application/json' content type")
	}
	return payload.data.Set(key, val)
}

/*
SetData sets the full data payload for non-JSON content types
*/
func (payload Payload) SetData(val interface{}) error {
	switch payload.ContentType {
	case "application/json":
		return payload.data.UnmarshalJSON(val.([]byte))
	default:
		payload.data.Set("data", val.([]byte))
		return nil
	}
}

/*
FromBase64 implements jose.Decoder
This implements RawURLEncodingBase64 -> JSON -> T
*/
func (payload *Payload) FromBase64(bytes []byte) error {
	if "application/json" == payload.ContentType {
		return payload.data.FromBase64(bytes)
	}
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(bytes)))
	base64.RawURLEncoding.Decode(bytes, buf)
	payload.data.Set("data", buf)
	return nil

}

/*
ToBase64 implements jose.Encoder
This implements T -> JSON -> RawURLEncodingBase64
*/
func (payload Payload) ToBase64() ([]byte, error) {
	if len(payload.data) == 0 {
		return nil, nil
	}

	if "application/json" == payload.ContentType {
		return payload.data.ToBase64()
	}

	data, err := payload.data.Get("data")
	if nil != err {
		return nil, err
	}

	bytes, err := json.Marshal(data.([]byte))
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
func (payload Payload) MarshalJSON() ([]byte, error) {
	if len(payload.data) == 0 {
		return nil, nil
	}

	if "application/json" == payload.ContentType {
		return payload.data.MarshalJSON()
	}

	data, err := payload.data.Get("data")
	if nil != err {
		return nil, err
	}

	bytes, err := json.Marshal(data.([]byte))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (payload *Payload) UnmarshalJSON(bytes []byte) error {
	if bytes == nil {
		return nil
	}

	if "application/json" == payload.ContentType {
		return payload.data.UnmarshalJSON(bytes)
	}

	buf := make([]byte, 0)
	err := json.Unmarshal(bytes, buf)
	if nil != err {
		return err
	}

	payload.data.Set("data", buf)
	return nil
}
