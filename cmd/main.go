package main

import (
	"github.com/ThalesGroup/crypto11/internal"
	"log"
)

func main() {
	log.Print("Hello world !")
	keyLabel := "aes0"

	ctx, err := internal.ConfigureFromFile("../config")
	if err != nil {
		log.Panic(err)
	}

	attrs := internal.NewAttributeSet()
	_ = attrs.Set(internal.CkaLabel, keyLabel)
	keys, err := ctx.FindKeysWithAttributes(attrs)
	if err != nil {
		log.Panic(err)
	}

	var key *internal.SecretKey
	if len(keys) == 0 {
		log.Panicf("Found no keys with label '%s' in pkcs11 store", keyLabel)
	}

	key = keys[0]
	if key != nil {
		log.Printf("Found key with label '%s' in pkcs11 store", keyLabel)
	}

}
