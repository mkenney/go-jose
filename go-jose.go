package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/mkenney/go-jose/crypto"
	"github.com/mkenney/go-jose/jws"
)

func main() {
	log.Debugf(crypto.HS256.Alg())

	header := jws.NewProtectedHeader()
	log.Debugf("Set result: %s", header.SetAlg("asdf"))

	tmp, err := header.MarshalJSON()
	log.Debugf("Claims: %s, %s", tmp, err)
}
