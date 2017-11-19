package jose

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
