package jws

import (
	"net/url"

	"github.com/mkenney/go-jose/jose"
)

/*
NewHeader returns a pointer to a new Header instance
*/
func NewHeader() *Header {
	header := new(Header)
	header.Claims = jose.NewClaims()
	return header
}

/*
NewProtectedHeader returns a pointer to a new ProtectedHeader instance
*/
func NewProtectedHeader() *ProtectedHeader {
	protected := new(ProtectedHeader)
	protected.Claims = jose.NewClaims()
	return protected
}

/*
Header implements a JOSE Header with the addition and methods
*/
type Header struct {
	Claims jose.Claims
}

/*
ProtectedHeader headers are automatically base64-encoded after they're marshaled
into JSON
*/
type ProtectedHeader struct {
	Header
}

/*
DelAlg deletes the "alg" Claim
*/
func (header Header) DelAlg(val string) error {
	return header.Claims.Delete("alg")
}

/*
DelCrit deletes the "crit" Claim
*/
func (header Header) DelCrit(val []string) error {
	return header.Claims.Delete("crit")
}

/*
DelCty deletes the "cty" Claim
*/
func (header Header) DelCty(val string) error {
	return header.Claims.Delete("cty")
}

/*
DelEnc deletes the "enc" Claim
*/
func (header Header) DelEnc(val string) error {
	return header.Claims.Delete("enc")
}

/*
DelJku deletes the "jku" Claim
*/
func (header Header) DelJku(val string) error {
	return header.Claims.Delete("jku")
}

/*
DelJwk deletes the "jwk" Claim
*/
func (header Header) DelJwk(val string) error {
	return header.Claims.Delete("jwk")
}

/*
DelKid deletes the "kid" Claim
*/
func (header Header) DelKid(val string) error {
	return header.Claims.Delete("kid")
}

/*
DelTyp deletes the "typ" Claim
*/
func (header Header) DelTyp(val string) error {
	return header.Claims.Delete("typ")
}

/*
DelX5c deletes the "x5c" Claim
*/
func (header Header) DelX5c(val string) error {
	return header.Claims.Delete("x5c")
}

/*
DelX5t deletes the "x5t" Claim
*/
func (header Header) DelX5t(val string) error {
	return header.Claims.Delete("x5t")
}

/*
DelX5tS256 deletes the "x5t#S256" Claim
*/
func (header Header) DelX5tS256(val string) error {
	return header.Claims.Delete("x5t#S256")
}

/*
DelX5u deletes the "x5u" Claim
*/
func (header Header) DelX5u(val string) error {
	return header.Claims.Delete("x5u")
}

/*
DelZip deletes the "zip" Claim
*/
func (header Header) DelZip(val string) error {
	return header.Claims.Delete("zip")
}

/*
GetAlg gets the "alg" Claim
*/
func (header Header) GetAlg(val string) (interface{}, error) {
	return header.Claims.Get("alg")
}

/*
GetCrit gets the "crit" Claim
*/
func (header Header) GetCrit(val []string) (interface{}, error) {
	return header.Claims.Get("crit")
}

/*
GetCty gets the "cty" Claim
*/
func (header Header) GetCty(val string) (interface{}, error) {
	return header.Claims.Get("cty")
}

/*
GetEnc gets the "enc" Claim
*/
func (header Header) GetEnc(val string) (interface{}, error) {
	return header.Claims.Get("enc")
}

/*
GetJku gets the "jku" Claim
*/
func (header Header) GetJku(val string) (interface{}, error) {
	return header.Claims.Get("jku")
}

/*
GetJwk gets the "jwk" Claim
*/
func (header Header) GetJwk(val string) (interface{}, error) {
	return header.Claims.Get("jwk")
}

/*
GetKid gets the "kid" Claim
*/
func (header Header) GetKid(val string) (interface{}, error) {
	return header.Claims.Get("kid")
}

/*
GetTyp gets the "typ" Claim
*/
func (header Header) GetTyp(val string) (interface{}, error) {
	return header.Claims.Get("typ")
}

/*
GetX5c gets the "x5c" Claim
*/
func (header Header) GetX5c(val string) (interface{}, error) {
	return header.Claims.Get("x5c")
}

/*
GetX5t gets the "x5t" Claim
*/
func (header Header) GetX5t(val string) (interface{}, error) {
	return header.Claims.Get("x5t")
}

/*
GetX5tS256 gets the "x5t#S256" Claim
*/
func (header Header) GetX5tS256(val string) (interface{}, error) {
	return header.Claims.Get("x5t#S256")
}

/*
GetX5u gets the "x5u" Claim
*/
func (header Header) GetX5u(val string) (interface{}, error) {
	return header.Claims.Get("x5u")
}

/*
GetZip gets the "zip" Claim
*/
func (header Header) GetZip(val string) (interface{}, error) {
	return header.Claims.Get("zip")
}

/*
HasAlg checks for the "alg" Claim
*/
func (header Header) HasAlg(val string) bool {
	return header.Claims.Has("alg")
}

