package main

import (
	"log"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = api.Generate(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
