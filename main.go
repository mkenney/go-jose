package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/mkenney/go-jose/crypto"
)

func main() {
	level, _ := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	log.SetLevel(level)
	log.Debugf(crypto.HS256.Alg())
}
