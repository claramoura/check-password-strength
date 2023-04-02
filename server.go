package main

import (
	"log"
	"net/http"
	"os"
	"safepasswordverification/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