/*
HasCrit checks for the "crit" Claim
*/
func (header Header) HasCrit(val []string) bool {
	return header.Claims.Has("crit")
}

/*
HasCty checks for the "cty" Claim
*/
func (header Header) HasCty(val string) bool {
	return header.Claims.Has("cty")
}

/*
HasEnc checks for the "enc" Claim
*/
func (header Header) HasEnc(val string) bool {
	return header.Claims.Has("enc")
}

/*
HasJku checks for the "jku" Claim
*/
func (header Header) HasJku(val string) bool {
	return header.Claims.Has("jku")
}

/*
HasJwk checks for the "jwk" Claim
*/
func (header Header) HasJwk(val string) bool {
	return header.Claims.Has("jwk")
}

/*
HasKid checks for the "kid" Claim
*/
func (header Header) HasKid(val string) bool {
	return header.Claims.Has("kid")
}

/*
HasTyp checks for the "typ" Claim
*/
func (header Header) HasTyp(val string) bool {
	return header.Claims.Has("typ")
}

/*
HasX5c checks for the "x5c" Claim
*/
func (header Header) HasX5c(val string) bool {
	return header.Claims.Has("x5c")
}

/*
HasX5t checks for the "x5t" Claim
*/
func (header Header) HasX5t(val string) bool {
	return header.Claims.Has("x5t")
}

/*
HasX5tS256 checks for the "x5t#S256" Claim
*/
func (header Header) HasX5tS256(val string) bool {
	return header.Claims.Has("x5t#S256")
}

/*
HasX5u checks for the "x5u" Claim
*/
func (header Header) HasX5u(val string) bool {
	return header.Claims.Has("x5u")
}

/*
HasZip checks for the "zip" Claim
*/
func (header Header) HasZip(val string) bool {
	return header.Claims.Has("zip")
}

/*
SetAlg sets the "alg" Claim
*/
func (header Header) SetAlg(val string) error {
	return header.Claims.Set("alg", val)
}

/*
SetCrit sets the "crit" Claim
*/
func (header Header) SetCrit(val []string) error {
	return header.Claims.Set("crit", val)
}

/*
SetCty sets the "cty" Claim
*/
func (header Header) SetCty(val string) error {
	return header.Claims.Set("cty", val)
}

/*
SetEnc sets the "enc" Claim
*/
func (header Header) SetEnc(val string) error {
	return header.Claims.Set("enc", val)
}

/*
SetJku sets the "jku" Claim
*/
func (header Header) SetJku(val string) error {
	_, err := url.Parse(val)
	if nil != err {
		return err
	}
	return header.Claims.Set("jku", val)
}

/*
SetJwk sets the "jwk" Claim
*/
func (header Header) SetJwk(val string) error {
	return header.Claims.Set("jwk", val)
}

/*
SetKid sets the "kid" Claim
*/
func (header Header) SetKid(val string) error {
	return header.Claims.Set("kid", val)
}

/*
SetTyp sets the "typ" Claim
*/
func (header Header) SetTyp(val string) error {
	return header.Claims.Set("typ", val)
}

/*
SetX5c sets the "x5c" Claim
*/
func (header Header) SetX5c(val string) error {
	return header.Claims.Set("x5c", val)
}

/*
SetX5t sets the "x5t" Claim
*/
func (header Header) SetX5t(val string) error {
	return header.Claims.Set("x5t", val)
}

/*
SetX5tS256 sets the "x5t#S256" Claim
*/
func (header Header) SetX5tS256(val string) error {
	return header.Claims.Set("x5t#S256", val)
}

/*
SetX5u sets the "x5u" Claim
*/
func (header Header) SetX5u(val string) error {
	_, err := url.Parse(val)
	if nil != err {
		return err
	}
	return header.Claims.Set("x5u", val)
}

/*
SetZip sets the "zip" Claim
*/
func (header Header) SetZip(val string) error {
	return header.Claims.Set("zip", val)
}

/*
MarshalJSON implements json.Marshaler for Protected.
*/
func (header Header) MarshalJSON() ([]byte, error) {
	return header.Claims.MarshalJSON()
}

/*
UnmarshalJSON implements json.Unmarshaler for header.
*/
func (header *Header) UnmarshalJSON(bytes []byte) error {
	return header.Claims.UnmarshalJSON(bytes)
}

/*
MarshalJSON implements json.Marshaler for ProtectedHeader
ProtectedHeader is automatically base64 encoded after marshaling
*/
func (protected ProtectedHeader) MarshalJSON() ([]byte, error) {
	return protected.Claims.ToBase64()
}

/*
UnmarshalJSON implements json.Unmarshaler for ProtectedHeader
ProtectedHeader is automatically base64 decoded before unmarshaling
*/
func (protected *ProtectedHeader) UnmarshalJSON(bytes []byte) error {
	return protected.Claims.FromBase64(bytes)
}
