package jwa

/*
String implements the Stringer interface for CompressionAlgorithm
*/
func (alg CompressionAlgorithm) String() string {
	return string(alg)
}

/*
String implements the Stringer interface for ContentEncryptionAlgorithm
*/
func (alg ContentEncryptionAlgorithm) String() string {
	return string(alg)
}

/*
String implements the Stringer interface for EllipticCurveAlgorithm
*/
func (crv EllipticCurveAlgorithm) String() string {
	return string(crv)
}

/*
Size returns the size of the EllipticCurveAlgorithm
*/
func (crv EllipticCurveAlgorithm) Size() int {
	switch crv {
	case P256:
		return 32
	case P384:
		return 48
	case P521:
		return 66
	}
	return 0
}

/*
String implements the Stringer interface for KeyEncryptionAlgorithm
*/
func (alg KeyEncryptionAlgorithm) String() string {
	return string(alg)
}

/*
String implements the Stringer interface for KeyType
*/
func (kty KeyType) String() string {
	return string(kty)
}

/*
String implements the Stringer interface for SignatureAlgorithm
*/
func (alg SignatureAlgorithm) String() string {
	return string(alg)
}
