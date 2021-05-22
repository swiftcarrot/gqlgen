package main

import (
	"log"
	"net/http"

	"github.com/swiftcarrot/gqlgen/example/dataloader"
	"github.com/swiftcarrot/gqlgen/graphql/handler"
	"github.com/swiftcarrot/gqlgen/graphql/playground"
)

func main() {
	router := http.NewServeMux()

	router.Handle("/", playground.Handler("Dataloader", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		dataloader.NewExecutableSchema(dataloader.Config{Resolvers: &dataloader.Resolver{}}),
	))

	log.Println("connect to http://localhost:8082/ for graphql playground")
	log.Fatal(http.ListenAndServe(":8082", dataloader.LoaderMiddleware(router)))
}
