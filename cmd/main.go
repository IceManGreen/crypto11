package main

import (
	"github.com/ThalesGroup/crypto11/internal"
	"log"
)

func main() {
	log.Print("Hello world !")

	ctx, err := internal.ConfigureFromFile("../config")
	if err != nil {
		log.Panic(err)
	}

	attrs := internal.NewAttributeSet()
	_ = attrs.Set(internal.CkaLabel, "aes0")
	keys, err := ctx.FindKeysWithAttributes(attrs)
	if err != nil {
		log.Panic(err)
	}

	var key *internal.SecretKey
	key = keys[0]
	log.Print(key)

	if key != nil {
		log.Print("OK")
	}

}
