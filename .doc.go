/*
Package go-jose contains tools for implementing the various JOSE technologies
such as JWT, JWS, JWE, etc in Go.

JWS (https://tools.ietf.org/html/rfc7515)
JWE (https://tools.ietf.org/html/rfc7516)
JWK (https://tools.ietf.org/html/rfc7517)
JWA (https://tools.ietf.org/html/rfc7518)
JWT (https://tools.ietf.org/html/rfc7519)
Examples (https://tools.ietf.org/html/rfc7520)

The primary goal of this library is to provide a complete, easy to use,
Javascript Object Signing and Encryption (JOSE) implementation. Existing
libraries are either incomplete, poorly documented, heavily dependent on
3rd-party libraries, simply or difficult to implement.

	https://github.com/dgrijalva/jwt-go
	https://github.com/lestrrat/go-jwx
	https://github.com/SermoDigital/jose
	https://github.com/square/go-jose
*/
package go-jose

/*
Version describes the version of this library.
*/
const Version = "0.0.1"